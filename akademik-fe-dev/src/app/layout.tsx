import type { Metadata } from "next";
import { Poppins } from "next/font/google";

import authOptions from "@/config/next-auth";
import { metadataConfig } from "@/config/metadata";
import { QueryProvider } from "@/providers/query-client";
import { AuthProvider } from "../providers/auth-provider";
import { getServerSession, Session } from "next-auth";
import { ModalConfirmationProvider } from "@/providers/modal-confirmation-provider";
import { ModalProvider } from "@/providers/modal-provider";
import { ModalSuccessConfirmation } from "@/components/ui/modal-success-confirmation";
import { ModalErrorConfirmation } from "@/components/ui/modal-error-confirmation";

import "./globals.css";
import "@/assets/scss/themes.scss";
import { ToolbarContext } from "@/lib/contexts/toolbar-context";
import { LexicalProvider } from "@/lib/contexts/lexical-context";

const poppins = Poppins({
  subsets: ["latin"],
  weight: ["400", "500", "600", "700"],
  display: "swap",
  preload: false,
  adjustFontFallback: false,
});

export function generateMetadata(): Metadata {
  return {
    ...metadataConfig,
  };
}

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const session = await getServerSession(authOptions);
  return (
    <html
      lang="en"
      data-layout-style="default"
      data-sidebar-size="lg"
      data-sidebar="dark"
      data-layout-width="fluid"
      data-bs-theme="light"
      data-layout-position="fixed"
      data-topbar="light"
      data-layout="vertical"
      data-sidebar-image="none"
      data-sidebar-visibility="show"
    >
      <body suppressHydrationWarning className={poppins.className}>
        <AuthProvider session={session as Session}>
          <QueryProvider>
            <LexicalProvider>
              <ToolbarContext>
                <ModalProvider>
                  <ModalConfirmationProvider>
                    <ModalSuccessConfirmation />
                    <ModalErrorConfirmation />
                    {children}
                  </ModalConfirmationProvider>
                </ModalProvider>
              </ToolbarContext>
            </LexicalProvider>
          </QueryProvider>
        </AuthProvider>
      </body>
    </html>
  );
}
