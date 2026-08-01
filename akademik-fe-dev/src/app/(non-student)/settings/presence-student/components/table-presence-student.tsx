"use client";

import { Col, Row } from "reactstrap";
import { usePathname, useRouter, useSearchParams } from "next/navigation";
import { signOut } from "next-auth/react";

import { AddIcon } from "@/components/icons/add";
import DataTables from "@/components/ui/datatable";
import { useColumnDefinitionPresence } from "./column-definition-presence-presence";
import { useGetSearchAcademicPeriod } from "@/services/api/data-referensi/academic-period/use-get-search-academic-period";
import { useGetUnsiaStudyProgram } from "@/services/api/data-referensi/study-program/use-get-unsia-study-program";
import { SelectComponent } from "@/components/ui/select";
import { FilterListIcon } from "@/components/icons/filter-list";
import { ModalDuplicatePresenceStudent } from "./modal-duplicate-presense-student";
import { useEffect, useState } from "react";

export const TablePresenceStudent = ({
  data,
}: {
  data: ApiResponse<PaginationData<Presences[]>>;
}) => {
  const [studyProgramId, setStudyProgramId] = useState("");
  const pathname = usePathname();
  const router = useRouter();
  const searchParams = useSearchParams();
  const params = new URLSearchParams(searchParams);

  const academicPeriodParam = searchParams.get("academicPeriod");
  const studyProgramParam = searchParams.get("handleFilterStudyProgram");
  const [academicPeriodState, setAcademicPeriodState] = useState<
    OptionType | undefined
  >(undefined);
  const [studyProgramState, setStudyProgramState] = useState<
    OptionType | undefined
  >(undefined);

  const { data: academicPeriod, isLoading: isLoadingAcademicPeriod } =
    useGetSearchAcademicPeriod();

  const academicPeriodOptions = academicPeriod?.data?.map((opt) => ({
    label: opt.fullname,
    value: opt.id,
  }));

  const { data: studyProgram, isLoading: isLoadingStudyProgram } =
    useGetUnsiaStudyProgram();

  const studyProgramOptions = studyProgram?.data?.map((opt) => ({
    label: opt.name,
    value: opt.id,
  }));

  const { columns } = useColumnDefinitionPresence({ setStudyProgramId });

  const handlePagination = (newPage: number) => {
    if (newPage) {
      params.set("page", (newPage + 1).toString());
    } else {
      params.delete("page");
    }

    router.push(`${pathname}?${params.toString()}`);
  };

  const handleFilterAcademicPeriod = (value: OptionType) => {
    if (value) {
      params.set("academicPeriod", value.value);
    } else {
      params.delete("academicPeriod");
    }

    router.replace(`${pathname}?${params.toString()}`);
  };

  const handleFilterStudyProgram = (value: OptionType) => {
    if (value) {
      params.set("studyProgram", value.value);
    } else {
      params.delete("studyProgram");
    }

    router.replace(`${pathname}?${params.toString()}`);
  };

  useEffect(() => {
    if (searchParams && academicPeriodParam) {
      const findCurriculum = academicPeriodOptions?.find(
        (data) => data.value === academicPeriodParam
      );
      setAcademicPeriodState(findCurriculum as OptionType);
    } else {
      setAcademicPeriodState(undefined);
    }
  }, [searchParams, academicPeriod]);

  useEffect(() => {
    if (searchParams && studyProgramParam) {
      const findStudyProgram = studyProgramOptions?.find(
        (data) => data.value === studyProgramParam
      );
      setStudyProgramState(findStudyProgram);
    } else {
      setStudyProgramState(undefined);
    }
  }, [searchParams, studyProgram]);

  if (data.status === 401) {
    signOut();
  }
  return (
    <>
      <div className="d-flex align-items-center justify-content-between px-3">
        <ModalDuplicatePresenceStudent studyProgramId={studyProgramId} />
        <Row style={{ width: "50%" }}>
          <Col sm={5}>
            <div className="form-icon">
              <SelectComponent
                options={academicPeriodOptions as OptionType[]}
                id="filter"
                placeholder="Filter Periode Akademik"
                hasIcon
                onChange={(value) => handleFilterAcademicPeriod(value)}
                isLoading={isLoadingAcademicPeriod}
                isDisabled={isLoadingAcademicPeriod}
                isClearable
                value={academicPeriodState}
              />
              <i>
                <FilterListIcon />
              </i>
            </div>
          </Col>
          <Col sm={5}>
            <div className="form-icon">
              <SelectComponent
                options={studyProgramOptions as OptionType[]}
                id="filter"
                placeholder="Filter Program Studi"
                hasIcon
                isLoading={isLoadingStudyProgram}
                isDisabled={isLoadingStudyProgram}
                onChange={(value) => handleFilterStudyProgram(value)}
                isClearable
                value={studyProgramState}
              />
              <i>
                <FilterListIcon />
              </i>
            </div>
          </Col>
          {/* <Col sm={2} className="p-0">
            <Button color="danger" onClick={() => handleClearFilter()}>
              <CloseIcon color="white" />
            </Button>
          </Col> */}
        </Row>
        <div className="d-flex gap-2">
          <button
            className="btn-outline h-100 py-2 px-3 text-primary"
            onClick={() => router.push(`${pathname}/add`)}
          >
            <AddIcon color="#10487A" />
            Tambah
          </button>
        </div>
      </div>
      <Col className="table-responsive px-3 mt-3" sm={12}>
        <DataTables
          columns={columns}
          data={data?.data}
          pageCount={data?.data?.metadata?.total_page}
          pagination={data?.data?.metadata}
          setPagination={handlePagination}
          total={data?.data?.metadata?.total_data}
        />
      </Col>
    </>
  );
};
