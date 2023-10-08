import MyAppBar from "@/app/components/AppBar";
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
      <CssBaseline />
      <body className={inter.className}>
        <MyAppBar />
        {children}
      </body>
    </html>
  );
}