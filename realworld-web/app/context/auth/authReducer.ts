import { AuthAction, AuthActionType } from "./authAction";
import { AuthState } from "./authContext";

export function authReducer(state: AuthState, action: AuthAction): AuthState {
  switch (action.type) {
    case AuthActionType.SIGN_IN:
      return { ...state, user: action.payload.user };
    case AuthActionType.SIGN_OUT:
      const { user: _, ...rest } = state;
      return { ...rest };
    case AuthActionType.START_SESSION_CHECK:
      return { ...state, checkSession: true };
    case AuthActionType.FINISH_SESSION_CHECK:
      return { ...state, checkSession: false };
  }
}
