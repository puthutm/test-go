"use client";

import { Col, Row } from "reactstrap";
import { useDebouncedCallback } from "use-debounce";
import { usePathname, useRouter, useSearchParams } from "next/navigation";
import { signOut } from "next-auth/react";
import { useEffect, useState } from "react";

import DataTables from "@/components/ui/datatable";
import { SearchIcon } from "@/components/icons/search";
import { useColumnPresenceClass } from "./column-definition-presence-class";
import FormPresenceComponent from "./form-presence-component";

export const TablePresenceClass = ({
  data,
  presenceComponent,
}: {
  data: ApiResponse<PaginationData<ClassPresence>> | undefined;
  presenceComponent: ApiResponse<PresenceComponent> | undefined;
}) => {
  const [input, setInput] = useState("");
  const pathname = usePathname();
  const router = useRouter();
  const searchParams = useSearchParams();
  const params = new URLSearchParams(searchParams);
  const { columns } = useColumnPresenceClass();

  const handlePagination = (newPage: number) => {
    if (newPage) {
      params.set("page", (newPage + 1).toString());
    } else {
      params.delete("page");
    }

    router.replace(`${pathname}?${params.toString()}`);
  };

  const handleSearch = useDebouncedCallback((value: string) => {
    if (value) {
      params.set("q", value);
    } else {
      params.delete("q");
    }

    params.set("page", "1");
    router.replace(`${pathname}?${params.toString()}`);
  }, 1000);

  useEffect(() => {
    if (searchParams) {
      setInput(searchParams.get("q") ?? "");
    } else {
      setInput("");
    }
  }, [searchParams]);

  if (data?.status === 401) {
    signOut();
  }

  return (
    <>
      <div className="d-flex align-items-center justify-content-between px-3 border-bottom border-2 pb-3">
        <Row className="w-100 gap-2">
          <Col sm={12} lg={3} className="px-0">
            <div className="form-icon">
              <input
                className={`form-control form-control-icon`}
                id="no_kk"
                placeholder="Cari Mata Kuliah"
                onChange={(e) => {
                  setInput(e.target.value);
                  handleSearch(e.target.value);
                }}
                value={input}
              />
              <i>
                <SearchIcon />
              </i>
            </div>
          </Col>
        </Row>
      </div>
      <Col className=" ps-0 mt-3" sm={12}>
        <Row>
          <Col sm={7} className="table-responsive">
            <DataTables
              columns={columns}
              data={data?.data}
              pageCount={data?.data?.metadata?.total_page as number}
              pagination={data?.data?.metadata}
              setPagination={handlePagination}
              total={data?.data?.metadata?.total_data as number}
            />
          </Col>
          <Col sm={5}>
            <div className="border rounded-3 p-3">
              <FormPresenceComponent presenceComponent={presenceComponent} />
            </div>
          </Col>
        </Row>
      </Col>
    </>
  );
};
