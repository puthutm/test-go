"use client";

import { useState } from "react";
import { Button, Col, Row } from "reactstrap";
import { debounce } from "lodash";
import { usePathname, useRouter, useSearchParams } from "next/navigation";
import { signOut } from "next-auth/react";

import DataTables from "@/components/ui/datatable";
import { useTableStudentColumns } from "./column-table-student";
import { SelectComponent } from "@/components/ui/select";
import { FilterListIcon } from "@/components/icons/filter-list";
import { SearchIcon } from "@/components/icons/search";
// import { ImportExportIcon } from "@/components/icons/import-export-icon";
import { useModalContext } from "@/lib/hooks/use-modal";
import { ModalStudentAccount } from "./modal-student-account";
import { AddIcon } from "@/components/icons/add";
import { FileUploadIcon } from "@/components/icons/file-upload";
import { ModalImportStudent } from "./modal-import-student";

export const TableStudent = ({
  data,
}: {
  data: ApiResponse<PaginationData<PortalStudent>> | any;
}) => {
  const pathname = usePathname();
  const router = useRouter();
  const searchParams = useSearchParams();
  const params = new URLSearchParams(searchParams);
  const { columns } = useTableStudentColumns();
  const { setModalState } = useModalContext();

  const [isOpenImport, setIsOpenImport] = useState(false);

  const handlePagination = (newPage: number) => {
    if (newPage) {
      params.set("page", (newPage + 1).toString());
    } else {
      params.delete("page");
    }

    router.push(`${pathname}?${params.toString()}`);
  };

  const handleSearch = debounce((value: string) => {
    if (value) {
      params.set("q", value);
    } else {
      params.delete("q");
    }

    params.set("page", "1");
    router.push(`${pathname}?${params.toString()}`);
  }, 1000);

  const handleAdd = () => {
    setModalState({
      open: true,
      state: "add",
    });
  };

  if (data?.status === 401) {
    signOut();
  }

  return (
    <>
      <div className="d-flex justify-content-between align-items-center">
        <Row className="d-flex">
          <Col sm={6}>
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
          </Col>
          <Col sm={6} className="px-0">
            <div className="form-icon">
              <input
                className={`form-control form-control-icon`}
                id="no_kk"
                placeholder="Keywoard"
                onChange={(e) => handleSearch(e.target.value)}
              />
              <i>
                <SearchIcon />
              </i>
            </div>
          </Col>
        </Row>
        <div className="d-flex gap-2">
          <Button
            className="btn btn-primary px-4"
            type="button"
            onClick={handleAdd}
            color="success"
          >
            <AddIcon color="#ffffff" /> Tambah
          </Button>
          <Button
            className="btn btn-primary px-4"
            type="button"
            onClick={() => setIsOpenImport(true)}
            color="info"
          >
            <FileUploadIcon color="#ffffff" /> Import
          </Button>
          {/* <button className="btn btn-outline-primary px-4">
            <ImportExportIcon />
            Export
          </button> */}
        </div>
      </div>
      <div className="table-responsive">
        <DataTables
          columns={columns}
          data={data?.data || []}
          pageCount={data?.data?.metadata.total_page || 1}
          pagination={data?.data?.metadata || 1}
          setPagination={handlePagination || null}
          total={data?.data?.metadata.total_data || 1}
        />
      </div>

      <ModalStudentAccount />
      <ModalImportStudent
        isOpen={isOpenImport}
        toggle={() => setIsOpenImport(!isOpenImport)}
      />
    </>
  );
};
