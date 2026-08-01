"use client";

import React from "react";
import { useGradeScaleTrashColumns } from "./columns";
import { useSearchParams, usePathname, useRouter } from "next/navigation";
import { useDebouncedCallback } from "use-debounce";
import { Card, CardBody, CardHeader, Col, Input, Row } from "reactstrap";
import DataTables from "@/components/ui/datatable";
import { SearchIcon } from "@/components/icons/search";
import Link from "next/link";

const GradeScaleTrashPageClient = ({
  dataTrashGradeScale
}:{
  dataTrashGradeScale: ApiResponse<PaginationData<IGradeScale>>;
}) => {
  const router = useRouter();
  const pathname = usePathname();
  const searchParams = useSearchParams();
  const params = new URLSearchParams(searchParams);
  const { columns } = useGradeScaleTrashColumns();

  //! handle search
  const handleSearch = useDebouncedCallback((search: string) => {
    if (search.trim().length > 0) {
      params.set("search", search);
    } else {
      params.delete("search");
    }
    params.set("page", "1"); // Reset to first page on search
    router.push(`${pathname}?${params.toString()}`);
  }, 1000);


    //! event handle page pagination
  const handlePagination = (newPage: number) => {
    if (newPage) {
      params.set("page", (newPage + 1).toString());
    } else {
      params.delete("page");
    }

    router.push(`${pathname}?${params.toString()}`);
  };
  return (
    <Row>
      <Col>
        <Card className="p-0">
          <CardHeader>
            <Row className="row-gap-3">
              <Col sm={12} md={8} lg={9}>
                <div className="w-auto d-flex align-items-center gap-2">
                <Link
                href={`/settings/grade-scale`}
                  className="btn w-auto  d-flex align-items-center gap-2"
                  style={{
                    color: "#10487A",
                    border: "1px solid #10487A",
                  }}
                  color="transparent"
                >
                  Kembali
                </Link>
                </div>
              </Col>
              <Col className="d-flex gap-2 ms-auto" sm={12} md={4} lg={3}>
                <div className="form-icon w-100 ms-auto">
                  <Input
                    type="text"
                    className="form-control form-control-icon"
                    id="iconInput"
                    placeholder="Search"
                    onChange={(e) => handleSearch(e.target.value)}
                  />
                  <i >
                    <SearchIcon  />
                  </i>
                </div>
              </Col>
            </Row>
          </CardHeader>

          <CardBody>
            <Col className="table-responsive" sm={12}>
              <DataTables
                columns={columns}
                data={dataTrashGradeScale?.data}
                pageCount={dataTrashGradeScale?.data?.metadata.total_page}
                pagination={dataTrashGradeScale?.data?.metadata}
                setPagination={handlePagination}
                total={dataTrashGradeScale?.data?.metadata.total_data}
              />
            </Col>
          </CardBody>
        </Card>
      </Col>
    </Row>
  );
};




export default GradeScaleTrashPageClient;
