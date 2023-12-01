function getItemPriority(charIndex: number): number {
  let index = 0;
  if (charIndex >= 65 && charIndex <= 90) {
    index = charIndex - 64 + 26;
  } else {
    index = charIndex - 96;
  }
  return index;
}

export function part1(data: string[]): number {
  let sum: number = 0;
  const rucksack: number[] = new Array(53);
  data.forEach(items => {
    rucksack.fill(0);

    for (let l = 0; l < items.length / 2; l++) {
      const leftIdx = getItemPriority(items.charCodeAt(l));
      rucksack[leftIdx]++;
    }

    let result: number = 0;
    for (let r = items.length / 2; r < items.length; r++) {
      const rightIdx = getItemPriority(items.charCodeAt(r));
      if (rucksack[rightIdx] > 0) {
        result = rightIdx;
        break;
      }
    }
    sum += result;
  });

  return sum;
}

export function part2(data: string[]): number {

  const rucksack: number[] = new Array(53).fill(0);
  let sum = 0;
  let result: number = 0;
  for (let i = 0; ; i++) {
    if (i % 3 == 0) {

      sum += result;
      rucksack.fill(0);
      result = 0;
    }
    if (i === data.length) {
      break;
    }

    const items = data[i];
    const set = new Set<number>();
    for (let j = 0; j < items.length; j++) {
      const idx = getItemPriority(items.charCodeAt(j));
      set.add(idx);
    }
    set.forEach(val => {
      rucksack[val!]++;

      if (rucksack[val] === 3) {
        result = val;
      }
    });


  }

  return sum;
}

