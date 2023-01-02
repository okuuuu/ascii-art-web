FROM golang:latest

LABEL maintainer="https://01.kood.tech/git/fpetuhov"
LABEL authors="okuu, jjelisav"
LABEL description="ASCII Art Web Dockerize project for kood/Jõhvi"

COPY ./app /app
WORKDIR /app

RUN go build -o /ascii-art-web

EXPOSE 8080

CMD [ "/ascii-art-web" ]