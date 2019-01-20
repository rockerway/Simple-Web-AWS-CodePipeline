FROM golang:alpine AS build-env
ADD ./app /src
RUN cd /src && go build -o app

FROM alpine
WORKDIR /app
COPY --from=build-env /src/app /app/
ENTRYPOINT ./app
