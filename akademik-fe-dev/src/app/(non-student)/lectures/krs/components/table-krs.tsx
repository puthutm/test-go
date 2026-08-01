"use client";

import { Col, Row } from "reactstrap";
import { useDebouncedCallback } from "use-debounce";
import { usePathname, useRouter, useSearchParams } from "next/navigation";
import { signOut } from "next-auth/react";

import DataTables from "@/components/ui/datatable";
import { SelectComponent } from "@/components/ui/select";
import { FilterListIcon } from "@/components/icons/filter-list";
import { SearchIcon } from "@/components/icons/search";
import { useTableDefinitionKrsColumn } from "./column-definition-krs";
import { useGetSearchAcademicPeriod } from "@/services/api/data-referensi/academic-period/use-get-search-academic-period";
import { useEffect, useState } from "react";

type InputFilter = {
  academicPeriod: OptionType | undefined;
  search: string | undefined;
};

export const TableKrs = ({
  data,
}: {
  data: ApiResponse<PaginationData<any[]>>;
}) => {
  const [inputFilter, setInputFilter] = useState<InputFilter>({
    academicPeriod: undefined,
    search: "",
  });
  const pathname = usePathname();
  const router = useRouter();
  const searchParams = useSearchParams();
  const params = new URLSearchParams(searchParams);
  const { columns } = useTableDefinitionKrsColumn();

  const academicPeriodParam = searchParams.get("academicPeriod");
  const searchParam = searchParams.get("q");

  const { data: academicPeriod, isLoading: isLoadingAcademicPeriod } =
    useGetSearchAcademicPeriod();

  const academicPeriodOptions = academicPeriod?.data?.map((opt) => ({
    label: opt.fullname,
    value: opt.id,
  }));

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

  const handleFilterAcademicPeriod = (value: OptionType) => {
    if (value) {
      params.set("academicPeriod", value.value);
    } else {
      params.delete("academicPeriod");
    }

    router.replace(`${pathname}?${params.toString()}`);
  };

  useEffect(() => {
    if (searchParams && academicPeriodParam) {
      const findCurriculum = academicPeriodOptions?.find(
        (data) => data.value === academicPeriodParam
      );
      setInputFilter((prev) => ({
        ...prev,
        academicPeriod: findCurriculum as OptionType,
        search: searchParam ?? "",
      }));
    } else {
      setInputFilter((prev) => ({
        ...prev,
        academicPeriod: undefined,
      }));
    }
  }, [searchParams, academicPeriod]);

  if (data.status === 401) {
    signOut();
  }

  return (
    <>
      <div className="d-flex align-items-center justify-content-between px-3 border-bottom border-2 pb-3">
        <Row className="w-100 gap-2">
          <Col sm={12} lg={3} className="ps-0">
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
                value={inputFilter.academicPeriod}
              />
              <i>
                <FilterListIcon />
              </i>
            </div>
          </Col>
          <Col sm={12} lg={3} className="px-0">
            <div className="form-icon">
              <input
                className={`form-control form-control-icon`}
                id="no_kk"
                placeholder="Cari Mata Kuliah"
                onChange={(e) => {
                  handleSearch(e.target.value);
                  setInputFilter((prev) => ({
                    ...prev,
                    search: e.target.value,
                  }));
                }}
                value={inputFilter.search}
              />
              <i>
                <SearchIcon />
              </i>
            </div>
          </Col>
        </Row>
      </div>
      <Col className="table-responsive ps-0 mt-3" sm={12}>
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
