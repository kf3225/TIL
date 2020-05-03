import { parse } from "query-string";
import * as React from "react";
import { RouteComponentProps, withRouter } from "react-router";
import { Redirect } from "react-router-dom";
import { Button } from "@material-ui/core";
import { Home } from "@material-ui/icons";

import { characterData } from "../../characterData";
import Spinner from "../common/Spinner";
import { Helmet } from "react-helmet";
import CharacterList from "./CharacterList";
import "./index.css";

type CharacterListProps =　{} & RouteComponentProps<{ code: string }>;

const Characters: React.SFC<CharacterListProps> = ({
  history,
  location,
  match,
}) => {
  const codes = Object.keys(characterData);
  const targetCode = match.params.code;
  const isLoading = parse(location.search).loading === "true";

  return codes.includes(targetCode) ? (
    <>
      <Helmet>
        <title>Characters</title>
      </Helmet>
      <header>
        <h1>CharacterList / {characterData[targetCode].school}</h1>
      </header>
      {isLoading ? (
        <Spinner />
      ) : (
        <CharacterList key={characterData[targetCode].school}
          school={characterData[targetCode].school}
          characters={characterData[targetCode].players}
        />
      )}
      <Button onClick={() => history.push("/")}>
        <Home />
        Home
      </Button>
    </>
  ) : (
    <Redirect to="/" />
  );
};

export default withRouter(Characters);
