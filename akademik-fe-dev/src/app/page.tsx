"use client";

import React, { useState } from "react";
import { signIn } from "next-auth/react";
import { useRouter } from "next/navigation";
import {
  Card,
  CardBody,
  Button,
  Form,
  FormGroup,
  Label,
  Input,
  Alert,
  Spinner,
  Badge,
} from "reactstrap";
import { MAHASISWA, AKADEMIK, KAPRODI, DOSEN } from "@/lib/constants/role";

export default function Home() {
  const router = useRouter();
  const [username, setUsername] = useState("200101001");
  const [password, setPassword] = useState("password123");
  const [role, setRole] = useState<string>(MAHASISWA);
  const [useSSO, setUseSSO] = useState<boolean>(false);
  const [loading, setLoading] = useState(false);
  const [errorMsg, setErrorMsg] = useState<string | null>(null);

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setErrorMsg(null);

    try {
      if (useSSO) {
        const response = await signIn("sso", {
          callbackUrl: role === MAHASISWA ? "/student" : "/dashboard",
          redirect: false,
        });

        if (!response?.ok || response.error) {
          setErrorMsg("Gagal terhubung ke server SSO. Menggunakan login langsung...");
          setUseSSO(false);
          setLoading(false);
          return;
        }
      } else {
        const res = await signIn("credentials", {
          username: username.trim(),
          password: password.trim(),
          role: role,
          redirect: false,
        });

        if (res?.ok) {
          if (role === MAHASISWA) {
            router.push("/student");
          } else {
            router.push("/dashboard");
          }
        } else {
          setErrorMsg("Login gagal. Silakan periksa kembali username & password.");
        }
      }
    } catch (err: any) {
      setErrorMsg(err?.message || "Terjadi kesalahan saat melakukan autentikasi.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <main
      className="main-root p-3 d-flex justify-content-center align-items-center"
      style={{
        minHeight: "100vh",
        background: "linear-gradient(135deg, #0f172a 0%, #1e1b4b 50%, #0f172a 100%)",
        fontFamily: "'Inter', sans-serif",
      }}
    >
      <Card
        className="shadow-lg border-0"
        style={{
          width: "100%",
          maxWidth: "440px",
          borderRadius: "16px",
          background: "rgba(30, 41, 59, 0.85)",
          backdropFilter: "blur(12px)",
          border: "1px solid rgba(255, 255, 255, 0.1)",
          color: "#f8fafc",
        }}
      >
        <CardBody className="p-4 p-md-5">
          <div className="text-center mb-4">
            <div
              className="d-inline-flex align-items-center justify-content-center mb-3"
              style={{
                width: "60px",
                height: "60px",
                borderRadius: "14px",
                background: "linear-gradient(135deg, #6366f1 0%, #4f46e5 100%)",
                boxShadow: "0 8px 16px rgba(99, 102, 241, 0.3)",
              }}
            >
              <i className="ri-graduation-cap-line text-white" style={{ fontSize: "28px" }}></i>
            </div>
            <h4 className="fw-bold text-white mb-1">SIA AKADEMIK</h4>
            <p className="text-muted small mb-0">Universitas Siber Asia (UNSIA)</p>

            <div className="mt-3">
              <Badge
                color={useSSO ? "warning" : "info"}
                pill
                className="px-3 py-2 fw-normal mb-2"
                style={{ fontSize: "11px", letterSpacing: "0.5px" }}
              >
                {useSSO ? "🔒 Mode SSO Active" : "⚡ Direct Login (SSO Disabled)"}
              </Badge>
            </div>
          </div>

          <div className="mb-4">
            <Label className="small text-muted fw-medium d-block text-center mb-2" style={{ fontSize: "11px" }}>
              PILIK SEEDER AKUN DEMO (QUICK LOGIN)
            </Label>
            <div className="d-flex flex-wrap gap-1 justify-content-center">
              <Button
                size="sm"
                color={role === MAHASISWA ? "primary" : "outline-light"}
                style={{ fontSize: "11px", borderRadius: "6px" }}
                onClick={() => {
                  setRole(MAHASISWA);
                  setUsername("200101001");
                  setPassword("password123");
                }}
              >
                👨‍🎓 Mahasiswa
              </Button>
              <Button
                size="sm"
                color={role === DOSEN ? "primary" : "outline-light"}
                style={{ fontSize: "11px", borderRadius: "6px" }}
                onClick={() => {
                  setRole(DOSEN);
                  setUsername("0401018501");
                  setPassword("password123");
                }}
              >
                👨‍🏫 Dosen
              </Button>
              <Button
                size="sm"
                color={role === KAPRODI ? "primary" : "outline-light"}
                style={{ fontSize: "11px", borderRadius: "6px" }}
                onClick={() => {
                  setRole(KAPRODI);
                  setUsername("0415088202");
                  setPassword("password123");
                }}
              >
                👔 Kaprodi
              </Button>
              <Button
                size="sm"
                color={role === AKADEMIK ? "primary" : "outline-light"}
                style={{ fontSize: "11px", borderRadius: "6px" }}
                onClick={() => {
                  setRole(AKADEMIK);
                  setUsername("adminakademik");
                  setPassword("password123");
                }}
              >
                🏢 Admin
              </Button>
            </div>
          </div>

          {errorMsg && (
            <Alert color="danger" className="py-2 px-3 small border-0 mb-3" style={{ borderRadius: "8px" }}>
              {errorMsg}
            </Alert>
          )}

          <Form onSubmit={handleLogin}>
            <FormGroup className="mb-3">
              <Label className="small text-light fw-medium mb-1">Pilih Peran (Role)</Label>
              <Input
                type="select"
                value={role}
                onChange={(e: any) => setRole(e.target.value)}
                className="bg-dark text-white border-secondary"
                style={{ borderRadius: "8px", padding: "10px 14px" }}
              >
                <option value={MAHASISWA}>Mahasiswa</option>
                <option value={DOSEN}>Dosen</option>
                <option value={KAPRODI}>Ketua Program Studi (Kaprodi)</option>
                <option value={AKADEMIK}>Admin Akademik</option>
              </Input>
            </FormGroup>

            <FormGroup className="mb-3">
              <Label className="small text-light fw-medium mb-1">Username / NIM / NIDN</Label>
              <Input
                type="text"
                placeholder="Masukkan NIM/NIDN/Username"
                value={username}
                onChange={(e: any) => setUsername(e.target.value)}
                required
                className="bg-dark text-white border-secondary"
                style={{ borderRadius: "8px", padding: "10px 14px" }}
              />
            </FormGroup>

            <FormGroup className="mb-4">
              <Label className="small text-light fw-medium mb-1">Password</Label>
              <Input
                type="password"
                placeholder="Masukkan Password"
                value={password}
                onChange={(e: any) => setPassword(e.target.value)}
                required
                className="bg-dark text-white border-secondary"
                style={{ borderRadius: "8px", padding: "10px 14px" }}
              />
            </FormGroup>

            <Button
              type="submit"
              color="primary"
              disabled={loading}
              className="w-100 py-2 fw-semibold border-0 shadow-sm"
              style={{
                borderRadius: "8px",
                background: "linear-gradient(135deg, #4f46e5 0%, #4338ca 100%)",
                fontSize: "15px",
              }}
            >
              {loading ? (
                <span>
                  <Spinner size="sm" className="me-2" /> Masuk...
                </span>
              ) : (
                "Masuk Ke Sistem"
              )}
            </Button>
          </Form>

          <div className="mt-4 pt-2 border-top border-secondary text-center">
            <button
              type="button"
              onClick={() => setUseSSO(!useSSO)}
              className="btn btn-link text-muted p-0 text-decoration-none small"
              style={{ fontSize: "12px" }}
            >
              {useSSO ? "Ganti ke Direct Login (Lokal)" : "Gunakan SSO Service OAuth"}
            </button>
          </div>
        </CardBody>
      </Card>
    </main>
  );
}
