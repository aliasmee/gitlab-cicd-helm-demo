# Multi stage build reduced image size
FROM golang:alpine AS builder
LABEL author="aliasmee"
RUN mkdir /build
ADD . /build
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o minions .

# CMD [ "./minions" ]

FROM scratch
COPY --from=builder /build/minions /app/
ADD . /app
WORKDIR /app
CMD ["./minions"]