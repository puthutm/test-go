"use client";
import React from "react";

// import third party component
import { Button } from "reactstrap";
import Link from "next/link";
// import component
import { EditIcon } from "@/components/icons/edit";
import DataTables from "@/components/ui/datatable";

import useColumnEvaluationPlan from "../_columns/column-definition-evaluation-plan";

export interface IDummyValueTable {
  id: string;
  basis_evaluation: string;
  component_evaluation: string;
  deskripsi_ind: string;
  deskripsi_eng: string;
  bobot: string;
}

const dummyValueTable: PaginationData<IDummyValueTable> = {
  metadata: {
    page: 1,
    size: 1,
    total_data: 1,
    total_page: 1,
  },
  data: [
    {
      id: "1",
      basis_evaluation: "Aktivitas Partisipatif",
      component_evaluation: "-",
      deskripsi_ind: "text",
      deskripsi_eng: "text",
      bobot: "100",
    },
  ],
};

function SectionEvaluationPlan() {
  // const [queryParams, setQueryParams] = useState<QueryParam>({
  //   page: 1,
  //   filter: null,
  //   page_size: null,
  //   sort_by: null,
  //   sort_direction: null,
  // });

  //   columns
  const { columns } = useColumnEvaluationPlan();
  return (
    <>
      <section className="mt-3 position-relative ">
        {/*//! action */}
        <section className="d-flex justify-content-end  align-items-center gap-2">
          {/*//! edit btn */}
          <Button
            className="btn  d-flex align-items-center gap-2"
            color="transparent"
            style={{
              color: "#10487A",
              padding: "5.5px 12px",
              border: "1px solid #10487A",
            }}
          >
            <EditIcon color="#10487A" width="15" height={"15"} />
            Edit
          </Button>
        </section>
        {/*//! tables  */}

        <div className="table-responsive mt-2">
          <DataTables
            columns={columns}
            data={dummyValueTable}
            pageCount={0}
            pagination={null}
            setPagination={null}
            isLoading={false}
            total={0}
            isPaginate={false}
          />
        </div>

        {/*//! total */}
        <div
          className="position-relative fw-semibold py-3 px-5 border-t border-4 rounded-1 d-flex align-items-center gap-2"
          style={{
            background: "#FFE91D33",
            color: "#495057",
          }}
        >
          <p className="m-0 p-0">Total :</p>
          <p className="m-0 p-0">0</p>
        </div>
      </section>

      {/*//! section button back */}
      <section className="position-relative mt-3 d-flex justify-content-starts">
        <Link
          href={"?tab=study"}
          className="btn  d-flex align-items-center gap-2"
          color="transparent"
          style={{ color: "#10487A", border: "1px solid #10487A" }}
        >
          Kembali
        </Link>
      </section>
    </>
  );
}

export default SectionEvaluationPlan;
