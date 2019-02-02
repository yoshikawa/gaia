import { FETCH_USERS_BYID } from "../../constants/users";
import FetchUser from "../../services/application-Services/FetchUser";

export const fetchUserByID = (id: any) => (dispatch: any) => {
  const api = new FetchUser();
  api
    .fetchApi()
    .then(response => response.json())
    .then(user => {
      dispatch({
        type: FETCH_USERS_BYID,
        payload: user
      });
    });
};
