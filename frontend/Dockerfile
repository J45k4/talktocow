FROM node:24-alpine AS builder

WORKDIR /usr/src

COPY ./package.json ./package.json
COPY ./package-lock.json ./package-lock.json

RUN npm ci

COPY ./index.html ./index.html
COPY ./tsconfig.json ./tsconfig.json
COPY ./vite.config.ts ./vite.config.ts
COPY ./eslint.config.mjs ./eslint.config.mjs
COPY ./pages ./pages
COPY ./src ./src
COPY ./public ./public
COPY ./styles ./styles

RUN npm run build

CMD ["npm", "start"]
