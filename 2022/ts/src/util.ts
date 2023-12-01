import { resolve } from 'node:path'
import { readFileSync } from 'node:fs'

export const readData = (file: string): Array<string> => {
  return readFileSync(resolve(__dirname, file))
    .toString()
    .trim()
    .split('\n')
}
