import * as React from "react";
import { withRouter, RouteComponentProps } from "react-router-dom";
import Modal from "./Modal";

export interface Props extends RouteComponentProps<any> {
  fetchUserByID: (id: number) => any;
  userByID: any;
  item: any;
}

interface State {
  show: boolean;
}

class UserContent extends React.Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = {
      show: false
    };
  }
  componentWillMount() {
    const id = this.props.match.params.id;
    this.props.fetchUserByID(id);
  }
  gotoHome = () => {
    this.props.history.push("/");
  };
  modalHandler = () => {
    this.setState({ show: true });
  };
  closeHandler = () => {
    this.setState({ show: false });
  };

  render() {
    const userData = this.props.userByID.item || [];
    return (
      <div className="user-content">
        <Modal
          SourceUrl={userData.url}
          Show={this.state.show}
          closeHandler={this.closeHandler}
        />
        <div className="user-header">
          <h2>{userData.title}</h2>
          <h4>author : {userData.author}</h4>
          <h4>published At : {userData.publishedAt}</h4>
          <div className="under-line" />
        </div>
        <div className="user-glance">
          <h3>{userData.description}</h3>
        </div>
        <div className="view-more">
          <button onClick={() => this.modalHandler()}>
            View More from Source
          </button>
        </div>
        <div className="view-home">
          <button onClick={() => this.gotoHome()}>Home</button>
        </div>
      </div>
    );
  }
}

export default withRouter(UserContent);
