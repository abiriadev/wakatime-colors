import { writeFile } from 'node:fs/promises'
import { fetchJson } from '@/index.js'
import { argv, stdout } from 'node:process'

const res = JSON.stringify(await fetchJson(), null, 4)

const outFile = argv[2] ?? './colors.json'

if (outFile === '-') {
	stdout.write(res)
} else {
	writeFile(outFile, res)
}
