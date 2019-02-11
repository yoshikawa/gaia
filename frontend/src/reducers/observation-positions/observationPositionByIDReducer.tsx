import { FETCH_OBSERVATION_POSITION_BYID } from "../../constants/observationPositions";

const initialState = {
  item: {}
};

const observationPositionByIDReducer = (state = initialState, action: any) => {
  switch (action.type) {
    case FETCH_OBSERVATION_POSITION_BYID:
      return (Object as any).assign({}, state, {
        item: action.payload
      });
    default:
      return state;
  }
};

export default observationPositionByIDReducer;
