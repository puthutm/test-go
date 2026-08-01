"use client";

import { AutoCompleteInput } from "@/components/ui/auto-complete";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import {
  FormClassLecturerSchema,
  FormClassLecturerSchemaType,
} from "@/lib/validations/settings/academic-period/form-class-lecturer";
import { useSearchLecturer } from "@/services/api/sdm/biodata/use-get-search-lecture";
import { addClassLecturer } from "@/services/api/settings/academic-period/class-lecturer/add-class-lecturer";
import { updateClassLecturer } from "@/services/api/settings/academic-period/class-lecturer/update-class-lecturer";
import { zodResolver } from "@hookform/resolvers/zod";
import { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { Col, Label, Modal, ModalBody, ModalHeader, Row } from "reactstrap";
import { useDebouncedCallback } from "use-debounce";

export const ModalClassLecturer = ({
  data,
  classId,
}: {
  data: ApiResponse<ClassLecturer>;
  classId: string;
}) => {
  const [isModalOpen, setIsModalOpen] = useState<boolean>(false);
  const [search, setSearch] = useState("");

  const { setModalConfirmationState } = useModalConfirmationContext();
  const {
    formState: { errors, isSubmitting },
    handleSubmit,
    register,
    setValue,
    clearErrors,
    watch,
  } = useForm({
    resolver: zodResolver(FormClassLecturerSchema),
    defaultValues: {
      lecturer_id: "",
      subtitute_lecturer_id: null,
    },
  });

  const lecturer = watch("lecturer_id");

  const { data: dataLecturer, isLoading: isLoadingLecturer } =
    useSearchLecturer({ page: 1, search });

  const lecturerOptions = dataLecturer?.data?.data.map((lecture) => ({
    label: lecture.name_of_user,
    value: lecture.id,
  }));

  const onSubmit = async (payload: FormClassLecturerSchemaType) => {
    try {
      const response = data.data
        ? await updateClassLecturer(classId, data.data.id, payload)
        : await addClassLecturer(classId, payload);

      if (response.error) {
        return setModalConfirmationState((prev) => ({
          ...prev,
          open: true,
          state: "failed",
          message: response.message,
        }));
      }

      toggleModal();

      return setModalConfirmationState((prev) => ({
        ...prev,
        open: true,
        message: `Dosen pengajar berhasil ${data.data ? "diubah" : "ditambah"}`,
        state: "success",
      }));
    } catch (error: any) {
      setModalConfirmationState((prev) => ({
        ...prev,
        open: true,
        state: "failed",
        message: error.toString(),
      }));
    }
  };

  const toggleModal = () => {
    setIsModalOpen(!isModalOpen);
    setSearch("");
  };

  const handleSearch = useDebouncedCallback((value: string) => {
    if (value) {
      setSearch(value);
    } else {
      setSearch("");
    }
  }, 500);

  useEffect(() => {
    if (lecturer) {
      setSearch("");
    }
  }, [lecturer]);

  useEffect(() => {
    if (data.data) {
      setValue("lecturer_id", data?.data?.lecturer_id);
      setValue("subtitute_lecturer_id", data?.data?.subtitute_lecturer_id);
    }
  }, [data]);
  return (
    <>
      <button className="btn btn-primary" onClick={toggleModal}>
        {data.data ? "Ubah" : "Tambah"}
      </button>
      <Modal isOpen={isModalOpen} centered>
        <ModalHeader toggle={toggleModal}>
          {data.data ? "Ubah" : "Tambah"} Dosen Pengajar
        </ModalHeader>
        <ModalBody>
          <form
            onSubmit={handleSubmit(onSubmit)}
            className="d-flex flex-column gap-3"
            autoComplete="off"
          >
            <Row className="gap-2">
              {/* lecturer */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="lecturer_id"
                      className="form-label mb-0 fw-medium"
                    >
                      Dosen Pengajar
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <AutoCompleteInput
                      clearErrors={clearErrors}
                      data={lecturerOptions}
                      id="lecturer_id"
                      register={register}
                      setValue={setValue}
                      placeholder="Cari Dosen Pengajar"
                      onSearch={handleSearch}
                      errors={errors.lecturer_id}
                      isLoading={isLoadingLecturer}
                    />
                  </Col>
                </Row>
              </Col>
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="subtitute_lecturer_id"
                      className="form-label mb-0 fw-medium optional"
                    >
                      Dosen Pengganti
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <AutoCompleteInput
                      clearErrors={clearErrors}
                      data={lecturerOptions}
                      id="subtitute_lecturer_id"
                      register={register}
                      setValue={setValue}
                      placeholder="Cari Dosen Pengganti"
                      onSearch={handleSearch}
                      errors={errors.subtitute_lecturer_id}
                      isLoading={isLoadingLecturer}
                    />
                  </Col>
                </Row>
              </Col>
            </Row>
            <div className="d-flex justify-content-end mt-3 gap-2">
              <button
                className="bg-transparent text-primary rounded px-3 py-2"
                type="button"
                style={{ border: "1px solid #10487A" }}
                onClick={toggleModal}
                disabled={isSubmitting}
              >
                Batal
              </button>
              <button className="btn btn-primary" disabled={isSubmitting}>
                {data.data ? "Ubah" : "Tambah"}
              </button>
            </div>
          </form>
        </ModalBody>
      </Modal>
    </>
  );
};
