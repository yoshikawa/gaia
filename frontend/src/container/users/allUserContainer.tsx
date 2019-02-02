import { connect } from "react-redux";
import { fetchUser } from "../../actions/users/userActions";
import { bindActionCreators } from "redux";
import UserHome from "../../components/users/UserHome";
interface MyStateProps {
  user: Array<any>;
}

function mapStateToProps(state: MyStateProps) {
  return {
    user: state.user
  };
}

function matchDispatchToProps(dispatch: any) {
  return bindActionCreators({ fetchUser: fetchUser }, dispatch);
}

export default connect(
  mapStateToProps,
  matchDispatchToProps
)(UserHome);
