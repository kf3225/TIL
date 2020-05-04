import React, { FC } from "react";
import {
  red,
  orange,
  yellow,
  green,
  teal,
  blue,
  purple,
  pink,
  brown,
  grey,
} from "@material-ui/core/colors";
import { Container, makeStyles } from "@material-ui/core";
import { Brightness1 } from "@material-ui/icons";


const useStyles = makeStyles({
  beadsbox: {
    marginTop: "40px",
  },
});

const range = (n: number) => (n < 0 ? [] : Array.from(Array(n), (_, i) => i));
const colorArr: string[] = [
  red[500],
  orange[500],
  yellow[500],
  green[500],
  teal[500],
  blue[500],
  purple[500],
  pink[500],
  brown[500],
  grey[500],
];

export interface ColorfulBeadsProps {
  count?: number;
}

const ColorfulBeads: FC<ColorfulBeadsProps> = ({ count = 0 }) => {
  const classes = useStyles();

  return (
    <Container className={classes.beadsbox}>
      {range(count).map((i: number) => (
        <Brightness1
          style={{ color: colorArr[i % colorArr.length] }}
          key={i.toString()}
        />
      ))}
    </Container>
  );
};

export default ColorfulBeads;
