import * as React from "react";
import { Component } from "react";
import { Route, HashRouter, Switch } from "react-router-dom";
import { Provider } from "react-redux";
import store from "../../store/store";
import UserContainer from "../../container/users/allUserContainer";
import DescriptioContainer from "../../container/users/descriptionContainer";

class App extends Component {
  render() {
    return (
      <div className="App">
        <Provider store={store}>
          <HashRouter>
            <div>
              <Switch>
                <Route path="/" exact component={UserContainer} />
                <Route path="/:id" exact component={DescriptioContainer} />
              </Switch>
            </div>
          </HashRouter>
        </Provider>
      </div>
    );
  }
}

export default App;
