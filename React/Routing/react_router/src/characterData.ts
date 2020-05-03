interface Character {
  id: string;
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
        id: "1",
        name: "Taro",
        age: 15,
        height: 170,
      },
      {
        id: "2",
        name: "Jiro",
        age: 16,
        height: 171,
      },
      {
        id: "3",
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
        id: "4",
        name: "Shiro",
        age: 16,
        height: 173,
      },
      {
        id: "5",
        name: "Goro",
        age: 15,
      },
    ],
  },
  CCC: {
    school: "CCC",
    players: [
      {
        id: "6",
        name: "Rokuro",
        age: 16,
        height: 173,
      },
      {
        id: "7",
        name: "Shichiro",
        age: 15,
      },
      {
        id: "8",
        name: "Hachiro",
        age: 16,
      },
    ],
  },
};
