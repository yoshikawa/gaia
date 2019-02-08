import * as React from "react";
import { withRouter, RouteComponentProps } from "react-router-dom";
import Modal from "./Modal";

export interface Props extends RouteComponentProps<any> {
  fetchOrganizationByID: (id: number) => any;
  organizationByID: any;
  item: any;
}

interface State {
  show: boolean;
}

class OrganizationContent extends React.Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = {
      show: false
    };
  }

  componentWillMount() {
    const id = this.props.match.params.id;
    this.props.fetchOrganizationByID(id);
  }

  gotoHome = () => {
    this.props.history.push("/organizations");
  };

  closeHandler = () => {
    this.setState({ show: false });
  };

  render() {
    const organizationData = this.props.organizationByID.item || [];

    return (
      <div className="organization-content">
        <Modal
          SourceUrl={organizationData.url}
          Show={this.state.show}
          closeHandler={this.closeHandler}
        />
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
