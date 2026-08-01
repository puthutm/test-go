"use client";

import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { fetchApiSso } from "@/lib/utils/fetch-server";

export const useLogout = () => {
  const { setModalConfirmationState } = useModalConfirmationContext();

  const handleLogout = async () => {
    try {
      const { data } = await fetchApiSso("/auth/logout");
      return data;
    } catch (error: any) {
      setModalConfirmationState((prev) => ({
        ...prev,
        state: "failed",
        message: error.message,
      }));
      throw new Error(error.response.data.message || "Something went wrong");
    }
  };

  return { handleLogout };
};
