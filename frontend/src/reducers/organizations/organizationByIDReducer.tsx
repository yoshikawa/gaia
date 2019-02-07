import { FETCH_ORGANIZATION_BYID } from "../../constants/organizations";

const initialState = {
  item: {}
};

const organizationByidReducer = (state = initialState, action: any) => {
  switch (action.type) {
    case FETCH_ORGANIZATION_BYID:
      return (Object as any).assign({}, state, {
        item: action.payload
      });
    default:
      return state;
  }
};

export default organizationByidReducer;
