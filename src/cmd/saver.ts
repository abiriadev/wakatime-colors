import { writeFile } from 'node:fs/promises'
import { fetchJson } from '@/index.js'

await writeFile(
	'./colors.json',
	JSON.stringify(await fetchJson(), null, 4),
)
