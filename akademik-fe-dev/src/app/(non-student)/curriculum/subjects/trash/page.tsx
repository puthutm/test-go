import { Metadata } from "next";
import React from "react";
import SubjectTrashClientPage from "./_components/client";

export const metadata: Metadata = {
  title: "Mata Kuliah (Sampah)",
};

export default function SubjectTrashPage() {
  return <SubjectTrashClientPage />;
}
