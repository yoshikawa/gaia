import * as React from "react";
import OrganizationTile from "./OrganizationTile";
import OrganizationList from "./OrganizationList";

export interface Props {
  fetchOrganization: () => any;
  organization: any;
  id: number;
  organizationID: number;
  name: string;
  email: string;
  country: string;
  administrator: boolean;
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
        <div className="side-bar">
          <div className="shortLine">
            <h4>Organization Name</h4>
          </div>
          <div className="OrganizationList">
            <OrganizationList organizationData={organizationData} />
          </div>
        </div>

        <div className="tiles-container">
          <div className="organization-menu">
            <h3>Organization list</h3>
          </div>
          {(organizationData || []).map(
            (organizationData: any, index: number) => {
              return (
                <div key={index}>
                  <OrganizationTile
                    index={index}
                    id={organizationData.id}
                    name={organizationData.name}
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
export default OrganizationHome;
