import MyAppBar from "@/app/components/AppBar";
import type { Metadata } from "next";
import { Inter } from "next/font/google";
import ThemeRegistry from "./components/ThemeRegistry";

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
        <ThemeRegistry options={{ key: "mui" }}>
          <MyAppBar />
          {children}
        </ThemeRegistry>
      </body>
    </html>
  );
}
