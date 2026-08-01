"use client";

import { createOrUpdateStudentPresenceByComponent } from "@/services/api/settings/presence/lecturer/create-or-update-student-presence-by-component";
import { useMutation } from "@tanstack/react-query";
import { useParams } from "next/navigation";

interface Props {
  defaultValue: boolean;
  studentId: string;
  presenceType: StudentPresenceComponent;
}

export const FormComponentStudentPresence = ({
  defaultValue,
  presenceType,
  studentId,
}: Props) => {
  const params = useParams();

  const sessionId = params.sessionId as string;
  const { mutateAsync, isPending } = useMutation({
    mutationFn: createOrUpdateStudentPresenceByComponent,
  });
  return (
    <select
      name="open_session"
      id="open_session"
      className="form-select"
      defaultValue={String(defaultValue)}
      onChange={async (e) =>
        await mutateAsync({
          sessionId,
          payload: {
            presenceStatus: e.target.value,
            presenceType,
            studentId,
          },
        })
      }
      disabled={isPending}
    >
      <option value="true">Hadir</option>
      <option value="false">Tidak Hadir</option>
    </select>
  );
};
