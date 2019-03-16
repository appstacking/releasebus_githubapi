FROM alpine:latest

WORKDIR /app

ADD ./releasebus_service_githubapi ./

RUN ls -la

EXPOSE 12001

CMD ["./releasebus_service_githubapi"]
