FROM node:12.6.0-alpine as builder

WORKDIR /web/
COPY . .

RUN npm config set registry https://registry.npm.taobao.org
RUN npm install
RUN npm run build

FROM nginx:alpine
LABEL MAINTAINER="lirui@thooh.com"

COPY .docker-compose/nginx/ /etc/nginx/
COPY --from=builder /web/dist /var/www/html
WORKDIR /var/www/html
