"use client";

import { useState } from "react";
import { Card, Col, Row } from "reactstrap";
import { usePathname, useRouter, useSearchParams } from "next/navigation";
import { signOut } from "next-auth/react";
import { useDebouncedCallback } from "use-debounce";

import { SearchIcon } from "@/components/icons/search";
import { AddIcon } from "@/components/icons/add";
import { useAcademicYearColumn } from "./column-academic-year";
import { deleteAcademicYear } from "@/services/api/data-referensi/academic-year/delete-academic-year";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { useModalContext } from "@/lib/hooks/use-modal";
import { ModalAcademicYear } from "./modal-academic-year";
import { ModalDeleteConfirmation } from "@/components/ui/modal-delete-confirmation";
import DataTables from "@/components/ui/datatable";

export const TableAcademicYear = ({
  data,
}: {
  data: ApiResponse<PaginationDataReferensi<AcademicYear>>;
}) => {
  const [loadingDelete, setLoadingDelete] = useState<boolean>(false);
  const pathname = usePathname();
  const router = useRouter();
  const searchParams = useSearchParams();
  const params = new URLSearchParams(searchParams);

  const { setModalState } = useModalContext();
  const { modalConfirmationState, setModalConfirmationState } =
    useModalConfirmationContext();

  const { columns } = useAcademicYearColumn();

  const handleToggleModal = () => {
    setModalState((prev) => ({
      ...prev,
      open: !prev.open,
      state: "add",
    }));
  };

  const onDeleteAcademicYear = async (id: string) => {
    setLoadingDelete(true);
    const response = await deleteAcademicYear(id);

    if (!response.error) {
      setModalConfirmationState((prev) => ({
        ...prev,
        id: undefined,
        open: !prev.open,
      }));

      setModalConfirmationState((prev) => ({
        ...prev,
        open: !prev.open,
        state: "success",
        message: "Data berhasil dihapus",
      }));
    }

    setLoadingDelete(false);
  };

  const handlePagination = (newPage: number) => {
    if (newPage) {
      params.set("page", (newPage + 1).toString());
    } else {
      params.delete("page");
    }

    router.push(`${pathname}?${params.toString()}`);
  };

  const handleSearch = useDebouncedCallback((value: string) => {
    if (value) {
      params.set("filter", value);
    } else {
      params.delete("filter");
    }

    params.set("page", "1");
    router.push(`${pathname}?${params.toString()}`);
  }, 1000);

  if (data.status === 401) {
    signOut();
  }

  return (
    <Row className="p-2">
      <Col>
        <ModalAcademicYear />
        <ModalDeleteConfirmation
          onDelete={async () =>
            await onDeleteAcademicYear(modalConfirmationState.id as string)
          }
          isLoading={loadingDelete}
        />

        <Card>
          <div className="gap-2 d-flex flex-column w-100 gap-3 p-3">
            <h1 className="fs-4 mb-0 fw-medium" style={{ color: "#3A3A3A" }}>
              Tahun Ajaran
            </h1>
            <div className="d-flex align-items-center justify-content-between">
              <Row>
                <Col sm={12}>
                  <div className="form-icon">
                    <input
                      className={`form-control form-control-icon`}
                      id="no_kk"
                      placeholder="Pencarian"
                      onChange={(e) => handleSearch(e.target.value)}
                    />
                    <i>
                      <SearchIcon />
                    </i>
                  </div>
                </Col>
              </Row>
              <button
                className="btn-outline h-100 py-2 px-3 text-primary"
                onClick={handleToggleModal}
              >
                <AddIcon color="#10487A" />
                Tambah
              </button>
            </div>
            <Col className="table-responsive" sm={12}>
              <DataTables
                columns={columns}
                data={data?.data}
                pageCount={data?.data?.metadata.total_pages}
                pagination={data?.data?.metadata}
                setPagination={handlePagination}
                total={data?.data?.metadata.total}
              />
            </Col>
          </div>
        </Card>
      </Col>
    </Row>
  );
};
