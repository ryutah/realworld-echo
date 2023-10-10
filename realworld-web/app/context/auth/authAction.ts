import { UserState } from "./authContext";

export const AuthActionType = {
  SIGN_IN: "auth/sign_in",
  SIGN_OUT: "auth/sign_out",
  START_SESSION_CHECK: "auth/start_session_check",
  FINISH_SESSION_CHECK: "auth/finish_session_check",
} as const;
export type AuthActionType =
  (typeof AuthActionType)[keyof typeof AuthActionType];

export type AuthAction =
  | {
      type: typeof AuthActionType.SIGN_IN;
      payload: { user: UserState };
    }
  | {
      type: typeof AuthActionType.SIGN_OUT;
    }
  | {
      type: typeof AuthActionType.START_SESSION_CHECK;
    }
  | {
      type: typeof AuthActionType.FINISH_SESSION_CHECK;
    };

export const signIn = (user: UserState) => ({
  type: AuthActionType.SIGN_IN,
  payload: { user },
});

export const signOut = () => ({
  type: AuthActionType.SIGN_OUT,
});

export const startSessionCheck = () => ({
  type: AuthActionType.START_SESSION_CHECK,
});

export const finishSessionCheck = () => ({
  type: AuthActionType.FINISH_SESSION_CHECK,
});
