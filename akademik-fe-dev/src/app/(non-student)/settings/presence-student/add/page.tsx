import React from "react";
import { Card, Col, Row } from "reactstrap";
import { FormPresenceStudent } from "../components/form-presence-student";
import { getServerSession } from "next-auth";
import authOptions from "@/config/next-auth";
import { AKADEMIK } from "@/lib/constants/role";
import { notFound } from "next/navigation";

export default async function CreatePresenceStudent() {
  const session = await getServerSession(authOptions);
  const role = session?.user?.role_name;

  if (role !== AKADEMIK) notFound();

  return (
    <Row>
      <Col>
        <Card className="p-3">
          <div className="d-flex flex-column w-100 gap-3 p-3 border rounded-3">
            <h1
              className="fs-4 mb-0 pb-3 fw-medium border-bottom"
              style={{ color: "#3A3A3A" }}
            >
              Tambah Presensi Mahasiswa
            </h1>
            <FormPresenceStudent />
          </div>
        </Card>
      </Col>
    </Row>
  );
}
