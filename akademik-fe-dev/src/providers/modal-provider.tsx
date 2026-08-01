"use client";

import { useState } from "react";

import { ModalContext } from "@/lib/contexts/modal-context";

export const ModalProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [modalState, setModalState] = useState<ModalState>({
    open: false,
    state: "add",
    id: null,
  });

  return (
    <ModalContext.Provider value={{ modalState, setModalState }}>
      {children}
    </ModalContext.Provider>
  );
};
