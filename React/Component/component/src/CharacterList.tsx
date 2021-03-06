import React, { Component } from 'react';
import { Header, Icon, Item } from 'semantic-ui-react';

export interface Character {
  id: number;
  name: string;
  age: number;
  height?: number;
}

interface CharacterListProps {
  school: string;
  characters: Character[];
}

class CharacterList extends Component<CharacterListProps> {
  render() {
    const { school, characters } = this.props;

    return (
      <>
        <Header as="h2">{school}</Header>
        <Item.Group>
          {characters.map(c => (
          /* eslint-disable react/jsx-key */
          <Item key={c.id}>
            <Icon name="user circle" size="huge" />
            <Item.Content>
              <Item.Header>name : {c.name}</Item.Header>
              <Item.Meta>age    : {c.age}</Item.Meta>
              <Item.Meta>height : {c.height ? c.height : '???'} cm</Item.Meta>
            </Item.Content>
          </Item>
          /* eslint-enable */
          ))}
        </Item.Group>
      </>
    );
  }
}

export default CharacterList;
