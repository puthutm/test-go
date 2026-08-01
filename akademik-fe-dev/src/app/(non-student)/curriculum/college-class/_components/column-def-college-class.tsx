

'use client' 
import { ColumnDef } from "@tanstack/react-table" 
    // import component 
 import Link from "next/link" 
//  import { EditIcon } from "@/components/icons/edit" 
 import { EyeAkademikIcon } from "@/components/icons/eye-akademik"

// import hook 
import { IDummyValueTable } from "./client"

interface iColumnsParams {     ():{ columns: ColumnDef<IDummyValueTable>[] } }

const useColumnDefcollegeClass : iColumnsParams = ()=>{


    const columns : ColumnDef<IDummyValueTable>[] = [
        {
            header: 'No',
            accessorKey: "no",
            enableColumnFilter: false,
            cell: ({row}) => {
              return <p className="m-0 p-0 text-center">{(row.index+1)}</p>;
            },
        },
        {
            header: "Periode",
            accessorKey: "periode",
            enableColumnFilter: false,
            cell: ({row}) => {
              return <p className="m-0 p-0 text-center">{row.original?.periode ?? '-'}</p>;
            },
        },
        {
            header: "Kur .",
            accessorKey: "kurikulum",
            enableColumnFilter: false,
            cell: ({row}) => {
              return <p className="m-0 p-0 text-center">{row.original?.kurikulum ?? '-'}</p>;
            },
        },
        {
            header: "Kode",
            accessorKey: "kode",
            enableColumnFilter: false,
            cell: ({row}) => {
              return <p className="m-0 p-0 text-center">{row.original?.kode ?? '-'}</p>;
            },
        },
        {
            header: "Nama Mata Kuliah",
            accessorKey: "nama_matkul",
            cell: ({row}) => {
              return <p className="m-0 p-0 text-center">{row.original?.nama_matkul ?? '-'}</p>;
            },
        },
        {
            header: "Kapasitas",
            accessorKey: "kapasitas",
            cell: ({row}) => {
              return <p className="m-0 p-0 text-center">{row.original?.kapasitas ?? '-'}</p>;
            },
        },
        {
            header: "Peserta",
            accessorKey: "peserta",
            cell: ({row}) => {
              return <p className="m-0 p-0 text-center">{row.original?.peserta ?? '-'}</p>;
            },
        },
        {
            header: "Dosen Pengajar",
            accessorKey: "dosen_pengajar",
            cell: ({row}) => {
              return <p className="m-0 p-0 text-center">{row.original?.dosen_pengajar ?? '-'}</p>;
            },
        },
        {
            header: "Action",
            enableSorting: false,
            cell: ({ row }) => {
                return (
                  <>
                    <div className="d-flex gap-2 justify-content-center align-items-center">
                      {/*//! action edit */}
                      {/* <Link href="subjects/edit/123"
                      className="bg-transparent border-0 text-black p-0">
                        <EditIcon color="#0AB39C" width='20' height='20'/>
                      </Link> */}
                      {/*//! action view ap */}
                      <Link href={`college-class/detail/${row.original.id}?tab=lecturer-teaching`}
                       className="bg-transparent border-0 text-black p-0">
                        <EyeAkademikIcon color='#2E3192' width='20' height='20'/>
                      </Link>
                      {/*//! action delete */}
                      {/* <Button
                        onClick={()=>{
                          setModalConfirmationState(()=>({
                            open:true,
                            state:'confirm',
                            message:'hapus data Mata Kuliah',
                            id:row.original.id,
                          }))
                        }}
                        className="bg-transparent border-0 text-black p-0">
                        <DeleteIcon color='#F06548' width='20' height='20'/>
                      </Button> */}
                    </div>
                  </>
                );
              },
      
        }    
    ]

    return {columns}

}


export default useColumnDefcollegeClass