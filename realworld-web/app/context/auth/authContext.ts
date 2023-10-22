export type UserState = {
  userName: string;
};

export type AuthState = {
  checkSession: boolean;
  user?: UserState;
};
