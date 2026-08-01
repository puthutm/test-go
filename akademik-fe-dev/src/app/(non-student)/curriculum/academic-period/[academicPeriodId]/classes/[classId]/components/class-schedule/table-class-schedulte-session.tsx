"use client";

import { usePathname, useRouter } from "next/navigation";
import { useState } from "react";
import { Spinner } from "reactstrap";

import DataTables from "@/components/ui/datatable";
import { useColumnClassScheduleSession } from "./column-class-schedule-session";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { generateClassSchedule } from "@/services/api/curriculum/academic-period/class-schedule/generate-class-schedule";

export const TableClassScheduleSession = ({
  data,
  searchParams,
  classId,
  isDetail,
}: {
  data: ApiResponse<PaginationData<ClassSchedule>>;
  searchParams: { [key: string]: string | string[] | undefined };
  classId: string;
  isDetail?: boolean;
}) => {
  const [loadingGenerate, setLoadingGenerate] = useState<boolean>(false);
  const pathname = usePathname();
  const router = useRouter();
  const searchParam = new URLSearchParams(searchParams as any);
  const { columns } = useColumnClassScheduleSession();

  const { setModalConfirmationState } = useModalConfirmationContext();

  const handlePagination = (newPage: number) => {
    if (newPage) {
      searchParam.set("page", (newPage + 1).toString());
    } else {
      searchParam.delete("page");
    }

    router.replace(`${pathname}?${searchParam.toString()}`);
  };

  const handleGenerateClassSchedule = async () => {
    try {
      setLoadingGenerate(true);
      const response = await generateClassSchedule({
        classId,
      });

      if (response.error) {
        return setModalConfirmationState((prev) => ({
          ...prev,
          open: true,
          state: "failed",
          message: response.message,
        }));
      }

      return setModalConfirmationState((prev) => ({
        ...prev,
        open: true,
        message: `Jadwal berhasil di-generate`,
        state: "success",
      }));
    } catch (error: any) {
      setModalConfirmationState((prev) => ({
        ...prev,
        open: true,
        state: "failed",
        message: error.toString(),
      }));
    } finally {
      setLoadingGenerate(false);
    }
  };
  return (
    <div className="d-flex flex-column gap-2">
      <div className={`d-flex align-items-center justify-content-end`}>
        {!isDetail ? (
          <button
            className="btn-outline px-3 text-primary"
            onClick={() => handleGenerateClassSchedule()}
            disabled={loadingGenerate}
            style={{ minWidth: "100px", height: "2rem" }}
          >
            {loadingGenerate ? <Spinner size={"sm"} /> : "Generate"}
          </button>
        ) : null}
      </div>
      <div className="table-responsive">
        <DataTables
          columns={columns}
          data={data?.data}
          pageCount={data?.data.metadata.total_page}
          pagination={data?.data?.metadata}
          setPagination={handlePagination}
          total={data.data.metadata.total_data}
        />
      </div>
    </div>
  );
};
