export const splitFileNameUploaded = (filePath: string) => {
  const parts = filePath?.split("_fn_");
  if (parts?.length > 1) {
    return parts?.[1].replace(".pdf", "");
  }
};
