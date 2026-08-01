interface Props {
  data: ApiResponse<KrsDetail>;
}

export const StudentInfo = ({ data }: Props) => {
  return (
    <div className="p-3" style={{ backgroundColor: "#FAFCFF" }}>
      <div className="d-flex flex-column flex-md-row">
        <div className="d-flex flex-column gap-3 w-50">
          <div className="d-flex flex-column">
            <p>Nama</p>
            <p>{data?.data?.student_name}</p>
          </div>
          <div className="d-flex flex-column">
            <p>Periode Akademik</p>
            <p>{data?.data?.academic_periode_name}</p>
          </div>
          <div className="d-flex flex-column">
            <p>Semester</p>
            <p>-</p>
          </div>
        </div>
        <div className="d-flex flex-column gap-3 w-50">
          <div className="d-flex flex-column">
            <p>NIM</p>
            <p>{data?.data?.student_nim}</p>
          </div>
          <div className="d-flex flex-column">
            <p>Program Studi</p>
            <p>{data?.data?.study_program_name}</p>
          </div>
          <div className="d-flex flex-column">
            <p>SKS</p>
            <p>{data?.data?.total_sks_taken}</p>
          </div>
        </div>
      </div>
    </div>
  );
};
