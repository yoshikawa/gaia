import * as React from "react";
import { withRouter, RouteComponentProps } from "react-router-dom";
import Header from "../base/Header";

export interface Props extends RouteComponentProps<any> {
  fetchOrganizationByID: (id: number) => any;
  organizationByID: any;
  item: any;
}

class OrganizationContent extends React.Component<Props> {
  constructor(props: Props) {
    super(props);
  }

  componentWillMount() {
    const id = this.props.match.params.id;
    this.props.fetchOrganizationByID(id);
  }

  gotoHome = () => {
    this.props.history.push("/organizations");
  };

  render() {
    const organizationData = this.props.organizationByID.item || [];

    return (
      <div className="organization-content">
        <Header title="gaia" />
        <div className="organization-header">
          <h2>{organizationData.name}</h2>
          <div className="under-line" />
        </div>
        <div className="view-home">
          <button onClick={() => this.gotoHome()}>Home</button>
        </div>
      </div>
    );
  }
}

export default withRouter(OrganizationContent);
