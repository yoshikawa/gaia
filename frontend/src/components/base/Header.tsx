import * as React from "react";
import { NavLink, withRouter, RouteComponentProps } from "react-router-dom";
// import * as PropTypes from 'prop-types';

interface HeaderProps {
  title: string;
}

export default class Header extends React.Component<HeaderProps> {
  render() {
    return (
      <header id="header" className="header">
        <div className="header-title">
          <ul className="title-name">
            <li className="name">
              <h1>
                <a href="/">{this.props.title}</a>
              </h1>
            </li>
          </ul>
        </div>
      </header>
    );
  }
}
