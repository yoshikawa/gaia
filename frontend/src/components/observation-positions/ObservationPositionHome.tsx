import * as React from "react";
import ObservationPositionTile from "./ObservationPositionTile";

export interface Props {
  fetchObservationPosition: () => any;
  observationPosition: any;
  id: number;
  field_id: number;
  plant_id: number;
  name: string;
  point: number;
  altitude: number;
}
class ObservationPositionHome extends React.Component<Props> {
  componentWillMount() {
    this.props.fetchObservationPosition();
  }
  componentDidMount() {
    // this.props.changeLoaderState();
  }

  render() {
    let observationPositionData = this.props.observationPosition.items || [];

    return (
      <div className="container">
        <div className="tiles-container">
          <div className="observationPosition-menu">
            <h3>ObservationPosition list</h3>
          </div>
          {(observationPositionData || []).map(
            (observationPositionData: any) => {
              return (
                <div key={observationPositionData.id}>
                  <ObservationPositionTile
                    id={observationPositionData.id}
                    field_id={observationPositionData.field_id}
                    plant_id={observationPositionData.plant_id}
                    name={observationPositionData.name}
                    point={observationPositionData.point}
                    altitude={observationPositionData.altitude}
                  />
                </div>
              );
            }
          )}
        </div>
      </div>
    );
  }
}
export default ObservationPositionHome;
