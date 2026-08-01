"use client";

import { Col, Row } from "reactstrap";
import { useDebouncedCallback } from "use-debounce";
import { usePathname, useRouter, useSearchParams } from "next/navigation";
import { signOut } from "next-auth/react";

import DataTables from "@/components/ui/datatable";
import { SelectComponent } from "@/components/ui/select";
import { FilterListIcon } from "@/components/icons/filter-list";
import { SearchIcon } from "@/components/icons/search";
import { useTableAcademicPeriodColumns } from "./column-table-academic-period";
import { RefreshIcon } from "@/components/icons/refresh";

export const TableAcademicPeriod = ({
  data,
}: {
  data: ApiResponse<PaginationDataReferensi<AcademicPeriod[]>>;
}) => {
  const pathname = usePathname();
  const router = useRouter();
  const searchParams = useSearchParams();
  const params = new URLSearchParams(searchParams);
  const { columns } = useTableAcademicPeriodColumns();


  const handlePagination = (newPage: number) => {
    if (newPage) {
      params.set("page", (newPage + 1).toString());
    } else {
      params.delete("page");
    }

    router.replace(`${pathname}?${params.toString()}`);
  };

  const handleSearch = useDebouncedCallback((value: string) => {
    if (value) {
      params.set("q", value);
    } else {
      params.delete("q");
    }

    params.set("page", "1");
    router.replace(`${pathname}?${params.toString()}`);
  }, 1000);

  if (data.status === 401) {
    signOut();
  }
  return (
    <>
      <div className="row">
        <Col sm={4}>
          <Row>
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
            <Col sm={5} className="px-0">
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
            <Col sm={1}>
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
        </Col>
      </div>
      <div className="table-responsive">
        <DataTables
          columns={columns}
          data={data?.data}
          pageCount={data?.data?.metadata.total_pages}
          pagination={data?.data?.metadata}
          setPagination={handlePagination}
          total={data?.data?.metadata.total}
        />
      </div>
    </>
  );
};
