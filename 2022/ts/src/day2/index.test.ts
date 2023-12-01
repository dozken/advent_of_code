import { expect, describe, it } from 'vitest'
import { part1, part2 } from '.'
import { readData } from '../util'

describe('day 2', () => {

  const sampleData = readData('./day2/sampleInput.txt');
  const data = readData('./day2/input.txt');
  it('part 1', () => {
    expect(part1(sampleData)).toBe(15);
    expect(part1(data)).toBe(14297);
  })

  it('part 2', () => {
    expect(part2(sampleData)).toBe(12);
    expect(part2(data)).toBe(10498);
  })
})
