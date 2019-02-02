import { userURL } from "../../constants/urls";

class FetchUser {
  fetchApi = () => {
    return fetch(userURL);
  };
}

export default FetchUser;
