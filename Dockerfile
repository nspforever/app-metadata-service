FROM golang:1.17.2-buster AS base

RUN mkdir -p /go/src/github.com/nspforever/app-metadata-service
WORKDIR /go/src/github.com/nspforever/app-metadata-service

COPY . .

RUN make build

# runnable
ADD ./scripts/start.sh .
RUN chmod +x ./start.sh
CMD ["./start.sh"]

EXPOSE 9999