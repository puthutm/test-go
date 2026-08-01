"use client";

import { usePathname, useRouter, useSearchParams } from "next/navigation";
import { useDebouncedCallback } from "use-debounce";
import { useState } from "react";

import { AddIcon } from "@/components/icons/add";
import { SearchIcon } from "@/components/icons/search";
import DataTables from "@/components/ui/datatable";
import { useModalContext } from "@/lib/hooks/use-modal";
import { Col, Row } from "reactstrap";
import { useClassParticipantColumnTable } from "./column-table-class-participant";
import { ModalAddClassParticipant } from "./modal-add-class-participant";
import { deleteClassParticipant } from "@/services/api/settings/academic-period/class-participant/delete-class-participant";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { ModalDeleteConfirmation } from "@/components/ui/modal-delete-confirmation";
import { ModalDetailClassParticipant } from "./modal-detail-class-participant";

export const TableClassParticipant = ({
  studyProgramId,
  classId,
  participant,
  isDetail,
}: {
  studyProgramId: string;
  classId: string;
  participant: ApiResponse<PaginationData<ClassParticipant>>;
  isDetail?: boolean;
}) => {
  const [loadingDelete, setLoadingDelete] = useState<boolean>(false);
  const pathname = usePathname();
  const router = useRouter();
  const searchParams = useSearchParams();
  const params = new URLSearchParams(searchParams);
  const { setModalState } = useModalContext();
  const { modalConfirmationState, setModalConfirmationState } =
    useModalConfirmationContext();

  const { columns } = useClassParticipantColumnTable(isDetail);

  const onDeleteAcademicSubject = async () => {
    setLoadingDelete(true);
    const response = await deleteClassParticipant(
      classId,
      modalConfirmationState?.id as string
    );

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

  const handleSearch = useDebouncedCallback((value: string) => {
    if (value) {
      params.set("q", value);
    } else {
      params.delete("q");
    }

    params.set("page", "1");
    router.push(`${pathname}?${params.toString()}`);
  }, 1000);

  const handlePagination = (newPage: number) => {
    if (newPage) {
      params.set("page", (newPage + 1).toString());
    } else {
      params.delete("page");
    }

    router.push(`${pathname}?${params.toString()}`);
  };
  return (
    <div className="d-flex flex-column gap-3">
      <ModalAddClassParticipant
        studyProgramId={studyProgramId}
        classId={classId}
      />
      <ModalDeleteConfirmation
        onDelete={onDeleteAcademicSubject}
        isLoading={loadingDelete}
      />
      <ModalDetailClassParticipant classId={classId} />
      <div className="d-flex justify-content-between">
        <Row>
          <Col>
            <div className="form-icon">
              <input
                className={`form-control form-control-icon`}
                placeholder="Cari peserta kelas"
                onChange={(e) => handleSearch(e.target.value)}
              />
              <i>
                <SearchIcon />
              </i>
            </div>
          </Col>
        </Row>
        {!isDetail ? (
          <button
            className="btn-outline rounded-xl rounded-3 text-primary px-3"
            onClick={() =>
              setModalState((prev) => ({
                ...prev,
                open: true,
                state: "add",
              }))
            }
          >
            <AddIcon color="#10487A" />
            Tambah
          </button>
        ) : null}
      </div>
      <DataTables
        data={participant.data}
        columns={columns}
        pageCount={participant?.data?.metadata.total_page}
        pagination={participant?.data?.metadata}
        setPagination={handlePagination}
        total={participant?.data?.metadata.total_data}
      />
    </div>
  );
};
