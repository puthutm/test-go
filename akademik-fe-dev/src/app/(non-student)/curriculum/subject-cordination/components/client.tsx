"use client";
import { useRef } from "react";
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
import { SelectComponent } from "@/components/ui/select";
import { ReplayIcon } from "@/components/icons/replay";
import { SearchIcon } from "@/components/icons/search";
import DataTables from "@/components/ui/datatable";

import useColumnSubjectCordination from "./column-definition-subject-cordination";
import { signOut } from "next-auth/react";
import { useSearchParams } from "next/navigation";
import { useRouter } from "next/navigation";
import { usePathname } from "next/navigation";

interface IPropsParamsPageSubjectCordination {
  dataSubjectCordination: ApiResponse<
    PaginationData<ILectureSubjectsCordinator>
  >;
}

function PageSubjectCordination({
  dataSubjectCordination,
}: IPropsParamsPageSubjectCordination) {
  const router = useRouter();
  const pathname = usePathname();
  const searchParams = useSearchParams();
  const params = new URLSearchParams(searchParams);
  const inputSearch = useRef<HTMLInputElement | null>(null);

  // dumy data
  const dummyData = [
    {
      value: "All",
      label: "Semua",
    },
  ];

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

  //! event handle page pagination
  const handlePagination = (newPage: number) => {
    if (newPage) {
      params.set("page", (newPage + 1).toString());
    } else {
      params.delete("page");
    }

    router.push(`${pathname}?${params.toString()}`);
  };

  //! event handle clear
  const handleClear = () => {
    if (inputSearch.current) {
      inputSearch.current.value = "";
    }
    router.push(`${pathname}`);
  };

  if (dataSubjectCordination.status === 401) {
    signOut();
  }

  const { columns } = useColumnSubjectCordination();
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
                Kordinator Mata Kuliah
              </h2>
            </section>

            {/*//! action */}
            <section className="position-relative gap-2 d-flex align-items-center flex-column flex-md-row w-100 my-3">
              {/*//! left */}
              <div className="position-relative flex-grow-1 d-flex gap-2 flex-wrap">
                {/*//! select semester*/}
                <SelectComponent
                  options={dummyData}
                  placeholder="Semester"
                  id={"selectFilterSemester"}
                  isClearable
                  onChange={(value) => {
                    console.log(value);
                  }}
                />

                {/*//! input search */}
                <div className="form-icon">
                  <Input
                    innerRef={inputSearch}
                    autoComplete="off"
                    type="text"
                    name="inputSearchKordinator"
                    className={`form-control form-control-icon py-2 `}
                    id="inputSearchKordinator"
                    placeholder="Cari Kordinator Mata Kuliah"
                  />

                  <i className="">
                    <SearchIcon />
                  </i>
                </div>

                {/*//! button search */}
                <Button
                  className="btn  d-flex align-items-center gap-2"
                  color="transparent"
                  onClick={handleSearch}
                  style={{ color: "#10487A", border: "1px solid #10487A" }}
                >
                  Cari
                </Button>

                {/*//! button clear */}
                <Button
                  className="btn  d-flex align-items-center gap-2"
                  color="transparent"
                  disabled={
                    inputSearch.current == null ||
                    inputSearch.current?.value.trim().length === 0
                  }
                  onClick={handleClear}
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
                data={dataSubjectCordination?.data}
                pageCount={
                  dataSubjectCordination.data?.metadata?.total_page as number
                }
                pagination={dataSubjectCordination.data?.metadata}
                setPagination={handlePagination}
                isLoading={false}
                total={
                  dataSubjectCordination.data.metadata?.total_data as number
                }
              />
            </section>
          </CardBody>
        </Card>
      </Col>
    </Row>
  );
}

export default PageSubjectCordination;
