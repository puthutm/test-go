"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

interface QueryParamKrs extends QueryParam {
    academic_periode_id: string;
}

export const getAllSavedKrs = async (queryParam: QueryParamKrs): Promise<ApiResponse<KrsItem[]>> => {
    const params = new URLSearchParams();

    if (queryParam.search) {
        params.append("search", queryParam.search);
    }

    if (queryParam.page !== undefined) {
        params.append("page", String(queryParam.page));
    }

    if (queryParam.limit !== undefined) {
        params.append("limit", String(queryParam.limit));
    }

    if (queryParam.sort_by) {
        params.append("sort_by", queryParam.sort_by);
    }

    if (queryParam.sort_direction) {
        params.append("sort_direction", queryParam.sort_direction);
    }

    if (queryParam.academic_periode_id) {
        params.append("academic_periode_id", queryParam.academic_periode_id);
    }

    try {
        const response = await fetchApi(`/student/academic/filling-krs/saved?${params.toString()}`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
            },
        });

        return response;
    } catch (error) {
        throw new Error(
            error instanceof Error ? error?.message : "Something went wrong"
        );
    }
};
