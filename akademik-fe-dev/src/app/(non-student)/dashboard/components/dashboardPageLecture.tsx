"use client";
import React from "react";

import {
  Row,
  Col,
  Card,
  CardHeader,
  CardBody,
  Progress,
  Alert,
} from "reactstrap";
import { CoPresentIcon } from "@/components/icons/co-present";
import { WorkIcon } from "@/components/icons/work";
import { Deversity3Icon } from "@/components/icons/diversity-3";
import { InventoryIcon } from "@/components/icons/inventory";
import { ScoreIcon } from "@/components/icons/score";
import { AutoStoriesIcon } from "@/components/icons/auto-stories";

import Link from "next/link";

// === Recharts ===
import {
  ResponsiveContainer,
  ComposedChart,
  Line,
  Bar,
  XAxis,
  YAxis,
  Tooltip,
  Legend,
  CartesianGrid,
  PieChart,
  Pie,
  Cell,
} from "recharts";

function DashboardPageLecture() {
  // --- Dummy data (bisa diganti dari API) ---
  const weeklyTeachData = [
    { week: "W1", pertemuan: 5, sks: 8 },
    { week: "W2", pertemuan: 7, sks: 10 },
    { week: "W3", pertemuan: 6, sks: 9 },
    { week: "W4", pertemuan: 8, sks: 12 },
  ];

  const guidanceStatus = [
    { name: "Terjadwal", value: 120 },
    { name: "Menunggu", value: 80 },
    { name: "Selesai", value: 100 },
    { name: "Dibatalkan", value: 20 },
  ];

  return (
    <section className="position-relative">
      {/*//! TOP SECTION */}
      <Row className="pb-0" style={{ gap: "10px" }}>
        {/*//! card teaching activity */}
        <Col className="p-0 m-0" sm={3}>
          <Card className="p-3 m-0">
            <CardHeader className="p-0 d-flex align-items-center gap-2 border-0">
              <div
                className="rounded-circle p-2 d-flex justify-content-center align-items-center"
                style={{ background: "#DEE5EC", width: "36px", height: "36px" }}
              >
                <CoPresentIcon />
              </div>
              <h2
                className=" m-0 p-0 flex-grow-1 fw-bold"
                style={{ color: "#3A3A3A", fontSize: "14px" }}
              >
                Aktivitas Mengajar
              </h2>
            </CardHeader>
            <CardBody className="pt-2 pb-0 px-0">
              <p
                className="m-0 p-0"
                style={{ color: "#495057", fontSize: "14px" }}
              >
                Anda mengajar 7 kelas di semester ini dengan total 20 SKS.
              </p>
              <button
                className="btn w-100 mt-2"
                style={{ border: "1px solid #10487A", color: "#10487A" }}
              >
                Lihat Detail
              </button>
            </CardBody>
          </Card>
        </Col>

        {/*//! card  teaching weigth*/}
        <Col className="p-0 m-0" sm={3}>
          <Card className="p-3 m-0">
            <CardHeader className="p-0 d-flex align-items-center gap-2 border-0">
              <div
                className="rounded-circle p-2 d-flex justify-content-center align-items-center"
                style={{ background: "#DEE5EC", width: "36px", height: "36px" }}
              >
                <WorkIcon />
              </div>
              <h2
                className=" m-0 p-0 flex-grow-1 fw-bold"
                style={{ color: "#3A3A3A", fontSize: "14px" }}
              >
                Beban Mengajar
              </h2>
            </CardHeader>
            <CardBody className="pt-2 pb-0 px-0">
              <p
                className="m-0 p-0 text-center"
                style={{ color: "#495057", fontSize: "14px" }}
              >
                10 dari 16 SKS
              </p>
              <Progress
                max="16"
                value="10"
                barStyle={{ background: "#FFA21D" }}
                barClassName="rounded-pill"
                className="rounded-pill"
                style={{ padding: "3px", height: "16px" }}
              />
              <button
                className="btn w-100 mt-2"
                style={{ border: "1px solid #10487A", color: "#10487A" }}
              >
                Lihat Detail
              </button>
            </CardBody>
          </Card>
        </Col>

        {/*//! card  bimbingan*/}
        <Col className="p-0 m-0">
          <Card className="p-3 m-0">
            <CardHeader className="p-0 d-flex align-items-center gap-2 border-0">
              <div
                className="rounded-circle p-2 d-flex justify-content-center align-items-center"
                style={{ background: "#DEE5EC", width: "36px", height: "36px" }}
              >
                <Deversity3Icon />
              </div>
              <h2
                className=" m-0 p-0 flex-grow-1 fw-bold"
                style={{ color: "#3A3A3A", fontSize: "14px" }}
              >
                Mahasiswa Bimbingan
              </h2>
            </CardHeader>
            <CardBody className="pt-2 pb-0 px-0">
              <section className="d-flex gap-2 align-items-center">
                <div className="flex-grow-1">
                  <h2
                    className="m-0 p-0 fw-semibold"
                    style={{ color: "#495057", fontSize: "12px" }}
                  >
                    Total: 320 Mahasiswa
                  </h2>
                  <Progress multi style={{ padding: "3px", height: "16px" }}>
                    <Progress
                      bar
                      value="50"
                      multi
                      barStyle={{
                        background: "#FFA21D",
                        borderTopLeftRadius: "10px",
                        borderBottomLeftRadius: "10px",
                      }}
                    />
                    <Progress
                      bar
                      value="70"
                      multi
                      barStyle={{
                        background: "#10487A",
                        borderTopRightRadius: "10px",
                        borderBottomRightRadius: "10px",
                      }}
                    />
                  </Progress>
                </div>
                <div className="d-flex flex-column gap-1">
                  <div className="d-flex align-items-center gap-1">
                    <div
                      className="rounded-1"
                      style={{ background: "#FFA21D", width: 16, height: 16 }}
                    />
                    <h2
                      className="m-0 p-0 fw-semibold"
                      style={{ color: "#495057", fontSize: "12px" }}
                    >
                      Mhs wali
                    </h2>
                  </div>
                  <div className="d-flex align-items-center gap-1">
                    <div
                      className="rounded-1"
                      style={{ background: "#10487A", width: 16, height: 16 }}
                    />
                    <h2
                      className="m-0 p-0 fw-semibold"
                      style={{ color: "#495057", fontSize: "12px" }}
                    >
                      Mhs TA
                    </h2>
                  </div>
                </div>
              </section>

              <section className="w-100 d-flex gap-1">
                <button
                  className="btn w-100 flex-grow-1 mt-2"
                  style={{ border: "1px solid #10487A", color: "#10487A" }}
                >
                  Detail Mahasiswa Wali
                </button>
                <button
                  className="btn w-100 flex-grow-1 mt-2"
                  style={{ border: "1px solid #10487A", color: "#10487A" }}
                >
                  Detail Mahasiswa TA
                </button>
              </section>
            </CardBody>
          </Card>
        </Col>
      </Row>

      {/*//! BOTTOM SECTION */}
      <Row className="mt-3" style={{ gap: "10px" }}>
        {/*//! left content */}
        <Col sm={8} className="p-0">
          {/* === Charts Section === */}
          <Card className="p-3 m-0">
            <CardHeader className="m-0 p-0 border-0 d-flex align-items-center">
              <h2
                className="m-0 p-0 flex-grow-1 fw-semibold"
                style={{ color: "#3A3A3A", fontSize: "16px" }}
              >
                Aktivitas Mengajar (Mingguan)
              </h2>
            </CardHeader>
            <CardBody className="px-0">
              <div style={{ width: "100%", height: 260 }}>
                <ResponsiveContainer>
                  <ComposedChart data={weeklyTeachData}>
                    <CartesianGrid strokeDasharray="3 3" />
                    <XAxis dataKey="week" />
                    <YAxis />
                    <Tooltip />
                    <Legend />
                    {/* warna bar */}
                    <Bar
                      dataKey="pertemuan"
                      name="Pertemuan"
                      fill="#FFA21D"
                      radius={[6, 6, 0, 0]}
                    />
                    {/* warna line */}
                    <Line
                      type="monotone"
                      dataKey="sks"
                      name="SKS"
                      stroke="#10487A"
                      strokeWidth={2}
                      dot={{ fill: "#10487A" }}
                    />
                  </ComposedChart>
                </ResponsiveContainer>
              </div>
            </CardBody>
          </Card>
          <Card className="p-3 m-0 mt-3">
            <CardHeader className="m-0 p-0 border-0 d-flex align-items-center">
              <h2
                className="m-0 p-0 flex-grow-1 fw-semibold"
                style={{ color: "#3A3A3A", fontSize: "16px" }}
              >
                Status Bimbingan
              </h2>
            </CardHeader>
            <CardBody className="px-0">
              <div style={{ width: "100%", height: 260 }}>
                <ResponsiveContainer>
                  <PieChart>
                    <Tooltip />
                    <Legend />
                    <Pie
                      data={guidanceStatus}
                      dataKey="value"
                      nameKey="name"
                      cx="50%"
                      cy="50%"
                      outerRadius={90}
                      label
                    >
                      {guidanceStatus.map((_, i) => {
                        const COLORS = [
                          "#FFA21D",
                          "#10487A",
                          "#0AB39C",
                          "#F06548",
                        ];
                        return (
                          <Cell
                            key={`cs-${i}`}
                            fill={COLORS[i % COLORS.length]}
                          />
                        );
                      })}
                    </Pie>
                  </PieChart>
                </ResponsiveContainer>
              </div>
            </CardBody>
          </Card>
        </Col>

        {/*//! rigth content */}
        <Col className="p-0">
          <section className="position-relative">
            {/*//! card calendar information */}
            <Card className="p-3 m-0">
              <CardHeader className="m-0 p-0 border-0 d-flex align-items-center">
                <h2
                  className="m-0 p-0 flex-grow-1 fw-semibold"
                  style={{ color: "#3A3A3A", fontSize: "16px" }}
                >
                  Kalendar
                </h2>
                <Link href={"/dashboard"} className="m-0 p-0 fw-semibold">
                  Lihat Detail
                </Link>
              </CardHeader>

              <CardBody className="m-0 p-0 mt-3 d-flex flex-column gap-3 ">
                <Alert
                  style={{ background: "#6CBE401A" }}
                  className="w-100 py-2 px-3  m-0 border-0 rounded-2 overflow-hidden"
                  fade={false}
                >
                  <div
                    style={{
                      position: "absolute",
                      width: "2px",
                      left: "0",
                      top: "0",
                      bottom: "0",
                      background: "#6CBE40",
                    }}
                  />
                  <h2
                    className="m-0 p-0 fw-semibold"
                    style={{ fontSize: "16px" }}
                  >
                    Hari Ini
                  </h2>
                  <div
                    className="mt-2 d-flex flex-column gap-1"
                    style={{ color: "#909090" }}
                  >
                    <p className="m-0 p-0">
                      10.00 - 12.00 : Pengantar Sist. Informasi (W3)
                    </p>
                    <p className="m-0 p-0">
                      10.00 - 12.00 : Pengantar Sist. Informasi (W3)
                    </p>
                  </div>
                </Alert>

                <Alert
                  style={{ background: "#0AB39C1A" }}
                  className="w-100 py-2 px-3  m-0 border-0 rounded-2 overflow-hidden"
                  fade={false}
                >
                  <div
                    style={{
                      position: "absolute",
                      width: "2px",
                      left: "0",
                      top: "0",
                      bottom: "0",
                      background: "#0AB39C",
                    }}
                  />
                  <h2
                    className="m-0 p-0 fw-semibold"
                    style={{ fontSize: "16px" }}
                  >
                    Minggu Ini
                  </h2>
                  <div
                    className="mt-2 d-flex flex-column gap-1"
                    style={{ color: "#909090" }}
                  >
                    <p className="m-0 p-0">
                      21 - 25 April : Validasi KRS Mahasiswa
                    </p>
                  </div>
                </Alert>

                <Alert
                  style={{ background: "#F065481A" }}
                  className="w-100 py-2 px-3  m-0 border-0 rounded-2 overflow-hidden"
                  fade={false}
                >
                  <div
                    style={{
                      position: "absolute",
                      width: "2px",
                      left: "0",
                      top: "0",
                      bottom: "0",
                      background: "#F06548",
                    }}
                  />
                  <h2
                    className="m-0 p-0 fw-semibold"
                    style={{ fontSize: "16px" }}
                  >
                    Bulan Ini
                  </h2>
                  <div
                    className="mt-2 d-flex flex-column gap-1"
                    style={{ color: "#909090" }}
                  >
                    <p className="m-0 p-0">18 April : Wafat Isa Almasih</p>
                  </div>
                </Alert>
              </CardBody>
            </Card>

            {/*//! card tanggungan dosen */}
            <Card className="p-3 m-0 mt-3">
              <CardHeader className="m-0 p-0 border-0 d-flex align-items-center">
                <h2
                  className="m-0 p-0 flex-grow-1 fw-semibold"
                  style={{ color: "#3A3A3A", fontSize: "16px" }}
                >
                  Tanggungan Dosen
                </h2>
              </CardHeader>

              <CardBody className="m-0 p-0 mt-3 d-flex flex-column gap-3 ">
                <Card className="p-3 m-0 border-1">
                  <CardHeader className="p-0 d-flex align-items-center gap-2 border-0">
                    <div className="rounded-circle d-flex justify-content-center align-items-center px-1">
                      <AutoStoriesIcon />
                    </div>
                    <h2
                      className=" m-0 p-0 flex-grow-1 fw-bold"
                      style={{ color: "#3A3A3A", fontSize: "14px" }}
                    >
                      Realisasi Perkuliahan
                    </h2>
                  </CardHeader>
                  <CardBody className="pt-2 pb-0 px-0">
                    <p
                      className="m-0 p-0"
                      style={{ color: "#495057", fontSize: "14px" }}
                    >
                      Terdapat 7 pertemuan yang belum Anda isi realisasinya.
                    </p>
                    <button
                      className="btn w-100 mt-2"
                      style={{
                        border: "1px solid #495057",
                        color: "#495057",
                        background: "#F3F6F9",
                      }}
                    >
                      Lihat Detail
                    </button>
                  </CardBody>
                </Card>

                <Card className="p-3 m-0 border-1">
                  <CardHeader className="p-0 d-flex align-items-center gap-2 border-0">
                    <div className="rounded-circle d-flex justify-content-center align-items-center px-1">
                      <InventoryIcon />
                    </div>
                    <h2
                      className=" m-0 p-0 flex-grow-1 fw-bold"
                      style={{ color: "#3A3A3A", fontSize: "14px" }}
                    >
                      Presensi Mahasiswa
                    </h2>
                  </CardHeader>
                  <CardBody className="pt-2 pb-0 px-0">
                    <p
                      className="m-0 p-0"
                      style={{ color: "#495057", fontSize: "14px" }}
                    >
                      Presensi belum diisi untuk 8 kelas pertemuan.
                    </p>
                    <button
                      className="btn w-100 mt-2"
                      style={{
                        border: "1px solid #495057",
                        color: "#495057",
                        background: "#F3F6F9",
                      }}
                    >
                      Lihat Detail
                    </button>
                  </CardBody>
                </Card>

                <Card className="p-3 m-0 border-1">
                  <CardHeader className="p-0 d-flex align-items-center gap-2 border-0">
                    <div className="rounded-circle d-flex justify-content-center align-items-center px-1">
                      <ScoreIcon />
                    </div>
                    <h2
                      className=" m-0 p-0 flex-grow-1 fw-bold"
                      style={{ color: "#3A3A3A", fontSize: "14px" }}
                    >
                      Pengisian Nilai
                    </h2>
                  </CardHeader>
                  <CardBody className="pt-2 pb-0 px-0">
                    <p
                      className="m-0 p-0"
                      style={{ color: "#495057", fontSize: "14px" }}
                    >
                      Ada 8 nilai yang belum dikunci.
                    </p>
                    <button
                      className="btn w-100 mt-2"
                      style={{
                        border: "1px solid #495057",
                        color: "#495057",
                        background: "#F3F6F9",
                      }}
                    >
                      Lihat Detail
                    </button>
                  </CardBody>
                </Card>
              </CardBody>
            </Card>
          </section>
        </Col>
      </Row>
    </section>
  );
}

export default DashboardPageLecture;
