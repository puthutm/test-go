import React from "react";

export default function ThesisProposalStudentInfoSlot() {
  const infoData = [
    {
      title: "Nama",
      value: "Muhammad Farhan",
    },
    {
      title: "NIM",
      value: "230401020086",
    },
    {
      title: "Program Studi",
      value: "PJJ Informatika ",
    },

    {
      title: "Angkatan",
      value: "Ganjil 2020/2021",
    },
    {
      title: "Jalur Masuk",
      value: "Reguler",
    },
  ];

  return (
    <div className="row row-cols-2 row-gap-3">
      {infoData.map((data, index) => (
        <div key={index} className="col">
          <div className="fw-semibold">{data.title}</div>
          <div>{data.value}</div>
        </div>
      ))}
    </div>
  );
}
