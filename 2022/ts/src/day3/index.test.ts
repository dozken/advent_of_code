import { describe, expect, it } from "vitest";
import { part1, part2 } from ".";
import { readData } from "../util";

describe('day 3', () => {
  const sampleData = readData('./day3/sampleInput.txt');
  const data = readData('./day3/input.txt');

  it('part 1, sampleData', () => {
    expect(part1(sampleData)).toBe(157);
  })

  it('part 1, data', () => {
    expect(part1(data)).toBe(7908);
  })

  it('part 2, sampleData', () => {
    expect(part2(sampleData)).toBe(70);
  })

  it('part 2, data', () => {
    expect(part2(data)).toBe(2838);
  })
})
