import { describe, expect, it } from "vitest";
import { part1, part2 } from ".";
import { readData } from "../util";

describe('day 4', () => {
  const sampleData = readData('./day4/sampleInput.txt');
  const data = readData('./day4/input.txt');

  it('part 1, sampleData', () => {
    expect(part1(sampleData)).toBe(2);
  })

  it('part 1, data', () => {
    expect(part1(data)).toBe(605);
  })

  it('part 2, sampleData', () => {
    expect(part2(sampleData)).toBe(4);
  })

  it('part 2, data', () => {
    expect(part2(data)).toBe(914);
  })
})
