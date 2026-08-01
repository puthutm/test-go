"use client";
import { useState, useRef } from "react";
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
import { FileDownloadIcon } from "@/components/icons/file-download";
import { FilterListIcon } from "@/components/icons/filter-list";
// import { SelectComponent } from "@/components/ui/select";
import { ReplayIcon } from "@/components/icons/replay";
import { SearchIcon } from "@/components/icons/search";
import DataTables from "@/components/ui/datatable";
import ModalFilterCollegeClass from "./modal-filter-college-class";

import useColumnCollegeClass from "./column-definittion-college-class";
import { signOut } from "next-auth/react";
import { useSearchParams } from "next/navigation";
import { useRouter } from "next/navigation";
import { usePathname } from "next/navigation";

export interface IModalManipulationFilterCollegeClass {
  status: boolean;
  title: "Filter";
  data?: string;
}

interface IPropsParamsPageSubject {
  dataClassSchedule: ApiResponse<PaginationData<ClassSchedule>>;
}
function PageCollegeClass({ dataClassSchedule }: IPropsParamsPageSubject) {
  const router = useRouter();
  const pathname = usePathname();
  const searchParams = useSearchParams();
  const params = new URLSearchParams(searchParams);
  const inputSearch = useRef<HTMLInputElement | null>(null);

  //! state show Modal FILTER
  const [showModalFilter, setShowModalFilter] =
    useState<IModalManipulationFilterCollegeClass>({
      status: false,
      title: "Filter",
    });


  // column
  const { columns } = useColumnCollegeClass();

  //! event handle page pagination
  const handlePagination = (newPage: number) => {
    if (newPage) {
      params.set("page", (newPage + 1).toString());
    } else {
      params.delete("page");
    }

    router.push(`${pathname}?${params.toString()}`);
  };

  //! event handle search
  const handleSearch = () => {
    if (inputSearch.current) {
      if (inputSearch.current.value.trim().length !== 0) {
        params.set("search", inputSearch.current.value);
      } else {
        params.delete("search");
      }

      params.set("page", "1");
      router.push(`${pathname}?${params.toString()}`);
    }
  };

  //! event handle clear
  const handleClear = () => {
    if (inputSearch.current) {
      inputSearch.current.value = "";
    }
    router.push(`${pathname}`);
  };
  if (dataClassSchedule.status === 401) {
    signOut();
  }
  return (
    <>
      <ModalFilterCollegeClass
        showModal={showModalFilter}
        setShowModal={setShowModalFilter}
      />
      <Row>
        <Col>
          <Card className="p-3 rounded-3 bg-white">
            {/*//! card header */}
            <CardHeader className="p-0 m-0">
              {/*//! TITLE */}
              <section className="d-flex align-items-center">
                {/* //! text */}
                <h2 className="m-0 p-0 fs-5 fw-semibold flex-grow-1">
                  Kelas Kuliah
                </h2>
              </section>

              {/*//! action */}
              <section className="position-relative gap-2 d-flex align-items-center flex-column flex-md-row w-100 my-4">
                {/*//! left */}
                <div className="position-relative flex-grow-1 d-flex gap-2 flex-wrap">
                  {/*//! button filter */}
                  <Button
                    className="border d-flex align-items-center gap-2"
                    color="transparent"
                    style={{ color: "#909090" }}
                    onClick={() => {
                      setShowModalFilter(() => ({
                        status: true,
                        title: "Filter",
                      }));
                    }}
                  >
                    <FilterListIcon className="fs-4" color="#909090" />
                    Filter
                  </Button>

                  {/*//! select component */}
                  {/* <SelectComponent
                    options={dummyData}
                    placeholder="Semua"
                    // isDisabled={!isEdit || isLoadingRegistrantBiodataSystem}
                    // isError={!!errors.gender}
                    id={"selectFilter"}
                    isClearable
                    onChange={(value) => {
                      console.log(value);
                    }}
                  /> */}


                  {/*//! input search */}
                  <div className="form-icon">
                    <Input
                      innerRef={inputSearch}
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
                    onClick={() => {
                      handleSearch();
                    }}
                  >
                    Cari
                  </Button>

                  {/*//! button clear */}
                  <Button
                    className="btn  d-flex align-items-center gap-2"
                    color="transparent"
                    onClick={() => {
                      handleClear();
                    }}
                    disabled={
                      inputSearch.current == null ||
                      inputSearch.current?.value.trim().length === 0
                    }
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
                  data={dataClassSchedule?.data}
                  pageCount={
                    dataClassSchedule?.data.metadata?.total_page as number
                  }
                  pagination={dataClassSchedule.data.metadata}
                  setPagination={handlePagination}
                  isLoading={false}
                  total={dataClassSchedule?.data.metadata?.total_data as number}
                />
              </section>
            </CardBody>
          </Card>
        </Col>
      </Row>
    </>
  );
}

export default PageCollegeClass;
