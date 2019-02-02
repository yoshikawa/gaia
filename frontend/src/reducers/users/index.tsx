import { combineReducers } from "redux";
import userReducer from "./userReducer";
import loaderReducer from "./loaderReducer";
import userByIDReducer from "./userByIDReducer";
import { routerReducer } from "react-router-redux";

const allReducers = combineReducers({
  user: userReducer,
  routing: routerReducer,
  loader: loaderReducer,
  userByID: userByIDReducer
});

export default allReducers;
