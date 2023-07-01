import {
  createContext,
  Dispatch,
  useContext,
  useReducer,
  ReactNode,
} from "react";
import userReducer, { UserAction } from "@/reducers/userReducer";

import { UserState } from "@/state/userState";

const UserContext = createContext<UserState | undefined>(undefined);

/**
 * useUserContext returns user contexts
 * @returns user contexts
 */
export function useUserContext() {
  const context = useContext(UserContext);
  if (!context) {
    throw new Error("useUserContext must be used within a UserProvider");
  }
  return context;
}

type UserDispatch = Dispatch<UserAction>;

const UserDispatchContext = createContext<UserDispatch | undefined>(undefined);

/**
 * useUserDispatchContext returns user dispatch contexts
 * @returns user dispatch contexts
 */
export function useUserDispatchContext() {
  const context = useContext(UserDispatchContext);
  if (!context) {
    throw new Error(
      "useUserDispatchContext must be used within a UserProvider"
    );
  }
  return context;
}

export default function UserProvider({ children }: { children: ReactNode }) {
  const [state, dispatch] = useReducer(userReducer, new UserState());

  return (
    <UserDispatchContext.Provider value={dispatch}>
      <UserContext.Provider value={state}>{children}</UserContext.Provider>
    </UserDispatchContext.Provider>
  );
}
