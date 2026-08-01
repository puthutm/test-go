"use client";

import { saveAs } from "file-saver";
import { useState } from "react";

import { splitFileNameUploaded } from "@/lib/utils/split-filename-upload";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";

export const useGetFileStorage = () => {
  const [loading, setLoading] = useState(false);
  const { setModalConfirmationState } = useModalConfirmationContext();
  const getFileStorage = async (path: string) => {
    try {
      setLoading(true);
      const response = await fetch(
        `/api/download?path=${encodeURIComponent(path)}`
      );

      if (!response.ok) {
        return setModalConfirmationState((prev) => ({
          ...prev,
          state: "failed",
          open: true,
          message: "Gagal mengunduh file",
        }));
      }

      const data = await response.blob();

      saveAs(data, splitFileNameUploaded(path));
      return data;
    } catch (error: any) {
      console.log(error.message);

      throw new Error(error?.message);
    } finally {
      setLoading(false);
    }
  };

  return { getFileStorage, loading };
};
