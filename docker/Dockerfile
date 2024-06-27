FROM alpine:latest

RUN apk add --no-cache bash curl openjdk11-jre

COPY download_detect.sh /opt/

RUN chmod +x /opt/download_detect.sh

RUN /opt/download_detect.sh

COPY release/linux/amd64/blackduck-plugin /bin/

ENTRYPOINT ["/bin/blackduck-plugin"]