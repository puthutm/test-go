import { Metadata } from "next";

import { HeaderBiodata } from "@/components/ui/header-biodata";
import { BannerProfile } from "@/components/ui/profile-banner";
import { Tabs } from "./components/tabs";

export const metadata: Metadata = {
  title: "Mahasiswa",
};

export default function StudentBiodataLayou({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <>
      <BannerProfile />
      <div className="container">
        <HeaderBiodata />
        <Tabs />
        {children}
      </div>
    </>
  );
}
