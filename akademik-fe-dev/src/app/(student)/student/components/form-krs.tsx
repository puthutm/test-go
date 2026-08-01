"use client";

import { ClockIcon } from "@/components/icons/clock";
import { WarningIcon } from "@/components/icons/warning";
import styles from "@/styles/krs.module.css";
import { useEffect, useState } from "react";
import { Button } from "reactstrap";

interface CourseWithGrade {
  kode: string;
  mataKuliah: string;
  jadwal: string;
  dosenPengampu: string;
  kelas: string;
  sks: number | string;
  kuota: string;
  nilai: string;
  remidi?: boolean;
  pilihMK: boolean;
}

interface CourseWithGradeData {
  semester: string;
  courses: CourseWithGrade[];
}

const useCountdown = (targetDate: string | Date) => {
  // State to store countdown values
  const [countdownTime, setCountdownTime] = useState({
    hours: 0,
    minutes: 0,
    seconds: 0,
    isExpired: false,
  });

  useEffect(() => {
    // Convert targetDate to milliseconds
    const targetTime =
      typeof targetDate === "string"
        ? new Date(targetDate).getTime()
        : targetDate.getTime();

    // Function to calculate the time difference
    const calculateTimeLeft = () => {
      const now = new Date().getTime();
      const difference = targetTime - now;

      // If countdown expired
      if (difference <= 0) {
        return {
          hours: 0,
          minutes: 0,
          seconds: 0,
          isExpired: true,
        };
      }

      // Calculate hours, minutes, seconds
      const hours = Math.floor(difference / (1000 * 60 * 60));
      const minutes = Math.floor((difference % (1000 * 60 * 60)) / (1000 * 60));
      const seconds = Math.floor((difference % (1000 * 60)) / 1000);

      return {
        hours,
        minutes,
        seconds,
        isExpired: false,
      };
    };

    // Initial calculation
    setCountdownTime(calculateTimeLeft());

    // Update countdown every second
    const timer = setInterval(() => {
      setCountdownTime(calculateTimeLeft());
    }, 1000);

    // Clean up the interval on unmount
    return () => clearInterval(timer);
  }, [targetDate]);

  return countdownTime;
};

