import { combineReducers } from "redux";
import observationPositionReducer from "./observationPositionReducer";
import loaderReducer from "../loaderReducer";
import observationPositionByIDReducer from "./observationPositionByIDReducer";
import { routerReducer } from "react-router-redux";

const allObservationPositionReducers = combineReducers({
  observationPosition: observationPositionReducer,
  routing: routerReducer,
  loader: loaderReducer,
  observationPositionByID: observationPositionByIDReducer
});

export default allObservationPositionReducers;
