import Link from "next/link";

const Unauthenticated = () => (
  <ul className="nav navbar-nav pull-xs-right">
    <li className="nav-item">
      <Link className="nav-link active" href="/">
        Home
      </Link>
    </li>
    <li className="nav-item">
      <Link className="nav-link" href="/login">
        Sign in
      </Link>
    </li>
    <li className="nav-item">
      <Link className="nav-link" href="/register">
        Sign up
      </Link>
    </li>
  </ul>
);

const Authenticated = ({ user }: { user: User }) => (
  <ul className="nav navbar-nav pull-xs-right">
    <li className="nav-item">
      <Link className="nav-link active" href="/">
        Home
      </Link>
    </li>
    <li className="nav-item">
      <Link className="nav-link" href="/editor">
        {" "}
        <i className="ion-compose"></i> New Article{" "}
      </Link>
    </li>
    <li className="nav-item">
      <Link className="nav-link" href="/settings">
        {" "}
        <i className="ion-gear-a"></i> Settings{" "}
      </Link>
    </li>
    <li className="nav-item">
      <Link className="nav-link" href="/profile/foobar">
        <img src={user.avatar} className="user-pic" />
        {user.name}
      </Link>
    </li>
  </ul>
);

export type User = {
  name: string;
  avatar: string;
};

export default function Nav({
  user,
  isLoggedIn = false,
}: {
  user?: User;
  isLoggedIn?: boolean;
}) {
  return (
    <nav className="navbar navbar-light">
      <div className="container">
        <Link className="navbar-brand" href="/">
          conduit
        </Link>
        {isLoggedIn && user ? (
          <Authenticated user={user} />
        ) : (
          <Unauthenticated />
        )}
      </div>
    </nav>
  );
}
