import React from 'react'
import { Metadata } from 'next'

import { Row,Col } from 'reactstrap'
import SidebarTabs from './components/sidebar-tabs'

export const metadata : Metadata = {
    title:'Kelas Kuliah'
}
function layout({
    children,
}:{
    children: React.ReactNode
}) {
  return (
    <section className="position-relative w-full mt-3">
        <Row className='gap-3  m-0 p-0'>
            {/*//! left  */}
            <Col md={3} className='g-0  m-0 p-0 '>
                <section className="w-100">
                    <SidebarTabs/>
                </section>
            </Col>
            {/*//! rigth */}
            <Col  className='g-0  m-0 p-0'>
                    {
                        children
                    }
            </Col>
        </Row>
      
    </section>
  )
}

export default layout