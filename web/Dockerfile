FROM docker.uclv.cu/node:19.9-alpine

WORKDIR /web

COPY ./ ./

# RUN npm i

EXPOSE 80
EXPOSE 41234/udp

CMD ["npm", "run", "dev"]