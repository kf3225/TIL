import * as React from "react";
import { Redirect, Route, Switch } from "react-router-dom";
import "./App.css";

import Characters from "./components/Characters";
import Home from "./components/Home";

const App: React.FC<{}> = () => (
  <div className="container">
      <Switch>
        <Route path="/characters/:code" component={Characters} />
        <Route path="/" component={Home} />
        <Redirect to="/" />
      </Switch>
  </div>
);

export default App;
