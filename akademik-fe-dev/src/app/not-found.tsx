import Link from "next/link";
import { Card, CardBody, Col, Container, Row } from "reactstrap";

export default function NotFound() {
  return (
    <div className="auth-page-wrapper auth-bg-cover py-5 d-flex justify-content-center align-items-center min-vh-100">
      <div className="bg-overlay"></div>
      <div className="auth-page-content overflow-hidden pt-lg-5">
        <Container>
          <Row className="justify-content-center">
            <Col xl={5}>
              <Card className="overflow-hidden">
                <CardBody className="p-4">
                  <div className="text-center">
                    <i className="ri-bard-line display-5 text-success"></i>
                    <h1 className="text-primary mb-4">Ups !</h1>
                    <h4 className="text-uppercase">
                      Maaf, halaman tidak ditemukan 😭
                    </h4>
                    <p className="text-muted mb-4">
                      {/* The page you are looking for not available! */}
                      Halaman yang kamu cari tidak tersedia!
                    </p>
                    <Link href="/" className="btn btn-success">
                      <i className="mdi mdi-home me-1"></i>Kembali ke Beranda
                    </Link>
                  </div>
                </CardBody>
              </Card>
            </Col>
          </Row>
        </Container>
      </div>
    </div>
  );
}
