import { CalendarTodayIcon } from "@/components/icons/calendar-today";
import { DatePicker } from "@/components/ui/date-picker";
import { FormDescription } from "@/components/ui/form-description";
import { FormErrorMessage } from "@/components/ui/form-error-message";
import { Controller, useForm } from "react-hook-form";
import { Button, Col, Input, Label, Row } from "reactstrap";

export const FormInfoKp = ({
  isEdit,
  setIsEdit,
}: {
  hasProposal: boolean;
  isEdit: boolean;
  setIsEdit: any;
}) => {
  const {
    control,
    handleSubmit,
    formState: { errors },
    reset,
  } = useForm();
  return (
    <form
      className="mt-3 mb-2"
      onSubmit={handleSubmit((data) => console.log(data))}
    >
      <Row className="gap-1 gap-lg-0">
        {/* left section */}
        <Col md={12} lg={6}>
          <Row className="gap-2">
            {/* instansi tujuan */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="instansi"
                    className="form-label mb-0 fw-medium"
                  >
                    Instansi Tujuan
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="instansi"
                    control={control}
                    render={({ field }) => (
                      <Input
                        className={`form-control form-control-icon ${
                          errors.instansi ? "border border-danger" : ""
                        }`}
                        id="instansi"
                        placeholder="Masukkan Instansi Tujuan"
                        readOnly={!isEdit}
                        {...field}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.instansi} />
                </Col>
              </Row>
            </Col>
            {/* judul */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="instansi"
                    className="form-label mb-0 fw-medium"
                  >
                    Judul Kegiatan Kerja Praktik
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="instansi"
                    control={control}
                    render={({ field }) => (
                      <Input
                        className={`form-control form-control-icon ${
                          errors.instansi ? "border border-danger" : ""
                        }`}
                        id="instansi"
                        placeholder="Masukkan Judul Kegiatan Kerja Praktik"
                        readOnly={!isEdit}
                        {...field}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.instansi} />
                </Col>
              </Row>
            </Col>
            {/* surat pengantar */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="instansi"
                    className="form-label mb-0 fw-medium optional"
                  >
                    Surat Pengantar Kerja Praktik
                  </Label>
                </Col>
                <Col sm={12}>
                  <div className="d-flex gap-2">
                    <Controller
                      name="fileBpjsKesehatan"
                      control={control}
                      render={({ field: { onChange, value, ...field } }) => (
                        <div className="position-relative w-100">
                          <Input
                            type="text"
                            className={`form-control ${
                              errors.fileBpjsKesehatan
                                ? "border border-danger"
                                : ""
                            }`}
                            placeholder="Pilih file"
                            readOnly
                            value={value ? value.name : ""}
                            disabled={!isEdit}
                            {...field}
                          />
                          {isEdit && (
                            <Input
                              type="file"
                              id="fileBpjsKesehatan"
                              className={`form-control form-control-icon ${
                                errors.fileBpjsKesehatan
                                  ? "border border-danger"
                                  : ""
                              }`}
                              onChange={(e) => {
                                const file = e.target.files?.[0];
                                if (file) onChange(file);
                              }}
                              hidden
                              accept=".pdf"
                            />
                          )}
                          <FormDescription message="File dalam bentuk .pdf max 10mb." />
                          <FormErrorMessage errors={errors.fileBpjsKesehatan} />
                        </div>
                      )}
                    />
                    {isEdit && (
                      <div className="position-relative">
                        <label
                          htmlFor="fileBpjsKesehatan"
                          className="btn btn-light d-flex align-items-center mb-0"
                          style={{ whiteSpace: "nowrap" }}
                        >
                          Upload File
                        </label>
                        <a
                          href="#"
                          className="position-absoulte text-primary d-block"
                          style={{
                            fontSize: "10px",
                            fontStyle: "italic",
                            textAlign: "right",
                          }}
                        >
                          Template surat 2
                        </a>
                      </div>
                    )}
                    {!isEdit && (
                      <div>
                        <button
                          className="btn d-flex align-items-center btn-light"
                          style={{ whiteSpace: "nowrap" }}
                        >
                          Lihat File
                        </button>
                        <a
                          href="#"
                          className="position-absoulte text-primary"
                          style={{
                            fontSize: "10px",
                            fontStyle: "italic",
                            textAlign: "right",
                          }}
                        >
                          Template surat
                        </a>
                      </div>
                    )}
                  </div>
                </Col>
              </Row>
            </Col>
          </Row>
        </Col>
        {/* rigth section */}
        <Col md={12} lg={6}>
          <Row className="gap-2">
            {/* surat pengantar */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="instansi"
                    className="form-label mb-0 fw-medium"
                  >
                    Surat Keterangan Instansi
                  </Label>
                </Col>
                <Col sm={12}>
                  <div className="d-flex gap-2">
                    <Controller
                      name="instansi"
                      control={control}
                      render={({ field: { onChange, value, ...field } }) => (
                        <div className="position-relative w-100">
                          <Input
                            type="text"
                            className={`form-control ${
                              errors.instansi ? "border border-danger" : ""
                            }`}
                            placeholder="Pilih file"
                            readOnly
                            value={value ? value.name : ""}
                            disabled={!isEdit}
                            {...field}
                          />
                          {isEdit && (
                            <Input
                              type="file"
                              id="instansi"
                              className={`form-control form-control-icon ${
                                errors.instansi ? "border border-danger" : ""
                              }`}
                              onChange={(e) => {
                                const file = e.target.files?.[0];
                                if (file) onChange(file);
                              }}
                              hidden
                              accept=".pdf"
                            />
                          )}
                          <FormDescription message="File dalam bentuk .pdf max 10mb." />
                          <FormErrorMessage errors={errors.instansi} />
                        </div>
                      )}
                    />
                    {isEdit && (
                      <label
                        htmlFor="instansi"
                        className="btn btn-light d-flex align-items-center mb-0"
                        style={{ whiteSpace: "nowrap" }}
                      >
                        Upload File
                      </label>
                    )}
                    {!isEdit && (
                      <button
                        className="btn d-flex align-items-center btn-light"
                        style={{ whiteSpace: "nowrap" }}
                      >
                        Lihat File
                      </button>
                    )}
                  </div>
                </Col>
              </Row>
            </Col>
            {/* tanggal pelaksanaan */}
            {/* Tanggal Lahir */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="tanggal_pelaksanaan"
                    className="form-label mb-0 fw-medium"
                  >
                    Tanggal Pelaksanaan
                  </Label>
                </Col>
                <Col sm={12}>
                  <div className="form-icon">
                    <Controller
                      name="birth_date"
                      control={control}
                      render={({ field }) => {
                        return (
                          <DatePicker
                            onChange={(e) => field.onChange(e)}
                            className={`p-0 ${
                              errors.birth_date ? "border border-danger" : ""
                            }`}
                            classNameFlatpickr={`form-control form-control-icon disabled-input`}
                            options={{
                              mode: "single",
                              dateFormat: "d F Y",
                            }}
                            disabled={!isEdit}
                          />
                        );
                      }}
                    />
                    <i style={{ left: "15px" }}>
                      <CalendarTodayIcon color="#878A99" />
                    </i>
                  </div>
                  <FormErrorMessage errors={errors.birth_date} />
                </Col>
              </Row>
            </Col>
          </Row>
        </Col>
      </Row>
      {isEdit && (
        <div className="d-flex justify-content-between mt-3 gap-3">
          <button
            onClick={() => {
              setIsEdit(!isEdit);
              reset();
            }}
            className="bg-transparent text-primary rounded px-3"
            type="button"
            style={{ border: "1px solid #10487A" }}
          >
            <span>Batal</span>
          </button>
          <Button
            color="primary"
            className="d-flex flex-grow-0 justify-content-center align-items-center"
          >
            <span>Update</span>
          </Button>
        </div>
      )}
    </form>
  );
};
