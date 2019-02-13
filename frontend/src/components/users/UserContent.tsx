import * as React from "react";
import { withRouter, RouteComponentProps } from "react-router-dom";
import Header from "../base/Header";

export interface Props extends RouteComponentProps<any> {
  fetchUserByID: (id: number) => any;
  userByID: any;
  item: any;
}

class UserContent extends React.Component<Props> {
  constructor(props: Props) {
    super(props);
  }

  componentWillMount() {
    const id = this.props.match.params.id;
    this.props.fetchUserByID(id);
  }

  gotoHome = () => {
    this.props.history.push("/users");
  };

  render() {
    const userData = this.props.userByID.item || [];
    return (
      <div className="user-content">
        <Header title="gaia" />
        <div className="user-header">
          <h2>{userData.id}</h2>
          <h4>name : {userData.name}</h4>
          <div className="under-line" />
        </div>
        <div className="view-home">
          <button onClick={() => this.gotoHome()}>Home</button>
        </div>
      </div>
    );
  }
}

export default withRouter(UserContent);
