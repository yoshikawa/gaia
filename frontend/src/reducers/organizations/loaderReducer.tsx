import { LOADER_STATE } from "../../constants/organizations";

const initialState = {
  State: true
};

const loaderReducer = (state = initialState, action: any) => {
  switch (action.type) {
    case LOADER_STATE:
      return (Object as any).assign({}, state, {
        State: action.payload
      });
    default:
      return state;
  }
};

export default loaderReducer;
