"use client";

import { useState } from "react";

import { useModalContext } from "@/lib/hooks/use-modal";
import { useColumnClassScheduleTemplate } from "./column-class-schedule-template";
import { AddIcon } from "@/components/icons/add";
import { ModalClassScheduleTemplate } from "./modal-class-schedule-template";
import DataTables from "@/components/ui/datatable";
import { ModalDeleteConfirmation } from "@/components/ui/modal-delete-confirmation";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { deleteClassScheduleTemplate } from "@/services/api/curriculum/academic-period/class-schedule/delete-class-schedulte-template";

export const TableClassScheduleTemplate = ({
  data,
  classId,
  isDetail,
}: {
  data: ApiResponse<ClassScheduleTemplate>;
  classId: string;
  isDetail?: boolean;
}) => {
  const [loadingDelete, setLoadingDelete] = useState(false);

  const { setModalState } = useModalContext();

  const tableData = data.data ? [data.data] : null;

  const { modalConfirmationState, setModalConfirmationState } =
    useModalConfirmationContext();
  const { columns } = useColumnClassScheduleTemplate(isDetail);

  const onDeleteAcademicSubject = async () => {
    setLoadingDelete(true);
    const response = await deleteClassScheduleTemplate({
      classId: classId,
      classScheduleTemplateId: modalConfirmationState.id as string,
    });

    if (!response.error) {
      setModalConfirmationState((prev) => ({
        ...prev,
        id: undefined,
        open: !prev.open,
      }));
    }

    setModalConfirmationState((prev) => ({
      ...prev,
      open: !prev.open,
      state: "success",
      message: "Data berhasil dihapus",
    }));

    setLoadingDelete(false);
  };

  const handleToggleModal = () => {
    setModalState((prev) => ({
      ...prev,
      open: !prev.open,
      state: "add",
    }));
  };
  return (
    <div className="d-flex flex-column gap-2">
      <div
        className={`d-flex align-items-center justify-content-between ${
          isDetail ? "mb-2" : ""
        }`}
      >
        <p className="fw-medium fs-5" style={{ color: "#3A3A3A" }}>
          Jadwal Perkuliahan
        </p>
        <ModalDeleteConfirmation
          onDelete={onDeleteAcademicSubject}
          isLoading={loadingDelete}
        />
        <ModalClassScheduleTemplate />
        {!data?.data?.id && !isDetail ? (
          <>
            <button
              className="btn-outline h-100 py-2 px-3 text-primary"
              onClick={handleToggleModal}
            >
              <AddIcon color="#10487A" />
              Tambah
            </button>
          </>
        ) : null}
      </div>
      <div className="table-responsive">
        <DataTables
          columns={columns}
          data={{ data: tableData }}
          pageCount={0}
          pagination={1}
          setPagination={() => {}}
          total={0}
          isPaginate={false}
        />
      </div>
    </div>
  );
};
