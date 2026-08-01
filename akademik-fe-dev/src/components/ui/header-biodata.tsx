"use client";

import { Col, Row } from "reactstrap";

import styles from "@/styles/card-biodata-information.module.css";
import { ImageComponent } from "./image";
import { useGetProfile } from "@/services/api/sso/profile/use-get-profile";

export const HeaderBiodata: React.FC = () => {
  const { data: profile } = useGetProfile();
  return (
    <div className="pt-3 mb-3 mb-lg-3">
      <Row className="g-4 mt-1 gap-5 ps-3">
        <Col xs={3} md={2} lg={1} className=" mt-4">
          <ImageComponent
            src={profile?.data?.avatar as string}
            className="avatar-lg rounded-circle border border-white"
            alt="profile"
            width={50}
            height={50}
          />
        </Col>
        <Col xs={7} md={6} lg={10} className="px-0">
          <Row className="justify-content-between">
            <Col
              className="text-white d-flex flex-column"
              lg={6}
              style={{ gap: "12px" }}
            >
              {/* NAME */}
              <Row>
                <Col xs={3}>Nama</Col>
                <Col
                  xs={1}
                  className="d-flex justify-content-end p-0"
                  style={{ width: "fit-content" }}
                >
                  <span>:</span>
                </Col>
                <Col>{profile?.data?.name}</Col>
              </Row>
              {/* NIM */}
              <Row>
                <Col xs={3}>NIM</Col>
                <Col
                  xs={1}
                  className="d-flex justify-content-end p-0"
                  style={{ width: "fit-content" }}
                >
                  <span>:</span>
                </Col>
                <Col>Lorem ipsum dolor sit amet.</Col>
              </Row>
              {/* ANGKATAN */}
              <Row>
                <Col xs={3}>Angkatan</Col>
                <Col
                  xs={1}
                  className="d-flex justify-content-end p-0"
                  style={{ width: "fit-content" }}
                >
                  <span>:</span>
                </Col>
                <Col>Lorem ipsum dolor sit amet.</Col>
              </Row>
              {/* JALUR MASUK */}
              <Row>
                <Col xs={3}>Jalur Masuk</Col>
                <Col
                  xs={1}
                  className="d-flex justify-content-end p-0"
                  style={{ width: "fit-content" }}
                >
                  <span>:</span>
                </Col>
                <Col>Lorem ipsum dolor sit amet.</Col>
              </Row>
              {/* DOSEN PA */}
              <Row>
                <Col xs={3}>Dosen PA</Col>
                <Col
                  xs={1}
                  className="d-flex justify-content-end p-0"
                  style={{ width: "fit-content" }}
                >
                  <span>:</span>
                </Col>
                <Col>Lorem ipsum dolor sit amet.</Col>
              </Row>
            </Col>
            <Col lg={5} className="d-none d-lg-block">
              <div className={styles.card_header}>
                <p>Jurusan</p>
                <span>Aktif</span>
              </div>
              <div className={styles.card_info_wrapper}>
                <div className={styles.info}>
                  <div>
                    <p>4.00</p>
                    <span>IPS</span>
                  </div>
                  <div>
                    <p>12</p>
                    <span>SKSS</span>
                  </div>
                  <div>
                    <p>4</p>
                    <span>Semester</span>
                  </div>
                </div>
              </div>
              <div className={styles.card_info_wrapper}>
                <div className={styles.info}>
                  <div>
                    <p>4.00</p>
                    <span>IPK</span>
                  </div>
                  <div>
                    <p>60</p>
                    <span>SKSK</span>
                  </div>
                  <div>
                    <p>4</p>
                    <span>MK Ulang</span>
                  </div>
                </div>
              </div>
            </Col>
          </Row>
        </Col>
      </Row>
    </div>
  );
};
