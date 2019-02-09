import * as React from "react";
import { NavLink, withRouter, RouteComponentProps } from "react-router-dom";

export interface Props extends RouteComponentProps<any> {
  id: number;
  field_id: number;
  plant_id: number;
  name: string;
  point: number;
  altitude: number;
}

class ObservationPositionTile extends React.Component<Props> {
  gotoContent = () => {
    this.props.history.push("/observation-positions/" + this.props.id);
  };

  render() {
    return (
      <div className="observationPosition-container">
        <div className="content">
          <div className="observationPosition-header">
            <h4>{this.props.id}</h4>
            <div className="source">name: {this.props.name}</div>
          </div>
          <button onClick={() => this.gotoContent()}>readMore >></button>
        </div>
      </div>
    );
  }
}

export default withRouter(ObservationPositionTile);
