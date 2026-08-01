'use client'
import React from 'react'
import { Button } from 'reactstrap'
import { AddIcon } from '@/components/icons/add'
import DataTables from '@/components/ui/datatable'
import useColumnDefLectureTeaching from './column-def-lecture-teaching'

export interface IDummyValueTable {
  id:string,
  dosen_pengajar:string,
  dosen_pengganti:string,
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
      dosen_pengajar:'Novi Dian Natashia, S.Kom., MMSI ',
      dosen_pengganti:'Novi Dian Natashia, S.Kom., MMSI ',
    },
]
}

function SectionLectureTeaching() {

    const { columns } = useColumnDefLectureTeaching()


  return (
    <section className="position-relative p-3 bg-white rounded-3">
        {/*//! */}
        <h2 className="m-0 p-0 fs-5 fw-semibold">
            Dosen Pengajar 
        </h2>

        {/*//! section data */}
        <section className="position-relative mt-2 ">
            {/*//! head */}
            <section className="d-flex justify-content-end">
                {/*//! button add */}
                <Button className='btn  d-flex align-items-center gap-2 '
                  color='transparent'
                  disabled
                  style={{color:'#10487A',border:'1px solid #10487A'}}
                  >
                    <AddIcon color='#10487A'/>
                    Tambah
                </Button>
            </section>

            {/*//! table */}
            <section className="table-responsive mt-2">
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
        </section>
    </section>
  )
}

export default SectionLectureTeaching