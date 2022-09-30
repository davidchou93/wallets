FROM golang:1.18-alpine

RUN mkdir -p /app/
WORKDIR /app/
COPY ./bin/* /app/

EXPOSE 8080
CMD ["/app/server"]