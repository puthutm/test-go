"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { signOut } from "next-auth/react";
import { useCallback, useEffect } from "react";
import { useForm, Controller } from "react-hook-form";
import { Button, Col, Input, Row } from "reactstrap";
import { $getRoot } from "lexical";

import { SaveIcon } from "@/components/icons/save";
import { UploadIcon } from "@/components/icons/upload";
import { FormDescription } from "@/components/ui/form-description";
import { FormErrorMessage } from "@/components/ui/form-error-message";
import Editor from "@/components/ui/text-editor/index";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { splitFileNameUploaded } from "@/lib/utils/split-filename-upload";
import {
  FormClassContractSchema,
  FormClassContractSchemaType,
} from "@/lib/validations/settings/academic-period/form-class-contract";
import { useGetFileStorage } from "@/services/api/file-storage/use-get-file-storage";
import { updateClassContract } from "@/services/api/settings/academic-period/class-contract/update-class-contract";
import { useLexicalComposerContext } from "@lexical/react/LexicalComposerContext";

export const FormClassContract = ({
  classId,
  detailClass,
  isDetail,
}: {
  classId: string;
  detailClass: ApiResponse<Class>;
  isDetail?: boolean;
}) => {
  const { setModalConfirmationState } = useModalConfirmationContext();
  const [editor] = useLexicalComposerContext();

  const {
    handleSubmit,
    control,
    setValue,
    formState: { errors, isSubmitting },
  } = useForm({
    resolver: zodResolver(FormClassContractSchema),
  });

  const { getFileStorage, loading } = useGetFileStorage();

  const setFormValue = useCallback(() => {
    if (!detailClass?.data?.contract_description) {
      editor.update(() => {
        const root = $getRoot();
        root.clear();
      });

      return;
    }
    setValue(
      "contract_description",
      JSON.parse(detailClass?.data?.contract_description)
    );
  }, [setValue, detailClass, editor]);

  useEffect(() => {
    setFormValue();
  }, [setValue, detailClass?.data?.contract_description]);

  if (detailClass?.status === 401) return signOut();

  const onSubmit = async (data: FormClassContractSchemaType) => {
    try {
      const payload = new FormData();

      payload.append("contract_file", data.contract_file);
      payload.append(
        "contract_description",
        JSON.stringify(data.contract_description)
      );

      const response = await updateClassContract(classId, payload);

      if (response.error) {
        return setModalConfirmationState((prev) => ({
          ...prev,
          open: true,
          state: "failed",
          message: response.message,
        }));
      }

      return setModalConfirmationState((prev) => ({
        ...prev,
        open: true,
        message: "Data berhasil di-update",
        state: "success",
      }));
    } catch (error: any) {
      setModalConfirmationState((prev) => ({
        ...prev,
        open: true,
        state: "failed",
        message: error.toString(),
      }));
    }
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className={isDetail ? "mb-3" : ""}>
      <Row className="gap-2">
        <Col sm={12}>
          <div className="px-3 border rounded-3 pt-3 pb-4">
            <h2 className="fs-5 fw-semibold mb-1" style={{ color: "#3A3A3A" }}>
              Kontrak Kuliah
            </h2>
            <Controller
              name="contract_description"
              control={control}
              render={({ field }) => (
                <Editor
                  onChange={field.onChange}
                  value={field.value}
                  isError={!!errors.contract_description}
                  placeholder="Masukkan deskripsi kontrak kuliah"
                  disabled={isDetail}
                />
              )}
            />
            <FormErrorMessage errors={errors.contract_description} />
          </div>
        </Col>

        <Col sm={12} style={{ position: "unset" }}>
          <label className="fs-5 fw-semibold mb-1" style={{ color: "#3A3A3A" }}>
            Dokumen Kontrak Kuliah
          </label>
          <Controller
            name="contract_file"
            control={control}
            render={({ field }) => (
              <div className="d-flex gap-2">
                <label
                  htmlFor="contract_file"
                  className={`form-control form-control-icon mb-0 ${
                    errors.contract_file ? "border border-danger" : ""
                  }`}
                >
                  {field.value instanceof File
                    ? field.value.name
                    : splitFileNameUploaded(
                        detailClass?.data?.contract_file_path
                      ) || "Belum ada berkas yang dipilih"}
                </label>
                <Input
                  type="file"
                  id="contract_file"
                  accept="application/pdf"
                  onChange={(e) => {
                    const file = e.target.files?.[0];
                    field.onChange(file);
                  }}
                  hidden
                  disabled={isDetail}
                />
                {detailClass?.data?.contract_file_path ? (
                  <Button
                    color="light"
                    className="flex-shrink-0 rounded"
                    type="button"
                    onClick={() =>
                      getFileStorage(detailClass?.data?.contract_file_path)
                    }
                    disabled={loading}
                  >
                    {loading ? "Mengunduh" : "Lihat File"}
                  </Button>
                ) : (
                  <label
                    htmlFor="contract_file"
                    className={`btn d-flex align-items-center btn-light mb-0 ${
                      isDetail ? "pe-none" : ""
                    }`}
                    style={{ whiteSpace: "nowrap" }}
                  >
                    <UploadIcon /> Upload File
                  </label>
                )}
              </div>
            )}
          />
          {errors.contract_file ? (
            <FormErrorMessage errors={errors.contract_file} />
          ) : (
            <FormDescription message="File dengan format .pdf maksimal 10mb" />
          )}
        </Col>
      </Row>
      {!isDetail && (
        <div className="d-flex justify-content-end mt-3">
          <Button color="success" disabled={isSubmitting}>
            <SaveIcon /> Simpan
          </Button>
        </div>
      )}
    </form>
  );
};
