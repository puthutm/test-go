"use client"
import React from 'react'
import { Card, CardBody, CardTitle, CardText } from 'reactstrap';
import { VideoConferenceIcon } from "@/components/icons/video-conference";
import { AttachIcon } from '@/components/icons/attach';
import { InsertDriveFileIcon } from '@/components/icons/inser-drive-file';
import { CloseIcon } from '@/components/icons/close';

interface ChatPembimbingProps {
    // Definisikan tipe dari props di sini
    isClicked: boolean;
}


export const ChatPembimbing: React.FC<ChatPembimbingProps> = ({ isClicked }) => {

    if (!isClicked) return null;
    return (

        <>
            {/* Tombol Konsultasi dan Text Pembimbing Akademik */}
            <div className="d-flex border-2 border-bottom justify-content-between align-items-center pb-3 ">


                <div className="d-flex align-items-center">
                    <h5 className="card-title mb-1 me-2" >
                        Pembimbing Akademik
                    </h5>
                    <span className="me-2">&gt;</span>
                    <h5 className="card-title mb-1" style={{ fontWeight: "500" }}>
                        Konsultasi KRS
                    </h5>
                </div>
                <button
                    className="bg-primary rounded px-4 d-flex gap-1 align-items-center justify-content-center text-white py-3"
                    style={{ border: "1px solid #10487A", background: "#10487A", paddingBlock: "8px" }}
                >
                    <VideoConferenceIcon />
                    <span>Video Conference</span>
                </button>
            </div>

            {/* Komponen Chat */}
            <div className="mt-4">
                <div>
                    <div className="d-flex justify-content-end">
                        <Card className="mb-3" style={{ backgroundColor: "#f0fff0", width: "75%" }}>
                            <CardBody>
                                <CardTitle tag="h6" className="mb-2 text-right">
                                    230401020075 - Haerunnisa
                                </CardTitle>
                                <div className="d-flex justify-content-end">
                                    <small className="text-muted">6 Mar 2025, 10:15 WIB</small>
                                </div>
                                <CardText className="mt-2">
                                    Lorem ipsum dolor sit amet consectetur. Risus tortor arcu ut fringilla lobortis.. Et mauris dictum in urna consectetur nibh morbi.. Phasellus at tortor aenean vitae nisi eu nam semper tincidunt.. Id mauris tortor in massa morbi mauris vitae condimentum..
                                </CardText>
                            </CardBody>
                        </Card>
                    </div>

                    <div className="d-flex justify-content-start">
                        <Card className="mb-3" style={{ backgroundColor: "#fffacd", width: "75%" }}>
                            <CardBody>
                                <CardTitle tag="h6" className="mb-2 text-right">
                                    Ir. Henny Yulianti, M.M., M.Kom
                                </CardTitle>
                                <div className="d-flex justify-content-end">
                                    <small className="text-muted">6 Mar 2025, 20:00 WIB</small>
                                </div>
                                <CardText className="mt-2">
                                    Lorem ipsum dolor sit amet consectetur. Risus tortor arcu ut fringilla lobortis.. Et mauris dictum in urna consectetur nibh morbi.. Phasellus at tortor aenean vitae nisi eu nam semper tincidunt.. Id mauris tortor in massa morbi mauris vitae condimentum...
                                </CardText>
                            </CardBody>
                        </Card>
                    </div>
                </div>

                {/* Formulir Pesan */}
                <div className="d-flex  align-items-center">

                    {/* Tombol Upload */}
                    <label
                        htmlFor="uploadDokumen"
                        className="d-flex justify-content-center align-items-center border border-2 rounded me-2 py-2 px-3"
                        style={{ width: "48px", height: "45px", cursor: "pointer" }}
                    >
                        <div>
                            <AttachIcon />
                        </div>
                    </label>

                    {/* Input file yang tersembunyi */}
                    <input
                        type="file"
                        accept=".pdf"
                        id="uploadDokumen"
                        className="d-none"
                    />

                    {/* Input Text */}
                    <label htmlFor="text" className="w-100 me-2">
                        <input
                            type="text"
                            className="form-control me-2"
                            placeholder="Text"
                            style={{ height: "45px" }}
                        />
                    </label>

                    {/* Tombol Kirim */}
                    <label htmlFor="confirm">
                        <button
                            className="btn btn-primary"
                            style={{ width: "120px", height: "45px" }}
                        >
                            Kirim
                        </button>
                    </label>

                </div>
                <label
                    htmlFor="uploadDokumen"
                    className="d-flex align-items-center border border-2 rounded m-0"
                    style={{
                        width: "142px",
                        height: "25px",
                        cursor: "pointer",
                        borderRadius: "4px",
                        padding: "4px 8px",
                    }}
                >
                    <div className="d-flex align-items-center" style={{ gap: "8px"  }}>
                        <InsertDriveFileIcon style={{ color: "#909090" }} /> {/* Atur ukuran ikon */}
                        <span style={{ color: "#909090", fontSize: "12px", lineHeight: "16px", fontStyle: "italic" }}>Namafile.pdf</span> {/* Sesuaikan ukuran teks */}
                        <CloseIcon />
                    </div>
                </label>

            </div>
        </>
    )


}

export default ChatPembimbing