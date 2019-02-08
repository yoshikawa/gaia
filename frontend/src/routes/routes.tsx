import * as React from "react";
import { Component } from "react";
import { Route, HashRouter, Switch } from "react-router-dom";
import { Provider } from "react-redux";
import store from "../store/store";
import UserContainer from "../container/users/allUserContainer";
import UserDescriptionContainer from "../container/users/descriptionContainer";
import OrganizationContainer from "../container/organizations/allOrganizationContainer";
import OrganizationDescriptionContainer from "../container/organizations/descriptionContainer";

export default class Routes extends Component {
  render() {
    return (
      <div className="App">
        <Provider store={store}>
          <HashRouter>
            <div>
              <Switch>
                <Route path="/users" exact component={UserContainer} />
                <Route
                  path="/users/:id"
                  exact
                  component={UserDescriptionContainer}
                />
                <Route
                  path="/organizations"
                  exact
                  component={OrganizationContainer}
                />
                <Route
                  path="/organizations/:id"
                  exact
                  component={OrganizationDescriptionContainer}
                />
              </Switch>
            </div>
          </HashRouter>
        </Provider>
      </div>
    );
  }
}
