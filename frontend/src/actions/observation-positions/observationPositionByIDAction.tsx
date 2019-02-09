import { FETCH_OBSERVATION_POSITION_BYID } from "../../constants/observationPositions";
import FetchObservationPosition from "../../services/application-services/FetchObservationPosition";

export const fetchObservationPositionByID = (id: any) => (dispatch: any) => {
  const api = new FetchObservationPosition();

  api
    .fetchApi(id)
    .then(response => response.json())
    .then(observationPosition => {
      dispatch({
        type: FETCH_OBSERVATION_POSITION_BYID,
        payload: observationPosition
      });
    });
};
