"use client";
import { useState } from "react";
// import third part component
import {
  Row,
  Col,
  Card,
  CardHeader,
  CardBody,
  Button,
  Input,
} from "reactstrap";

// import component
import { FileDownloadIcon } from "@/components/icons/file-download";
import { SelectComponent } from "../../../../components/ui/select";
import { ReplayIcon } from "@/components/icons/replay";
import { SearchIcon } from "@/components/icons/search";
import DataTables from "@/components/ui/datatable";

// import hooks
import useColumnLecturerCollegeCurriculum from "./components/column-definition-college-curriculum";

export interface IDummyValueTable {
  id: string;
  is_active: boolean;
  year_curriculum: string;
  code: string;
  curriculum_name: string;
  sks: string;
  type_subjects: string;
  teaching_programs: string;
}

export interface IModalManipulationFilterCollegeCurriculum {
  status: boolean;
  title: "Filter";
  data?: string;
}

const dummyValueTable: PaginationData<IDummyValueTable> = {
  metadata: {
    page: 1,
    size: 1,
    total_data: 1,
    total_page: 1,
  },
  data: [
    {
      id: "1",
      is_active: true,
      year_curriculum: "2022",
      code: "IF3110",
      curriculum_name: "Sistem Informasi",
      sks: "4",
      type_subjects: "kuliah",
      teaching_programs: "PJJ management",
    },
    {
      id: "2",
      is_active: false,
      year_curriculum: "2021",
      code: "IF3110",
      curriculum_name: "Information Technology",
      sks: "4",
      type_subjects: "kuliah",
      teaching_programs: "PJJ management",
    },
  ],
};

function CollegeCurriculumPage() {
  const [queryParams] = useState<QueryParam>({
    page: 1,
  });

  //! state show Modal FILTER
  // const [showModalFilter, setShowModalFilter] =
  //   useState<IModalManipulationFilterCollegeCurriculum>({
  //     status: false,
  //     title: "Filter",
  //   });

  //! column
  const { columns } = useColumnLecturerCollegeCurriculum();
  // dumy data
  const dummyData = [
    {
      value: "All",
      label: "Semua",
    },
  ];

  return (
    <Row>
      <Col>
        <Card className="p-3 rounded-3 bg-white">
          {/*//! card header */}
          <CardHeader className="p-0 m-0">
            {/*//! TITLE */}
            <section className="d-flex align-items-center">
              {/* //! text */}
              <h2 className="m-0 p-0 fs-5 fw-semibold flex-grow-1">
                Kurikulum Kuliah
              </h2>
            </section>

            {/*//! action */}
            <section className="position-relative gap-2 d-flex align-items-center flex-column flex-md-row w-100 my-3">
              {/*//! left */}
              <div className="position-relative flex-grow-1 d-flex gap-2 flex-wrap">
                {/*//! select component program studi*/}
                <SelectComponent
                  options={dummyData}
                  placeholder="Program studi"
                  // isDisabled={!isEdit || isLoadingRegistrantBiodataSystem}
                  // isError={!!errors.gender}
                  id={"selectFilterProgramStudi"}
                  isClearable
                  onChange={(value) => {
                    console.log(value);
                  }}
                />
                {/*//! select component tahun ajara */}
                <SelectComponent
                  options={dummyData}
                  placeholder="Tahun ajaran"
                  // isDisabled={!isEdit || isLoadingRegistrantBiodataSystem}
                  // isError={!!errors.gender}
                  id={"selectFilterTahunAjaran"}
                  isClearable
                  onChange={(value) => {
                    console.log(value);
                  }}
                />

                {/*//! input search */}
                <div className="form-icon">
                  <Input
                    autoComplete="off"
                    type="text"
                    name="inputSearchSubject"
                    className={`form-control form-control-icon py-2 `}
                    id="inputNamaLengkap"
                    placeholder="Cari Mata Kuliah"
                  />

                  <i className="">
                    <SearchIcon />
                  </i>
                </div>

                {/*//! button search */}
                <Button
                  className="btn  d-flex align-items-center gap-2"
                  color="transparent"
                  style={{ color: "#10487A", border: "1px solid #10487A" }}
                >
                  Cari
                </Button>

                {/*//! button clear */}
                <Button
                  className="btn  d-flex align-items-center gap-2"
                  color="transparent"
                  style={{ color: "#10487A", border: "1px solid #10487A" }}
                >
                  <ReplayIcon color="#10487A" />
                </Button>
              </div>

              {/*//! right */}
              <div className="position-relative flex-shrink-0 gap-2 d-flex">
                {/*//! button downlod */}
                <Button
                  className="btn btn-primary d-flex align-items-center gap-2"
                  color="#10487A"
                >
                  <FileDownloadIcon color="#fff" />
                  Download
                </Button>
              </div>
            </section>
          </CardHeader>

          {/*//! card body */}
          <CardBody className="p-0 m-0">
            <section className="table-responsive">
              <DataTables
                columns={columns}
                data={dummyValueTable}
                pageCount={0}
                pagination={queryParams}
                setPagination={() => {}}
                isLoading={false}
                total={1}
              />
            </section>
          </CardBody>
        </Card>
      </Col>
    </Row>
  );
}

export default CollegeCurriculumPage;
