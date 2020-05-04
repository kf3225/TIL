import React, { FC } from "react";
import { Card, CardContent, Typography, Button } from "@material-ui/core";

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
    <CardContent>
      <Typography component="h1">count</Typography>
      <Typography component="h5">{count}</Typography>
      <Button onClick={decrement}>-1</Button>
      <Button onClick={increment}>+1</Button>
      <Button onClick={() => add(10)}>+10</Button>
    </CardContent>
  </Card>
);

export default Counter;
