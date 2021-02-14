# FROM golang:1.14 AS go

# WORKDIR /src
# COPY . /src
# RUN go mod download && \
#     go build -o main . && \
#     chmod +x ./main
# EXPOSE 8000
# CMD ["/main"]
# # ENTRYPOINT [ "./main" ]
# HEALTHCHECK --interval=1m --timeout=3s CMD wget --quiet --tries=1 --spider http://localhost:8000/posts || exit 1
# ----------
# FROM golang:alpine
# RUN apk add --no-cache git
# RUN apk add --no-cache sqlite-libs sqlite-dev
# RUN apk add --no-cache build-base
# WORKDIR /go/src/app
# COPY . .
# RUN go mod download && \
#     go build -o main .
# CMD ["./main"]
# ---------------
FROM golang:1.14.13-alpine3.12 as builder

WORKDIR /go/src/app
COPY . .
RUN apk --no-cache add make git gcc libtool musl-dev ca-certificates dumb-init 

RUN go mod download && \
    go build -o main .

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /go/src/app .
# ENTRYPOINT ./app
LABEL Name=cloud-native-go Version=0.0.1
EXPOSE 8000
CMD ["./main"]
