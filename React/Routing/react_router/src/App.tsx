import * as React from "react";
import { Redirect, Route, Switch, BrowserRouter } from "react-router-dom";
import "./App.css";

import Characters from "./components/Characters";
import Home from "./components/Home";

const App: React.SFC<{}> = () => (
  <div className="container">
    <BrowserRouter>
      <Switch>
        <Route path="/characters/:code" component={Characters} />
        <Route path="/" component={Home} />
        <Redirect to="/" />
      </Switch>
    </BrowserRouter>
  </div>
);

export default App;
