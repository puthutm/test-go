"use client";

import { useEffect, useState } from "react";
import { useQueryClient } from "@tanstack/react-query";
import { useSearchParams } from "next/navigation";
import { Input } from "reactstrap";

import DataTables from "@/components/ui/datatable";
import { useColumnCurriculumStudyProgram } from "./column-table-curriculum-study-program";
import { useGetCurriculumStudyProgramBySemesterIdForProgramHead } from "@/services/api/curriculum/curriculum-study-program/program-head/use-get-curriculum-study-program-by-semester-id-for-program-hear";
import { AKADEMIK, KAPRODI } from "@/lib/constants/role";
import { useGetCurriculumStudyProgramBySemesterIdForAcademic } from "@/services/api/curriculum/curriculum-study-program/academic/use-get-curriculum-study-program-for-academic";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { updatePackageCurriculumStudyProgramForAcademic } from "@/services/api/curriculum/curriculum-study-program/academic/update-package-curriculum-study-program-for-academic";
import { updatePackageCurriculumStudyProgramForProgramHead } from "@/services/api/curriculum/curriculum-study-program/program-head/update-package-curriculum-study-program-for-program-head";

export const TableCurriculumStudyProgram = ({
  semester,
  role,
  studyProgramId,
}: {
  semester: SemesterNumberOptions;
  role: string;
  studyProgramId?: string;
}) => {
  const [isLoading, setIsLoading] = useState(false);
  const searchParams = useSearchParams();
  const curriculumYearParam = searchParams.get("curriculum_year");
  const { columns } = useColumnCurriculumStudyProgram();

  const queryClient = useQueryClient();

  const { setModalConfirmationState } = useModalConfirmationContext();

  let query;

  const curriculumStudyProgramAcademic =
    useGetCurriculumStudyProgramBySemesterIdForAcademic({
      semesterNumberId: semester.id,
      studyProgramId: studyProgramId as string,
      curriculumYearId: curriculumYearParam as string,
      role,
    });

  const curriculumStudyProgramForProgramHead =
    useGetCurriculumStudyProgramBySemesterIdForProgramHead({
      curriculumYearId: curriculumYearParam as string,
      semesterNumberId: semester.id,
      role: KAPRODI,
    });

  if (role === AKADEMIK) {
    query = curriculumStudyProgramAcademic;
  } else if (role === KAPRODI) {
    query = curriculumStudyProgramForProgramHead;
  }

  const handlePackageCurriculumStudyProgram = async ({
    isPackage,
  }: {
    isPackage: boolean;
  }) => {
    try {
      setIsLoading(true);
      const response =
        role === AKADEMIK
          ? await updatePackageCurriculumStudyProgramForAcademic({
              isPackage,
              semesterNumberId: semester.id,
              studyProgramId: studyProgramId as string,
              curriculumYearId: curriculumYearParam as string,
            })
          : await updatePackageCurriculumStudyProgramForProgramHead({
              isPackage,
              semesterNumberId: semester.id,
              curriculumYearId: curriculumYearParam as string,
            });

      if (response.error) {
        return setModalConfirmationState((prev) => ({
          ...prev,
          open: true,
          state: "failed",
          message: response.message,
        }));
      }

      queryClient.resetQueries({
        queryKey: [
          "curriculum-study-program-for-academic",
          semester.id,
          studyProgramId,
          curriculumYearParam,
          AKADEMIK,
        ],
      });

      queryClient.refetchQueries({
        queryKey: [
          "curriculum-study-program-for-program-head",
          semester.id,
          curriculumYearParam,
        ],
      });
    } catch (error: any) {
      setModalConfirmationState((prev) => ({
        ...prev,
        open: true,
        state: "failed",
        message: error.toString(),
      }));
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    if (curriculumYearParam && studyProgramId && role === AKADEMIK) {
      curriculumStudyProgramAcademic.refetch();
    }

    if (curriculumYearParam && role === KAPRODI) {
      curriculumStudyProgramForProgramHead.refetch();
    }
  }, [searchParams, role, studyProgramId]);

  return (
    <div className="d-flex flex-column gap-2">
      <div
        className="position-relative text-center"
        style={{
          backgroundColor: "#DEE5EC",
          padding: "15px 5px",
        }}
      >
        <h2
          className="m-0 p-0 fs-5 fw-medium rounded-2"
          style={{ color: "#495057" }}
        >
          Semester {semester.semester_number}
        </h2>
        <div className="position-absolute top-50 end-0 translate-middle-y me-3">
          <label
            className="form-check-label me-1"
            htmlFor={`is_package-${semester.id}`}
            style={{ fontWeight: 300 }}
          >
            Paket
          </label>
          <Input
            className="form-check-input"
            type="checkbox"
            id={`is_package-${semester.id}`}
            onChange={async (e) => {
              handlePackageCurriculumStudyProgram({
                isPackage: e.target.checked,
              });
            }}
            checked={query?.data?.data?.[0]?.is_package ?? false}
            disabled={isLoading}
            key={semester.id}
          />
        </div>
      </div>
      <div className="d-flex flex-column table-responsive">
        <DataTables
          columns={columns}
          data={query?.data?.data}
          pageCount={1}
          pagination={0}
          setPagination={() => {}}
          isLoading={query?.isLoading}
          total={1}
          isPaginate={false}
          key={semester.id}
        />
        <div
          className="text-center"
          style={{
            backgroundColor: "#DEE5EC",
            padding: "15px 5px",
            marginTop: "-10px",
          }}
        >
          <div
            className="d-flex gap-5 justify-content-center align-items-center fw-semibold"
            style={{ color: "#495057" }}
          >
            <div className="d-flex gap-3">
              <p>Total</p>
              <span>{query?.data?.data?.total?.sks} SKS</span>
            </div>
            <div className="d-flex flex-column text-start">
              <p>
                Wajib: <span>{query?.data?.data?.total?.mandatory}</span>
              </p>
              <p>
                Pilihan: <span>{query?.data?.data?.total?.no_mandatory}</span>
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};
