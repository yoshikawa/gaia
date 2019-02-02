import * as React from "react";
import * as ReactDOM from "react-dom";
import App from "./components/users/App";
import { HashRouter as Router } from "react-router-dom";

ReactDOM.render(
  <Router>
    <App />
  </Router>,
  document.getElementById("root") as HTMLElement
);
