import React, { FC } from "react";
import { Header, Icon, Item } from "semantic-ui-react";

export interface Character {
  id: number;
  name: string;
  age: number;
  height?: number;
}

interface CharacterListProps {
  school: string;
  characterList: Character[];
}

const CharacterList: FC<CharacterListProps> = ({
  school = "unknown",
  characterList,
}) => (
  <>
    <Header as="h2">{school}</Header>
    <Item.Group>
      {characterList.map((c) => (
        <Item>
          <Icon name="user circle" size="huge" />
          <Item.Content>
            <Item.Header>name : {c.name}</Item.Header>
            <Item.Meta>age : {c.age}</Item.Meta>
            <Item.Meta>height : {c.height ? c.height : "???"} cm</Item.Meta>
          </Item.Content>
        </Item>
      ))}
    </Item.Group>
  </>
);

export default CharacterList;
