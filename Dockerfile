FROM node:latest as builder

ADD . /srv
WORKDIR /srv/web/app
RUN npm i
RUN npm run build

FROM --platform=linux/amd64 golang:1.18-alpine as gobuilder
WORKDIR /app
ADD . /app
RUN go mod vendor 
RUN go build -buildvcs=false -o /app-linux


FROM --platform=linux/amd64 alpine:latest
RUN mkdir -p /root/web/app/.output/public
COPY --from=gobuilder /app-linux /root/app
COPY --from=builder /srv/web/app/.output/public /root/web/app/.output/public

RUN chmod +x /root/app
RUN apk add ca-certificates

WORKDIR /root

ENTRYPOINT [ "./app" ]

ADD resources /root/resources