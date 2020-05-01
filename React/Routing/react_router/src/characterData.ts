interface Character {
  name: string;
  age: number;
  height?: number;
}

export interface Characters {
  [code: string]: {
    school?: string;
    players: Character[];
  };
}

export const characterData: Characters = {
  AAA: {
    school: "AAA",
    players: [
      {
        name: "Taro",
        age: 15,
        height: 170,
      },
      {
        name: "Jiro",
        age: 16,
        height: 171,
      },
      {
        name: "Saburo",
        age: 15,
        height: 172,
      },
    ],
  },
  BBB: {
    school: "BBB",
    players: [
      {
        name: "Shiro",
        age: 16,
        height: 173,
      },
      {
        name: "Goro",
        age: 15,
      },
    ],
  },
  CCC: {
    school: "CCC",
    players: [
      {
        name: "Rokuro",
        age: 16,
        height: 173,
      },
      {
        name: "Shichiro",
        age: 15,
      },
      {
        name: "Hachiro",
        age: 16,
      },
    ],
  },
};
