"use client";

import { Col, Row } from "reactstrap";
import { debounce } from "lodash";
import { usePathname, useRouter, useSearchParams } from "next/navigation";
import { signOut } from "next-auth/react";

import DataTables from "@/components/ui/datatable";
import { useTableLecturersColumns } from "./column-table-lecturer";
import { SelectComponent } from "@/components/ui/select";
import { FilterListIcon } from "@/components/icons/filter-list";
import { SearchIcon } from "@/components/icons/search";
import { ImportExportIcon } from "@/components/icons/import-export-icon";

export const TableLecturers = ({
  data,
}: {
  data: ApiResponse<PaginationData<PortalLecturer>> | any;
}) => {
  const pathname = usePathname();
  const router = useRouter();
  const searchParams = useSearchParams();
  const params = new URLSearchParams(searchParams);
  const { columns } = useTableLecturersColumns();

  const handlePagination = (newPage: number) => {
    if (newPage) {
      params.set("page", newPage.toString());
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

  if (data?.status === 401) {
    signOut();
  }
  return (
    <>
      <div className="d-flex justify-content-between">
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
        <button className="btn-outline text-primary px-4">
          <ImportExportIcon />
          Export
        </button>
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
    </>
  );
};
