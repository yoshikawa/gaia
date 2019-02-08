import { FETCH_USERS } from "../../constants/users";
import FetchUser from "../../services/application-services/FetchUser";

export const fetchUser = () => (dispatch: any) => {
  const api = new FetchUser();

  api
    .fetchApi(null)
    .then(response => response.json())
    .then(user => {
      dispatch({
        type: FETCH_USERS,
        payload: user
      });
    });
};
