# Wakatime colors

## Overview

A list of programming languages and their primary colors used in wakatime as machine-readable format.

The list is updated automatically periodically.

## How to read the data?

You can directly fetch [the raw content](https://github.com/abiriadev/wakatime-colors/raw/main/colors.json) to get the latest data. Otherwise, you can import this repository as npm module then call [`fetchJson()`](https://github.com/abiriadev/wakatime-colors/blob/main/src/index.ts#L3) function.

## Manually build

```sh
# checkout this repository
$ git https://github.com/abiriadev/wakatime-colors && cd wakatime-colors

# install dependencies
$ corepack enable
$ pnpm install

# build source
$ pnpm build

# then, run below command to fetch and save data
$ pnpm run fetch
```

Then, you will get data in `colors.json`.

### Render SVG preview image

Assume you have `colors.json` file at the root.

```sh
$ go run .
```

Then, you will get rendered SVG file in `colors.svg`.

## Preview

![](./colors.svg)
