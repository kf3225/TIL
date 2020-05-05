import { Reducer } from "redux";
import { CounterAction, CounterActionType } from "./actions/counter";

export interface CounterState {
  count: number;
}

export const initialState: CounterState = { count: 0 };

const counterReducer: Reducer<CounterState, CounterAction> = (
  state: CounterState = initialState,
  action: CounterAction
): CounterState => {
  switch (action.type) {
    case CounterActionType.ADD:
      const addAns = state.count + (action.amount || 0);
      return {
        ...state,
        count: addAns >= 0 ? addAns : 0,
      };
    case CounterActionType.DECREMENT:
      const decrementAns = state.count - 1
      return {
        ...state,
        count: decrementAns >= 0 ? decrementAns : 0,
      };
    case CounterActionType.INCREMENT:
      return {
        ...state,
        count: state.count + 1,
      };
    default: {
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
      const _: never = action.type;

      return state;
    }
  }
};

export default counterReducer;
