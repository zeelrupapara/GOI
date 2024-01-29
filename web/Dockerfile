FROM node:16.13.1-alpine

WORKDIR /usr/src/app

COPY ./package*.json ./

RUN npm install

COPY . ./

RUN mv .env.docker .env

RUN npm run build

CMD ["npm", "run", "start"]
