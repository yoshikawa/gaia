import { observationPositionURL } from "../../constants/urls";

export default class FetchObservationPosition {
  fetchApi = (id: any) => {
    if (id == null) {
      return fetch(observationPositionURL);
    } else {
      return fetch(observationPositionURL + "/" + id);
    }
  };
}