export default function FormKrs() {
  const courseDataWithGrade: CourseWithGradeData[] = [
    {
      semester: "Semester 3",
      courses: [
        {
          kode: "KPIS502",
          mataKuliah: "Kalkulus",
          nilai: "D",
          jadwal: "Sen, 10.00-14.00 WIB",
          dosenPengampu: "Vika febri muliati, S.KOM, M.Kom",
          kelas: "402",
          sks: 3,
          kuota: "15/50",
          pilihMK: true,
          remidi: true,
        },
        {
          kode: "Text",
          mataKuliah: "Cloud Computing",
          nilai: "B-",
          jadwal: "Text",
          dosenPengampu: "Text",
          kelas: "Text",
          sks: "Text",
          kuota: "Text",
          pilihMK: false,
        },
      ],
    },
    {
      semester: "Semester 1",
      courses: [
        {
          kode: "KPIS502",
          mataKuliah: "Kalkulus",
          nilai: "D",
          jadwal: "Sen, 10.00-14.00 WIB",
          dosenPengampu: "Vika febri muliati, S.KOM, M.Kom",
          kelas: "402",
          sks: 3,
          kuota: "15/50",
          pilihMK: true,
        },
        {
          kode: "Text",
          mataKuliah: "Cloud Computing",
          nilai: "B-",
          jadwal: "Text",
          dosenPengampu: "Text",
          kelas: "Text",
          sks: "Text",
          kuota: "Text",
          pilihMK: false,
        },
      ],
    },
  ];
  // const coursesData = [
  //   {
  //     semester: "Semester 5",
  //     courses: [
  //       {
  //         kode: "KPIS502",
  //         mataKuliah: "Teori Perilaku Organisasi",
  //         jadwal: "Sen, 10.00-14.00 WIB",
  //         dosenPengampu: "Vika febri muliati, S.KOM, M.Kom",
  //         kelas: "402",
  //         sks: 3,
  //         kuota: "15/50",
  //         pilihMK: true,
  //       },
  //       {
  //         kode: "Text",
  //         mataKuliah: "Cloud Computing",
  //         jadwal: "Text",
  //         dosenPengampu: "Text",
  //         kelas: "Text",
  //         sks: "Text",
  //         kuota: "Text",
  //         pilihMK: false,
  //       },
  //       // Additional rows for Semester 5 with placeholder "Text" values
  //       {
  //         kode: "Text",
  //         mataKuliah: "Text",
  //         jadwal: "Text",
  //         dosenPengampu: "Text",
  //         kelas: "Text",
  //         sks: "Text",
  //         kuota: "Text",
  //         pilihMK: false,
  //       },
  //       {
  //         kode: "Text",
  //         mataKuliah: "Text",
  //         jadwal: "Text",
  //         dosenPengampu: "Text",
  //         kelas: "Text",
  //         sks: "Text",
  //         kuota: "Text",
  //         pilihMK: false,
  //       },
  //       {
  //         kode: "Text",
  //         mataKuliah: "Text",
  //         jadwal: "Text",
  //         dosenPengampu: "Text",
  //         kelas: "Text",
  //         sks: "Text",
  //         kuota: "Text",
  //         pilihMK: false,
  //       },
  //       {
  //         kode: "Text",
  //         mataKuliah: "Text",
  //         jadwal: "Text",
  //         dosenPengampu: "Text",
  //         kelas: "Text",
  //         sks: "Text",
  //         kuota: "Text",
  //         pilihMK: false,
  //       },
  //       {
  //         kode: "Text",
  //         mataKuliah: "Text",
  //         jadwal: "Text",
  //         dosenPengampu: "Text",
  //         kelas: "Text",
  //         sks: "Text",
  //         kuota: "Text",
  //         pilihMK: false,
  //       },
  //       {
  //         kode: "Text",
  //         mataKuliah: "Text",
  //         jadwal: "Text",
  //         dosenPengampu: "Text",
  //         kelas: "Text",
  //         sks: "Text",
  //         kuota: "Text",
  //         pilihMK: false,
  //       },
  //     ],
  //   },
  //   {
  //     semester: "Semester 7",
  //     courses: [
  //       {
  //         kode: "KPIS502",
  //         mataKuliah: "Kalkulus",
  //         jadwal: "Sen, 10.00-14.00 WIB",
  //         dosenPengampu: "Vika febri muliati, S.KOM, M.Kom",
  //         kelas: "402",
  //         sks: 3,
  //         kuota: "15/50",
  //         pilihMK: true,
  //       },
  //       {
  //         kode: "Text",
  //         mataKuliah: "Cloud Computing",
  //         jadwal: "Text",
  //         dosenPengampu: "Text",
  //         kelas: "Text",
  //         sks: "Text",
  //         kuota: "Text",
  //         pilihMK: false,
  //       },
  //       // Additional rows for Semester 7 with placeholder "Text" values
  //       {
  //         kode: "Text",
  //         mataKuliah: "Text",
  //         jadwal: "Text",
  //         dosenPengampu: "Text",
  //         kelas: "Text",
  //         sks: "Text",
  //         kuota: "Text",
  //         pilihMK: false,
  //       },
  //       {
  //         kode: "Text",
  //         mataKuliah: "Text",
  //         jadwal: "Text",
  //         dosenPengampu: "Text",
  //         kelas: "Text",
  //         sks: "Text",
  //         kuota: "Text",
  //         pilihMK: false,
  //       },
  //       {
  //         kode: "Text",
  //         mataKuliah: "Text",
  //         jadwal: "Text",
  //         dosenPengampu: "Text",
  //         kelas: "Text",
  //         sks: "Text",
  //         kuota: "Text",
  //         pilihMK: false,
  //       },
  //       {
  //         kode: "Text",
  //         mataKuliah: "Text",
  //         jadwal: "Text",
  //         dosenPengampu: "Text",
  //         kelas: "Text",
  //         sks: "Text",
  //         kuota: "Text",
  //         pilihMK: false,
  //       },
  //       {
  //         kode: "Text",
  //         mataKuliah: "Text",
  //         jadwal: "Text",
  //         dosenPengampu: "Text",
  //         kelas: "Text",
  //         sks: "Text",
  //         kuota: "Text",
  //         pilihMK: false,
  //       },
  //       {
  //         kode: "Text",
  //         mataKuliah: "Text",
  //         jadwal: "Text",
  //         dosenPengampu: "Text",
  //         kelas: "Text",
  //         sks: "Text",
  //         kuota: "Text",
  //         pilihMK: false,
  //       },
  //     ],
  //   },
  // ];
  const [isOpen] = useState(true);
  const [isActive, setIsActive] = useState<number>(0);
  const [isKrs, setIsKrs] = useState(0);
  const [course, setCourse] = useState([
    {
      semester: "Semester 5",
      courses: [
        {
          kode: "KPIS502",
          mataKuliah: "Teori Perilaku Organisasi",
          jadwal: "Sen, 10.00-14.00 WIB",
          dosenPengampu: "Vika febri muliati, S.KOM, M.Kom",
          kelas: "402",
          sks: 3,
          kuota: "15/50",
          pilihMK: true,
        },
        {
          kode: "Text",
          mataKuliah: "Cloud Computing",
          jadwal: "Text",
          dosenPengampu: "Text",
          kelas: "Text",
          sks: "Text",
          kuota: "Text",
          pilihMK: false,
        },
        // Additional rows for Semester 5 with placeholder "Text" values
        {
          kode: "Text",
          mataKuliah: "Text",
          jadwal: "Text",
          dosenPengampu: "Text",
          kelas: "Text",
          sks: "Text",
          kuota: "Text",
          pilihMK: false,
        },
        {
          kode: "Text",
          mataKuliah: "Text",
          jadwal: "Text",
          dosenPengampu: "Text",
          kelas: "Text",
          sks: "Text",
          kuota: "Text",
          pilihMK: false,
        },
        {
          kode: "Text",
          mataKuliah: "Text",
          jadwal: "Text",
          dosenPengampu: "Text",
          kelas: "Text",
          sks: "Text",
          kuota: "Text",
          pilihMK: false,
        },
        {
          kode: "Text",
          mataKuliah: "Text",
          jadwal: "Text",
          dosenPengampu: "Text",
          kelas: "Text",
          sks: "Text",
          kuota: "Text",
          pilihMK: false,
        },
        {
          kode: "Text",
          mataKuliah: "Text",
          jadwal: "Text",
          dosenPengampu: "Text",
          kelas: "Text",
          sks: "Text",
          kuota: "Text",
          pilihMK: false,
        },
        {
          kode: "Text",
          mataKuliah: "Text",
          jadwal: "Text",
          dosenPengampu: "Text",
          kelas: "Text",
          sks: "Text",
          kuota: "Text",
          pilihMK: false,
        },
      ],
    },
    {
      semester: "Semester 7",
      courses: [
        {
          kode: "KPIS502",
          mataKuliah: "Kalkulus",
          jadwal: "Sen, 10.00-14.00 WIB",
          dosenPengampu: "Vika febri muliati, S.KOM, M.Kom",
          kelas: "402",
          sks: 3,
          kuota: "15/50",
          pilihMK: true,
        },
        {
          kode: "Text",
          mataKuliah: "Cloud Computing",
          jadwal: "Text",
          dosenPengampu: "Text",
          kelas: "Text",
          sks: "Text",
          kuota: "Text",
          pilihMK: false,
        },
        // Additional rows for Semester 7 with placeholder "Text" values
        {
          kode: "Text",
          mataKuliah: "Text",
          jadwal: "Text",
          dosenPengampu: "Text",
          kelas: "Text",
          sks: "Text",
          kuota: "Text",
          pilihMK: false,
        },
        {
          kode: "Text",
          mataKuliah: "Text",
          jadwal: "Text",
          dosenPengampu: "Text",
          kelas: "Text",
          sks: "Text",
          kuota: "Text",
          pilihMK: false,
        },
        {
          kode: "Text",
          mataKuliah: "Text",
          jadwal: "Text",
          dosenPengampu: "Text",
          kelas: "Text",
          sks: "Text",
          kuota: "Text",
          pilihMK: false,
        },
        {
          kode: "Text",
          mataKuliah: "Text",
          jadwal: "Text",
          dosenPengampu: "Text",
          kelas: "Text",
          sks: "Text",
          kuota: "Text",
          pilihMK: false,
        },
        {
          kode: "Text",
          mataKuliah: "Text",
          jadwal: "Text",
          dosenPengampu: "Text",
          kelas: "Text",
          sks: "Text",
          kuota: "Text",
          pilihMK: false,
        },
        {
          kode: "Text",
          mataKuliah: "Text",
          jadwal: "Text",
          dosenPengampu: "Text",
          kelas: "Text",
          sks: "Text",
          kuota: "Text",
          pilihMK: false,
        },
      ],
    },
  ]);
  const [courseWithNilai, setCourseWithNilai] = useState(courseDataWithGrade);
  // ubah state dibawah ini menjadi false untuk melihat tampilan KRS paket
  const [isPaket] = useState(false);

  const buttons = ["Pilih Kelas", "KRS Tersimpan"];
  // Set target date - March 18, 2025, 23:59 WIB
  const targetDate = "2025-03-18T23:59:00+07:00";

  // Use the countdown hook
  const countdown = useCountdown(targetDate);

  // Format countdown time with leading zeros (e.g., 05:08:02)
  const formatTime = (time: number): string => {
    return time.toString().padStart(2, "0");
  };

  // Format the countdown display
  const formattedCountdown = countdown.isExpired
    ? "00 : 00 : 00"
    : `${formatTime(countdown.hours)} : ${formatTime(
        countdown.minutes
      )} : ${formatTime(countdown.seconds)}`;

  const handleButtonClick = (index: number) => {
    setIsActive(index);
  };
  const handleIsKrs = (label: string) => {
    if (label === "Pilih Kelas") {
      setIsKrs(0);
    } else {
      setIsKrs(1);
    }
  };
  const handleCombineFunction = async (index: number, label: string) => {
    try {
      handleButtonClick(index);
      handleIsKrs(label);
    } catch (error) {
      console.log(error);
    }
  };
  const handleCheckboxChange = (semesterIndex: number, courseIndex: number) => {
    const updatedCoursesData = [...course];
    updatedCoursesData[semesterIndex].courses[courseIndex].pilihMK =
      !updatedCoursesData[semesterIndex].courses[courseIndex].pilihMK;
    setCourse(updatedCoursesData);
  };
  const handleCourseWithNilaiCheckboxChange = (
    semesterIndex: number,
    courseIndex: number
  ) => {
    const updatedCourseWithNilai = [...courseWithNilai];
    updatedCourseWithNilai[semesterIndex].courses[courseIndex].pilihMK =
      !updatedCourseWithNilai[semesterIndex].courses[courseIndex].pilihMK;
    setCourseWithNilai(updatedCourseWithNilai);
  };
  const getSelectedCourses = () => {
    const selected = [];
    for (const semester of course) {
      for (const course of semester.courses) {
        if (course.pilihMK) {
          selected.push({
            ...course,
            semester: semester.semester,
          });
        }
      }
    }

    return selected;
  };

  return (
    <>
      <div className="d-flex border-2 border-bottom justify-content-between align-items-center">
        <h5 className="card-title py-3 mb-1" style={{ fontWeight: "500" }}>
          Pengisian KRS
        </h5>
      </div>
      <div
        className={`alert alert-warning alert-border-left ${styles.warning_container} ${styles.m_20} d-flex justify-content-between align-items-center`}
        role="alert"
      >
        <p className={`${styles.text_12}`}>
          Selesaikan pengisian KRS Anda sebelum 28 Februari 2025 Pukul 23.59
          WIB.
        </p>
        <button
          type="button"
          className={`btn btn-danger ${styles.clock_btn} waves-effect waves-light d-flex align-items-center justidfy-content-between`}
        >
          <ClockIcon />
          {/* show the time form countdown function*/}
          <p>{formattedCountdown}</p>
        </button>
      </div>
      {/* end alert kuning */}
      {/* start alert merah */}
      <div
        className={`alert alert-danger alert-border-left  fade show`}
        role="alert"
      >
        <div className="d-flex gap-2">
          <WarningIcon />
          <p className={styles.text_12}>
            Nilai mata kuliah berikut belum memenuhi standar. Mata kuliah
            tersebut bisa diulang di semester ini.
          </p>
        </div>
        <p className={styles.text_12}>KPIS502 - Kalkulus / Ganjil</p>
      </div>
      {/* end alert merah */}
      {/* start alert abu */}
      <div
        className={`alert alert-light container text-dark ${styles.p_alert_light}`}
        role="alert"
      >
        <div className="row">
          <p className="col">Semester saat ini : 5</p>
          <p className="col">Total SKS yang dapat diambil : 20</p>
        </div>
      </div>
      {isOpen ? (
        <>
          <div>
            {isPaket ? (
              <>
                <div className="table-responsive">
                  <table
                    className={`table table-striped table-bordered text-center ${styles.mb_0} `}
                  >
                    <thead className={`table-light ${styles.text_14}`}>
                      <tr>
                        <th scope="col" className={`${styles.p_table}`}>
                          Kode
                        </th>
                        <th scope="col" className={`${styles.p_table}`}>
                          Mata Kuliah
                        </th>
                        <th scope="col" className={`${styles.p_table}`}>
                          Jadwal
                        </th>
                        <th scope="col" className={`${styles.p_table}`}>
                          Dosen Pengampu
                        </th>
                        <th scope="col" className={`${styles.p_table}`}>
                          Kelas
                        </th>
                        <th scope="col" className={`${styles.p_table}`}>
                          SKS
                        </th>
                        <th scope="col" className={`${styles.p_table}`}>
                          Kuota
                        </th>
                        <th scope="col" className={`${styles.p_mk}`}>
                          Pilih MK
                        </th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr>
                        <td>KPIS502</td>
                        <td>Teori Perilaku Organisasi</td>
                        <td>Sen, 10.00-14.00 WIB</td>
                        <td>Vika febri muliati, S.KOM, M.Kom</td>
                        <td>402</td>
                        <td>3</td>
                        <td>15/50</td>
                        <td>
                          <input type="checkbox" />
                        </td>
                      </tr>
                      <tr>
                        <td>Text</td>
                        <td>Text</td>
                        <td>Text</td>
                        <td>Text</td>
                        <td>Text</td>
                        <td>Text</td>
                        <td>Text</td>
                        <td>
                          <input type="checkbox" />
                        </td>
                      </tr>
                      <tr>
                        <td>Text</td>
                        <td>Text</td>
                        <td>Text</td>
                        <td>Text</td>
                        <td>Text</td>
                        <td>Text</td>
                        <td>Text</td>
                        <td>
                          <input type="checkbox" />
                        </td>
                      </tr>
                      <tr>
                        <td>Text</td>
                        <td>Text</td>
                        <td>Text</td>
                        <td>Text</td>
                        <td>Text</td>
                        <td>Text</td>
                        <td>Text</td>
                        <td>
                          <input type="checkbox" />
                        </td>
                      </tr>
                      <tr>
                        <td>Text</td>
                        <td>Text</td>
                        <td>Text</td>
                        <td>Text</td>
                        <td>Text</td>
                        <td>Text</td>
                        <td>Text</td>
                        <td>
                          <input type="checkbox" />
                        </td>
                      </tr>
                    </tbody>
                  </table>
                  <div
                    className={`h5 fw-semibold py-3 px-2 m-0 row d-flex justify-content-between`}
                    style={{ backgroundColor: "#FFE91D33" }}
                  >
                    <p className="col">Total SKS</p>
                    <p className="col text-center">3/20</p>
                  </div>

                  <div
                    className={`alert border-0 text-center ${styles.alert_gray} ${styles.m_20}`}
                    role="alert"
                  >
                    <p> + Tambah Kelas</p>
                  </div>

                  <div className="d-flex justify-content-between mt-3 gap-3">
                    <button
                      onClick={() => {}}
                      className="bg-transparent text-primary rounded px-3"
                      type="button"
                      style={{ border: "1px solid #10487A" }}
                    >
                      <span>Batal</span>
                    </button>
                    <Button
                      color="primary"
                      className="d-flex flex-grow-0 justify-content-center align-items-center"
                    >
                      <span>Update</span>
                    </Button>
                  </div>
                </div>
              </>
            ) : (
              <>
                <div>
                  <div className="d-flex">
                    {buttons.map((label, index) => (
                      <button
                        className={`${styles.btn_table_krs} ${
                          isActive === index ? styles.active : ""
                        }`}
                        onClick={() => handleCombineFunction(index, label)}
                        key={index}
                      >
                        {label}
                      </button>
                    ))}
                  </div>
                  {isKrs === 0 ? (
                    <div>
                      {course.map((semester, semesterIndex) => (
                        <div
                          key={semester.semester}
                          className="table-responsive"
                        >
                          <p className="p-3">{semester.semester}</p>
                          <table
                            className={`table table-striped table-bordered text-center ${styles.mb_0} `}
                          >
                            <thead className={`table-light ${styles.text_14}`}>
                              <tr>
                                <th scope="col" className={`${styles.p_table}`}>
                                  Kode
                                </th>
                                <th scope="col" className={`${styles.p_table}`}>
                                  Mata Kuliah
                                </th>
                                <th scope="col" className={`${styles.p_table}`}>
                                  Jadwal
                                </th>
                                <th scope="col" className={`${styles.p_table}`}>
                                  Dosen Pengampu
                                </th>
                                <th scope="col" className={`${styles.p_table}`}>
                                  Kelas
                                </th>
                                <th scope="col" className={`${styles.p_table}`}>
                                  SKS
                                </th>
                                <th scope="col" className={`${styles.p_table}`}>
                                  Kuota
                                </th>
                                <th scope="col" className={`${styles.p_mk}`}>
                                  Pilih MK
                                </th>
                              </tr>
                            </thead>
                            <tbody>
                              {semester.courses.map((course, courseIndex) => (
                                <tr key={courseIndex}>
                                  <td>{course.kode}</td>
                                  <td>{course.mataKuliah}</td>
                                  <td>{course.jadwal}</td>
                                  <td>{course.dosenPengampu}</td>
                                  <td>{course.kelas}</td>
                                  <td>{course.sks}</td>
                                  <td>{course.kuota}</td>
                                  <td>
                                    <input
                                      type="checkbox"
                                      checked={course.pilihMK}
                                      onChange={() =>
                                        handleCheckboxChange(
                                          semesterIndex,
                                          courseIndex
                                        )
                                      }
                                    />
                                  </td>
                                </tr>
                              ))}
                            </tbody>
                          </table>
                        </div>
                      ))}
                      {courseWithNilai.map((semester, semesterIndex) => (
                        <div
                          key={semester.semester}
                          className="table-responsive"
                        >
                          <p className="p-3">{semester.semester}</p>
                          <table
                            className={`table table-striped table-bordered text-center ${styles.mb_0} `}
                          >
                            <thead className={`table-light ${styles.text_14}`}>
                              <tr>
                                <th scope="col" className={`${styles.p_table}`}>
                                  Kode
                                </th>
                                <th scope="col" className={`${styles.p_table}`}>
                                  Mata Kuliah
                                </th>
                                <th scope="col" className={`${styles.p_table}`}>
                                  Nilai
                                </th>
                                <th scope="col" className={`${styles.p_table}`}>
                                  Jadwal
                                </th>
                                <th scope="col" className={`${styles.p_table}`}>
                                  Dosen Pengampu
                                </th>
                                <th scope="col" className={`${styles.p_table}`}>
                                  Kelas
                                </th>
                                <th scope="col" className={`${styles.p_table}`}>
                                  SKS
                                </th>
                                <th scope="col" className={`${styles.p_table}`}>
                                  Kuota
                                </th>
                                <th scope="col" className={`${styles.p_mk}`}>
                                  Pilih MK
                                </th>
                              </tr>
                            </thead>
                            <tbody>
                              {semester.courses.map((course, courseIndex) => (
                                <tr
                                  key={courseIndex}
                                  className={
                                    course.remidi ? "table-danger" : ""
                                  }
                                >
                                  <td>{course.kode}</td>
                                  <td>{course.mataKuliah}</td>
                                  {course.nilai ? (
                                    <td>{course.nilai}</td>
                                  ) : null}
                                  <td>{course.jadwal}</td>
                                  <td>{course.dosenPengampu}</td>
                                  <td>{course.kelas}</td>
                                  <td>{course.sks}</td>
                                  <td>{course.kuota}</td>
                                  <td>
                                    <input
                                      type="checkbox"
                                      checked={course.pilihMK}
                                      onChange={() =>
                                        handleCourseWithNilaiCheckboxChange(
                                          semesterIndex,
                                          courseIndex
                                        )
                                      }
                                    />
                                  </td>
                                </tr>
                              ))}
                            </tbody>
                          </table>
                        </div>
                      ))}
                      <div
                        className={`h5 fw-semibold py-3 px-2 m-0 row d-flex justify-content-between`}
                        style={{ backgroundColor: "#FFE91D33" }}
                      >
                        <p className="col">Total SKS</p>
                        <p className="col text-center">3/20</p>
                      </div>
                      <div className={`d-grid gap-2 ${styles.mt_20}`}>
                        <button className="btn btn-primary" type="button">
                          Validasi Krs
                        </button>
                      </div>
                    </div>
                  ) : (
                    <div className="table-responsive">
                      <p className="p-3">KRS Tersimpan</p>
                      <table
                        className={`table table-striped table-bordered text-center ${styles.mb_0} `}
                      >
                        <thead className={`table-light ${styles.text_14}`}>
                          <tr>
                            <th scope="col" className={`${styles.p_table}`}>
                              Kode
                            </th>
                            <th scope="col" className={`${styles.p_table}`}>
                              Mata Kuliah
                            </th>
                            <th scope="col" className={`${styles.p_table}`}>
                              Jadwal
                            </th>
                            <th scope="col" className={`${styles.p_table}`}>
                              Dosen Pengampu
                            </th>
                            <th scope="col" className={`${styles.p_table}`}>
                              Kelas
                            </th>
                            <th scope="col" className={`${styles.p_table}`}>
                              SKS
                            </th>
                            <th scope="col" className={`${styles.p_table}`}>
                              Kuota
                            </th>
                            <th scope="col" className={`${styles.p_mk}`}>
                              Pilih MK
                            </th>
                          </tr>
                        </thead>
                        <tbody>
                          {getSelectedCourses().map((course, index) => (
                            <tr key={index}>
                              <td>{course.kode}</td>
                              <td>{course.mataKuliah}</td>
                              <td>{course.jadwal}</td>
                              <td>{course.dosenPengampu}</td>
                              <td>{course.kelas}</td>
                              <td>{course.sks}</td>
                              <td>{course.kuota}</td>
                              <td>
                                <input
                                  type="checkbox"
                                  checked={course.pilihMK}
                                />
                              </td>
                            </tr>
                          ))}
                        </tbody>
                      </table>
                      <div
                        className={`h5 fw-semibold py-3 px-2 m-0 row d-flex justify-content-between`}
                        style={{ backgroundColor: "#FFE91D33" }}
                      >
                        <p className="col">Total SKS</p>
                        <p className="col text-center">3/20</p>
                      </div>

                      <div
                        className={`alert border-0 text-center ${styles.alert_gray} ${styles.m_20}`}
                        role="alert"
                      >
                        <p> + Tambah Kelas</p>
                      </div>

                      <div className="d-flex justify-content-between mt-3 gap-3">
                        <button
                          onClick={() => {}}
                          className="bg-transparent text-primary rounded px-3"
                          type="button"
                          style={{ border: "1px solid #10487A" }}
                        >
                          <span>Batal</span>
                        </button>
                        <Button
                          color="primary"
                          className="d-flex flex-grow-0 justify-content-center align-items-center"
                        >
                          <span>Update</span>
                        </Button>
                      </div>
                    </div>
                  )}
                </div>
              </>
            )}
          </div>
        </>
      ) : (
        <div
          className={`alert alert-warning alert-border-left  fade show d-flex gap-2 ${styles.warning_container} ${styles.m_20}
        `}
          role="alert"
        >
          <WarningIcon />
          <p className={styles.text_12}>
            Pengisian KRS semester Ganjil belum dibuka. Silakan cek jadwal
            akademik untuk mengetahui tanggal mulai pengisian.
          </p>
        </div>
      )}
    </>
  );
}
