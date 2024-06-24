# Blackduck-plugin

- [Synopsis](#Synopsis)
- [Plugin Image](#Plugin-Image)
- [Parameters](#Parameters)
- [Building](#building)
- [Examples](#Examples)


## Synopsis

This repository contains a plugin for running Synopsys Detect or Black Duck scans using a Go application packaged in a Docker container.


## Plugin Image

The plugin `anshikaanand/blackduck-plugin` is available for the following architectures:

| OS            | Tag                 |
|---------------|---------------------|
| latest        | `linux-amd64/arm64` |
| linux/amd64   | `linux-amd64`       |
| linux/arm64   | `linux-arm64`       |
| windows/amd64 | `windows-amd64`     |


## Parameters

| Parameter                                                                | Comments                                                                                                                           |
|:-------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------|
| BlackduckURL <span style="font-size: 10px"><br/>`required`</span>        | The URL of the Black Duck server. |
| BlackduckToken <span style="font-size: 10px"><br/>`required`</span>      | The API token for accessing the Black Duck server.                                                              |
| BLackduckProperties <span style="font-size: 10px"><br/>`optional`</span> | Additional properties to pass to the Synopsys Detect script.                                                      |


## Building

Build the plugin image:

```text
./scripts/build.sh
```

## Examples

```
docker run --rm \
    -e PLUGIN_BLACKDUCK_URL="$PLUGIN_BLACKDUCK_URL" \
    -e PLUGIN_BLACKDUCK_TOKEN="$PLUGIN_BLACKDUCK_TOKEN" \
    -e PLUGIN_PROPERTIES="$PLUGIN_PROPERTIES" \
    blackduck-plugin

```

```
# Plugin YAML
- step:
    type: Plugin
    name: blackduck-plugin-arm64
    identifier: blackduck-plugin-arm64
    spec:
        connectorRef: harness-docker-connector
        image: anshikaanand/blackduck-plugin:linux-arm64
       

- step:
    type: Plugin
    name: blackduck-plugin-amd64
    identifier: blackduck-plugin-amd64
    spec:
        connectorRef: harness-docker-connector
        image: anshikaanand/blackduck-plugin:linux-amd64
        