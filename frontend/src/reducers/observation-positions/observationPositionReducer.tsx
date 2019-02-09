import { FETCH_OBSERVATION_POSITIONS } from "../../constants/observationPositions";

const initialState = {
  items: []
};

const observationPositionReducer = (state = initialState, action: any) => {
  switch (action.type) {
    case FETCH_OBSERVATION_POSITIONS:
      return (Object as any).assign({}, state, {
        items: action.payload
      });
    default:
      return state;
  }
};

export default observationPositionReducer;
