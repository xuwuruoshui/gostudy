FROM golang:alpine

RUN  mkdir -p /drone-test
ADD  ./main /drone-test
WORKDIR /drone-test


EXPOSE 4567

ENTRYPOINT ["/drone-test/main"]