import { FooterStudent } from "@/components/ui/footer";
import { NavbarStudent } from "@/components/ui/navbar";

export default function StudentLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <>
      <NavbarStudent />
      <div className="page-content px-0" style={{ minHeight: "100dvh" }}>
        {children}
      </div>
      <FooterStudent />
    </>
  );
}
