import * as React from "react";
import UserTile from "./UserTile";

export interface Props {
  fetchUser: () => any;
  user: any;
  id: number;
  organizationID: number;
  name: string;
  email: string;
  country: string;
  administrator: boolean;
  index: number;
}
class UserHome extends React.Component<Props> {
  componentWillMount() {
    this.props.fetchUser();
  }
  componentDidMount() {
    // this.props.changeLoaderState();
  }

  render() {
    let userData = this.props.user.items || [];

    return (
      <div className="container">
        <div className="tiles-container">
          <div className="user-menu">
            <h3>User list</h3>
          </div>
          {(userData || []).map((userData: any, index: number) => {
            return (
              <div key={index}>
                <UserTile
                  index={index}
                  id={userData.id}
                  organizationID={userData.organizationID}
                  name={userData.name}
                  email={userData.email}
                  country={userData.country}
                  administrator={userData.administrator}
                />
              </div>
            );
          })}
        </div>
      </div>
    );
  }
}
export default UserHome;
