export function part1(data: string[]): number {
  let result = 0;

  data.forEach(line => {

    let [f1, f2, s1, s2] = line.split(/,|-/);
    f1 = parseInt(f1);
    f2 = parseInt(f2);
    s1 = parseInt(s1);
    s2 = parseInt(s2);

    if ((f1 >= s1 && f2 <= s2)
      || (s1 >= f1 && s2 <= f2)
    ) {
      result++;
    }
  });

  return result;
}

export function part2(data: string[]): number {
  let result = 0;

  data.forEach(line => {

    let [f1, f2, s1, s2] = line.split(/,|-/);
    f1 = parseInt(f1);
    f2 = parseInt(f2);
    s1 = parseInt(s1);
    s2 = parseInt(s2);

    if ((f2 >= s1 && f2 <= s2)
      || (s2 >= f1 && s2 <= f2)
    ) {
      result++;
    }
  });

  return result;
}

