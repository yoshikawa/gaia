import * as React from "react";
import * as ReactDOM from "react-dom";
import App from "./App";
import { HashRouter as Router } from "react-router-dom";
import "./stylesheets/style.scss";

ReactDOM.render(
  <Router>
    <App />
  </Router>,
  document.getElementById("root") as HTMLElement
);
