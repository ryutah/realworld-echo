export class Profile {
  readonly userName: string;
  readonly bio: string;
  readonly image: string;
  readonly following: boolean;

  constructor(data: any) {
    this.userName = data.userName;
    this.bio = data.bio;
    this.image = data.image;
    this.following = data.following;
  }
}
