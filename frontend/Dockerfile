FROM node:14.16 as builder

WORKDIR /usr/src

ADD ./package.json ./package.json
ADD ./package-lock.json ./package-lock.json

RUN npm ci

ADD ./tsconfig.json ./tsconfig.json
ADD ./next.config.js ./next.config.js
ADD ./worker.js ./worker.js
COPY ./pages ./pages
COPY ./src ./src
COPY ./public ./public
COPY ./styles ./styles

RUN npm run build

CMD ["npm", "start"]