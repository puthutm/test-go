"use client";

import { useState } from "react";

import { ModalConfirmationContext } from "@/lib/contexts/modal-confirmation-context";

export const ModalConfirmationProvider: React.FC<{
  children: React.ReactNode;
}> = ({ children }) => {
  const [modalConfirmationState, setModalConfirmationState] =
    useState<ModalConfirmationState>({
      open: false,
      state: "success",
      message: null,
    });

  return (
    <ModalConfirmationContext.Provider
      value={{ modalConfirmationState, setModalConfirmationState }}
    >
      {children}
    </ModalConfirmationContext.Provider>
  );
};
