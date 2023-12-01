export function part1(data: string[]): number {
  let max = 0;
  let current = 0;
  for (let i = 0; i < data.length; i++) {
    if (data[i])
      current += parseInt(data[i]);

    if (!data[i] || i === data.length - 1) {
      max = Math.max(max, current);
      current = 0;
    }
  }
  return max;
}

export function part2(data: string[]): number {
  const top3: number[] = [0, 0, 0];
  let current = 0;
  for (let i = 0; i < data.length; i++) {
    if (data[i])
      current += parseInt(data[i]);

    if (!data[i] || i === data.length - 1) {
      top3[2] = Math.max(top3[2], current);
      current = 0;
      top3.sort((a, b) => b - a);
    }
  }

  return top3.reduce((acc, val) => acc + val);
}

