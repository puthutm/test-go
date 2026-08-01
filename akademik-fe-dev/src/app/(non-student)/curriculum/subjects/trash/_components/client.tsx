"use client";

import { useGetAllSubjectTrash } from "@/services/api/curriculum/subjects/trash";
import { useSubjectTrashColumn } from "./column";
import { useDebouncedCallback } from "use-debounce";
import { useEffect, useState } from "react";
import { Card, CardBody, CardHeader, Col, Input, Row } from "reactstrap";
import DataTables from "@/components/ui/datatable";

const SubjectTrashClientPage = () => {
  const [queryParams, setQueryParams] = useState<QueryParam>({
    page: 1,
    search: null,
    limit: null,
    sort_by: null,
    sort_direction: null,
  });

  const {
    data: trashSubjectData,
    isLoading: isTrashSubjectLoading,
    isSuccess: isTrashSubjectSuccess,
  } = useGetAllSubjectTrash(queryParams);

  const { columns } = useSubjectTrashColumn();

  const handleSearch = useDebouncedCallback((value) => {
    setQueryParams((prev) => ({ ...prev, search: value, page: 1 }));
  }, 1000);

  const setPage = (value: number) => {
    setQueryParams((prev) => ({ ...prev, page: value + 1 }));
  };

  useEffect(() => {
    if (
      (trashSubjectData?.data?.metadata?.total_page as number) <
      queryParams.page
    ) {
      setQueryParams((prev) => ({
        ...prev,
        page: trashSubjectData?.data?.metadata?.total_page as number,
      }));
    }
  }, [isTrashSubjectSuccess, trashSubjectData?.data?.metadata?.total_page]);

  return (
    <Row>
      <Col>
        <Card className="p-0">
          <CardHeader>
            <Row>
              <Col className="d-flex gap-2 ms-auto" sm={12} md={4} lg={2}>
                <div className="form-icon w-100 ms-auto">
                  <Input
                    type="text"
                    className="form-control form-control-icon"
                    id="iconInput"
                    placeholder="Search"
                    onChange={(e) => handleSearch(e.target.value)}
                  />
                  <i className="ri-search-line"></i>
                </div>
              </Col>
            </Row>

            <CardBody>
              <Col className="table-responsive" sm={12}>
                <DataTables
                  columns={columns}
                  data={trashSubjectData?.data}
                  pageCount={
                    trashSubjectData?.data?.metadata.total_page as number
                  }
                  pagination={queryParams}
                  setPagination={setPage}
                  isLoading={isTrashSubjectLoading}
                  total={trashSubjectData?.data.metadata.total_data as number}
                />
              </Col>
            </CardBody>
          </CardHeader>
        </Card>
      </Col>
    </Row>
  );
};

export default SubjectTrashClientPage;
