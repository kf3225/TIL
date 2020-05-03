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
  makeStyles,
} from "@material-ui/core";
import { Favorite, Share } from "@material-ui/icons";

export interface Character {
  id: string;
  name: string;
  age: number;
  height?: number;
}

interface CharacterListProps {
  school?: string;
  characters: Character[];
}

const useStyles = makeStyles({
  container: {
    minWidth: 275,
    maxWidth: 275,
    margin: "0 2px",
    paddingLeft: "10%",
  },
  media: {
    height: 0,
    paddingTop: "56.25%", // 16:9
    backgroundColor: "gray",
  },
  card: {
    flexDirection: "column",
    marginBottom: "5%",
  },
});

const CharacterList: React.SFC<CharacterListProps> = ({
  school = "Unknown",
  characters,
}) => {
  const classes = useStyles();
  const [fav, isFav] = React.useState(false);

  const isPressFavoriteButton = () => {
    isFav(!fav);
  };

  return (
    <div className={classes.container}>
      {characters.map((c) => (
        <Card key={c.id} variant="outlined" className={classes.card}>
          <CardHeader
            avatar={<Avatar>{c.name.split("")[0]}</Avatar>}
            title={c.name}
          />
          <CardMedia className={classes.media}></CardMedia>
          <CardContent>
            <Typography variant="body2" color="textSecondary" component="p">
              Age : {c.age}
              <br />
              Height : {c.height ? c.height : "???"}
            </Typography>
          </CardContent>
          <CardActions disableSpacing>
            <IconButton
              aria-label="add to favorites"
              onClick={isPressFavoriteButton}
              key={c.id}
            >
              <Favorite color={fav ? "secondary" : "action"} />
            </IconButton>
            <IconButton aria-label="share">
              <Share />
            </IconButton>
          </CardActions>
        </Card>
      ))}
    </div>
  );
};

export default CharacterList;
