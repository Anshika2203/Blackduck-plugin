FROM alpine:latest

RUN apk add --no-cache bash curl openjdk11-jre

COPY release/linux/amd64/blackduck-plugin /bin/

ENTRYPOINT ["/bin/blackduck-plugin"]