import { combineReducers } from "redux";
import loaderReducer from "./loaderReducer";
import userReducer from "./users/userReducer";
import userByIDReducer from "./users/userByIDReducer";
import organizationReducer from "./organizations/organizationReducer";
import organizationByIDReducer from "./organizations/organizationByIDReducer";
import { routerReducer } from "react-router-redux";

const allRootReducers = combineReducers({
  routing: routerReducer,
  loader: loaderReducer,
  user: userReducer,
  organization: organizationReducer,
  userByID: userByIDReducer,
  organizationByID: organizationByIDReducer
});

export default allRootReducers;
