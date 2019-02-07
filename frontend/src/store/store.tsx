import { createStore } from "redux";
import thunk from "redux-thunk";
import rootReducer from "../reducers/index";
import { applyMiddleware } from "redux";

const initialState = {};
const middleware = [thunk];
const store = createStore(
  rootReducer,
  initialState,
  applyMiddleware(...middleware)
);

export default store;
