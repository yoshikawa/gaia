import { FETCH_OBSERVATION_POSITIONS } from "../../constants/observationPositions";
import FetchObservationPosition from "../../services/application-services/FetchObservationPosition";

export const fetchObservationPosition = () => (dispatch: any) => {
  const api = new FetchObservationPosition();

  api
    .fetchApi(null)
    .then(response => response.json())
    .then(observationPosition => {
      dispatch({
        type: FETCH_OBSERVATION_POSITIONS,
        payload: observationPosition
      });
    });
};
