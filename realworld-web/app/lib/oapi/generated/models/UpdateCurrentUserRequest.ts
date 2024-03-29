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

import type { UpdateUser } from "./UpdateUser";
import { UpdateUserFromJSON, UpdateUserToJSON } from "./UpdateUser";

/**
 *
 * @export
 * @interface UpdateCurrentUserRequest
 */
export interface UpdateCurrentUserRequest {
  /**
   *
   * @type {UpdateUser}
   * @memberof UpdateCurrentUserRequest
   */
  user: UpdateUser;
}

/**
 * Check if a given object implements the UpdateCurrentUserRequest interface.
 */
export function instanceOfUpdateCurrentUserRequest(value: object): boolean {
  let isInstance = true;
  isInstance = isInstance && "user" in value;

  return isInstance;
}

export function UpdateCurrentUserRequestFromJSON(
  json: any
): UpdateCurrentUserRequest {
  return UpdateCurrentUserRequestFromJSONTyped(json, false);
}

export function UpdateCurrentUserRequestFromJSONTyped(
  json: any,
  ignoreDiscriminator: boolean
): UpdateCurrentUserRequest {
  if (json === undefined || json === null) {
    return json;
  }
  return {
    user: UpdateUserFromJSON(json["user"]),
  };
}

export function UpdateCurrentUserRequestToJSON(
  value?: UpdateCurrentUserRequest | null
): any {
  if (value === undefined) {
    return undefined;
  }
  if (value === null) {
    return null;
  }
  return {
    user: UpdateUserToJSON(value.user),
  };
}
