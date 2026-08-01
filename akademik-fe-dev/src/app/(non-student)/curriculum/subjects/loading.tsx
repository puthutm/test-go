import React from 'react'

import { Row,Col,Card,Spinner } from 'reactstrap'

function LoadingSubjects() {
  return (
        <Row>
            <Col>
                <Card className='p-3 rounded-3 bg-white d-flex mx-auto'>
                        <Spinner className='mx-auto'/>
                </Card>
            </Col>
        </Row>
  )
}

export default LoadingSubjects