"use client";

import React, { createContext } from "react";

type ModalContextType = {
  modalState: ModalState;
  setModalState: React.Dispatch<React.SetStateAction<ModalState>>;
};

export const ModalContext = createContext<ModalContextType | undefined>(
  undefined
);
