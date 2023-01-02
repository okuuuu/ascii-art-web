FROM golang:latest
LABEL maintainer="https://01.kood.tech/git/fpetuhov"
LABEL authors="okuu, jjelisav"
WORKDIR /app
COPY . /app
RUN go build -o /web-dockerize
EXPOSE 8080
CMD [ "/web-dockerize" ]