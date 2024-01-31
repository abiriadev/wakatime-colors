import { load } from 'cheerio'

const $ = load(
	Buffer.from(
		await (
			await fetch(
				`https://wakatime.com/colors/languages`,
			)
		).arrayBuffer(),
	),
)


