import * as React from "react";
import { CircularProgress } from "@material-ui/core";
import { Skeleton } from "@material-ui/lab";

const Spinner: React.SFC = () => (
  <Skeleton animation="wave">
    <CircularProgress />
    Now Loading...
  </Skeleton>
);

export default Spinner;
