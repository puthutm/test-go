import { Table } from "reactstrap";

import { VisibilityIcon } from "@/components/icons/visibility";
import { useModalContext } from "@/lib/hooks/use-modal";
import { formatDate } from "@/lib/utils/format-date";

export const TableProposal = ({
  data,
}: {
  data: FinalProjectProposalStudent[];
}) => {
  const { setModalState } = useModalContext();
  return (
    <div className="table-responsive">
      <Table
        hover
        style={{ tableLayout: "auto" }}
        className="mt-3 align-center"
      >
        <thead className="table-light text-center">
          <tr>
            <th scope="col" style={{ maxWidth: "126px" }}>
              Tgl Pengajuan
            </th>
            <th scope="col">Judul Proposal</th>
            <th scope="col">Status Pengajuan</th>
            <th scope="col">Aksi</th>
          </tr>
        </thead>
        <tbody>
          {data?.map((item, index) => (
            <tr key={index}>
              <td>{formatDate(item?.date)}</td>
              <td>{item.title_id}</td>
              {/* <td>{item.field}</td> */}
              <td className="text-center">
                <span
                  className="px-2 py-1 rounded badge"
                  style={{
                    backgroundColor:
                      item.status === 0
                        ? "#F7B84B3B"
                        : item?.status === 1
                        ? "#6CBE401A"
                        : item.status
                        ? "#C9C0BE"
                        : "#F065481A",
                    color:
                      item.status === 0
                        ? "#F7B84B"
                        : item.status === 1
                        ? "#6CBE40"
                        : item.status === 2
                        ? "#A14835"
                        : "#F06548",
                    fontSize: "12px",
                  }}
                >
                  {item.status === 0 ? "Menunggu Persetujuan" : null}
                  {item.status === 1 ? "Disetujui" : null}
                  {item.status === 2 ? "Revisi" : null}
                  {item.status === 3 ? "Ditolak" : null}
                </span>
              </td>
              <td className="text-center">
                <VisibilityIcon
                  onClick={() => {
                    setModalState((prev) => ({
                      ...prev,
                      open: true,
                      data: item,
                      id: item.id,
                      state: "detail",
                    }));
                  }}
                  className="cursor-pointer"
                />
              </td>
            </tr>
          ))}
        </tbody>
      </Table>
    </div>
  );
};
