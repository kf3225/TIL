import * as React from "react";
import { Helmet } from "react-helmet";
import { Link } from "react-router-dom";
import {
  Container,
  List,
  ListItem,
  ListItemText,
  Typography,
  ListItemIcon,
} from "@material-ui/core";
import { Folder, ExpandLess, ExpandMore, StarBorder } from "@material-ui/icons";

import { characterData } from "../../characterData";

const codes = Object.keys(characterData);

const Home: React.SFC<{}> = () => {
  const [dense, setDense] = React.useState(false);
  const [secondary, setSecondary] = React.useState(false);

  const generate = (element: React.ReactElement) => {
    return codes.map((code) =>
      React.cloneElement(element, {
        key: characterData[code].school,
      })
    );
  };

  return (
    <>
      <Helmet>
        <title>Characters</title>
      </Helmet>
      <header>
        <h1>Home</h1>
      </header>
      <Container>
        <p>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</p>
      </Container>
      <Typography variant="h6">Schools</Typography>
      <div>
        <List dense={dense}>
          {codes.map((code) =>
            React.cloneElement(
              <ListItem button component={Link} to={`characters/${characterData[code].school}`}>
                <ListItemIcon>
                  <Folder />
                </ListItemIcon>
                <ListItemText primary={code} />
              </ListItem>,
              {
                key: code,
              }
            )
          )}
        </List>
      </div>
    </>
  );
};

export default Home;
