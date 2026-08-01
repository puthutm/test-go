"use client";
import { useState } from "react";
import { useSearchParams, usePathname, useRouter } from "next/navigation";

import DataTables from "@/components/ui/datatable";
import { Button, Card, CardBody, CardHeader, Col, Row } from "reactstrap";
import ModalGradeScale from "./ModalGradeScale";
import { SelectComponent } from "@/components/ui/select";
import { AddIcon } from "@/components/icons/add";
import { FilterListIcon } from "@/components/icons/filter-list";
import { DeleteIcon } from "@/components/icons/delete";
import Link from "next/link";

import { useModalContext } from "@/lib/hooks/use-modal";
import { useGradeScaleColumns } from "./datatable/columns";
import { deleteGradeScale } from "@/services/api/settings/grade-scale/delete-grade-scale";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { useGetUnsiaStudyProgram } from "@/services/api/data-referensi/study-program/use-get-unsia-study-program";
import { signOut } from "next-auth/react";

import { ModalDeleteConfirmation } from "@/components/ui/modal-delete-confirmation";

const GradeScaleClientPage = ({
  dataGradeScale,
}: {
  dataGradeScale: ApiResponse<PaginationData<IGradeScale>>;
}) => {
  const router = useRouter();
  const pathname = usePathname();
  const searchParams = useSearchParams();
  const params = new URLSearchParams(searchParams);

  const [isLoadingDelete, setIsLoadingDelete] = useState<boolean>(false);

  // dumy data
  const { data: programStudy, isLoading: isLoadingProgramStudy } =
    useGetUnsiaStudyProgram();

  //! mpaing option study program
  const mapingStudyProgramData =
    programStudy?.data?.map((el: UnsiaStudyProgram) => {
      return {
        label: el.name,
        value: el.id,
      };
    }) ?? [];

  const { columns } = useGradeScaleColumns();

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

  const handleToggleModal = (state: "add" | "edit" | "detail", id: string | null) => {
    setModalState((prev) => ({
      ...prev,
      open: !prev.open,
      state,
      id,
    }));
  };

  const handleDelete = async (idGradeScale: string) => {
    setIsLoadingDelete(true);
    try {
      const response = await deleteGradeScale(idGradeScale);
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

  if (dataGradeScale.status === 401) {
    signOut();
  }

  return (
    <Row>
      <Col>
        <ModalGradeScale />
        <ModalDeleteConfirmation
          isLoading={isLoadingDelete}
          onDelete={async () => {
            await handleDelete(modalConfirmationState.id as string);
          }}
        />

        <Card className="p-3 rounded-3 bg-white">
          <CardHeader className="p-0 m-0">
            {/*//! TITLE */}
            <section className="d-flex align-items-center">
              {/* //! text */}
              <h2 className="m-0 p-0 fs-5 fw-semibold flex-grow-1">
                Skala Nilai
              </h2>
            </section>

            {/*//! action */}
            <section className="position-relative gap-2 d-flex align-items-center flex-column flex-md-row w-100 my-2">
              {/*//! left */}
              <div className="position-relative flex-grow-1 d-flex gap-2 flex-wrap">
                {/*//! select component program studi */}
                <div className="form-icon">
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
                </div>
              </div>

              {/*//! right */}
              <div className="position-relative flex-shrink-0 gap-2 d-flex">
                {/*//! trash */}
                {/* Trash */}
                <Link
                  href={"/settings/grade-scale/trash"}
                  className="btn btn-danger"
                >
                 <DeleteIcon color="#fff"/>
                  Trash
                </Link>
                {/*//! button tambah */}
                <Button
                  className="btn  d-flex align-items-center gap-2"
                  style={{
                    color: "#10487A",
                    border: "1px solid #10487A",
                  }}
                  color="transparent"
                  onClick={() => handleToggleModal("add",null)}
                >
                  <AddIcon color="#10487A" />
                  Tambah
                </Button>
              </div>
            </section>
          </CardHeader>

          <CardBody className="p-0 m-0 mt-2">
            <Col>
              <DataTables
                columns={columns}
                data={dataGradeScale?.data}
                pageCount={dataGradeScale?.data?.metadata.total_page}
                pagination={dataGradeScale?.data?.metadata}
                setPagination={handlePagination}
                total={dataGradeScale?.data?.metadata.total_data}
              />
            </Col>
          </CardBody>
        </Card>
      </Col>
    </Row>
  );
};

export default GradeScaleClientPage;
