import { useQuery } from "@tanstack/react-query";
import { getBatches } from "./get-batches";

export const useGetBatches = () => {
  return useQuery({
    queryKey: ["batches"],
    queryFn: () => getBatches(),
  });
};
