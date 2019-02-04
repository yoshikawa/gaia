import * as React from "react";
import { NavLink, withRouter, RouteComponentProps } from "react-router-dom";
// import * as PropTypes from 'prop-types';

export interface Props extends RouteComponentProps<any> {
  index: number;
  id: number;
  name: string;
}

class OrganizationTile extends React.Component<Props> {
  gotoContent = () => {
    this.props.history.push("/organizations/" + this.props.id);
  };
  render() {
    return (
      <div className="organization-container">
        <div className="content">
          <div className="organization-header">
            <h4>{this.props.id}</h4>
            <div className="source">name:{this.props.name}</div>
          </div>
          <button onClick={() => this.gotoContent()}>readMore >></button>
        </div>
      </div>
    );
  }
}

export default withRouter(OrganizationTile);
