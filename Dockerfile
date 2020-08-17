FROM node:lts-alpine

WORKDIR /opt/karina-dash

#ENV

COPY . .

RUN npm install

RUN npm install @vue/cli -g

EXPOSE 8080

CMD ["npm", "run", "serve"]
