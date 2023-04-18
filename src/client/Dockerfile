FROM docker.uclv.cu/node:19.9-alpine

WORKDIR /app

COPY ./src/client/ ./

RUN npm i

EXPOSE 3000

CMD ["npm", "run", "dev"]