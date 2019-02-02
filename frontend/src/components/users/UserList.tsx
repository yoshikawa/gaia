import * as React from "react";

export interface Props {
  userData: any;
}

class UserList extends React.Component<Props> {
  render() {
    return (
      <ul>
        {(this.props.userData || []).map((userData: any, k: number) => {
          return (
            <li key={k}>
              <h6>{userData.title}</h6>
            </li>
          );
        })}
      </ul>
    );
  }
}

export default UserList;
