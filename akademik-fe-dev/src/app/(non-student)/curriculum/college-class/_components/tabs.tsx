'use client'
// import { Button } from "reactstrap"

interface ITabs {
    id:string,
    title:string,
}
function TabsSectionLecture() {

  const listTabs : ITabs[] =[
        {
            id:'1',
            title:"All",
        },
        {
            id:'2',
            title:"S1 PJJ Informatika",
        },
        {
            id:'3',
            title:"S1 PJJ Sistem Informasi",
        },
        {
            id:'4',
            title:"S1 PJJ Management",
        },
        {
            id:'5',
            title:"S1 PJJ Akutansi",
        },
        {
            id:'6',
            title:"S1 PJJ Komunikasi",
        },
   ]

  return (
    <section className="position-relative  mt-3 mb-2 d-flex flex-wrap d-flex  ">
        {
            listTabs.map((tab:ITabs)=>{
                return (
                <button 
                key={tab.id}
                    disabled={tab.id !== '1' ? true : false}
                    className={`btn-tabs-lecturer-subject  flex-grow-1 px-2 fw-medium 
                    ${tab.id === '1' ? 'btn-tabs-lecturer-subject-active' : "btn-tabs-lecturer-subject-not-active"}    
                        `}
                    color='transparent'
                    >
                        {tab.title}
                </button>
                )
            })
        }
    </section>
  )
}

export default TabsSectionLecture
