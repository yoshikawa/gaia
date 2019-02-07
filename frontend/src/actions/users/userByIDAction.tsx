import { FETCH_USER_BYID } from "../../constants/users";
import FetchUser from "../../services/application-services/FetchUser";

export const fetchUserByID = (id: any) => (dispatch: any) => {
  var api = new FetchUser();

  api
    .fetchApi(id)
    .then(response => response.json())
    .then(user => {
      dispatch({
        type: FETCH_USER_BYID,
        payload: user
      });
    });
};
