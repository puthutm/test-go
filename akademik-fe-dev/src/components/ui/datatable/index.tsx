"use client";
import {
  ColumnDef,
  flexRender,
  getCoreRowModel,
  getSortedRowModel,
  useReactTable,
} from "@tanstack/react-table";
import { Table } from "reactstrap";
import classnames from "classnames";
import { Pagination } from "./Pagination";

interface DataTableProps<TData, TValue> {
  columns: ColumnDef<TData, TValue>[];
  data: TData[] | any;
  pageCount: number;
  pagination: React.ComponentState;
  setPagination: React.ComponentState;
  total: number;
  isLoading?: boolean;
  isPaginate?: boolean;
  rowSelection?: Record<string, boolean>;
  onRowSelectionChange?: React.Dispatch<
    React.SetStateAction<Record<string, boolean>>
  >;
}

interface PageChangeEvent {
  selected: number;
}

const DataTables = <TData, TValue>({
  columns,
  data,
  pageCount,
  pagination,
  setPagination,
  total,
  isLoading,
  isPaginate = true,
  rowSelection,
  onRowSelectionChange,
}: DataTableProps<TData, TValue>) => {
  const table = useReactTable({
    data: data?.data ?? [],
    columns,
    getCoreRowModel: getCoreRowModel(),
    rowCount: total,
    pageCount: pageCount || 1,
    manualPagination: true,
    manualSorting: true,
    enableSortingRemoval: true,
    sortDescFirst: true,
    getSortedRowModel: getSortedRowModel(),
    ...(rowSelection !== undefined && {
      enableRowSelection: true,
      state: { rowSelection },
      onRowSelectionChange,
    }),
  });

  const handlePageChange = (e: PageChangeEvent) => {
    table.setPageIndex(e.selected);
    setPagination(e.selected);
  };

  return (
    <>
      <Table hover>
        <thead className="table-light">
          {table.getHeaderGroups().map((headerGroup) => (
            <tr key={headerGroup.id}>
              {headerGroup.headers.map((header, index) => (
                <th
                  key={`${header.id}-${index}`}
                  className={classnames(
                    index === 0 && "index-table-width",
                    headerGroup.headers.length - 1 === index &&
                      "action-table-width",
                    "text-center"
                  )}
                  style={{ color: "#495057" }}
                >
                  {header.isPlaceholder
                    ? null
                    : flexRender(
                        header.column.columnDef.header,
                        header.getContext()
                      )}
                </th>
              ))}
            </tr>
          ))}
        </thead>
        <tbody>
          {isLoading ? (
            Array.from({ length: columns.length }).map((_, index) => (
              <tr key={index}>
                {columns?.map((item, index) => (
                  <td
                    key={`${item.id}-${index + 1}`}
                    className="placeholder-glow"
                  >
                    <span className="w-100 d-block placeholder rounded py-2"></span>
                  </td>
                ))}
              </tr>
            ))
          ) : table.getRowModel().rows?.length ? (
            table.getRowModel().rows.map((row, index) => (
              <tr
                key={`${row.id}-${index}`}
                data-state={row.getIsSelected() && "selected"}
                className="even:bg-[#F4F3FF] text-center"
              >
                {row.getVisibleCells().map((cell, index) => (
                  <td
                    key={`${cell}-${index}`}
                    className={classnames(
                      index === 0 && "border border-start-0 text-center",
                      index === row.getVisibleCells().length - 1 &&
                        "border border-end-0",
                      "border align-middle text-start"
                    )}
                  >
                    {flexRender(cell.column.columnDef.cell, cell.getContext())}
                  </td>
                ))}
              </tr>
            ))
          ) : (
            <tr>
              <td colSpan={columns.length} className="h-24 text-center">
                Tidak ada data
              </td>
            </tr>
          )}
        </tbody>
      </Table>
      {isPaginate && (
        <div className="d-flex mt-2 align-items-center justify-content-between gap-2 flex-wrap">
          {data?.data && (
            <>
              <div className="d-flex gap-2 text-primary">
                <p className="mb-0 text-dark">{data?.data?.length || 1}</p>
                <p className="mb-0 text-dark">of</p>
                <p className="mb-0 text-dark">{total}</p>
              </div>
              <Pagination
                pageCount={pageCount}
                pageOffset={pagination.page}
                onPageChange={handlePageChange}
              />
            </>
          )}
        </div>
      )}
    </>
  );
};

export default DataTables;
