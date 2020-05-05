import React, { FC } from "react";
import {
  Card,
  CardContent,
  Typography,
  Button,
} from "@material-ui/core";

export interface CounterProps {
  count?: number;
  add?: (amount: number) => void;
  decrement: () => void;
  increment: () => void;
}

const Counter: FC<CounterProps> = ({
  count = 0,
  add = () => undefined,
  decrement = () => undefined,
  increment = () => undefined,
}) => (
  <Card>
    <Typography variant="h3">Counter</Typography>
    <Typography variant="h4">{count}</Typography>
    <CardContent>
      <Button variant="contained" color="primary" onClick={decrement}>
        -1
      </Button>
      <Button variant="contained" color="secondary" onClick={increment}>
        +1
      </Button>
      <Button variant="contained" color="primary" onClick={() => add(10)}>
        +10
      </Button>
      <Button variant="contained" color="secondary" onClick={() => add(-10)}>
        -10
      </Button>
    </CardContent>
  </Card>
);

export default Counter;
