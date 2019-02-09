import { connect } from "react-redux";
import { fetchObservationPositionByID } from "../../actions/observation-positions/observationPositionByIDAction";
import { bindActionCreators } from "redux";
import ObservationPositionContent from "../../components/observation-positions/ObservationPositionContent";

interface MyStateProps {
  observationPositionByID: any;
}

function mapStateToProps(state: MyStateProps) {
  return {
    observationPositionByID: state.observationPositionByID
  };
}

function matchDispatchToProps(dispatch: any) {
  return bindActionCreators(
    { fetchObservationPositionByID: fetchObservationPositionByID },
    dispatch
  );
}

export default connect(
  mapStateToProps,
  matchDispatchToProps
)(ObservationPositionContent) as any;
