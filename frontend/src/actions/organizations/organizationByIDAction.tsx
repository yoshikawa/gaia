import { FETCH_ORGANIZATIONS_BYID } from "../../constants/organizations";
import FetchOrganization from "../../services/application-Services/FetchOrganization";

export const fetchOrganizationByID = (id: any) => (dispatch: any) => {
  const api = new FetchOrganization();
  api
    .fetchApi()
    .then(response => response.json())
    .then(organization => {
      dispatch({
        type: FETCH_ORGANIZATIONS_BYID,
        payload: organization
      });
    });
};
