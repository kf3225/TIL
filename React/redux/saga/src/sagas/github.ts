import { getMembers } from "../actions/github";
import { getMembersFactory } from "../services/github/api";
import { call, put, takeLatest, fork, all } from "redux-saga/effects";
import * as Action from "../actions/githubConstants";

function* runGetMembers(action: ReturnType<typeof getMembers.start>) {
  const { companyName } = action.payload;

  try {
    const api = getMembersFactory();
    const users = yield call(api, companyName);

    yield put(getMembers.succeed({ companyName }, { users }));
  } catch (error) {
    yield put(getMembers.fail({ companyName }, error));
  }
}

export function* watchGetMembers() {
  yield takeLatest(Action.GET_MEMBERS_START, runGetMembers);
}

export default function* rootSaga() {
  yield all([fork(watchGetMembers)]);
}
