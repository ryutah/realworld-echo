import { UserState } from "@/state/userState";

/**
 * UserActionType is a type of user action
 */
export const UserActionType = {
  REGISTER: "register",
  LOGIN: "login",
  LOGOUT: "logout",
  EMAIL_DUPLICATE_CHECK: "email_duplicate_check",
} as const;

type UserActionType = (typeof UserActionType)[keyof typeof UserActionType];

/**
 * RegisterPayload is a type of register payload
 */
export type RegisterPayload = {
  name: string;
  email: string;
  password: string;
};

/**
 * LoginPayload is a type of login payload
 */
export type LoginPayload = {
  email: string;
  password: string;
};

export type EmailDuplicateCheckPayload = {
  email: string;
};

/**
 * UserAction is a type of user action
 */
export type UserAction =
  | { type: typeof UserActionType.REGISTER; payload: RegisterPayload }
  | { type: typeof UserActionType.LOGIN; payload: LoginPayload }
  | { type: typeof UserActionType.LOGOUT }
  | {
      type: typeof UserActionType.EMAIL_DUPLICATE_CHECK;
      payload: EmailDuplicateCheckPayload;
    };

export default function userReducer(
  state: UserState,
  action: UserAction
): UserState {
  switch (action.type) {
    case UserActionType.REGISTER:
      return new UserState({
        uid: "xxxx",
        name: action.payload.name,
        avatar: "",
      });
    case UserActionType.LOGIN:
      console.log("aaa");
      return new UserState({
        uid: "xxxx",
        name: "xxxxxxxxx",
        avatar: "",
      });
    case UserActionType.LOGOUT:
      console.log("logout start");
      return state.logout();
    case UserActionType.EMAIL_DUPLICATE_CHECK:
      alert("email duplicate check");
      return state;
  }
}
