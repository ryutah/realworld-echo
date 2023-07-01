import Head from "@/components/Head";
import Footer from "@/components/Footer";
import Nav from "@/components/Nav";

type LayoutProps = {
  children: React.ReactNode;
  user: {
    name: string;
    avatar: string;
    isLoggedIn: boolean;
  };
};

export default function Layout(props: LayoutProps) {
  return (
    <>
      <Head />
      <Nav
        user={{
          name: props.user.name,
          avatar: props.user.avatar,
        }}
        isLoggedIn={props.user.isLoggedIn}
      />
      <main>{props.children}</main>
      <Footer />
    </>
  );
}
