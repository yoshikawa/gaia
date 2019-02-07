import { FETCH_ORGANIZATIONS } from "../../constants/organizations";
import FetchOrganization from "../../services/application-services/FetchOrganization";

export const fetchOrganization = () => (dispatch: any) => {
  const api = new FetchOrganization();

  api
    .fetchApi(null)
    .then(response => response.json())
    .then(organization => {
      dispatch({
        type: FETCH_ORGANIZATIONS,
        payload: organization
      });
    });
};
