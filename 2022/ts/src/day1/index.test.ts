import { expect, describe, it } from 'vitest'
import { part1, part2 } from '.'
import { readData } from '../util'

describe('day 1', () => {

  const sampleData = readData('./day1/sampleInput.txt');
  const data = readData('./day1/input.txt');
  it('part 1', () => {
    expect(part1(sampleData)).toBe(24000);
    expect(part1(data)).toBe(68442);
  })

  it('part 2', () => {
    expect(part2(sampleData)).toBe(45000);
    expect(part2(data)).toBe(204837);
  })
})
