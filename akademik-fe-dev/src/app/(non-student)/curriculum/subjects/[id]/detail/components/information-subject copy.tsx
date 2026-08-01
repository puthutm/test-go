'use client'
import React from 'react'
import { Row,Col } from 'reactstrap'

function InformationSubject() {
  return (
    <div className="p-3 rounded-3" style={{background:'#FAFCFF'}}> 
        <Row className='row-gap-3'>
            {/*//! code subject */}
            <Col sm={6}>
                <Row>
                    <Col sm={12}>
                        <h3 className="m-0 p-0 fw-semibold fs-6" style={{color:'#3A3A3A'}}>Kode Mata Kuliah</h3>
                        <p className="m-0 p-0 " style={{color:'#3A3A3A'}}>text here</p>
                    </Col>
                    <Col sm={12}>
                    </Col>
                </Row>
            </Col>
            {/*//! name subject */}
            <Col sm={6}>
                <Row>
                    <Col sm={12}>
                        <h3 className="m-0 p-0 fw-semibold fs-6" style={{color:'#3A3A3A'}}>Mata Kuliah</h3>
                        <p className="m-0 p-0 " style={{color:'#3A3A3A'}}>text here</p>
                    </Col>
                    <Col sm={12}>
                    </Col>
                </Row>
            </Col>
            {/*//! year curriculum */}
            <Col sm={6}>
                <Row>
                    <Col sm={12}>
                        <h3 className="m-0 p-0 fw-semibold fs-6" style={{color:'#3A3A3A'}}>Tahun Kurikulum</h3>
                        <p className="m-0 p-0 " style={{color:'#3A3A3A'}}>text here</p>
                    </Col>
                    <Col sm={12}>
                    </Col>
                </Row>
            </Col>
            {/*//! sks */}
            <Col sm={6}>
                <Row>
                    <Col sm={12}>
                        <h3 className="m-0 p-0 fw-semibold fs-6" style={{color:'#3A3A3A'}}>SKS</h3>
                        <p className="m-0 p-0 " style={{color:'#3A3A3A'}}>text here</p>
                    </Col>
                    <Col sm={12}>
                    </Col>
                </Row>
            </Col>
        </Row>
    </div>
  )
}

export default InformationSubject