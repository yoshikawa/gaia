import * as React from "react";
import { Component } from "react";
import { Route, HashRouter, Switch } from "react-router-dom";
import { Provider } from "react-redux";
import store from "../store/store";
import Header from "../components/base/Header";
import UserContainer from "../container/users/allUserContainer";
import UserDescriptionContainer from "../container/users/descriptionContainer";
import OrganizationContainer from "../container/organizations/allOrganizationContainer";
import OrganizationDescriptionContainer from "../container/organizations/descriptionContainer";
import ObservationPositionsContainer from "../container/observation-positions/allObservationPositionContainer";
import ObservationPositionsDescriptionContainer from "../container/observation-positions/descriptionContainer";

export default class Routes extends Component {
  render() {
    return (
      <div className="App">
        <Header title="field sensing" />
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
                <Route
                  path="/observation-positions"
                  exact
                  component={ObservationPositionsContainer}
                />
                <Route
                  path="/observation-positions/:id"
                  exact
                  component={ObservationPositionsDescriptionContainer}
                />
              </Switch>
            </div>
          </HashRouter>
        </Provider>
      </div>
    );
  }
}
