import { useUserContext } from "@/contexts/UserProvider";
import Layout from "@/components/Layout";
import { ReactNode } from "react";

export default function LayoutContainer({ children }: { children: ReactNode }) {
  const user = useUserContext();

  return (
    <Layout
      user={{
        name: user.name ?? "",
        avatar: user.avatar ?? "",
        isLoggedIn: user.isLoggedIn(),
      }}
    >
      {children}
    </Layout>
  );
}
