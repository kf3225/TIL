import * as React from "react";
import {
  Card,
  CardActions,
  CardContent,
  Typography,
  CardHeader,
  Avatar,
  IconButton,
  CardMedia,
  Icon,
} from "@material-ui/core";
import { Favorite, Share, Person } from "@material-ui/icons";

export interface Character {
  name: string;
  age: number;
  height?: number;
}

interface CharacterListProps {
  school: string;
  characters: Character[];
}

const CharacterList: React.SFC<CharacterListProps> = ({
  school = "Unknown",
  characters,
}) => (
  <>
    {characters.map((c, i) => (
      // eslint-disable-next-line @typescript-eslint/no-unused-expressions
      <Card id={i.toString()}>
        <CardHeader
          avatar={<Avatar>{c.name.split("")}</Avatar>}
          title={c.name}
        />
        <CardMedia>
          <Icon aria-label="person">
            <Person />
          </Icon>
        </CardMedia>
        <CardContent>
          <Typography variant="body2" color="textSecondary" component="h1">
            {c.name}
          </Typography>
          <Typography variant="body2" color="textSecondary" component="p">
            {c.age}
          </Typography>
          <Typography variant="body2" color="textSecondary" component="p">
            {c.height ? c.height : "???"}
          </Typography>
        </CardContent>
        <CardActions disableSpacing>
          <IconButton aria-label="add to favorites">
            <Favorite />
          </IconButton>
          <IconButton aria-label="share">
            <Share />
          </IconButton>
        </CardActions>
      </Card>
    ))}
  </>
);

export default CharacterList;
