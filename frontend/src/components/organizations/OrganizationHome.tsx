import * as React from "react";
import OrganizationTile from "./OrganizationTile";

export interface Props {
  fetchOrganization: () => any;
  organization: any;
  id: number;
  name: string;
  index: number;
}
class OrganizationHome extends React.Component<Props> {
  componentWillMount() {
    this.props.fetchOrganization();
  }
  componentDidMount() {
    // this.props.changeLoaderState();
  }

  render() {
    let organizationData = this.props.organization.items || [];

    return (
      <div className="container">
        <div className="tiles-container">
          <div className="organization-menu">
            <h3>Organization list</h3>
          </div>
          {(organizationData || []).map((organizationData: any) => {
            return (
              <div key={organizationData.id}>
                <OrganizationTile
                  id={organizationData.id}
                  name={organizationData.name}
                />
              </div>
            );
          })}
        </div>
      </div>
    );
  }
}
export default OrganizationHome;
