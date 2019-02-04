import { FETCH_ORGANIZATIONS } from "../../constants/organizations";
import FetchUser from "../../services/application-Services/FetchUser";

export const FetchOrganization = () => (dispatch: any) => {
  const api = new FetchOrganization();
  api
    .fetchApi()
    .then(response => response.json())
    .then(organization => {
      dispatch({
        type: FETCH_ORGANIZATIONS,
        payload: organization
      });
    });
};
