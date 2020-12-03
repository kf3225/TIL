import { AxiosError } from "axios";

import { User } from "../services/github/models";
import * as ActionType from "./githubConstants";

interface GetMembersParams {
  companyName: string;
}

interface GetMembersResult {
  users: User[];
}

export const getMembers = {
  start: (params: GetMembersParams) =>
    ({
      type: ActionType.GET_MEMBERS_START,
      payload: params,
    } as const),

  succeed: (params: GetMembersParams, result: GetMembersResult) =>
    ({
      type: ActionType.GET_MEMBERS_SUCCEED,
      payload: { params, result },
    } as const),

  fail: (params: GetMembersParams, error: AxiosError) =>
    ({
      type: ActionType.GET_MEMBERS_FAIL,
      payload: { params, error },
      error: true,
    } as const),
};

export type GithubAction =
  | ReturnType<typeof getMembers.start>
  | ReturnType<typeof getMembers.succeed>
  | ReturnType<typeof getMembers.fail>;
