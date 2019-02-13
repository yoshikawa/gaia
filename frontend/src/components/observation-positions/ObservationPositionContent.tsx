import * as React from "react";
import { withRouter, RouteComponentProps } from "react-router-dom";
import Header from "../base/Header";

export interface Props extends RouteComponentProps<any> {
  fetchObservationPositionByID: (id: number) => any;
  observationPositionByID: any;
  item: any;
}

class ObservationPositionContent extends React.Component<Props> {
  constructor(props: Props) {
    super(props);
  }

  componentWillMount() {
    const id = this.props.match.params.id;
    this.props.fetchObservationPositionByID(id);
  }

  gotoHome = () => {
    this.props.history.push("/observation-positions");
  };

  render() {
    const observationPositionsData =
      this.props.observationPositionByID.item || [];

    return (
      <div className="observationPosition-content">
        <Header title="gaia" />
        <div className="observationPosition-header">
          <h2>{observationPositionsData.name}</h2>
          <div className="under-line" />
        </div>
        <div className="view-home">
          <button onClick={() => this.gotoHome()}>Home</button>
        </div>
      </div>
    );
  }
}

export default withRouter(ObservationPositionContent);
