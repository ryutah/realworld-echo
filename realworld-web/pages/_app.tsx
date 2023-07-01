import UserProvider from "@/contexts/UserProvider";
import LayoutContainer from "@/containers/LayoutContainer";
import type { AppProps } from "next/app";

export default function App({ Component, pageProps }: AppProps) {
  return (
    <UserProvider>
      <LayoutContainer>
        <Component {...pageProps} />
      </LayoutContainer>
    </UserProvider>
  );
}
