import { Col, Row } from "reactstrap";

export const Footer = () => {
  return (
    <footer className={`footer ms-0 m-0`} style={{ fontSize: "13px" }}>
      <Row>
        <Col sm={6} style={{ color: "#545454" }}>
          {new Date().getFullYear()} © Copyright BPPTI Universitas Siber Asia
        </Col>
        <Col sm={6}>
          <div
            className="text-sm-end d-none d-sm-block"
            style={{ color: "#545454" }}
          >
            Design & Develop by BPPTI
          </div>
        </Col>
      </Row>
    </footer>
  );
};

export const FooterStudent = () => {
  return (
    <footer className="container-fluid  py-3 border-top border border-3 bg-white">
      <Row className="container contiainer-sm container-md  mx-auto  p-0">
        <Col sm={12} md={6} style={{ color: "#98A6AD" }}>
          {new Date().getFullYear()} © Copyright BPPTI Universitas Siber Asia
        </Col>
        <Col sm={12} md={6}>
          <div
            className="text-sm-end d-none d-sm-block"
            style={{ color: "#98A6AD" }}
          >
            Design & Develop by BPPTI
          </div>
        </Col>
      </Row>
    </footer>
  );
};
