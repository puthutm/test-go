"use client";

import React, { useState } from "react";
import { Col, Row } from "reactstrap";
import { notFound, useParams, usePathname, useRouter } from "next/navigation";

import { SearchIcon } from "@/components/icons/search";
import { RefreshIcon } from "@/components/icons/refresh";
import { AddIcon } from "@/components/icons/add";
import { FileDownloadIcon } from "@/components/icons/file-download";
import { useModalContext } from "@/lib/hooks/use-modal";
import { ModalAddClass } from "./modal-add-class";
import { useDebouncedCallback } from "use-debounce";
import { useAcademicPeriodClassesColumns } from "./columns";
import DataTables from "@/components/ui/datatable";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { deleteClass } from "@/services/api/settings/academic-period/class/delete-class";
import { ModalDeleteConfirmation } from "@/components/ui/modal-delete-confirmation";

export const ClassTable = ({
  searchParams,
  data,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
  data: ApiResponse<PaginationData<Class[]>>;
}) => {
  const [isLoadingDelete, setIsLoadingDelete] = useState(false);
  // state show modal confirmation
  const { modalConfirmationState, setModalConfirmationState } =
    useModalConfirmationContext();
  const pathname = usePathname();
  const router = useRouter();
  const searchParam = new URLSearchParams(searchParams as any);
  const { setModalState } = useModalContext();
  const params = useParams();
  const academicPeriodId = params.academicPeriodId as string;
  const { columns } = useAcademicPeriodClassesColumns(academicPeriodId);

  const handlePagination = (newPage: number) => {
    if (newPage) {
      searchParam.set("page", (newPage + 1).toString());
    } else {
      searchParam.delete("page");
    }

    router.replace(`${pathname}?${searchParam.toString()}`);
  };

  const handleSearch = useDebouncedCallback((value: string) => {
    if (value) {
      searchParam.set("q", value);
    } else {
      searchParam.delete("q");
    }

    searchParam.set("page", "1");
    router.replace(`${pathname}?${searchParam.toString()}`);
  }, 1000);

  const handleDelete = async () => {
    setIsLoadingDelete(true);
    try {
      const response = await deleteClass(modalConfirmationState?.id as string);
      if (response.error) {
        setModalConfirmationState((prev) => ({
          ...prev,
          open: true,
          id: null,
          state: "failed",
          message: response.message || "Gagal menghapus data",
        }));

        return;
      }

      setModalConfirmationState((prev) => ({
        ...prev,
        open: true,
        state: "success",
        message: "Data berhasil dihapus",
      }));
    } catch (err: any) {
      throw new Error(err.message);
    } finally {
      setIsLoadingDelete(false);
    }
  };

  if (data.status === 401) {
    return notFound();
  }

  return (
    <div className="d-flex flex-column gap-3 bg-white p-2 rounded">
      <ModalDeleteConfirmation
        isLoading={isLoadingDelete}
        onDelete={async () => {
          await handleDelete();
        }}
      />
      <ModalAddClass academicPeriodId={academicPeriodId} />
      <div className="gap-2 d-flex align-items-center justify-content-between w-100">
        <Row className="w-50">
          {/* <Col sm={4}>
            <div className="form-icon">
              <SelectComponent
                options={[]}
                id="filter"
                placeholder="Filter"
                hasIcon
              />
              <i>
                <FilterListIcon />
              </i>
            </div>
          </Col> */}
          <Col sm={5}>
            <div className="form-icon">
              <input
                className={`form-control form-control-icon`}
                id="no_kk"
                placeholder="Cari Kelas"
                onChange={(e) => handleSearch(e.target.value)}
              />
              <i>
                <SearchIcon />
              </i>
            </div>
          </Col>
          <Col sm={2} className="ps-0">
            <button
              className="btn-outline text-primary"
              style={{ padding: "10px" }}
              onClick={() => router.refresh()}
            >
              <RefreshIcon />
            </button>
          </Col>
        </Row>

        <div className="d-flex gap-2">
          <button
            className="btn-outline text-primary px-3"
            onClick={() => {
              setModalState((prev) => ({
                ...prev,
                open: true,
                state: "add",
                id: academicPeriodId as string,
              }));
            }}
          >
            <AddIcon color="#10487A" />
            Tambah
          </button>
          <button className="btn btn-primary px-3">
            <FileDownloadIcon
              height="16"
              width="16"
              color="white"
              className="me-1"
            />
            Export
          </button>
        </div>
      </div>
      <DataTables
        columns={columns}
        data={data?.data}
        pageCount={data?.data?.metadata?.total_page}
        pagination={data?.data?.metadata}
        setPagination={handlePagination}
        total={data?.data?.metadata?.total_data}
      />
    </div>
  );
};
