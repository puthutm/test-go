"use client";

import { Col, Row } from "reactstrap";
import { debounce } from "lodash";
import { useParams, usePathname, useRouter } from "next/navigation";

import DataTables from "@/components/ui/datatable";
import { SearchIcon } from "@/components/icons/search";
import { useTableSubjectAcademicPeriodColumns } from "./column-table-subject-academic-period";
import { RefreshIcon } from "@/components/icons/refresh";
import { AddIcon } from "@/components/icons/add";
import { FileDownloadIcon } from "@/components/icons/file-download";
import { useModalContext } from "@/lib/hooks/use-modal";
import { ModalAddClass } from "./modal-add-class";

export const TableSubjectAcademicPeriod = ({
  searchParams,
  data,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
  data: ApiResponse<PaginationData<Class[]>>;
}) => {
  const { setModalState } = useModalContext();

  const params = useParams();
  const pathname = usePathname();
  const router = useRouter();
  const searchParam = new URLSearchParams(searchParams as any);
  const { columns } = useTableSubjectAcademicPeriodColumns(params);

  const handlePagination = (newPage: number) => {
    if (newPage) {
      searchParam.set("page", newPage.toString());
    } else {
      searchParam.delete("page");
    }

    router.push(`${pathname}?${searchParam.toString()}`);
  };

  const handleSearch = debounce((value: string) => {
    if (value) {
      searchParam.set("q", value);
    } else {
      searchParam.delete("q");
    }

    searchParam.set("page", "1");
    router.push(`${pathname}?${searchParam.toString()}`);
  }, 1000);
  return (
    <>
      <ModalAddClass academicPeriodId={params?.academicPeriodId as string} />
      <div className="d-flex justify-content-between">
        <Row className="w-50">
          <Col sm={5}>
            <div className="form-icon">
              <input
                className={`form-control form-control-icon`}
                id="no_kk"
                placeholder="Keywoard"
                onChange={(e) => handleSearch(e.target.value)}
              />
              <i>
                <SearchIcon />
              </i>
            </div>
          </Col>
          <Col sm={2}>
            <button
              className="btn-outline text-primary"
              style={{ padding: "10px" }}
            >
              <RefreshIcon />
            </button>
          </Col>
        </Row>
        <div className="d-flex gap-2">
          <button
            className="btn-outline text-primary px-3"
            onClick={() => {
              setModalState((prev) => ({
                ...prev,
                open: true,
                state: "add",
                id: params.academicPeriodId as string,
              }));
            }}
          >
            <AddIcon color="#10487A" />
            Tambah
          </button>
          <button className="btn btn-primary px-3">
            <FileDownloadIcon
              height="16"
              width="16"
              color="white"
              className="me-1"
            />
            Download
          </button>
        </div>
      </div>
      <div className="table-responsive">
        <DataTables
          columns={columns}
          data={data?.data}
          pageCount={data?.data?.metadata?.total_page}
          pagination={data?.data?.metadata}
          setPagination={handlePagination}
          total={data?.data?.metadata?.total_data}
        />
      </div>
    </>
  );
};
