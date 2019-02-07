import { FETCH_ORGANIZATIONS } from "../../constants/organizations";

const initialState = {
  items: []
};

const organizationReducer = (state = initialState, action: any) => {
  switch (action.type) {
    case FETCH_ORGANIZATIONS:
      return (Object as any).assign({}, state, {
        items: action.payload
      });
    default:
      return state;
  }
};

export default organizationReducer;
