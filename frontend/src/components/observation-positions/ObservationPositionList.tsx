import * as React from "react";

export interface Props {
  observationPositionData: any;
}

class ObservationPositionList extends React.Component<Props> {
  render() {
    return (
      <ul>
        {(this.props.observationPositionData || []).map(
          (observationPositionData: any, index: number) => {
            return (
              <li key={index}>
                <h6>{observationPositionData.name}</h6>
              </li>
            );
          }
        )}
      </ul>
    );
  }
}

export default ObservationPositionList;
