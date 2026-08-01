"use client";

import { ColumnDef } from "@tanstack/react-table";

import { CheckIcon } from "@/components/icons/check";
import { CloseIcon } from "@/components/icons/close";

export const useColumnTableClassScore = (
  gradeComposition: ApiResponse<PaginationData<IGradeComposition>> | undefined
) => {
  const scoreMapping: Record<string, keyof ClassScore> = {
    UTS: "uts_score",
    UAS: "uas_score",
    Tugas: "task_score",
    Kehadiran: "presence_score",
  };

  const dynamicColumns: ColumnDef<ClassScore>[] =
    gradeComposition?.data?.data?.map((item) => ({
      id: item.id,
      header: `${item.value_element_name} (${item.percentage}%)`,
      cell: ({ row }) => {
        const field = scoreMapping[item.value_element_name];
        const value = field ? row.original[field] : 0;

        return <p className="text-center">{value ?? 0}</p>;
      },
    })) ?? [];

  const columns: ColumnDef<ClassScore>[] = [
    {
      accessorKey: "nim",
      header: "NIM",
    },
    {
      accessorKey: "student_name",
      header: "Name",
    },
    ...dynamicColumns,
    {
      accessorKey: "final_score",
      header: "Score",
      cell: ({ row }) => (
        <p className="text-center">{row.original.final_score}</p>
      ),
    },
    {
      accessorKey: "grade_name",
      header: "Grade",
      cell: ({ row }) => (
        <p className="text-center">{row.original.grade_name}</p>
      ),
    },
    {
      accessorKey: "is_passed",
      header: "Passed",
      cell: ({ row }) => (
        <div className="d-flex justify-content-center">
          {row.original.is_passed ? (
            <CheckIcon color="#0AB39C" />
          ) : (
            <CloseIcon color="#F06548" />
          )}
        </div>
      ),
    },
    {
      accessorKey: "pass_note",
      header: "Note",
    },
  ];
  return { columns };
};
