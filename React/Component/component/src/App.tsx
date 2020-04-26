import React, { Component } from 'react';
import CharacterList, {Character} from './CharacterList';
import './App.css';

class App extends Component {
  render() {
    const characters: Character[] = [
      {
        id: 1,
        name: 'AAA',
        age: 20,
        height: 160,
      },
      {
        id: 2,
        name: 'BBB',
        age: 30,
        height: 170,
      },
      {
        id: 3,
        name: 'CCC',
        age: 40,
      },
    ];

    return(
      <div className="container">
        <header>
          <h1>characters</h1>
        </header>
        <CharacterList school="ABC" characters={characters} />
      </div>
    );
  }
}

export default App;
