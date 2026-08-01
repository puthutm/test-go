"use client";

import React, { createContext } from "react";

type ModalContextType = {
  modalConfirmationState: ModalConfirmationState;
  setModalConfirmationState: React.Dispatch<
    React.SetStateAction<ModalConfirmationState>
  >;
};

export const ModalConfirmationContext = createContext<
  ModalContextType | undefined
>(undefined);
