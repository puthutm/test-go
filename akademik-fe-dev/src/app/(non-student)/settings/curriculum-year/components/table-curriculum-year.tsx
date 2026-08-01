"use client";

import { useDebouncedCallback } from "use-debounce";
import { Card, Col, Row } from "reactstrap";
import { usePathname, useRouter, useSearchParams } from "next/navigation";
import { useState } from "react";
import { signOut } from "next-auth/react";

import { SearchIcon } from "@/components/icons/search";
import { AddIcon } from "@/components/icons/add";
import DataTables from "@/components/ui/datatable";
import { RefreshIcon } from "@/components/icons/refresh";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { useColumnDefinitionCurriculumYears } from "./column-definition-curriculum-year";
import { useModalContext } from "@/lib/hooks/use-modal";
import ModalCurriculumYear from "./modal-curriculum-year";
import { ModalDeleteConfirmation } from "@/components/ui/modal-delete-confirmation";
import { deleteCurriculumYear } from "@/services/api/data-referensi/curriculum-year/delete-curriculum-year";

export const TableCurriculumYear = ({
  data,
}: {
  data: ApiResponse<PaginationDataReferensi<CurriculumYear[]>>;
}) => {
  const [loadingDelete, setLoadingDelete] = useState<boolean>(false);
  const pathname = usePathname();
  const router = useRouter();
  const searchParams = useSearchParams();
  const params = new URLSearchParams(searchParams);

  const { modalConfirmationState, setModalConfirmationState } =
    useModalConfirmationContext();

  const { setModalState } = useModalContext();

  const onDeleteCurriculumYear = async () => {
    setLoadingDelete(true);
    const response = await deleteCurriculumYear(
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

  const { columns } = useColumnDefinitionCurriculumYears();

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
      params.set("q", value);
    } else {
      params.delete("q");
    }

    params.set("page", "1");
    router.push(`${pathname}?${params.toString()}`);
  }, 1000);

  if (data.status === 401) {
    signOut();
  }
  return (
    <Row>
      <ModalDeleteConfirmation
        onDelete={onDeleteCurriculumYear}
        isLoading={loadingDelete}
      />
      <ModalCurriculumYear />
      <Col>
        <Card>
          <div className="gap-2 d-flex flex-column w-100 gap-3 p-3">
            <h1 className="fs-4 mb-0 fw-medium" style={{ color: "#3A3A3A" }}>
              Tahun Kurikulum
            </h1>
            <div className="d-flex align-items-center justify-content-between">
              <Row>
                <Col sm={10}>
                  <div className="form-icon">
                    <input
                      className={`form-control form-control-icon`}
                      id="no_kk"
                      placeholder="Cari Tahun Kurikulum"
                      onChange={(e) => handleSearch(e.target.value)}
                    />
                    <i>
                      <SearchIcon />
                    </i>
                  </div>
                </Col>
                <Col sm={2} className="ps-0">
                  <button
                    className="d-flex align-items-center gap-2 bg-transparent p-2 rounded-3"
                    color="transparent"
                    onClick={() => router.refresh()}
                    style={{ color: "#10487A", border: "1px solid #10487A" }}
                  >
                    <RefreshIcon color="#10487A" />
                  </button>
                </Col>
              </Row>
              <div className="d-flex gap-2">
                <button
                  className="btn-outline h-100 py-2 px-3 text-primary"
                  onClick={() => setModalState({ open: true, state: "add" })}
                >
                  <AddIcon color="#10487A" />
                  Tambah
                </button>
              </div>
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
