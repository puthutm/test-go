"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useEffect, useRef, useState } from "react";
import { useForm } from "react-hook-form";
import {
  Button,
  Col,
  Label,
  Modal,
  ModalBody,
  ModalHeader,
  Row,
} from "reactstrap";
import { useDebouncedCallback } from "use-debounce";

import { useModalContext } from "@/lib/hooks/use-modal";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import {
  FormClassParticipantSchema,
  FormClassParticipantSchemaType,
} from "@/lib/validations/settings/academic-period/form-class-participant";
import { addClassParticipant } from "@/services/api/settings/academic-period/class-participant/add-class-participant";
import { useGetSearchStudentByStudyProgram } from "@/services/api/students/biodata/study-program/use-get-search-student-by-study-program";
import { AutoCompleteInput } from "@/components/ui/auto-complete";

export const ModalAddClassParticipant = ({
  studyProgramId,
  classId,
}: {
  studyProgramId: string;
  classId: string;
}) => {
  const inputRef = useRef<HTMLInputElement | null>(null);
  const [selectedStudent, setSelectedStudent] = useState<any>("");
  const [query, setQuery] = useState({
    search: "",
    studyProgramId,
  });
  const { setModalConfirmationState } = useModalConfirmationContext();
  const { modalState, setModalState } = useModalContext();

  const {
    formState: { errors, isSubmitting },
    handleSubmit,
    register,
    setValue,
    clearErrors,
  } = useForm({
    resolver: zodResolver(FormClassParticipantSchema),
    defaultValues: {
      student_id: "",
    },
  });

  const { data: dataStudent, isLoading: isLoadingStudent } =
    useGetSearchStudentByStudyProgram({ ...query });

  const studentOptions = dataStudent?.data?.data?.map((student) => ({
    label: `${student.student_nim} - ${student.student_name}`,
    value: student.student_id,
  }));

  const handleSearchStudent = useDebouncedCallback((value: string) => {
    if (value) {
      setQuery((prev) => ({
        ...prev,
        search: value,
      }));
      setSelectedStudent(value);
    } else {
      setSelectedStudent(value);
      setQuery((prev) => ({
        ...prev,
        search: "",
      }));
    }
  }, 500);

  const toggleModal = () => {
    setModalState((prev) => ({
      ...prev,
      open: !modalState.open,
    }));
    setSelectedStudent("");
    setQuery((prev) => ({
      ...prev,
      search: "",
    }));
  };

  const onSubmit = async (data: FormClassParticipantSchemaType) => {
    try {
      const response = await addClassParticipant(classId, data);

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
        message: `Peserta kelas berhasil ditambah`,
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

  useEffect(() => {
    if (selectedStudent && inputRef?.current?.value) {
      setValue("student_id", selectedStudent?.student_id);
      clearErrors("student_id");
    }
  }, [selectedStudent, inputRef]);

  return (
    <Modal isOpen={modalState.open && modalState.state === "add"} centered>
      <ModalHeader toggle={toggleModal}>
        <p>Tambah Peserta Kelas</p>
      </ModalHeader>
      <ModalBody>
        <form
          onSubmit={handleSubmit(onSubmit)}
          className="d-flex flex-column gap-3"
          autoComplete="off"
        >
          <Row className="gap-2">
            {/* participant */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="participant"
                    className="form-label mb-0 fw-medium"
                  >
                    Mahasiswa
                  </Label>
                </Col>
                <Col sm={12}>
                  <AutoCompleteInput
                    clearErrors={clearErrors}
                    data={studentOptions}
                    id="student_id"
                    register={register}
                    setValue={setValue}
                    placeholder="Cari mahasiswa"
                    onSearch={handleSearchStudent}
                    errors={errors.student_id}
                    isLoading={isLoadingStudent}
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
            >
              Batal
            </button>
            <Button color="primary" disabled={isSubmitting}>
              Tambah
            </Button>
          </div>
        </form>
      </ModalBody>
    </Modal>
  );
};
