import { organizationURL } from "../../constants/urls";

class FetchOrganization {
  fetchApi = () => {
    return fetch(organizationURL);
  };
}

export default FetchOrganization;
