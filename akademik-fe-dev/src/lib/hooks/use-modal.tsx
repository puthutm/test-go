"use client";

import { useContext } from "react";

import { ModalContext } from "../contexts/modal-context";

type UseModal = {
  modalState: ModalState;
  setModalState: React.Dispatch<React.SetStateAction<ModalState>>;
};

export const useModalContext = (): UseModal => {
  const context = useContext(ModalContext);
  if (!context) {
    throw new Error("useModalContext must be used within a ModalProvider");
  }
  return context;
};
