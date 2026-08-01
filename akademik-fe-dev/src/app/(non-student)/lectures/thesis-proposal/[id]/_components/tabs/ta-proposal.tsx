'use client'
import React from "react";
// import { Table } from "reactstrap";
import DataTables from "@/components/ui/datatable";

import useColumnDefTaProposal from "../columns/column-def-ta-proposal";

export interface IDummyValueTable {
  id:string,
  tanggal_pengajuan:string,
  judul_proposal:string,
  bidang:string,
  status_pengajuan:string,
}

const dummyValueTable : PaginationData<IDummyValueTable> = {
  metadata:{
    page: 1,
    size: 10,
    total_data: 10,
    total_page: 1,
  },
  data:[
     {
    id:'1',
    tanggal_pengajuan:'12 Mei 2025',
    judul_proposal:'Implementasi Otomatisasi CI/CD pada Pengembangan Aplikasi Web Menggunakan GitHub Actions dan Docker',
    bidang:'Pengembangan Aplikasi',
    status_pengajuan:'Aktif'
    },
  ]
}

const TAProposalTabContent = () => {

  const { columns } = useColumnDefTaProposal();

  return (
    <div>
      <div className="border-bottom border-3 mb-3">
        <h5 className="fw-semibold">Proposal Tugas Akhir</h5>
      </div>

            {/* //! tables */}
          <section className="table-responsive mt-3">
              <DataTables
                    columns={columns}
                    data={dummyValueTable}
                    pageCount={1}
                    pagination={dummyValueTable.metadata.page}
                    setPagination={() => {}}
                    isLoading={false}
                    total={1}
                  />
          </section>
    </div>
  );
};

export default TAProposalTabContent;
