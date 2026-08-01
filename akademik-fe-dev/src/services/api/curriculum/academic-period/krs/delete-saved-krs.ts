"use server"

import { revalidatePath } from "next/cache";
import { fetchApi } from "@/lib/utils/fetch-server"

export const deleteSavedKrs = async (
    krsItemId: string
) => {
    try {
        const response = await fetchApi(`/student/academic/filling-krs/saved/${krsItemId}`, {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json",
            }
        })

        revalidatePath(
            `/student/academic/filling-krs/saved`
        );

        return response
    } catch (error) {
        throw new Error(
            error instanceof Error ? error?.message : "Something went wrong"
        );
    }
}
