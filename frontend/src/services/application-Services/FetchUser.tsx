import { userURL } from "../../constants/urls";

export default class FetchUser {
  fetchApi = (id: any) => {
    if (id == null) {
      return fetch(userURL);
    } else {
      return fetch(userURL + "/" + id);
    }
  };
}
