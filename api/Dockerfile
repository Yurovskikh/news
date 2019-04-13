FROM golang:latest AS build
ADD . /src
RUN cd /src && go build -o api

FROM alpine
WORKDIR /bin
RUN ls
COPY --from=build /src/api /bin/
ENTRYPOINT ./api
EXPOSE 8080