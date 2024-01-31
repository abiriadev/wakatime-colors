# build stage
FROM node:20.11.0-alpine3.19 AS builder
WORKDIR /app

# install dependencies
COPY ./package.json ./pnpm-lock.yaml ./
RUN ["corepack", "enable"]
RUN ["pnpm", "install", "--frozen-lockfile"]

# copy source
COPY . .

# build package
RUN ["pnpm", "run", "build"]

# package codebase
RUN ["pnpm", "prune", "--prod"]

RUN ["mkdir", "pack"]
RUN ["mv", "package.json", "dist", "node_modules", "pack"]

# run stage
FROM node:20.11.0-slim AS runner
WORKDIR /home/node/app

# production setup
USER node
ENV NODE_ENV production

# labels
LABEL org.opencontainers.image.authors="Abiria <abiria.dev@gmail.com>"
LABEL org.opencontainers.image.url="https://github.com/abiriadev/wakatime-colors"
LABEL org.opencontainers.image.source="https://github.com/abiriadev/wakatime-colors"
LABEL org.opencontainers.image.licenses="MIT"

# copy packaged codebase
COPY --from=builder /app/pack ./

# run saver
ENTRYPOINT ["node", "./dist/src/cmd/saver.js", "-"]
