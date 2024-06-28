#!/bin/sh

# Set the source URL for the Synopsys Detect JAR file
DETECT_SOURCE=https://sig-repo.synopsys.com/artifactory/bds-integrations-release/com/synopsys/integration/synopsys-detect/9.7.0/synopsys-detect-9.7.0.jar

# Create the directory if it doesn't exist
mkdir -p /opt/jar

# Download the JAR file
curl -L -o /opt/jar/synopsys-detect-9.7.0.jar --progress-bar "${DETECT_SOURCE}"
