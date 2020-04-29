import React, { FC, useState, useEffect } from "react";
import { Button, Statistic, Card, Icon } from "semantic-ui-react";
import "semantic-ui-css/semantic.min.css";
import "./App.css";

let id: NodeJS.Timer;

const App: FC = () => {
  const LIMIT = 60;

  const [timeLeft, setTimeLeft] = useState(LIMIT);
  const [keep, setKeep] = useState(true);

  const tick = () => {
    setTimeLeft((prevState) => (prevState === 0 ? LIMIT : timeLeft - 1));
  };

  const reset = () => {
    setTimeLeft(LIMIT);
  };

  const stopOrRestartTimer = () => {
    setKeep(!keep);
  };

  useEffect(() => {
    if (keep) {
      id = setInterval(tick, 1000);
    }

    console.log(id);
    return () => clearInterval(id);
  });

  return (
    <>
      <div className="container">
        <header>
          <h2>Timer</h2>
        </header>
        <div className="two button ui">
          <Statistic>
            <Statistic.Label>Time</Statistic.Label>
            <Statistic.Value>{timeLeft}</Statistic.Value>
          </Statistic>
          <Card.Content>
            <Button color="red" onClick={reset}>
              <Icon name="redo alternate" />
              reset
            </Button>
            <Button color={keep ? "yellow" : "green"} onClick={stopOrRestartTimer}>
              <Icon name={keep ? "stop" : "refresh"} />
              {keep ? "stop" : "restart"}
            </Button>
          </Card.Content>
        </div>
      </div>
    </>
  );
};

export default App;
