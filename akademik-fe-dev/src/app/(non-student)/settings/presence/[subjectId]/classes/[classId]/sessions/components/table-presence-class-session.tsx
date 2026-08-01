"use client";

import { Col } from "reactstrap";

import DataTables from "@/components/ui/datatable";
import { useColumnPresenceClassSession } from "./column-definition-presence-class-session";

export const TablePresenceClassSession = ({
  data,
}: {
  data: ApiResponse<ClassPresenceSession[]> | undefined;
}) => {
  const { columns } = useColumnPresenceClassSession();
  return (
    <Col className="table-responsive ps-0 mt-3" sm={12}>
      <DataTables
        columns={columns}
        data={data}
        pageCount={1}
        pagination={1}
        setPagination={() => {}}
        total={1}
        isPaginate={false}
      />
    </Col>
  );
};
