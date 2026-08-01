"use client";
import React from "react";
import {
  Card,
  CardBody,
  CardHeader,
  Col,
  Row,
  Button,
  Input,
  Label,
} from "reactstrap";

import DataTables from "@/components/ui/datatable";
import { SearchIcon } from "@/components/icons/search";
import { RefreshIcon } from "@/components/icons/refresh";
import { FileDownloadIcon } from "@/components/icons/file-download";
import { SelectComponent } from "@/components/ui/select";
import { useColumnDefThesisProposal } from "./column-def-thesis-proposal";

function PageClientThesisProposal({
  data,
  role,
}: {
  data: ApiResponse<PaginationData<FinalProjectProposal[]>>;
  role: string;
}) {
  console.log({ role });

  // dumy data
  const dummyData = [
    {
      value: "All",
      label: "Semua",
    },
  ];

  const { columns } = useColumnDefThesisProposal();

  return (
    <Row>
      <Col>
        <Card className="p-3 rounded-3 bg-white">
          {/*//! card header */}
          <CardHeader className="p-0 m-0">
            {/*//! TITLE */}
            <section className="d-flex align-items-center">
              {/* //! text */}
              <h2 className="m-0 p-0 fs-4 fw-medium flex-grow-1">
                Proposal Tugas Akhir
              </h2>
            </section>

            {/*//!filter */}
            <Row className="mt-3 row-gap-3">
              {/*//! JENIS TA */}
              <Col md={6}>
                <Row>
                  <Col sm={12} md={5}>
                    <Label
                      className="m-0 p-0 fs-5 "
                      style={{ fontWeight: "400" }}
                      for="jenisTA"
                    >
                      Jenis TA
                    </Label>
                  </Col>
                  <Col sm={12} md={7}>
                    <SelectComponent
                      options={dummyData}
                      placeholder="jenis TA"
                      // isDisabled={!isEdit || isLoadingRegistrantBiodataSystem}
                      // isError={!!errors.gender}
                      id={"jenisTA"}
                      isClearable
                      onChange={(value) => {
                        console.log(value);
                      }}
                    />
                  </Col>
                </Row>
              </Col>
              {/*//! Status */}
              <Col md={6}>
                <Row>
                  <Col sm={12} md={5}>
                    <Label
                      className="m-0 p-0 fs-5 "
                      style={{ fontWeight: "400" }}
                      for="statusfilter"
                    >
                      Status
                    </Label>
                  </Col>
                  <Col sm={12} md={7}>
                    <SelectComponent
                      options={dummyData}
                      placeholder="status"
                      isDisabled
                      // isDisabled={!isEdit || isLoadingRegistrantBiodataSystem}
                      // isError={!!errors.gender}
                      id={"statusfilter"}
                      isClearable
                      onChange={(value) => {
                        console.log(value);
                      }}
                    />
                  </Col>
                </Row>
              </Col>
              {/*//! Program Studi */}
              <Col md={6}>
                <Row>
                  <Col sm={12} md={5}>
                    <Label
                      className="m-0 p-0 fs-5 "
                      style={{ fontWeight: "400" }}
                      for="prodifilter"
                    >
                      Program Studi
                    </Label>
                  </Col>
                  <Col sm={12} md={7}>
                    <SelectComponent
                      options={dummyData}
                      placeholder="Program Studi"
                      isDisabled
                      // isDisabled={!isEdit || isLoadingRegistrantBiodataSystem}
                      // isError={!!errors.gender}
                      id={"prodifilter"}
                      isClearable
                      onChange={(value) => {
                        console.log(value);
                      }}
                    />
                  </Col>
                </Row>
              </Col>
              {/*//! Angkatan */}
              <Col md={6}>
                <Row>
                  <Col sm={12} md={5}>
                    <Label
                      className="m-0 p-0 fs-5 "
                      style={{ fontWeight: "400" }}
                      for="angkatanfilter"
                    >
                      Angkatan
                    </Label>
                  </Col>
                  <Col sm={12} md={7}>
                    <SelectComponent
                      options={dummyData}
                      placeholder="Angkatan"
                      isDisabled
                      // isDisabled={!isEdit || isLoadingRegistrantBiodataSystem}
                      // isError={!!errors.gender}
                      id={"angkatanfilter"}
                      isClearable
                      onChange={(value) => {
                        console.log(value);
                      }}
                    />
                  </Col>
                </Row>
              </Col>
            </Row>

            {/*//! action */}
            <section className="position-relative gap-2 d-flex align-items-center flex-column flex-md-row w-100 mt-4 mb-2">
              {/*//! left */}
              <div className="position-relative flex-grow-1 d-flex gap-2 flex-wrap">
                {/*//! button filter */}
                {/* <Button
                  disabled
                  className='border d-flex align-items-center gap-2'
                  color='transparent'
                  style={{color:'#909090'}}
                  onClick={()=>{
                    setShowModalFilter(()=>({
                      status:true,
                      title:'Filter'
                    }))
                  }}
                  >
                    <FilterListIcon className='fs-4' color='#909090' />
                    Filter
                  </Button> */}

                {/*//! select component */}
                {/* <SelectComponent
                            options={dummyData}
                            placeholder="Semua"
                            isDisabled
                            isDisabled={!isEdit || isLoadingRegistrantBiodataSystem}
                            isError={!!errors.gender}
                            id={'selectFilter'}
                            isClearable
                            onChange={(value)=>{
                              console.log(value)
                            }}
                  /> */}

                {/*//! input search */}
                <div className="form-icon">
                  <Input
                    autoComplete="off"
                    type="text"
                    disabled
                    name="inputSearchSubject"
                    className={`form-control form-control-icon py-2 `}
                    id="inputNama"
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
                  disabled
                  style={{ color: "#10487A", border: "1px solid #10487A" }}
                >
                  Cari
                </Button>

                {/*//! button clear */}
                <Button
                  className="btn  d-flex align-items-center gap-2"
                  color="transparent"
                  disabled
                  style={{ color: "#10487A", border: "1px solid #10487A" }}
                >
                  <RefreshIcon color="#10487A" />
                </Button>
              </div>

              {/*//! right */}
              <div className="position-relative flex-shrink-0 gap-2 d-flex">
                {/*//! button add */}
                {/* <Button className='btn  d-flex align-items-center gap-2 '
                    color='transparent'
                    disabled
                    style={{color:'#10487A',border:'1px solid #10487A'}}
                    >
                      <AddIcon color='#10487A'/>
                      Tambah
                  </Button> */}
                {/*//! button downlod */}
                <Button
                  className="btn btn-primary d-flex align-items-center gap-2"
                  color="#10487A"
                  disabled
                >
                  <FileDownloadIcon color="#fff" />
                  Download
                </Button>
              </div>
            </section>
          </CardHeader>

          {/* //! card body */}
          <CardBody className="p-0 m-0">
            {/* //! tables */}
            <section className="table-responsive">
              <DataTables
                columns={columns}
                data={data?.data}
                pageCount={data?.data?.metadata?.total_page}
                pagination={data?.data?.metadata}
                setPagination={() => {}}
                isLoading={false}
                total={data?.data?.metadata?.total_data}
              />
            </section>
          </CardBody>
        </Card>
      </Col>
    </Row>
  );
}

export default PageClientThesisProposal;
