"use client";

import React, { useState } from "react";
import { signOut } from "next-auth/react";

import { useSearchParams, usePathname, useRouter } from "next/navigation";
import { useGradeCompositionColumns } from "./columns";
import { deleteGradeComposition } from "@/services/api/settings/grade-composition/delete-grade-composition";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { useModalContext } from "@/lib/hooks/use-modal";
import { useDebouncedCallback } from "use-debounce";

import { SearchIcon } from "@/components/icons/search";
import { AddIcon } from "@/components/icons/add";
import { DeleteIcon } from "@/components/icons/delete";
import { ContentCopyIcon } from "@/components/icons/content-copy";
import { Button, Card, CardBody, CardHeader, Col, Row,Input,Alert} from "reactstrap";
import ModalGradeComposition from "./ModalGradeComposition";
import { ModalDuplicateGradeComposition } from "./ModalDuplicateGradeComposition";
import { ModalDeleteConfirmation } from "@/components/ui/modal-delete-confirmation";
import { WarningIcon } from "@/components/icons/warning";


import Link from "next/link";
import DataTables from "@/components/ui/datatable";

const GradeCompositionClientPage = ({
  dataGradeComposition
}:{
  dataGradeComposition: ApiResponse<PaginationData<IGradeComposition>>;
}) => {
  const router = useRouter();
  const pathname = usePathname();
  const searchParams = useSearchParams();
  const params = new URLSearchParams(searchParams);

  const [isLoadingDelete, setIsLoadingDelete] = useState<boolean>(false);
  const { columns } = useGradeCompositionColumns();


  const { setModalState } = useModalContext();
  const { modalConfirmationState, setModalConfirmationState } =
    useModalConfirmationContext();


    //! event handle page pagination
  const handlePagination = (newPage: number) => {
    if (newPage) {
      params.set("page", (newPage + 1).toString());
    } else {
      params.delete("page");
    }

    router.push(`${pathname}?${params.toString()}`);
  };
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

  //! handle delete
  const handleDelete = async (idGradeComposition: string) => {
    setIsLoadingDelete(true);
    try {
      const response = await deleteGradeComposition(idGradeComposition);
      if (response.error) {
        setModalConfirmationState((prev) => ({
          ...prev,
          id: null,
          open: true,
          state: "failed",
          message: response.message || "Gagal menghapus data",
        }));
        return
      }

      setModalConfirmationState((prev) => ({
        ...prev,
          open: true,
          id: null,
        state: "success",
        message: "Data berhasil dihapus",
      }));
    } catch (err: any) {
      throw new Error(err.message);
    } finally {
      setIsLoadingDelete(false);
    }
  };
  if (dataGradeComposition.status === 401) {
    signOut();
  }
  return (
    <Row>
      <Col>
        <ModalGradeComposition />
        <ModalDuplicateGradeComposition />
        <ModalDeleteConfirmation
          isLoading={isLoadingDelete}
          onDelete={async () => {
            await handleDelete(
             modalConfirmationState.id as string
            );
          }}
        />

        <Card className="p-3 rounded-3 bg-white">
          <CardHeader className="p-0 m-0">
            {/*//! TITLE */}
            <section className="d-flex align-items-center">
              {/* //! text */}
              <h2 className="m-0 p-0 fs-5 fw-semibold flex-grow-1">
                Komposisi Nilai
              </h2>
            </section>

            {/*//! action */}
            <section className="position-relative gap-2 d-flex align-items-center flex-column flex-md-row w-100 my-2">
              {/*//! left */}
              <div className="position-relative flex-grow-1 d-flex gap-2 flex-wrap">
                                <div className="form-icon">
                  <Input
                    type="text"
                    className="form-control form-control-icon"
                    id="iconInput"
                    placeholder="Cari"
                    onChange={(event) =>
                      handleSearch((event.target as HTMLInputElement).value)
                    }
                  />
                  <i><SearchIcon /></i>
                </div>
                {/*//! select component */}
                {/* <div className="form-icon">
                  <SelectComponent
                    options={mapingStudyProgramData as OptionType[]}
                    placeholder="Program Studi"
                    isLoading={isLoadingProgramStudy}
                    // isDisabled={!isEdit || isLoadingRegistrantBiodataSystem}
                    // isError={!!errors.gender}
                    id={"selectFilterProgramStudi"}
                    hasIcon
                    isClearable
                    onChange={(value) => {
                      if (value != null) {
                        params.set("study_program_id", value.value);
                        params.set("page", '1');
                        router.push(`${pathname}?${params.toString()}`);
                        return;
                      }

                      router.push(`${pathname}`);
                    }}
                  />
                  <i>
                    <FilterListIcon color="#909090" />
                  </i>
                </div> */}
              </div>

              {/*//! right */}
              <div className="position-relative flex-shrink-0 gap-2 d-flex">
                {/*//! trash */}
                {/* Trash */}
                <Link
                  href={"/settings/grade-composition/trash"}
                  className="btn btn-danger"
                >
                 <DeleteIcon color="#fff"/>
                  Trash
                </Link>
                {/*//! button duplikasi */}
                <Button
                  className="btn  d-flex align-items-center gap-2"
                  style={{
                    color: "#10487A",
                    border: "1px solid #10487A",
                  }}
                  color="transparent"
                  onClick={() => {
                    setModalState((prev) => ({
                      ...prev,
                      open: true,
                      state: "duplicate",
                      id: null,
                    }));
                  }}
                >
                  <ContentCopyIcon color="#10487A" />
                  Duplikasi
                </Button>
                {/*//! button tambah */}
                <Button
                  className="btn  d-flex align-items-center gap-2"
                  style={{
                    color: "#10487A",
                    border: "1px solid #10487A",
                  }}
                  color="transparent"
                  onClick={() => {
                    setModalState((prev) => ({
                      ...prev,
                      open: true,
                      state: "add",
                      id:null
                    }));
                  }}
                >
                  <AddIcon color="#10487A" />
                  Tambah
                </Button>
              </div>
            </section>
          </CardHeader>


          <CardBody className="p-0 m-0 mt-2">
          {/*//! alert warning */}
          <Alert color="warning" className='w-100 py-2 px-3 d-flex gap-2 align-items-center m-0 mb-2' fade={false}>
            <WarningIcon/>
                    {/*line */}
                    <div style={{
                        position:'absolute',
                        width:'2px',
                        left:'0',
                        top:'0',
                        bottom:'0',
                        background:'#F7B84B'
                    }}/>
                    <p className="m-0 p-0 flex-grow-1">
                    Total Persentase sudah 100 %
                    </p>
          </Alert>
                
            {/*//! table */}
            <Col className="table-responsive" sm={12}>
              <DataTables
                columns={columns}
                data={dataGradeComposition?.data}
                pageCount={dataGradeComposition?.data?.metadata.total_page}
                pagination={dataGradeComposition?.data?.metadata}
                setPagination={handlePagination}
                total={dataGradeComposition?.data?.metadata.total_data}
              />
            </Col>
          </CardBody>
        </Card>
      </Col>
    </Row>
  );
};

export default GradeCompositionClientPage;
