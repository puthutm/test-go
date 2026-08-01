"use server"

import { revalidatePath } from "next/cache";
import { fetchApi } from "@/lib/utils/fetch-server"
import { FormKrsDetailType } from "../../../../../lib/validations/academic/settings/academic-period/form-krs-detail-schema";

export const takeClassForKrs = async (
    payload: FormKrsDetailType
) => {
    const reqBody = {
        class_id: payload.class_id,
    }

    try {
        const response = await fetchApi(`/student/academic/filling-krs/pick/take`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(reqBody),
        })

        revalidatePath(
            `/student/academic/filling-krs/pick/take`
        );

        return response
    } catch (error) {
        throw new Error(
            error instanceof Error ? error?.message : "Something went wrong"
        );
    }
}
