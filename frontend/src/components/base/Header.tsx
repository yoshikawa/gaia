import * as React from "react";
import { Link } from "react-router-dom";
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
                <Link to="/">{this.props.title}</Link>
              </h1>
            </li>
          </ul>
        </div>
      </header>
    );
  }
}
