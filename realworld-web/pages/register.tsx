import SignUp from "@/components/auth/SignUp";
import { useUserDispatchContext } from "@/contexts/UserProvider";
import { UserActionType } from "@/reducers/userReducer";

export default function Register() {
  const dispatch = useUserDispatchContext();
  const onSubmit = ({
    name,
    email,
    password,
  }: {
    name: string;
    email: string;
    password: string;
  }) => {
    dispatch({
      type: UserActionType.REGISTER,
      payload: { name, email, password },
    });
  };

  return (
    <>
      <SignUp onSubmit={onSubmit} />
    </>
  );
}
