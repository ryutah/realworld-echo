import { ChangeEvent, HTMLInputTypeAttribute, useState } from "react";

const Title = () => (
  <>
    <h1 className="text-xs-center">Sign up</h1>
    <p className="text-xs-center">
      <a href="">Have an account?</a>
    </p>
  </>
);

const DuplicateEmailAlert = () => (
  <ul className="error-messages">
    <li>That email is already taken</li>
  </ul>
);

type FormProps = {
  onChangeName: (name: string) => void;
  onChangeEmail: (email: string) => void;
  onChangePassword: (password: string) => void;
  onSubmit: () => void;
};

type InputProps = {
  type?: HTMLInputTypeAttribute;
  placeholder?: string;
  onChange: (e: ChangeEvent<HTMLInputElement>) => void;
};

const Input = ({ type, placeholder, onChange }: InputProps) => (
  <input
    onChange={(e) => onChange(e)}
    className="form-control form-control-lg"
    type={type}
    placeholder={placeholder}
  />
);

const Form = ({
  onChangeName,
  onChangeEmail,
  onChangePassword,
  onSubmit,
}: FormProps) => {
  return (
    <form>
      <fieldset className="form-group">
        <Input
          type="text"
          placeholder="Your Name"
          onChange={(e) => onChangeName(e.target.value)}
        />
      </fieldset>
      <fieldset className="form-group">
        <Input
          type="text"
          placeholder="Email"
          onChange={(e) => onChangeEmail(e.target.value)}
        />
      </fieldset>
      <fieldset className="form-group">
        <Input
          type="password"
          placeholder="Password"
          onChange={(e) => onChangePassword(e.target.value)}
        />
      </fieldset>
      <button
        className="btn btn-lg btn-primary pull-xs-right"
        onClick={(e) => {
          e.preventDefault();
          onSubmit();
        }}
      >
        Sign up
      </button>
    </form>
  );
};

type SignUpProps = {
  duplicateEmail?: boolean;
  onSubmit?: ({
    name,
    email,
    password,
  }: {
    name: string;
    email: string;
    password: string;
  }) => void;
};

export default function SignUp({
  duplicateEmail = false,
  onSubmit,
}: SignUpProps) {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  return (
    <div className="auth-page">
      <div className="container page">
        <div className="row">
          <div className="col-md-6 offset-md-3 col-xs-12">
            <Title />

            {duplicateEmail && <DuplicateEmailAlert />}

            <Form
              onChangeName={(name) => setName(name)}
              onChangeEmail={(email) => setEmail(email)}
              onChangePassword={(password) => setPassword(password)}
              onSubmit={() => onSubmit?.({ name, email, password })}
            />
          </div>
        </div>
      </div>
    </div>
  );
}
