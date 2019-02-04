import { connect } from "react-redux";
import { fetchOrganizationByID } from "../../actions/organizations/organizationByIDAction";
import { bindActionCreators } from "redux";
import OrganizationContent from "../../components/organizations/OrganizationContent";

interface MyStateProps {
  organizationByID: any;
}

function mapStateToProps(state: MyStateProps) {
  return {
    organizationByID: state.organizationByID
  };
}

function matchDispatchToProps(dispatch: any) {
  return bindActionCreators(
    { fetchOrganizationByID: fetchOrganizationByID },
    dispatch
  );
}

export default connect(
  mapStateToProps,
  matchDispatchToProps
)(OrganizationContent) as any;
