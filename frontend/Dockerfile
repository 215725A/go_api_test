FROM node:18.9.1

WORKDIR /app

COPY . .

RUN npm install

RUN npm install typescript -g

RUN npm run tsc

RUN npm run build

EXPOSE 3000

CMD ["npm", "start"]