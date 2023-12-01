import { describe, expect, it } from "vitest";
import { part1, part2 } from ".";
import { readData } from "../util";

describe.skip('day 5', () => {
  const sampleData = readData('./day5/sampleInput.txt');
  const data = readData('./day5/input.txt');

  it('part 1, sampleData', () => {
    expect(part1(sampleData)).toBe(2);
  })

  it.skip('part 1, data', () => {
    expect(part1(data)).toBe(605);
  })

  it.skip('part 2, sampleData', () => {
    expect(part2(sampleData)).toBe(4);
  })

  it.skip('part 2, data', () => {
    expect(part2(data)).toBe(914);
  })
})
