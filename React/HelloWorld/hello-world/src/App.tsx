import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import { render } from '@testing-library/react';

class App extends Component {
  render() {
    const logoAttrs = {
      alt: "logo",
      className: "App-logo",
      src: logo
    };
    const title = "";
    const targets = ["World", "Keisuke"];

    return (
      <div className="App">
        <header className="App-header">
          {
            // comments
          }
          <img {...logoAttrs} />
        {title ? <p>{title}</p> : <p>Hello React World!!</p>}
        {targets.map(target => (
          <p>Hello, {target}!</p>
        ))}
        </header>
      </div>
    );
  }
}

export default App;
