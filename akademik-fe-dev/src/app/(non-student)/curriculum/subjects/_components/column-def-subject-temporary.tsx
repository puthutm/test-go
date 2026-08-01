

'use client' 
import { ColumnDef } from "@tanstack/react-table" 
    // import component 
//  import Link from "next/link" 
//  import { EditIcon } from "@/components/icons/edit" 
//  import { EyeAkademikIcon } from "@/components/icons/eye-akademik"

// import hook 
import { IDummyValueTable } from "./client-temporary"

interface iColumnsParams {     ():{ columns: ColumnDef<IDummyValueTable>[] } }

const useColumnDefSubjectTemprary : iColumnsParams = ()=>{


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
            header: "Kode",
            accessorKey: "kode",
            enableColumnFilter: false,
            cell: ({row}) => {
              return <p className="m-0 p-0 text-center">{row.original?.kode ?? '-'}</p>;
            },
        },  
        {
            header: "Kur",
            accessorKey: "kurikulum",
            enableColumnFilter: false,
            cell: ({row}) => {
              return <p className="m-0 p-0 text-center">{row.original?.kurikulum ?? '-'}</p>;
            },
        },
        {
            header: "Nama Mata Kuliah",
            accessorKey: "nama_matkul",
            enableColumnFilter: false,
            cell: ({row}) => {
              return <p className="m-0 p-0 text-center">{row.original?.nama_matkul ?? '-'}</p>;
            },
        },
        {
            header: "Prodi Pengampu",
            accessorKey: "prodi_pengampu",
            enableColumnFilter: false,
            cell: ({row}) => {
              return <p className="m-0 p-0 text-center">{row.original?.prodi_pengampu ?? '-'}</p>;
            },
        },
        {
            header: "SKS",
            accessorKey: "sks",
            enableColumnFilter: false,
            cell: ({row}) => {
              return <p className="m-0 p-0 text-center">{row.original?.sks ?? '-'}</p>;
            },
        },
        {
            header: "Jenis MK",
            accessorKey: "jenis_matkul",
            enableColumnFilter: false,
            cell: ({row}) => {
              return <p className="m-0 p-0 text-center">{row.original?.jenis_matkul ?? '-'}</p>;
            },
        },
    ]

    return {columns}

}


export default useColumnDefSubjectTemprary