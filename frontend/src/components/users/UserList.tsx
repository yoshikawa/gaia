import * as React from "react";

export interface Props {
  userData: any;
}

class UserList extends React.Component<Props> {
  render() {
    return (
      <ul>
        {(this.props.userData || []).map((userData: any, index: number) => {
          return (
            <li key={index}>
              <h6>{userData.name}</h6>
            </li>
          );
        })}
      </ul>
    );
  }
}

export default UserList;
