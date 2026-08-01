// LaporanKerjaView.tsx
"use client";

import React, { useState } from "react";
import { FilledLaporanKP } from "../filled-laporan-kp";
import { EmptyStateLaporanKP } from "../empty-state-laporan-kp";
import { FormLaporanKP } from "../form-laporan-kp";

interface LaporanKP {
  judul: string;
  sertifikat: string;
  laporan: string;
}

export default function LaporanKerjaView() {
  const [hasLaporanKP, setHasLaporanKP] = useState<boolean>(true); // Set to true to show filled form by default
  const [isModalOpen, setIsModalOpen] = useState<boolean>(false);
  const [activeLaporan, setActiveLaporan] = useState<LaporanKP>({
    judul: "Text", // Default value matching the image
    sertifikat: "SertifikatPembicara2.pdf",
    laporan: "LaporanKP.pdf",
  });

  const toggleModal = () => {
    setIsModalOpen(!isModalOpen);
  };

  const handleSubmit = (data: {
    judul: string;
    laporan: File | null;
    sertifikat: File | null;
  }) => {
    // Process form data
    setActiveLaporan({
      judul: data.judul,
      laporan: data.laporan ? data.laporan.name : activeLaporan.laporan,
      sertifikat: data.sertifikat
        ? data.sertifikat.name
        : activeLaporan.sertifikat,
    });
    setHasLaporanKP(true);
  };

  return (
    <div className="container p-4 pt-0">
      <div className="border-2 border-bottom mb-4">
        <h2
          className={`card-title fw-medium py-3 mb-0`}
          style={{ color: "#495057" }}
        >
          Laporan Kerja Praktik
        </h2>
      </div>
      {hasLaporanKP ? (
        <FilledLaporanKP data={activeLaporan} openEditModal={toggleModal} />
      ) : (
        <EmptyStateLaporanKP openModal={toggleModal} />
      )}

      <FormLaporanKP
        isOpen={isModalOpen}
        toggle={toggleModal}
        onSubmit={handleSubmit}
        initialData={hasLaporanKP ? activeLaporan : undefined}
      />
    </div>
  );
}
