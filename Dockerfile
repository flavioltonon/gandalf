FROM golang:1.18.2-alpine3.14 as builder

### Install dependencies
RUN apk update && apk add make

### Start build flow
WORKDIR /app

# Copy all files from source and compile code
COPY . .

RUN make build

FROM alpine:3.14

# Copy required data
COPY --from=builder /app/bin /bin

EXPOSE 8080

CMD [ "gandalf" ]