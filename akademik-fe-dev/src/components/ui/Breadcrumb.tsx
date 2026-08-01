import React from "react";
import { Col, Row } from "reactstrap";

interface BreadCrumbProps {
  pageTitle: string;
  parentTitle: string;
}

const BreadCrumb: React.FC<BreadCrumbProps> = ({ pageTitle, parentTitle }) => {
  return (
    <React.Fragment>
      <Row>
        <Col xs={12}>
          <div className="page-title-box d-sm-flex align-items-center justify-content-between">
            <h4 className="mb-sm-0 fw-bolder" style={{ color: "#495057" }}>
              {pageTitle}
            </h4>

            <div className="page-title-right">
              <ol className="breadcrumb m-0">
                <li className="breadcrumb-item fw-normal">
                  <a href="#" className="text-decoration-none">
                    {parentTitle}
                  </a>
                </li>
                <li className="breadcrumb-item active">{pageTitle}</li>
              </ol>
            </div>
          </div>
        </Col>
      </Row>
    </React.Fragment>
  );
};

export default BreadCrumb;
