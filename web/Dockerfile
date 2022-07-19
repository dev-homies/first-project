FROM node:18-alpine

WORKDIR /web

COPY package-lock.json package-lock.json
COPY package.json package.json
RUN npm install

EXPOSE 3000
EXPOSE 9229

CMD [ "npm", "run", "debug" ]
