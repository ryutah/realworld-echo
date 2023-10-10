import MyAppBar from "@/app/common/components/AppBar";
import { CssBaseline } from "@mui/material";
import type { Metadata } from "next";
import { Inter } from "next/font/google";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Real World",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="ja">
      <body className={inter.className}>
        <CssBaseline />
        <MyAppBar />
        {children}
      </body>
    </html>
  );
}
