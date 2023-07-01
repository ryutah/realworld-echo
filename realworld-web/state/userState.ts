/**
 * UserState is a type of user state
 */

export class UserState {
  uid?: string;
  name?: string;
  avatar?: string;

  constructor(params?: { uid?: string; name?: string; avatar?: string }) {
    this.uid = params?.uid;
    this.name = params?.name;
    this.avatar = params?.avatar;
  }

  public logout(): UserState {
    return new UserState();
  }

  public isLoggedIn(): boolean {
    return this.uid !== undefined;
  }
}
