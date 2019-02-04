import * as React from "react";
import { Component } from "react";
import { Route, HashRouter, Switch } from "react-router-dom";
import { Provider } from "react-redux";
import store from "../../store/store";
import OrganizationContainer from "../../container/organizations/allOrganizationContainer";
import DescriptionContainer from "../../container/organizations/descriptionContainer";

class App extends Component {
  render() {
    return (
      <div className="App">
        <Provider store={store}>
          <HashRouter>
            <div>
              <Switch>
                <Route
                  path="/organizations"
                  exact
                  component={OrganizationContainer}
                />
                <Route
                  path="/organizations/:id"
                  exact
                  component={DescriptionContainer}
                />
              </Switch>
            </div>
          </HashRouter>
        </Provider>
      </div>
    );
  }
}

export default App;
