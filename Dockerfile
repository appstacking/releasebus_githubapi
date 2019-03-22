FROM alpine:latest

WORKDIR /app

ADD ./releasebus_githubapi ./

RUN ls -la

EXPOSE 12004

CMD ["./releasebus_githubapi"]
