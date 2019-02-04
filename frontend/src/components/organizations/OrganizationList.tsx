import * as React from "react";

export interface Props {
  organizationData: any;
}

class OrganizationList extends React.Component<Props> {
  render() {
    return (
      <ul>
        {(this.props.organizationData || []).map(
          (organizationData: any, k: number) => {
            return (
              <li key={k}>
                <h6>{organizationData.name}</h6>
              </li>
            );
          }
        )}
      </ul>
    );
  }
}

export default OrganizationList;
