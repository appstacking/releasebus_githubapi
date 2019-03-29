FROM centos:7

WORKDIR /app

ADD ./App ./

#RUN ls -la

EXPOSE 12004

CMD ["./App"]
