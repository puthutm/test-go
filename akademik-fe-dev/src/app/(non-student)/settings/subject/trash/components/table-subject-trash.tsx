"use client";

import { useDebouncedCallback } from "use-debounce";
import { Card, Col, Row } from "reactstrap";
import { usePathname, useRouter, useSearchParams } from "next/navigation";
import { signOut } from "next-auth/react";

import { useColumnDefinitionSubjectTrash } from "./column-definition-subject-trash";
import DataTables from "@/components/ui/datatable";
import { SearchIcon } from "@/components/icons/search";

export const TableSubjectTrash = ({
  data,
}: {
  data: ApiResponse<PaginationData<Subject[]>>;
}) => {
  const pathname = usePathname();
  const router = useRouter();
  const searchParams = useSearchParams();
  const params = new URLSearchParams(searchParams);

  const { columns } = useColumnDefinitionSubjectTrash();

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
      <Col>
        <Card>
          <div className="gap-2 d-flex flex-column w-100 gap-3 p-3">
            <h1 className="fs-4 mb-0 fw-medium" style={{ color: "#3A3A3A" }}>
              Trash Mata Kuliah
            </h1>
            <div className="d-flex align-items-center justify-content-between">
              <Row>
                <Col sm={10}>
                  <div className="form-icon">
                    <input
                      className={`form-control form-control-icon`}
                      id="no_kk"
                      placeholder="Cari Mata Kuliah"
                      onChange={(e) => handleSearch(e.target.value)}
                    />
                    <i>
                      <SearchIcon />
                    </i>
                  </div>
                </Col>
              </Row>
            </div>
            <Col className="table-responsive" sm={12}>
              <DataTables
                columns={columns}
                data={data?.data}
                pageCount={data?.data?.metadata.total_page}
                pagination={data?.data?.metadata}
                setPagination={handlePagination}
                total={data?.data?.metadata.total_data}
              />
            </Col>
          </div>
        </Card>
      </Col>
    </Row>
  );
};
