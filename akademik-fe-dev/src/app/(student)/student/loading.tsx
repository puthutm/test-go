import { Card, CardBody, Col, Row, Spinner } from "reactstrap";

export default function Loading() {
  // You can add any UI inside Loading, including a Skeleton.
  return (
    <Row className="pt-4">
      <Col lg={3}>
        <Card
          className="placeholder-glow"
          style={{
            position: "sticky",
            top: 80,
            borderRadius: "8px",
            height: "200px",
          }}
          aria-hidden="true"
        >
          <CardBody className="d-flex justify-content-center align-items-center">
            {/* <SideTab tabs={tabs} /> */}
            <Spinner className="mx-auto" />
          </CardBody>
        </Card>
      </Col>
      <Col lg={9}>
        <Card style={{ borderRadius: "8px", height: "200px" }}>
          <CardBody className="d-flex justify-content-center align-items-center">
            {/* <SideTab tabs={tabs} /> */}
            <Spinner className="mx-auto" />
          </CardBody>
        </Card>
      </Col>
    </Row>
  );
}
