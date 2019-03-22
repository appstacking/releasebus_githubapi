FROM alpine:latest

WORKDIR /app

ADD ./releasebus_githubapi ./

RUN ls -la

EXPOSE 12001

CMD ["./releasebus_githubapi"]
