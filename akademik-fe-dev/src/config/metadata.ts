import { Metadata } from "next";

export const metadataConfig: Metadata = {
  title: {
    default: "Akademik - Universitas Siber Asia",
    template: "%s | Akademik - UNSIA",
  },
  description: "Akademik Universitas Siber Asia",
  keywords: [
    "SIAKAD",
    "Sistem Informasi Akademik Unsia",
    "siakad unsia",
    "Sistem Informasi Akademik Universitas Siber Asia",
    "SIAKAD UNSIA",
    "Universitas Siber Asia",
  ],
  authors: {
    name: "Universitas Siber Asia",
  },
  creator: "DPS - Universitas Siber Asia",
  icons: [
    {
      url: "/logo-unsia.ico",
      href: "/logo-unsia.ico",
    },
  ],
  applicationName: "Siakad",
  openGraph: {
    title: "Siakad UNSIA",
    siteName: "Siakad UNSIA",
    locale: "id_Id",
    url: "https://www.unsia.ac.id",
    countryName: "Indonesia",
    type: "website",
    emails: "admission@unsia.ac.id",
    phoneNumbers: "(021) 278-061-89",
  },
  twitter: {
    title: "UNSIA",
    site: "@univsiberasia",
  },
  metadataBase: new URL("https://www.akademik.unsia.ac.id"),
};
