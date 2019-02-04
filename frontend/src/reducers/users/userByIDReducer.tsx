import { FETCH_USERS_BYID } from "../../constants/users";
const initialState = {
  item: {}
};
const userByidReducer = (state = initialState, action: any) => {
  switch (action.type) {
    case FETCH_USERS_BYID:
      return (Object as any).assign({}, state, {
        item: action.payload
      });
    default:
      return state;
  }
};
export default userByidReducer;
