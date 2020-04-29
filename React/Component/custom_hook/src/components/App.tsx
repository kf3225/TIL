import React, { FC } from "react";
import { Card, Statistic, Button, Icon } from "semantic-ui-react";
import "semantic-ui-css/semantic.min.css";
import "../App.css";

interface AppProps {
  timeLeft: number;
  keep: boolean;
  reset: () => void;
  stopOrRestart: () => void;
}

const AppComponent: FC<AppProps> = ({ timeLeft, keep, reset, stopOrRestart }) => (
  <div className="container">
    <div className="two button ui">
      <header>
        <h1>Timer</h1>
      </header>
      <Card>
        <Statistic>
          <Statistic.Label>Time</Statistic.Label>
          <Statistic.Value>{timeLeft}</Statistic.Value>
        </Statistic>
        <Card.Content>
          <Button color="red" onClick={reset}>
            <Icon name="redo alternate" />
            reset
          </Button>
          <Button color={keep ? "yellow" : "green"} onClick={stopOrRestart}>
            <Icon name={keep ? "stop circle" : "refresh"} />
            {keep ? "stop" : "restart"}
          </Button>
        </Card.Content>
      </Card>
    </div>
  </div>
);

export default AppComponent;
