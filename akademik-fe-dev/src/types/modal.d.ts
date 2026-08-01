interface ModalState {
    open: boolean;
    state: "add" | "edit" | "detail" | "duplicate";
    id?: string | null;
  }
  
  interface ModalConfirmationState {
    open: boolean;
    state: "confirm" | "success"| "failed";
    message: string | null;
    id?: string | null;
  }
  