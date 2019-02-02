import { connect } from "react-redux";
import { fetchUserByID } from "../../actions/users/userByIDAction";
import { bindActionCreators } from "redux";
import UserContent from "../../components/users/UserContent";

interface MyStateProps {
  userByid: any;
}

function mapStateToProps(state: MyStateProps) {
  return {
    userByid: state.userByid
  };
}

function matchDispatchToProps(dispatch: any) {
  return bindActionCreators({ fetchUserByID: fetchUserByID }, dispatch);
}

export default connect(
  mapStateToProps,
  matchDispatchToProps
)(UserContent) as any;
