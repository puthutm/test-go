"use client"

import { useQuery } from "@tanstack/react-query";
import { getAllListKrs } from "./get-all-list-krs";
import { getAllSavedKrs } from "./get-all-saved-krs";

interface QueryParamKrs extends QueryParam {
    academic_periode_id: string;
}

export const useGetListKrs = (queryParam: QueryParamKrs) => {
    return useQuery({
        queryKey: ["list-krs", queryParam.academic_periode_id, queryParam.search, queryParam.page, queryParam.limit, queryParam.sort_by, queryParam.sort_direction],
        queryFn: async () => {
            const data = await getAllListKrs(queryParam);
            return data;
        },
    });
};

export const useGetSavedKrs = (queryParam: QueryParamKrs) => {
    return useQuery({
        queryKey: ["saved-krs", queryParam.academic_periode_id],
        queryFn: async () => {
            const data = await getAllSavedKrs(queryParam);
            return data;
        },
    });
};