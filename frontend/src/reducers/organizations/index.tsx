import { combineReducers } from "redux";
import organizationReducer from "./organizationReducer";
import loaderReducer from "./loaderReducer";
import organizationByIDReducer from "./organizationByIDReducer";
import { routerReducer } from "react-router-redux";

const allReducers = combineReducers({
  organization: organizationReducer,
  routing: routerReducer,
  loader: loaderReducer,
  organizationByID: organizationByIDReducer
});

export default allReducers;
