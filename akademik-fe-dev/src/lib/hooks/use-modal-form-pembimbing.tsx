import { Button, Modal, ModalHeader, ModalBody, ModalFooter } from 'reactstrap';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { FormPembimbingSchema, FormPembimbingSchemaType } from '@/lib/validations/students/akademik/form-pembimbing-schema';
import { FormDescription } from '@/components/ui/form-description';
import { FormErrorMessage } from '@/components/ui/form-error-message';
import { UploadIcon } from '@/components/icons/upload';
import { useState } from 'react';


const KonsultasiModal: React.FC<any> = ({ isOpen, toggle }) => {
  const {
    formState: { errors },
    register,
    handleSubmit,
    reset,
  } = useForm<FormPembimbingSchemaType>({
    resolver: zodResolver(FormPembimbingSchema),
  });

  const [fileName, setFileName] = useState<string>("Upload File"); // State buat nyimpen nama file

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files && event.target.files.length > 0) {
      setFileName(event.target.files[0].name); // Simpan nama file yang diupload
    }
  };

  const onSubmit = (data: FormPembimbingSchemaType) => {
    console.log('Data yang dikirim:', data);
    toggle(); // Tutup modal setelah submit
  };

  const handleCancel = () => {
    reset(); // Reset form ketika tombol Batal ditekan
    setFileName("Upload File"); // Reset nama file
    toggle(); // Tutup modal setelah reset
  };

  return (
    <Modal isOpen={isOpen} toggle={toggle} centered>
      <ModalHeader toggle={toggle} className="d-flex justify-content-between align-items-center pb-3" style={{ position: 'relative' }}>
        <div style={{ fontSize: "14px" }}>
          Buat Konsultasi
        </div>
        <div style={{ position: 'absolute', bottom: 0, left: '4%', right: '4%', borderBottom: '2px solid #ccc' }} />
      </ModalHeader>


      <form onSubmit={handleSubmit(onSubmit)}>
        <ModalBody>
          <div className="mb-4">
            <label className="form-label">Topik Konsultasi</label>
            <input {...register("topik_konsultasi")} className="form-control" placeholder="Masukkan Topik" />
            <FormErrorMessage errors={errors.topik_konsultasi} />
          </div>

          <div className="mb-4">
            <label className="form-label">Isi Pesan</label>
            <textarea
              {...register("isi_pesan")}
              className="form-control"
              placeholder="Masukkan Pesan"
              style={{ height: "80px", resize: "none" }} // Mengatur tinggi dan ukuran font
            />
            <FormErrorMessage errors={errors.isi_pesan} />
          </div>

          <div className="mb-0">
            <label className="form-label">File Dokumen</label>
            <div className="d-flex justify-content-between align-items-center">

              {/* Label sebagai tombol upload */}
              <label
                htmlFor="uploadDokumen"
                className="d-flex align-items-center border border-2 rounded p-3 m-0"
                style={{ width: "330px", height: "45px", cursor: "pointer" }}
              >
                <span style={{ color: "#909090" }}>{fileName}</span> {/* Menampilkan nama file */}
              </label>

              {/* Tombol Upload */}
              <label
                htmlFor="uploadDokumen"
                className="d-flex justify-content-center align-items-center border border-2 rounded p-3 m-0"
                style={{ width: "120px", height: "45px", cursor: "pointer" }}
              >
                <UploadIcon />
                <span>Upload</span>
              </label>

              {/* Input file yang tersembunyi */}
              <input
                {...register("dokumen")}
                type="file"
                accept=".pdf"
                id="uploadDokumen"
                className="d-none"
                onChange={handleFileChange} // Tambahin event handler 
              />
            </div>
            <FormErrorMessage errors={errors.dokumen} />
            <FormDescription message="File dalam bentuk file .pdf max 10mb" />
          </div>
        </ModalBody>

        {/* ModalFooter di dalam form */}
        <ModalFooter className="d-flex justify-content-between mt-3">
          <Button
            onClick={handleCancel} // Menggunakan handleCancel untuk reset dan tutup modal
            className="bg-transparent text-primary rounded px-3"
            type="button"
            style={{ border: "1px solid #10487A" }}
          >
            <span>Batal</span>
          </Button>
          <Button
            color="primary"
            type="submit"
            className="d-flex flex-grow-0 justify-content-center align-items-center"
          >
            <span>Submit</span>
          </Button>
        </ModalFooter>
      </form>
    </Modal>


  );
};

export default KonsultasiModal;
