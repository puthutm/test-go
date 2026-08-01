"use client";
import React from "react";
import { useSearchParams } from "next/navigation";
import { Card, CardHeader, CardBody, Alert } from "reactstrap";
import TabsSectionLecture from "../tabs";
import InformationSubject from "../information-subject";
import { ErrorIcon } from "@/components/icons/error";

function LayoutDetailSubject({ children }: { children: React.ReactNode }) {
  const searchParams = useSearchParams();

  return (
    <section className="position-relative">
      <Card className="py-3 px-4 rounded-3 bg-white">
        {/*//! card header */}
        <CardHeader className="p-0 m-0 pb-2 pt-2 border-3 d-flex align-items-center gap-1">
          {/* //! text */}
          <h2 className="m-0 p-0 fs-5 fw-semibold flex-grow-1">
            Detail Mata Kuliah
          </h2>
        </CardHeader>

        {/*//! card body */}
        <CardBody className="position-relative px-0">
          {/* alert information detail subject */}
          {searchParams.get("tab") === "subject" && (
            <section className="position-relative   p-0 mb-2">
              <Alert
                color="warning"
                className="w-100 py-2 px-3 d-flex gap-2 align-items-center m-0"
                fade={false}
              >
                {/*line */}
                <div
                  style={{
                    position: "absolute",
                    width: "2px",
                    left: "0",
                    top: "0",
                    bottom: "0",
                    background: "#F7B84B",
                  }}
                />
                <div
                  className="p-0 m-0 d-flex align-items-center"
                  style={{ width: "16px", height: "16px" }}
                >
                  <ErrorIcon width="100%" height="100%" />
                </div>
                <p className="m-0 p-0 flex-grow-1">
                  Kode, Nama & Total SKS tidak bisa diubah karena mata kuliah
                  sudah digunakan di kurikulum S1 PJJ Sistem Informasi
                </p>
              </Alert>
            </section>
          )}
          {/*//!tabs */}
          <section className="position-relative ">
            <TabsSectionLecture />
          </section>
          {/*//! information subject */}
          {searchParams.get("tab") != null &&
            searchParams.get("tab") !== "subject" && (
              <section className="position-relative">
                <InformationSubject />
              </section>
            )}
          {/*//! Content */}
          {children}
        </CardBody>
      </Card>
    </section>
  );
}

export default LayoutDetailSubject;
