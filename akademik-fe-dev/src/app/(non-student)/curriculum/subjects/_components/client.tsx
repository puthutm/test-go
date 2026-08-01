"use client";

import { useGetAllSubject } from "@/services/api/curriculum/subjects/get-all";
import React, { useEffect, useState } from "react";
import { useSubjectColumns } from "./column";
import { useDeleteSubject } from "@/services/api/curriculum/subjects/delete";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { Card, CardBody, CardHeader, Col, Row } from "reactstrap";
import { ModalDeleteConfirmation } from "@/components/ui/modal-delete-confirmation";
import Link from "next/link";
import DataTables from "@/components/ui/datatable";

const SubjectClientPage = () => {
  const [queryParams, setQueryParams] = useState<QueryParam>({
    page: 1,
    search: null,
    limit: null,
    sort_by: null,
    sort_direction: null,
  });

  const { data: subjectData, isLoading } = useGetAllSubject(queryParams);
  const { columns } = useSubjectColumns();
  console.log(subjectData);
  const {
    mutateAsync: handleDeleteSubject,
    isPending: isDeleteSubjectLoading,
    isSuccess: isDeleteSubjectSuccess,
  } = useDeleteSubject();

  const { modalConfirmationState, setModalConfirmationState } =
    useModalConfirmationContext();

  useEffect(() => {
    if (isDeleteSubjectSuccess) {
      setModalConfirmationState((prev) => ({
        ...prev,
        open: false,
      }));
      setModalConfirmationState((prev) => ({
        ...prev,
        open: !prev.open,
        message: "Berhasil menghapus mata kuliah",
        state: "success",
      }));
    }

    if (Number(subjectData?.data.metadata.total_page) < queryParams.page) {
      setQueryParams((prev) => ({
        ...prev,
        page: Number(subjectData?.data.metadata.total_page),
      }));
    }
  }, [isDeleteSubjectSuccess, subjectData?.data.metadata.total_page]);

  return (
    <Row>
      <Col>
        <ModalDeleteConfirmation
          isLoading={isDeleteSubjectLoading}
          onDelete={async () => {
            await handleDeleteSubject(String(modalConfirmationState.id));
          }}
        />

        <Card className="p-0">
          <CardHeader>
            <div className="gap-2 d-flex align-items-center justify-content-between w-100">
              <div></div>

              <div className="d-flex gap-2">
                {/* Trash */}
                <Link
                  href={"/curriculum/subjects/trash"}
                  className="btn btn-danger"
                >
                  <i className="las la-trash-restore align-center me-1"></i>{" "}
                  Trash
                </Link>
                {/* Add Button */}
                <Link
                  href={"/curriculum/subjects/add"}
                  className="btn btn-success"
                >
                  <i className="ri-add-line align-bottom me-1" /> Mata Kuliah
                </Link>
              </div>
            </div>
          </CardHeader>

          <CardBody>
            <Col>
              <DataTables
                columns={columns}
                data={subjectData?.data}
                pageCount={subjectData?.data.metadata.total_page as number}
                pagination={subjectData?.data.metadata}
                setPagination={setQueryParams}
                total={subjectData?.data.metadata.total_data as number}
                isLoading={isLoading}
              />
            </Col>
          </CardBody>
        </Card>
      </Col>
    </Row>
  );
};

export default SubjectClientPage;
