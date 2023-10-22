/* tslint:disable */
/* eslint-disable */
/**
 * Conduit API
 * Conduit API documentation
 *
 * The version of the OpenAPI document: 1.0.0
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

/**
 *
 * @export
 * @interface User
 */
export interface User {
  /**
   *
   * @type {string}
   * @memberof User
   */
  email: string;
  /**
   *
   * @type {string}
   * @memberof User
   */
  token: string;
  /**
   *
   * @type {string}
   * @memberof User
   */
  username: string;
  /**
   *
   * @type {string}
   * @memberof User
   */
  bio: string;
  /**
   *
   * @type {string}
   * @memberof User
   */
  image: string;
}

/**
 * Check if a given object implements the User interface.
 */
export function instanceOfUser(value: object): boolean {
  let isInstance = true;
  isInstance = isInstance && "email" in value;
  isInstance = isInstance && "token" in value;
  isInstance = isInstance && "username" in value;
  isInstance = isInstance && "bio" in value;
  isInstance = isInstance && "image" in value;

  return isInstance;
}

export function UserFromJSON(json: any): User {
  return UserFromJSONTyped(json, false);
}

export function UserFromJSONTyped(
  json: any,
  ignoreDiscriminator: boolean
): User {
  if (json === undefined || json === null) {
    return json;
  }
  return {
    email: json["email"],
    token: json["token"],
    username: json["username"],
    bio: json["bio"],
    image: json["image"],
  };
}

export function UserToJSON(value?: User | null): any {
  if (value === undefined) {
    return undefined;
  }
  if (value === null) {
    return null;
  }
  return {
    email: value.email,
    token: value.token,
    username: value.username,
    bio: value.bio,
    image: value.image,
  };
}
