"use client";

import { Col, Input, Label, Row } from "reactstrap";
import { useForm, Controller } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

import {
    FormSeminarSchema,
    FormSeminarSchemaType,
} from "@/lib/validations/students/kp/form-seminar-schema";
import { FormErrorMessage } from "@/components/ui/form-error-message";
import { FormDescription } from "@/components/ui/form-description";
import { ErrorIcon } from "@/components/icons/error";
import { FileUploadIcon } from "@/components/icons/file-upload";

export const FormSeminar = () => {


    const {
        formState: { errors },
        handleSubmit,
        control,
    } = useForm<FormSeminarSchemaType>({
        resolver: zodResolver(FormSeminarSchema),
        defaultValues: {
            //   email_kampus: "",
            //   email_pribadi: "",
            //   no_hp: "",
        },
    });

    return (
        <>

            <div className="d-flex border-2 border-bottom justify-content-between align-items-center">
                <h5 className="card-title py-3 mb-1" style={{ fontWeight: "500" }}>
                    Seminar Kerja Praktik
                </h5>
            </div>

            {/* Warning message */}
            <div className="alert alert-warning alert-border-left d-inline-flex gap-2 mt-4" role="alert">
                <ErrorIcon color="#A66900 " style={{ width: "16px", height: "16px" }} />
                <p className="text-sm mb-0">Jadwal seminar akan ditentukan setelah semua dokumen seminar telah lengkap. Silakan periksa dan lengkapi data yang diperlukan</p>
            </div>

            <form
                onSubmit={handleSubmit((data) => console.log(data))}
                className="my-2"
            >
                <Row className="gap-1 gap-lg-0">
                    {/* left section */}
                    <Col md={12} lg={6}>
                        <Row className="gap-2">

                            {/* Lembar Pengesahan */}
                            <Col sm={12}>
                                <Row
                                    className="align-items-center gap-2"
                                    style={{ paddingBottom: "10px" }}
                                >
                                    <Col sm={12}>
                                        <Label
                                            htmlFor="lembar_penilaian"
                                            className="form-label mb-0 fw-medium"
                                        >
                                            Lembar Pengesahan Seminar KP
                                        </Label>
                                    </Col>
                                    <Col sm={12}>
                                        <div className="d-flex gap-2">
                                            <Controller
                                                name="lembar_pengesahan"
                                                control={control}
                                                render={({ field: { onChange, value } }) => (
                                                    <div className="position-relative w-100">
                                                        <label
                                                            htmlFor="lembar_pengesahan"
                                                            className="d-flex align-items-center border border-1 rounded p-3 m-0"
                                                            style={{ width: "100%", height: "45px", cursor: "pointer" }}
                                                        >
                                                            <span style={{ color: "#909090" }}>
                                                                {value ? value.name : "Upload File"}
                                                            </span>
                                                        </label>

                                                        <Input
                                                            type="file"
                                                            id="lembar_pengesahan"
                                                            accept=".pdf"
                                                            onChange={(e) => {
                                                                const file = e.target.files?.[0];
                                                                onChange(file);
                                                            }}
                                                            hidden
                                                        />

                                                        {errors.lembar_pengesahan ? (
                                                            <FormErrorMessage errors={errors.lembar_pengesahan} />
                                                        ) : (
                                                            <FormDescription message="File dengan format .pdf max 10MB" />

                                                        )}
                                                    </div>
                                                )}

                                            />
                                            <label
                                                htmlFor="lembar_pengesahan"
                                                className="btn d-flex align-items-center btn-light mb-0"
                                                style={{ whiteSpace: "nowrap" }}
                                            >
                                                <FileUploadIcon className="me-1" /> Upload
                                            </label>
                                        </div>
                                        <div className="d-flex justify-content-end fst-italic" style={{ fontSize: "10px", color: "#495057" }}>
                                            <a
                                                href="#"
                                                target="_blank"
                                                rel="noopener noreferrer"
                                                className="text-decoration-underline"
                                            >
                                                Template Lembar Pengesahan
                                            </a>
                                        </div>
                                    </Col>
                                </Row>
                            </Col>


                            {/* Kartu Konsultasi */}
                            <Col sm={12}>
                                <Row
                                    className="align-items-center gap-2"
                                    style={{ paddingBottom: "10px" }}
                                >
                                    <Col sm={12}>
                                        <Label
                                            htmlFor="kartu_konsultasi"
                                            className="form-label mb-0 fw-medium"
                                        >
                                            Kartu Konsultasi KP dari Instansi
                                        </Label>
                                    </Col>
                                    <Col sm={12}>
                                        <div className="d-flex gap-2">
                                            <Controller
                                                name="kartu_konsultasi"
                                                control={control}
                                                render={({ field: { onChange, value } }) => (
                                                    <div className="position-relative w-100">
                                                        <label
                                                            htmlFor="kartu_konsultasi"
                                                            className="d-flex align-items-center border border-1 rounded p-3 m-0"
                                                            style={{ width: "100%", height: "45px", cursor: "pointer" }}
                                                        >
                                                            <span style={{ color: "#909090" }}>
                                                                {value ? value.name : "Upload File"}
                                                            </span>
                                                        </label>

                                                        <Input
                                                            type="file"
                                                            id="kartu_konsultasi"
                                                            accept=".pdf"
                                                            onChange={(e) => {
                                                                const file = e.target.files?.[0];
                                                                onChange(file);
                                                            }}
                                                            hidden
                                                        />

                                                        {errors.kartu_konsultasi ? (
                                                            <FormErrorMessage errors={errors.lembar_penilaian} />
                                                        ) : (
                                                            <FormDescription message="File dengan format .pdf max 10MB" />
                                                        )}
                                                    </div>
                                                )}
                                            />

                                            <label
                                                htmlFor="lembar_penilaian"
                                                className="btn d-flex align-items-center btn-light mb-0"
                                                style={{ whiteSpace: "nowrap" }}
                                            >
                                                <FileUploadIcon className="me-1" /> Upload
                                            </label>

                                        </div>
                                        <div className="d-flex justify-content-end fst-italic" style={{ fontSize: "10px", color: "#495057" }}>
                                            <a
                                                href="#"
                                                target="_blank"
                                                rel="noopener noreferrer"
                                                className="text-decoration-underline"
                                            >
                                                Template Kartu Konsultasi KP
                                            </a>
                                        </div>
                                    </Col>
                                </Row>
                            </Col>
                        </Row>
                    </Col>
                    {/* right section */}
                    <Col md={12} lg={6}>
                        <Row className="gap-2">

                            {/* Lembar Penilaian */}
                            <Col sm={12}>
                                <Row
                                    className="align-items-center gap-2"
                                    style={{ paddingBottom: "10px" }}
                                >
                                    <Col sm={12}>
                                        <Label
                                            htmlFor="lembar_penilaian"
                                            className="form-label mb-0 fw-medium"
                                        >
                                            Lembar Penilaian
                                        </Label>
                                    </Col>
                                    <Col sm={12}>
                                        <div className="d-flex gap-2">
                                            <Controller
                                                name="lembar_penilaian"
                                                control={control}
                                                render={({ field: { onChange, value } }) => (
                                                    <div className="position-relative w-100">
                                                        <label
                                                            htmlFor="lembar_penilaian"
                                                            className="d-flex align-items-center border border-1 rounded p-3 m-0"
                                                            style={{ width: "100%", height: "45px", cursor: "pointer" }}
                                                        >
                                                            <span style={{ color: "#909090" }}>
                                                                {value ? value.name : "Upload File"}
                                                            </span>
                                                        </label>

                                                        <Input
                                                            type="file"
                                                            id="lembar_penilaian"
                                                            accept=".pdf"
                                                            onChange={(e) => {
                                                                const file = e.target.files?.[0];
                                                                onChange(file);
                                                            }}
                                                            hidden
                                                        />

                                                        {errors.lembar_penilaian ? (
                                                            <FormErrorMessage errors={errors.lembar_penilaian} />
                                                        ) : (
                                                            <FormDescription message="File dengan format .pdf max 10MB" />
                                                        )}
                                                    </div>
                                                )}
                                            />


                                            <label
                                                htmlFor="lembar_penilaian"
                                                className="btn d-flex align-items-center btn-light mb-0"
                                                style={{ whiteSpace: "nowrap" }}
                                            >
                                                <FileUploadIcon className="me-1" /> Upload
                                            </label>

                                        </div>
                                        <div className="d-flex justify-content-end fst-italic" style={{ fontSize: "10px", color: "#495057" }}>
                                            <a
                                                href="#"
                                                target="_blank"
                                                rel="noopener noreferrer"
                                                className="text-decoration-underline"
                                            >
                                                Template Lembar Penilaian
                                            </a>
                                        </div>
                                    </Col>
                                </Row>
                            </Col>






                        </Row>
                    </Col>
                </Row>
            </form>
        </>
    );
};
