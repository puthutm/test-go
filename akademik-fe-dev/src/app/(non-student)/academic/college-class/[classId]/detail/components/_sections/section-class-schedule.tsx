"use client";
import React, { useState } from "react";
import { useQueryClient } from "@tanstack/react-query";
import { useParams } from "next/navigation";
// import third party component
import { Card, CardBody, CardHeader } from "reactstrap";

// import component
import TabsSectionDetailCollegeClass from "../tab-detail-college-class";
import InformationCollegeClass from "../information-college-class";
import DataTables from "@/components/ui/datatable";
import { ModalDeleteConfirmation } from "@/components/ui/modal-delete-confirmation";
import { ModalSuccessConfirmation } from "@/components/ui/modal-success-confirmation";

import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import useColumnClassSchedule from "../_columns/column-definition-class-schedule";

import { useGetClassScheduleSubDetail } from "@/services/api/academic/lecturer/class-schedule/detail-class/class-schedule/use-get-class-schedule-sub-detail";
import { deleteClassScheduleSubDetail } from "@/services/api/academic/lecturer/class-schedule/detail-class/class-schedule/delete-class-schedule-sub-detail";

function SectionClassSchedule() {
  const queryclient = useQueryClient();
  const params = useParams();
  const [isLoadingDelete, setIsLoadingDelete] = useState(false);

  // state show modal confirmation
  const { modalConfirmationState, setModalConfirmationState } =
    useModalConfirmationContext();

  const { columns } = useColumnClassSchedule();

  //! get data
  const {
    data: dataCourseClassScheduleSubDetail,
    isLoading: isLoadingCourseClassScheuuleSUbDetail,
  } = useGetClassScheduleSubDetail(params.classId as string);

  const handleDelete = async (classId: string, class_schedule_id: string) => {
    setIsLoadingDelete(true);
    try {
      const response = await deleteClassScheduleSubDetail(
        classId,
        class_schedule_id
      );
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
      queryclient.invalidateQueries({
        queryKey: ["get-class-schedule-sub-detail"],
      });
    } catch (err: any) {
      throw new Error(err.message);
    } finally {
      setIsLoadingDelete(false);
    }
  };

  return (
    <section className="position-relative">
      <ModalSuccessConfirmation />

      {/* modal delete confirm */}
      <ModalDeleteConfirmation
        isLoading={isLoadingDelete}
        onDelete={async () => {
          await handleDelete(
            params.classId as string,
            modalConfirmationState.id as string
          );
        }}
      />

      {/*//! INFORMATIONAL */}
      <Card className="py-3 px-4 rounded-3 bg-white">
        {/*//! card header */}
        <CardHeader className="p-0 m-0 border-0">
          {/*//! TITLE */}
          <section className="d-flex align-items-center gap-1 ">
            <h2 className="m-0 p-0 fs-5 fw-semibold flex-grow-1 mt-1">
              Jadwal Perkuliahan
            </h2>
          </section>
        </CardHeader>

        {/*//! card body */}
        <CardBody className="position-relative px-0">
          {/*//!tabs */}
          <section
            className="position-relative mt-2"
            style={{ borderBottom: "2px solid #ddd" }}
          >
            <TabsSectionDetailCollegeClass />
          </section>

          {/*//! informational college class */}
          <section className=" mt-3">
            <InformationCollegeClass />
          </section>
        </CardBody>
      </Card>

      {/*//! TABLES */}
      <Card className="p-4 rounded-3 bg-white">
        <div className="table-responsive mt-3">
          <DataTables
            columns={columns}
            data={dataCourseClassScheduleSubDetail?.data}
            pageCount={0}
            pagination={null}
            setPagination={() => {}}
            isLoading={isLoadingCourseClassScheuuleSUbDetail}
            total={0}
            isPaginate={false}
          />
        </div>
      </Card>
    </section>
  );
}

export default SectionClassSchedule;
