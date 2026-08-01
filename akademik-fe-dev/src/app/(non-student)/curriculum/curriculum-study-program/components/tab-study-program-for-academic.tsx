"use client";

import { usePathname, useRouter, useSearchParams } from "next/navigation";

export const TabStudyProgram = ({
  unsiaStudyProgram,
}: {
  unsiaStudyProgram: ApiResponse<UnsiaStudyProgram[]>;
}) => {
  const pathname = usePathname();
  const router = useRouter();
  const searchParams = useSearchParams();
  const params = new URLSearchParams(searchParams);

  let studyProgramParam = searchParams.get("study_program");

  if (!studyProgramParam) {
    studyProgramParam = unsiaStudyProgram?.data?.[0].id as string;
  }

  const handleTabChange = (tab: string) => {
    if (tab) {
      params.set("study_program", tab);
    } else {
      params.delete("study_program");
    }

    router.replace(`${pathname}?${params.toString()}`);
  };
  return (
    <div className="row gap-3 flex-wrap border-bottom pb-3 mx-0">
      {unsiaStudyProgram?.data?.map((tab) => (
        <button
          key={tab.id}
          className={`border-0 col rounded-top text-center py-2 px-4 fw-semibold ${
            studyProgramParam === tab.id
              ? "bg-primary text-white"
              : "bg-transparent"
          }`}
          style={{ color: "#909090", borderRadius: "4px 4px 0px 0px" }}
          onClick={() => handleTabChange(tab.id)}
        >
          {tab.name}
        </button>
      ))}
    </div>
  );
};
