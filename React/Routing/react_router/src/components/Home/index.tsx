import * as React from "react";
import {Helmet} from "react-helmet";
import { Link } from "react-router-dom";
import {
  Container,
  List,
  ListSubheader,
  ListItem,
  ListItemIcon,
  ListItemText,
  Collapse,
} from "@material-ui/core";
import { Folder, ExpandLess, ExpandMore, StarBorder } from "@material-ui/icons";

import { characterData } from "../../characterData";

const codes = Object.keys(characterData);

const Home: React.SFC<{}> = () => {
  const [open, isOpen] = React.useState(true);

  const handClick = () => {
    isOpen(!open);
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
        <p>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</p>
        <p>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</p>
        <p>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</p>
        <p>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</p>
        <p>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</p>
        <p>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</p>
        <p>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</p>
        <p>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</p>
      </Container>
      <List
        component="nav"
        aria-labelledby="nested-list-subheader"
        subheader={
          <ListSubheader component="div" id="nested-list-subheader">
            List Items
          </ListSubheader>
        }
      >
        <ListItem button>
          <ListItemIcon onClick={handClick}>
            <Folder />
          </ListItemIcon>
          <ListItemText primary="Schools" />
          {open ? <ExpandLess /> : <ExpandMore />}
        </ListItem>
        <Collapse in={open} timeout="auto" unmountOnExit>
          <List component="div" disablePadding>
            {codes.map((code) => {
              // eslint-disable-next-line @typescript-eslint/no-unused-expressions
              <ListItem>
                <Link to={`/characters/${code}`}>
                  <ListItemIcon>
                    <StarBorder />
                  </ListItemIcon>
                  <ListItemText primary={characterData[code].school} />
                </Link>
              </ListItem>;
            })}
          </List>
        </Collapse>
      </List>
    </>
  );
};

export default Home;
