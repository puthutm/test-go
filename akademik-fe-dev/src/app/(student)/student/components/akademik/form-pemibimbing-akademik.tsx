"use client";

import Image from 'next/image'
import Avatar8 from "@/assets/images/users/avatar-8.jpg";
import { konsultasiList } from "@/lib/constants/table-pembimbing-data";
import { useState } from "react";
import { EditAkademikIcon } from "@/components/icons/edit-akademik";
import { EyeAkademikIcon } from "@/components/icons/eye-akademik";
import KonsultasiModal from '@/lib/hooks/use-modal-form-pembimbing';
import ChatPembimbing from './chat-pembimbing-akademik';

export const FormPembimbing = () => {
  const [modalOpen, setModalOpen] = useState<boolean>(false); // State untuk modal
  const [isChatVisible, setIsChatVisible] = useState<boolean>(false); // State untuk chat
  const toggleModal = () => setModalOpen(!modalOpen); // Fungsi toggle modal

  const handleEyeClick = (event: React.MouseEvent<HTMLButtonElement>) => {
    event.preventDefault(); // Cegah reload form
    setIsChatVisible(!isChatVisible); // Toggle tampilan ChatPembimbing

  };

  return (
    <>
      {/* Conditional Rendering: ChatPembimbing hanya muncul jika isChatVisible = true */}
      {isChatVisible && <ChatPembimbing isClicked={isChatVisible} />}

      <KonsultasiModal isOpen={modalOpen} toggle={toggleModal} />
      {!isChatVisible && (
        <>
          {/* Tombol Konsultasi dan Text Pembimbing Akademik */}
          <div className="d-flex border-2 border-bottom justify-content-between align-items-center">
            <h5 className="card-title py-3 mb-1" style={{ fontWeight: "500" }}>
              Pembimbing Akademik
            </h5>
            <button
              className="bg-transparent rounded px-3 d-flex gap-1 align-items-center justify-content-center text-primary"
              style={{ border: "1px solid #10487A", paddingBlock: "8px" }}
              onClick={toggleModal}
            >
              <EditAkademikIcon />
              <span>Konsultasi</span>
            </button>
          </div>

          <form>
            {/* ini adalah card dosen */}
            <div className="card my-4" style={{ borderRadius: "8px", border: "1px solid #ccc" }}>
              <div className="py-3 px-4 d-flex align-items-center">
                <Image
                  src={Avatar8}
                  alt="Foto Profil"
                  width={45}
                  height={45}
                  className=" me-3" // Menambahkan margin kanan
                  style={{ objectFit: 'cover' }}
                />
                <div>
                  <p className="mb-0" style={{ fontWeight: "500", color: "#909090" }}>2320206034</p>
                  <h5 className="mb-0" style={{ fontWeight: "500", color: "#495057" }}>Ir. Henny Yulianti, M.M., M.Kom</h5>
                </div>
              </div>
            </div>

            {/* ini adalah table */}
            <div className="table-responsive">
              <table className="table" style={{ tableLayout: 'fixed' }}>
                <thead style={{ backgroundColor: '#DEE5EC' }}>
                  <tr>
                    <th scope="col" className='text-center' style={{ width: '150px' }}>Waktu Dibuat </th>
                    <th scope="col" className='text-center' >Topik Konsultasi </th>
                    <th scope="col" className='text-center' style={{ width: '100px' }}>Aksi</th>
                  </tr>
                </thead>
                <tbody>
                  {konsultasiList.map((konsultasi, index) => (
                    <tr key={index} className="">
                      <td className="mb-0 border-0 text-center">{konsultasi.waktuDibuat}</td>
                      <td className="border-0 text-start">{konsultasi.topikKonsultasi}</td>
                      <td className="border-0 text-center">
                        <button
                          className="border-0 bg-transparent p-0"
                          aria-label="Edit Akademik"
                          onClick={handleEyeClick} // Memanggil fungsi toggle chat
                        >
                          <EyeAkademikIcon />
                        </button>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </form>
        </>
      )}
    </>
  );
};


