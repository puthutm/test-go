"use client";

import {
  useParams,
  usePathname,
  useRouter,
  useSearchParams,
} from "next/navigation";
import { useDebouncedCallback } from "use-debounce";
import { useEffect, useState } from "react";
import { Col, Row } from "reactstrap";

// import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { useColumnTableClassScore } from "./column-table-class-score";
import { SearchIcon } from "@/components/icons/search";
import DataTables from "@/components/ui/datatable";
import { ButtonOpenCloseClassScore } from "./button-open-close-class-score";

interface Props {
  gradeComposition: ApiResponse<PaginationData<IGradeComposition>> | undefined;
  data: ApiResponse<PaginationData<ClassScore>> | undefined;
  statusLockClassScore: boolean;
}

export const TableClassScore = ({
  data,
  gradeComposition,
  statusLockClassScore,
}: Props) => {
  const pathname = usePathname();
  const params = useParams();
  const router = useRouter();
  const searchParams = useSearchParams();
  const queryParams = new URLSearchParams(searchParams);
  const academicPeriodId = params.academicPeriodId as string;
  const classId = params.classId as string;

  const initialSearch = searchParams.get("q") ?? "";
  const [search, setSearch] = useState(initialSearch);

  // const { modalConfirmationState, setModalConfirmationState } =
  //   useModalConfirmationContext();

  const { columns } = useColumnTableClassScore(gradeComposition);

  const handleSearch = useDebouncedCallback((value: string) => {
    if (value) {
      queryParams.set("q", value);
    } else {
      queryParams.delete("q");
    }

    queryParams.set("page", "1");
    router.push(`${pathname}?${queryParams.toString()}`);
  }, 1000);

  const handlePagination = (newPage: number) => {
    if (newPage) {
      queryParams.set("page", (newPage + 1).toString());
    } else {
      queryParams.delete("page");
    }

    router.push(`${pathname}?${queryParams.toString()}`);
  };

  useEffect(() => {
    if (initialSearch) {
      setSearch(initialSearch);
    }
  }, [initialSearch]);

  return (
    <div className="d-flex flex-column gap-3">
      <div className="d-flex justify-content-between align-items-center">
        <p className="fw-medium fs-5" style={{ color: "#3A3A3A" }}>
          Nilai Perkuliahan
        </p>
        <ButtonOpenCloseClassScore
          academicPeriodId={academicPeriodId}
          classId={classId}
          statusLock={statusLockClassScore}
        />
      </div>
      <Row>
        <Col sm={4}>
          <div className="form-icon">
            <input
              className={`form-control form-control-icon`}
              placeholder="Cari peserta kelas"
              onChange={(e) => {
                handleSearch(e.target.value);
                setSearch(e.target.value);
              }}
              value={search}
            />
            <i>
              <SearchIcon />
            </i>
          </div>
        </Col>
      </Row>
      <DataTables
        data={data?.data}
        columns={columns}
        pageCount={Number(data?.data?.metadata.total_page)}
        pagination={data?.data?.metadata}
        setPagination={handlePagination}
        total={Number(data?.data?.metadata.total_data)}
      />
    </div>
  );
};
