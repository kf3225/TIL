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
  makeStyles,
} from "@material-ui/core";
import { Folder } from "@material-ui/icons";

import { characterData } from "../../characterData";

const useStyles = makeStyles({
  list: {
    minWidth: 275,
    maxWidth: 275,
    margin: "0 2px",
    paddingLeft: "5%",
  }
})

const codes = Object.keys(characterData);

const Home: React.FC<{}> = () => {
  const classes = useStyles()
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
        <List className={classes.list}>
          {codes.map((code) =>
            React.cloneElement(
              <ListItem
                button
                component={Link}
                to={`characters/${characterData[code].school}`}
              >
                <ListItemIcon>
                  <Folder />
                </ListItemIcon>
                <ListItemText primary={code} />
              </ListItem>,
              characterData[code]
            )
          )}
        </List>
      </div>
    </>
  );
};

export default Home;
