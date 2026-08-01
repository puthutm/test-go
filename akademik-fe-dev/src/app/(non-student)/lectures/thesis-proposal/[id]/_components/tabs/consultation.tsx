'use client'
import React from "react";
import DataTables from "@/components/ui/datatable";

import useColumnDefConsulThesis from "../columns/column-def-consul-tesis";

export interface IDummyValueTable {
  id:string,
  waktu:string,
  dosen_pembimbing:string,
  topik:string,
  feedback:string,
  status:string,
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
    waktu:'16 Mei 2025',
    dosen_pembimbing:'Novi Dian Natashia, S.Kom., MMSI',
    topik:'Bab 1',
    feedback:'Lanjutkan',
    status:'Disetujui',
    },
     {
    id:'2',
    waktu:'17 Mei 2025',
    dosen_pembimbing:'Novi Dian Natashia, S.Kom., MMSI',
    topik:'Bab 2',
    feedback:'Bagus',
    status:'Disetujui',
    },
    {
    id:'2',
    waktu:'18 Mei 2025',
    dosen_pembimbing:'Novi Dian Natashia, S.Kom., MMSI',
    topik:'Bab 3',
    feedback:'Perbaikan pada Bab 3',
    status:'Revisi',
    }
  ]
}

const ConsultationTabContent = () => {

  const { columns } = useColumnDefConsulThesis();

  return (
    <div>
      <div className="border-bottom border-3 mb-2">
        <h5 className="fw-semibold">Proposal Tugas Akhir</h5>
      </div>
      <div>
        <p>Judul Tugas Akhir:</p>
        <p className="fw-semibold text-primary">{'Implementasi Otomatisasi CI/CD pada Pengembangan Aplikasi Web Menggunakan GitHub Actions dan Docker'}</p>
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

export default ConsultationTabContent;
