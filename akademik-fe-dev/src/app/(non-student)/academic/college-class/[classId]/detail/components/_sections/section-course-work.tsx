"use client";
import React from "react";
import { useParams } from "next/navigation";
// import third party component
import { Card, CardBody, CardHeader, Button } from "reactstrap";

// import component
import TabsSectionDetailCollegeClass from "../tab-detail-college-class";
import InformationCollegeClass from "../information-college-class";
import DataTables from "@/components/ui/datatable";
import { AddIcon } from "@/components/icons/add";
import ModalCourseAssigment from "./modal-course-assigment";

// import hook
import { useModalContext } from "@/lib/hooks/use-modal";
import useColumnCourseWork from "../_columns/column-definition-course-work";
import { useGetCourseAssignment } from "@/services/api/academic/lecturer/class-schedule/detail-class/course-assignment/use-get-course-assignment";

function SectionCourseWork() {
  //! columns
  const { columns } = useColumnCourseWork();
  const params = useParams();

  const {setModalState} = useModalContext()

  //! get data
  const { data: dataCourseAssignment, isLoading: isLoadingCourseAssignment } =
    useGetCourseAssignment(params.classId as string);
    const handleAdd = ()=>{
      setModalState({
        open:true,
        state:'add',
      })
    }

  return (
    <section className="position-relative">
      <ModalCourseAssigment/>
      {/*//! INFORMATIONAL */}
      <Card className="py-3 px-4 rounded-3 bg-white">
        {/*//! card header */}
        <CardHeader className="p-0 m-0 border-0">
          {/*//! TITLE */}
          <section className="d-flex align-items-center gap-1 ">
            <h2 className="m-0 p-0 fs-5 fw-semibold flex-grow-1 ">
              Tugas Kuliah
            </h2>

            <Button
              className="btn  d-flex align-items-center gap-2 "
              color="transparent"
              onClick={handleAdd}
              style={{
                color: "#10487A",
                padding: "6px 10px",
                border: "1px solid #10487A",
              }}
            >
              <AddIcon color="#10487A" />
              Tambah
            </Button>
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
            data={dataCourseAssignment}
            pageCount={0}
            pagination={null}
            setPagination={() => {}}
            isLoading={isLoadingCourseAssignment}
            total={0}
            isPaginate={false}
          />
        </div>
      </Card>
    </section>
  );
}

export default SectionCourseWork;
