import { organizationURL } from "../../constants/urls";

export default class FetchOrganization {
  fetchApi = (id: any) => {
    if (id == null) {
      return fetch(organizationURL);
    } else {
      return fetch(organizationURL + "/" + id);
    }
  };
}
