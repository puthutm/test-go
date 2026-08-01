import { z } from "zod";

export const academicPeriodSchema = ({ isEdit }: { isEdit: boolean }) => {
  return z.object({
    fullname: z
      .string({
        required_error: "Nama harus diisi",
      })
      .min(1, {
        message: "Nama harus diisi",
      }),
    code: z
      .string({
        required_error: "Kode harus diisi",
      })
      .min(1, {
        message: "Kode harus diisi",
      }),
    shortname: z
      .string({
        required_error: "Nama pendek harus diisi",
      })
      .min(1, {
        message: "Nama pendek harus diisi",
      }),
    academic_year_id: z.object(
      {
        label: z.string(),
        value: z.string(),
      },
      {
        message: "Tahun akademik harus dipilih",
      }
    ),
    semester_id: z.object(
      {
        label: z.string(),
        value: z.string(),
      },
      {
        message: "Semester harus dipilih",
      }
    ),
    start_date_of_college: z
      .array(z.date(), {
        message: "Tanggal mulai perkuliahan harus diisi",
      })
      .min(1, "Tanggal mulai perkuliahan wajib diisi"),
    end_date_of_college: z
      .array(z.date(), {
        message: "Tanggal selesai perkuliahan harus diisi",
      })
      .min(1, "Tanggal selesai perkuliahan wajib diisi"),
    start_date_of_uts: isEdit
      ? z
          .array(z.date(), {
            message: "Tanggal mulai uts harus diisi",
          })
          .min(1, "Tanggal mulai uts wajib diisi")
      : z.array(z.date()).nullable(),
    end_date_of_uts: isEdit
      ? z
          .array(z.date(), {
            message: "Tanggal selesai uts harus diisi",
          })
          .min(1, "Tanggal selesai uts wajib diisi")
      : z.array(z.date()).nullable(),
    start_date_of_uas: isEdit
      ? z
          .array(z.date(), {
            message: "Tanggal mulai uas harus diisi",
          })
          .min(1, "Tanggal mulai uas wajib diisi")
      : z.array(z.date()).nullable(),
    end_date_of_uas: isEdit
      ? z
          .array(z.date(), {
            message: "Tanggal selesai uas harus diisi",
          })
          .min(1, "Tanggal selesai uas wajib diisi")
      : z.array(z.date()).nullable(),
    number_of_lecture_meeting: z.object(
      {
        label: z.string(),
        value: z.string(),
      },
      {
        message: "Jumlah pertemuan harus dipilih",
      }
    ),
    is_active: isEdit ? z.boolean() : z.boolean().nullable(),
  });
};

export type AcademicPeriodFormType = z.infer<
  ReturnType<typeof academicPeriodSchema>
>;

export const AcademicPeriodInitValues = {
  code: "",
  is_active: false,
  fullname: "",
  shortname: "",
  start_date_of_college: [],
  end_date_of_college: [],
  start_date_of_uas: [],
  end_date_of_uas: [],
  start_date_of_uts: [],
  end_date_of_uts: [],
};
