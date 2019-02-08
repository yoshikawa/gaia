import { FETCH_ORGANIZATION_BYID } from "../../constants/organizations";
import FetchOrganization from "../../services/application-services/FetchOrganization";

export const fetchOrganizationByID = (id: any) => (dispatch: any) => {
  const api = new FetchOrganization();

  api
    .fetchApi(id)
    .then(response => response.json())
    .then(organization => {
      dispatch({
        type: FETCH_ORGANIZATION_BYID,
        payload: organization
      });
    });
};
