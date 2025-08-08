FROM golang:1.24-alpine AS build-stage

WORKDIR /app/url_shortener/

COPY . .

RUN go get
RUN go build -o ./shortener_exec .


FROM alpine AS final-stage

WORKDIR /app

COPY --from=build-stage /app/url_shortener/shortener_exec ./shortener

EXPOSE 8080
CMD ./shortener