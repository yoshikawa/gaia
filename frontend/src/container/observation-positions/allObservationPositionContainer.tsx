import { connect } from "react-redux";
import { fetchObservationPosition } from "../../actions/observation-positions/observationPositionActions";
import { bindActionCreators } from "redux";
import ObservationPositionHome from "../../components/observation-positions/ObservationPositionHome";

interface MyStateProps {
  observationPosition: Array<any>;
}

function mapStateToProps(state: MyStateProps) {
  return {
    observationPosition: state.observationPosition
  };
}

function matchDispatchToProps(dispatch: any) {
  return bindActionCreators(
    { fetchObservationPosition: fetchObservationPosition },
    dispatch
  );
}

export default connect(
  mapStateToProps,
  matchDispatchToProps
)(ObservationPositionHome);
