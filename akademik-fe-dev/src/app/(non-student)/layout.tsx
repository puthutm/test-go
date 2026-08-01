import Header from "@/components/layouts/header";
import Sidebar from "@/components/layouts/sidebar";

export const dynamic = "force-dynamic";

export default async function DashboardLayout({
  children,
  breadcrumbs,
}: {
  children: React.ReactNode;
  breadcrumbs: React.ReactNode;
}) {
  return (
    <div id="layout-wrapper">
      <Header />
      <Sidebar />
      <main className="main-content">
        <div className="page-content">
          <div className="container-fluid">
            {breadcrumbs}
            {children}
          </div>
        </div>
        {/* <Footer /> */}
      </main>
    </div>
  );
}
