"use client";

import { ClockIcon } from "@/components/icons/clock";
import { WarningIcon } from "@/components/icons/warning";
import styles from "@/styles/krs.module.css";
import { useEffect, useState } from "react";
import {
  useGetListKrs,
  useGetSavedKrs,
} from "@/services/api/curriculum/academic-period/krs/use-get-list-krs";
import { takeClassForKrs } from "@/services/api/curriculum/academic-period/krs/create-take-class";
import { deleteSavedKrs } from "@/services/api/curriculum/academic-period/krs/delete-saved-krs";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import PilihKelasTab from "./krs/pilih-kelas-tab";
import KrsTersimpanTab from "./krs/krs-tersimpan-tab";
import MkMengulangTab from "./krs/mk-mengulang-tab";
import { useGetKrsInfo } from "@/services/api/curriculum/academic-period/krs/use-get-krs-info";

const useCountdown = (targetDate: string | Date) => {
  const [countdownTime, setCountdownTime] = useState({
    hours: 0,
    minutes: 0,
    seconds: 0,
    isExpired: false,
  });

  useEffect(() => {
    const targetTime =
      typeof targetDate === "string"
        ? new Date(targetDate).getTime()
        : targetDate.getTime();

    const calculateTimeLeft = () => {
      const now = new Date().getTime();
      const difference = targetTime - now;

      if (difference <= 0) {
        return {
          hours: 0,
          minutes: 0,
          seconds: 0,
          isExpired: true,
        };
      }

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

    setCountdownTime(calculateTimeLeft());

    const timer = setInterval(() => {
      setCountdownTime(calculateTimeLeft());
    }, 1000);

    return () => clearInterval(timer);
  }, [targetDate]);

  return countdownTime;
};

export default function FormKrs() {
  const [isOpen, setIsOpen] = useState(false);
  const [targetDate, setTargetDate] = useState("2026-02-19T15:00:00+07:00");
  const [endDate, setEndDate] = useState("2026-02-19T15:00:00+07:00");
  const [isActive, setIsActive] = useState<number>(0);
  const [isKrs, setIsKrs] = useState(0);
  const [isPaket] = useState(false);
  const buttons = ["Pilih Kelas", "KRS Tersimpan", "M.K Mengulang"];
  const queryParam = { academic_periode_id: "", page: 1, limit: 100 };
  const {
    data: listData,
    refetch: refetchList,
    isLoading: isLoadingList,
  } = useGetListKrs(queryParam);
  const {
    data: savedData,
    refetch: refetchSaved,
    isLoading: isLoadingSaved,
  } = useGetSavedKrs(queryParam);
  const { data: krsInfo } = useGetKrsInfo();
  const { setModalConfirmationState } = useModalConfirmationContext();

  const refreshData = async () => {
    await Promise.all([refetchList(), refetchSaved()]);
  };

  useEffect(() => {
    const now = new Date();
    if (now > new Date(targetDate)) {
      setIsOpen(true);
    }
  }, [targetDate]);

  const activePeriod = listData?.data?.academic_periods?.find(
    (p: any) => p.is_active
  );
  const semesterName = activePeriod?.fullname || "Semester";

  useEffect(() => {
    const startDateMs = activePeriod?.start_date_of_college;
    const endDateMs = activePeriod?.end_date_of_college;
    if (startDateMs && endDateMs) {
      setTargetDate(new Date(startDateMs).toISOString());
      setEndDate(new Date(endDateMs).toISOString());
    }
  }, [activePeriod]);

  const classes = listData?.data?.classes ?? [];
  const savedKrsItems = savedData?.data ?? [];

  const countdown = useCountdown(isOpen ? endDate : targetDate);
  const formatTime = (time: number): string => {
    return time.toString().padStart(2, "0");
  };

  const formattedCountdown = countdown.isExpired
    ? "00 : 00 : 00"
    : `${formatTime(countdown.hours)} : ${formatTime(
        countdown.minutes
      )} : ${formatTime(countdown.seconds)}`;

  const handleTabChange = (index: number) => {
    setIsActive(index);
    setIsKrs(index);
  };

  const onSubmitTakeClass = async (classId: string) => {
    try {
      const response = await takeClassForKrs({ class_id: classId });
      if (!response || response.error) {
        setModalConfirmationState((prev) => ({
          ...prev,
          open: true,
          state: "failed",
          message: response?.message || "Gagal mengambil kelas.",
        }));
        return;
      }
      await refreshData();
    } catch (error: any) {
      setModalConfirmationState((prev) => ({
        ...prev,
        open: true,
        state: "failed",
        message:
          error?.message ||
          error?.toString() ||
          "Terjadi kesalahan saat mengambil kelas.",
      }));
    }
  };

  const handleDeleteKrs = async (krsItemId: string) => {
    try {
      const response = await deleteSavedKrs(krsItemId);
      if (!response || response.error) {
        setModalConfirmationState((prev) => ({
          ...prev,
          open: true,
          state: "failed",
          message: response?.message || "Gagal menghapus KRS.",
        }));
        return;
      }
      await refreshData();
    } catch (error: any) {
      setModalConfirmationState((prev) => ({
        ...prev,
        open: true,
        state: "failed",
        message:
          error?.message ||
          error?.toString() ||
          "Terjadi kesalahan saat menghapus KRS.",
      }));
    }
  };

  const renderActiveTab = () => {
    switch (isKrs) {
      case 0:
        return (
          <PilihKelasTab
            classes={classes}
            savedKrsItems={savedKrsItems}
            semesterName={semesterName}
            onTakeClass={onSubmitTakeClass}
            isLoading={isLoadingList}
          />
        );
      case 1:
        return (
          <KrsTersimpanTab
            savedKrsItems={savedKrsItems}
            onDeleteKrs={handleDeleteKrs}
            isLoading={isLoadingSaved}
          />
        );
      case 2:
        return <MkMengulangTab />;
      default:
        return null;
    }
  };

  return (
    <>
      <div className="d-flex border-2 border-bottom justify-content-between align-items-center">
        <h5
          className="card-title py-3 mb-1 font-poppins"
          style={{ fontWeight: "500" }}
        >
          Pengisian KRS
        </h5>
      </div>

      {isOpen ? (
        <>
          <div>
            <div
              className={`alert alert-warning alert-border-left ${styles.warning_container} ${styles.m_20} d-flex justify-content-between align-items-center`}
              role="alert"
            >
              <p className={`${styles.text_12}`}>
                Selesaikan pengisian KRS Anda sebelum{" "}
                {new Date(endDate).toLocaleDateString("id-ID", {
                  day: "numeric",
                  month: "long",
                  year: "numeric",
                })}{" "}
                Pukul{" "}
                {new Date(endDate).toLocaleTimeString("id-ID", {
                  hour: "2-digit",
                  minute: "2-digit",
                  hour12: false,
                })}{" "}
                WIB.
              </p>
              <button
                type="button"
                className={`btn btn-danger ${styles.clock_btn} waves-effect waves-light d-flex align-items-center justidfy-content-between`}
              >
                <ClockIcon />
                <p>{formattedCountdown}</p>
              </button>
            </div>
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
            <div
              className={`alert alert-light container text-dark ${styles.p_alert_light}`}
              role="alert"
            >
              <div className="row">
                <p className="col">Semester saat ini : {semesterName}</p>
                <p className="col">
                  Total SKS yang dapat diambil : {krsInfo?.data?.max_sks}
                </p>
              </div>
            </div>
            {isPaket ? (
              <></>
            ) : (
              <div>
                <div className="d-flex">
                  {buttons.map((label, index) => (
                    <button
                      className={`${styles.btn_table_krs} ${
                        isActive === index ? styles.active : ""
                      }`}
                      onClick={() => handleTabChange(index)}
                      key={index}
                    >
                      {label}
                    </button>
                  ))}
                </div>
                {renderActiveTab()}
              </div>
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
