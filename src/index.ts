import { cheerioJsonMapper } from 'cheerio-json-mapper'

export const fetchJson = async () =>
	(await cheerioJsonMapper(
		await (
			await fetch(
				`https://wakatime.com/colors/languages`,
			)
		).text(),
		[
			{
				$: '.editor-icons > .editor-icon',
				color: ':first',
				name: ':last',
			},
		],
	)) as Array<{
		name: string
		color: string
	}>
