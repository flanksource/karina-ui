FROM node: lts-alpine

WORKDIR /opt/karina-dash

#ENV

COPY . .

RUN nmp install

RUN nmp install @vue/cli@3.7.0 -g

EXPOSE 8080

CMD ["npm", "run", "serve"]
