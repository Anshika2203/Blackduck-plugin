$DETECT_SOURCE = "https://sig-repo.synopsys.com/artifactory/bds-integrations-release/com/synopsys/integration/synopsys-detect/9.7.0/synopsys-detect-9.7.0.jar"

New-Item -ItemType Directory -Path C:\opt\jar -Force

Invoke-WebRequest -Uri $DETECT_SOURCE -OutFile C:\opt\jar\synopsys-detect-9.7.0.jar -UseBasicParsing