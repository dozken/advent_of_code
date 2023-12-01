// A rock
// B paper
// C scissors

import { readData } from "../util";

// rock 1
// paper 2
// scissors 3
const shapePointPart1: Record<string, number> =
{
  X: 1,
  Y: 2,
  Z: 3,
}

const shapePointPart2: Record<string, number> =
{
  A: 1,
  B: 2,
  C: 3,
}

type Shape = 'A' | 'B' | 'C';
const loseMap: Record<Shape, Shape> = {
  A: 'C',
  B: 'A',
  C: 'B',
}



const winMap: Record<Shape, Shape> = {
  A: 'B',
  B: 'C',
  C: 'A'
}

// won 6
// draw 3
// lost 0
const statsPart1 = {
  "A": {
    X: 3,
    Y: 6,
    Z: 0,
  },
  "B": {
    X: 0,
    Y: 3,
    Z: 6,
  },
  "C": {
    X: 6,
    Y: 0,
    Z: 3,
  }
}

export function part1(data: string[]) {
  let score: number = 0;
  data.forEach((move: string) => {
    const [opponent, me] = move.split(" ");

    score += statsPart1[opponent as Shape][me as 'X' | 'Y' | 'Z'];
    score += shapePointPart1[me];
  });

  return score;
}

export function part2(data: string[]) {
  let score: number = 0;
  data.forEach((move: string) => {
    const [opponent, me] = move.split(" ");

    if (me === 'X') {
      score += shapePointPart2[loseMap[opponent as Shape]];
    } else if (me === 'Y') {
      score += 3 + shapePointPart2[opponent as Shape];
    } else if (me === 'Z') {
      score += 6 + shapePointPart2[winMap[opponent as Shape]]
    }
  });

  return score;
}

