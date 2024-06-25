FROM alpine:latest

RUN apk add --no-cache bash curl openjdk11-jre

ENV DETECT_SOURCE=https://sig-repo.synopsys.com/artifactory/bds-integrations-release/com/synopsys/integration/synopsys-detect/9.7.0/synopsys-detect-9.7.0.jar

RUN curl -L -o /opt/linuxamd64/synopsys-detect.jar --progress-bar "${DETECT_SOURCE}"

COPY release/linux/amd64/blackduck-plugin /bin/

ENTRYPOINT ["/bin/blackduck-plugin"]