import React, { FC } from "react";

import CharacterList, { Character } from "./CharacterList";
import "./App.css";

const App: FC<{}> = () => {
  const characters: Character[] = [
    {
      id: 1,
      name: "AAA",
      age: 10,
      height: 150,
    },
    {
      id: 2,
      name: "BBB",
      age: 20,
      height: 160,
    },
    {
      id: 3,
      name: "CCC",
      age: 30,
    },
  ];

  return (
    <div className="container">
      <header>
        <h1>Characters</h1>
      </header>
      <CharacterList characterList={characters} />
    </div>
  );
};

export default App;
