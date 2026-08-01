"use client";

import { useContext } from "react";

import { ModalConfirmationContext } from "../contexts/modal-confirmation-context";

type UseModalConfirmation = {
  modalConfirmationState: ModalConfirmationState;
  setModalConfirmationState: React.Dispatch<
    React.SetStateAction<ModalConfirmationState>
  >;
};

export const useModalConfirmationContext = (): UseModalConfirmation => {
  const context = useContext(ModalConfirmationContext);
  if (!context) {
    throw new Error(
      "useModalConfirmationContext must be used within a ModalConfirmationProvider"
    );
  }
  return context;
};
