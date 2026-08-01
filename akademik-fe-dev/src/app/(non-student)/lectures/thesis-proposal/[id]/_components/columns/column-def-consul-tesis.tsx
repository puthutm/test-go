

'use client' 
import { ColumnDef } from "@tanstack/react-table" 
    // import component 
//  import Link from "next/link" 
//  import { EditIcon } from "@/components/icons/edit" 
//  import { EyeAkademikIcon } from "@/components/icons/eye-akademik"

// import hook 
import { IDummyValueTable } from "../tabs/consultation"

interface iColumnsParams {     ():{ columns: ColumnDef<IDummyValueTable>[] } }

const useColumnDefConsulThesis : iColumnsParams = ()=>{


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
            header: "Waktu Bimbingan",
            accessorKey: "waktu",
            enableColumnFilter: false,
            cell: ({row}) => {
              return <p className="m-0 p-0 text-center">{row.original?.waktu ?? '-'}</p>;
            },
        },  
        {
            header: "Dosen Pembimbing",
            accessorKey: "dosen_pembimbing",
            enableColumnFilter: false,
            cell: ({row}) => {
              return <p className="m-0 p-0 text-center">{row.original?.dosen_pembimbing ?? '-'}</p>;
            },
        },   
        {
            header: "Feedback Pembimbing",
            accessorKey: "feedback",
            enableColumnFilter: false,
            cell: ({row}) => {
              return <p className="m-0 p-0 text-center">{row.original?.feedback ?? '-'}</p>;
            },
        },  
        {
            header: "Status",
            accessorKey: "status",
            enableColumnFilter: false,
            cell: ({row}) => {
              return <p className="m-0 p-0 text-center p-1  rounded-3 "
              style={{
                fontSize:'11px',
                color: row.original.status === 'Disetujui' ? '#6CBE40' : '#F06548',
                background: row.original.status === 'Disetujui' ? '#6CBE401A' : '#F065481A',
              }}
              >{row.original?.status ?? '-'}</p>;
            },
        },  
        // {
        //     header: "Action",
        //     enableSorting: false,
        //     cell: ({ row }) => {
        //         return (
        //           <>
        //             <div className="d-flex gap-2 justify-content-center align-items-center">
        //               {/*//! action edit */}
        //               {/* <Link href="subjects/edit/123"
        //               className="bg-transparent border-0 text-black p-0">
        //                 <EditIcon color="#0AB39C" width='20' height='20'/>
        //               </Link> */}
        //               {/*//! action view ap */}
        //               <Link href={`thesis-proposal/${row.original.id}?tab=final-grade`}
        //                className="bg-transparent border-0 text-black p-0">
        //                 <EyeAkademikIcon color='#2E3192' width='20' height='20'/>
        //               </Link>
        //               {/*//! action delete */}
        //               {/* <Button
        //                 onClick={()=>{
        //                   setModalConfirmationState(()=>({
        //                     open:true,
        //                     state:'confirm',
        //                     message:'hapus data Mata Kuliah',
        //                     id:row.original.id,
        //                   }))
        //                 }}
        //                 className="bg-transparent border-0 text-black p-0">
        //                 <DeleteIcon color='#F06548' width='20' height='20'/>
        //               </Button> */}
        //             </div>
        //           </>
        //         );
        //       },
      
        // }    
       
    ]

    return {columns}

}


export default useColumnDefConsulThesis