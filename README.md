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

| Parameter                                                                             | Comments                                                                                                                                                                                             |
|:--------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| blackduck_url <span style="font-size: 10px"><br/>`required`</span>                    | The URL of the Black Duck server.                                                                                                                                                                    |
| blackduck_token <span style="font-size: 10px"><br/>`required`</span>                  | The API token for accessing the Black Duck server.                                                                                                                                                   |
| blackduck_project  <span style="font-size: 10px"><br/>`required`</span>               | The name of the project in Blackduck.                                                                                                                                                                |
| blackduck_offline_mode <span style="font-size: 10px"><br/>`optional` `bool`</span>    | Offline Mode: This can disable Black Duck communication - if set to true, Synopsys Detect will not upload BDIO files, or check policies, and it will not download and install the signature scanner. |
| blackduck_test_connection <span style="font-size: 10px"><br/>`optional` `bool`</span> | Test the connection to Black Duck with the current configuration.                                                                                                                                    |
| blackduck_offline_bdio <span style="font-size: 10px"><br/>`optional` `bool`</span>    | This property will force Detect in offline mode to generate a BDIO even if no code locations were identified.                                                                                        |
| blackduck_trust_certs <span style="font-size: 10px"><br/>`optional` `bool`</span>     | If true, automatically trust the certificate for the current run of Detect only.                                                                                                                     |
| blackduck_timeout <span style="font-size: 10px"><br/>`optional` `int`</span>          | The amount of time in seconds Detect will wait for network connection, for scans to finish, and to generate reports.                                                                                 |
| blackduck_scan_mode <span style="font-size: 10px"><br/>`optional`</span>              | Set the Black Duck scanning mode of Detect. Acceptable Values: RAPID, STATELESS, INTELLIGENT.                                                                                                        |
| bLackduck_properties <span style="font-size: 10px"><br/>`optional`</span>             | Additional properties to pass to the Synopsys Detect script.                                                                                                                                         |


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
    -e PLUGIN_BLACKDUCK_PROJECT="$PLUGIN_BLACKDUCK_PROJECT" \
    -e PLUGIN_PROPERTIES="$PLUGIN_PROPERTIES" \
    synopsys-detect

```

```
# Plugin YAML
- step:
    type: Plugin
    name: synopsys-detect-plugin-arm64
    identifier: synopsys-detect-plugin-arm64
    spec:
        connectorRef: harness-docker-connector
        image: harnesscommunitytest/synopsys-detect:linux-arm64
        settings:
            blackduck_url: https://abcd.blackduck.com/
            blackduck_token: some_token
            blackduck_project: test
       

- step:
    type: Plugin
    name: synopsys-detect-plugin-amd64
    identifier: synopsys-detect-plugin-amd64
    spec:
        connectorRef: harness-docker-connector
        image: harnesscommunitytest/synopsys-detect:linux-amd64
        settings:
            blackduck_url: https://abcd.blackduck.com/
            blackduck_token: some_token
            blackduck_project: test