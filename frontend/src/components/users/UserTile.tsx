import * as React from "react";
import { NavLink, withRouter, RouteComponentProps } from "react-router-dom";
// import * as PropTypes from 'prop-types';
export interface Props extends RouteComponentProps<any> {
  index: number;
  id: number;
  organizationID: number;
  name: string;
  email: string;
  country: string;
  administrator: boolean;
}
class UserTile extends React.Component<Props> {
  gotoContent = () => {
    this.props.history.push("/users/" + this.props.id);
  };
  render() {
    return (
      <div className="user-container">
        <div className="content">
          <div className="user-header">
            <h4>{this.props.name}</h4>
            <div className="source">email:{this.props.email}</div>
          </div>
          <div className="user-glance">
            <p>{this.props.country}</p>
            <p>{this.props.administrator}</p>
          </div>
          <button onClick={() => this.gotoContent()}>readMore >></button>
        </div>
      </div>
    );
  }
}

export default withRouter(UserTile);
