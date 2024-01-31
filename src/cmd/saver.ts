import { writeFile } from 'node:fs/promises'
import { fetchJson } from '@/index.js'
import { argv, stdout } from 'node:process'

const res = JSON.stringify(await fetchJson(), null, 4)

// path to save output
const outFile = argv[2] ?? './colors.json'

// if the given path is -, prints to stdout.
if (outFile === '-') stdout.write(res)
else writeFile(outFile, res)
