"use client";

import { ColumnDef } from "@tanstack/react-table";
import { useParams } from "next/navigation";

import { useGetPresenceComponentBySessionId } from "@/services/api/settings/presence/lecturer/use-get-presence-component-by-session-id";
import { FormComponentStudentPresence } from "./form-component-presence";

export const useColumnStudentPresenceBySessionId = () => {
  const params = useParams();

  const sessionId = params.sessionId;

  const { data: presenceComponent } = useGetPresenceComponentBySessionId({
    sessionId: sessionId as string,
  });

  const dynamicColumns = [
    presenceComponent?.data.use_open_session
      ? {
          accessorKey: "open_session",
          header: `Buka Sesi (${presenceComponent?.data?.open_session_percentage}%)`,
          cell: ({ row }: { row: any }) => {
            const defaultValue = Boolean(row.original.open_session_percentage);
            return (
              <FormComponentStudentPresence
                defaultValue={defaultValue}
                presenceType={"open_session"}
                studentId={row.original.student_id}
              />
            );
          },
        }
      : undefined,
    presenceComponent?.data.use_document_material
      ? {
          accessorKey: "document_material",
          header: `View/Download Materi (${presenceComponent?.data?.document_material_percentage}%)`,
          cell: ({ row }: { row: any }) => {
            const defaultValue = Boolean(
              row.original.document_material_percentage
            );
            return (
              <FormComponentStudentPresence
                defaultValue={defaultValue}
                presenceType={"document_material"}
                studentId={row.original.student_id}
              />
            );
          },
        }
      : undefined,
    presenceComponent?.data.use_comment
      ? {
          accessorKey: "comment",
          header: `Komentar (${presenceComponent?.data?.comment_percentage}%)`,
          cell: ({ row }: { row: any }) => {
            const defaultValue = Boolean(row.original.comment_percentage);
            return (
              <FormComponentStudentPresence
                defaultValue={defaultValue}
                presenceType={"comment"}
                studentId={row.original.student_id}
              />
            );
          },
        }
      : undefined,
    presenceComponent?.data.use_video
      ? {
          accessorKey: "video",
          header: `Video (${presenceComponent?.data?.video_percentage}%)`,
          cell: ({ row }: { row: any }) => {
            const defaultValue = Boolean(row.original.video_percentage);
            return (
              <FormComponentStudentPresence
                defaultValue={defaultValue}
                presenceType={"video"}
                studentId={row.original.student_id}
              />
            );
          },
        }
      : undefined,
    presenceComponent?.data.use_task
      ? {
          accessorKey: "task",
          header: `Tugas (${presenceComponent?.data?.task_percentage}%)`,
          cell: ({ row }: { row: any }) => {
            const defaultValue = Boolean(row.original.task_percentage);
            return (
              <FormComponentStudentPresence
                defaultValue={defaultValue}
                presenceType={"task"}
                studentId={row.original.student_id}
              />
            );
          },
        }
      : undefined,
    presenceComponent?.data.use_quiz
      ? {
          accessorKey: "quiz",
          header: `Quis (${presenceComponent?.data?.quiz_percentage}%)`,
          cell: ({ row }: { row: any }) => {
            const defaultValue = Boolean(row.original.quiz_percentage);
            return (
              <FormComponentStudentPresence
                defaultValue={defaultValue}
                presenceType={"quiz"}
                studentId={row.original.student_id}
              />
            );
          },
        }
      : undefined,
    presenceComponent?.data.use_uts
      ? {
          accessorKey: "uts",
          header: `UTS (${presenceComponent?.data?.uts_percentage}%)`,
          cell: ({ row }: { row: any }) => {
            const defaultValue = Boolean(row.original.uts_percentage);
            return (
              <FormComponentStudentPresence
                defaultValue={defaultValue}
                presenceType={"uts"}
                studentId={row.original.student_id}
              />
            );
          },
        }
      : undefined,
    presenceComponent?.data.use_uas
      ? {
          accessorKey: "uas",
          header: `UAS (${presenceComponent?.data?.uas_percentage}%)`,
          cell: ({ row }: { row: any }) => {
            const defaultValue = Boolean(row.original.uas_percentage);
            return (
              <FormComponentStudentPresence
                defaultValue={defaultValue}
                presenceType={"uas"}
                studentId={row.original.student_id}
              />
            );
          },
        }
      : undefined,
  ].filter((col) => col !== undefined);

  const staticColumns: ColumnDef<StudentPresenceSession>[] = [
    {
      accessorKey: "student_id",
      header: () => "No.",
      cell: ({ row }) => {
        return row.index + 1;
      },
    },
    {
      accessorKey: "student_name",
      header: "Nama",
    },
    {
      accessorKey: "student_nim",
      header: "NIM",
    },
  ];

  const columns: ColumnDef<StudentPresenceSession>[] = [
    ...staticColumns,
    ...dynamicColumns,
    {
      accessorKey: "total_percentage",
      header: "Persentase",
      cell: ({ row }) => {
        const total = row.original.total_percentage;

        return `${total} %`;
      },
    },
    {
      accessorKey: "presence_flag",
      header: "Status",
      cell: ({ row }) => {
        const total = row.original.total_percentage;
        const isPresence = total === 100;
        return (
          <span
            className="py-1 px-2 rounded"
            style={{
              backgroundColor: isPresence ? "#6CBE401A" : "#F065481A",
              color: isPresence ? "#6CBE40" : "#F06548",
            }}
          >
            {isPresence ? "Hadir" : "Tidak Hadir"}
          </span>
        );
      },
    },
  ];

  return { columns };
};
