"use client";

import { Button } from "reactstrap";

import DataTables from "@/components/ui/datatable";
import { useTableDefinitionKrsDetailColumn } from "./column-definition-krs-detail";
import { ModalUpdateStatusKrs } from "./modal-update-status-krs";

interface Props {
  data: ApiResponse<KrsDetail>;
}

export const TableKrsDetail = ({ data }: Props) => {
  const dataTable = {
    data: data?.data?.krs_items,
  };
  const { columns } = useTableDefinitionKrsDetailColumn();
  return (
    <div className="d-flex flex-column gap-2">
      {data?.data?.krs_items.length ? (
        <div className="d-flex justify-content-end gap-2">
          <Button className="btn-success">Terima</Button>
          <Button className="btn-danger">Tidak</Button>
        </div>
      ) : null}
      <DataTables
        columns={columns}
        data={dataTable}
        pageCount={1}
        pagination={{}}
        setPagination={() => {}}
        total={10}
        isPaginate={false}
      />
      <ModalUpdateStatusKrs />
    </div>
  );
};
