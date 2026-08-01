import React from "react";
import { Col, Row, Table } from "reactstrap";

const ThesisTabContent = () => {
  const data = {
    title: "Pengembangan UI/UX Sistem Akademik Berbasis Mobile",
    details: {
      mentors: [
        {
          name: "Dr. Agus Saputra, M. Kom",
          type: "Pembimbing 1",
        },
        {
          name: "Dr. Rina Suryani, M. Kom",
          type: "Pembimbing 2",
        },
      ],
      topics: [
        {
          title: "Bab 1",
          name: "Pendahuluan",
          status: "approved",
        },
        {
          title: "Bab 2",
          name: "Kajian Pustaka",
          status: "on progress",
        },
        {
          title: "Bab 3",
          name: "Metode Penelitian",
          status: "pending",
        },
        {
          title: "Bab 4",
          name: "Hasil dan Pembahasan",
          status: "pending",
        },
        {
          title: "Bab 5",
          name: "Kesimpulan dan Saran",
          status: "pending",
        },
        {
          title: "Sidang Tugas Akhir",
          name: "",
          status: "pending",
        },
      ],
    },
  };

  return (
    <div>
         <div className="border-bottom border-3 mb-2">
        <h5 className="fw-semibold">Proposal Tugas Akhir</h5>
      </div>
      <Row className="mt-3">
        <Col sm={4}>
          {data.details.mentors.map((mentor, index) => (
            <div key={index} className="">
              <p>{mentor.type}</p>
              <p className="fw-semibold">{mentor.name}</p>
            </div>
          ))}
        </Col>
        <Col sm={8} className="table-responsive table-card">
          <Table striped className="table-bordered">
            <thead>
              <tr>
                <th>Topik</th>
                <th>Status</th>
                <th>Aksi</th>
              </tr>
            </thead>

            <tbody>
              {data.details.topics.map((topic, index) => (
                <tr key={index}>
                  <td>
                    <p>{topic.title}</p>
                    <p>{topic.name}</p>
                  </td>
                  <td>
                    <p
                      className={`text-capitalize badge ${
                        topic.status === "approved"
                          ? "bg-success-subtle text-success"
                          : topic.status === "on progress"
                          ? "bg-info-subtle text-info"
                          : "bg-warning-subtle text-warning"
                      }`}
                    >
                      {topic.status}
                    </p>
                  </td>
                  <td>
                    <button className="btn btn-primary">Detail</button>
                  </td>
                </tr>
              ))}
            </tbody>
          </Table>
        </Col>
      </Row>
    </div>
  );
};

export default ThesisTabContent;
