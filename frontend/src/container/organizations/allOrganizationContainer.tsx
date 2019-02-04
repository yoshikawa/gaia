import { connect } from "react-redux";
import { fetchOrganization } from "../../actions/organizations/organizationActions";
import { bindActionCreators } from "redux";
import OrganizationHome from "../../components/organizations/OrganizationHome";
interface MyStateProps {
  organization: Array<any>;
}

function mapStateToProps(state: MyStateProps) {
  return {
    organization: state.organization
  };
}

function matchDispatchToProps(dispatch: any) {
  return bindActionCreators({ fetchOrganization: fetchOrganization }, dispatch);
}

export default connect(
  mapStateToProps,
  matchDispatchToProps
)(OrganizationHome);
