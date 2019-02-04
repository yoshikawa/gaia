import { FETCH_USERS } from "../../constants/users";
const initialState = {
  items: []
};

const userReducer = (state = initialState, action: any) => {
  switch (action.type) {
    case FETCH_USERS:
      return (Object as any).assign({}, state, {
        items: action.payload
      });
    default:
      return state;
  }
};
export default userReducer;
