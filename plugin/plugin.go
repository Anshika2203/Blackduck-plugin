// Copyright 2024 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

type Plugin struct {
	BlackduckURL            string `envconfig:"PLUGIN_BLACKDUCK_URL"`
	BlackduckToken          string `envconfig:"PLUGIN_BLACKDUCK_TOKEN"`
	BlackduckProject        string `envconfig:"PLUGIN_BLACKDUCK_PROJECT"`
	BlackduckOfflineMode    bool   `envconfig:"PLUGIN_BLACKDUCK_OFFLINEMODE"`
	BlackduckTestConnection bool   `envconfig:"PLUGIN_BLACKDUCK_TEST_CONNECTION"`
	BlackduckOfflineBDIO    bool   `envconfig:"PLUGIN_BLACKDUCK_OFFLINE_BDIO"`
	BlackduckTrustCerts     bool   `envconfig:"PLUGIN_BLACKDUCK_TRUST_CERTS"`
	BlackduckTimeout        int    `envconfig:"PLUGIN_BLACKDUCK_TIMEOUT"`
	BlackduckScanMode       string `envconfig:"PLUGIN_BLACKDUCK_SCAN_MODE"`
	BLackduckProperties     string `envconfig:"PLUGIN_BLACKDUCK_PROPERTIES"`
	LogLevel                string `envconfig:"PLUGIN_LOG_LEVEL"`
}

func (p *Plugin) Exec(ctx context.Context) error {
	if err := runBlackDuckScan(p); err != nil {
		return err
	}

	return nil
}

func runBlackDuckScan(p *Plugin) error {
	bdURL := p.BlackduckURL
	bdToken := p.BlackduckToken
	bdProject := p.BlackduckProject

	if bdURL == "" || bdToken == "" || bdProject == "" {
		return fmt.Errorf("BLACKDUCK_URL, BLACKDUCK_TOKEN and BLACKDUCK_PROJECT environment variables must be set")
	}

	command := fmt.Sprintf("java -jar /opt/jar/synopsys-detect-9.7.0.jar --blackduck.url=\"%s\" --blackduck.api.token=\"%s\" --detect.project.name=\"%s\"", bdURL, bdToken, bdProject)

	if p.BlackduckOfflineMode {
		command += " --blackduck.offline.mode=" + strconv.FormatBool(p.BlackduckOfflineMode)
	}

	if p.BlackduckTestConnection {
		command += " --detect.test.connection=" + strconv.FormatBool(p.BlackduckTestConnection)
	}

	if p.BlackduckOfflineBDIO {
		command += " --blackduck.offline.mode.force.bdio=" + strconv.FormatBool(p.BlackduckOfflineBDIO)
	}

	if p.BlackduckTrustCerts {
		command += " --blackduck.trust.cert=" + strconv.FormatBool(p.BlackduckTrustCerts)
	}

	if p.BlackduckTimeout > 0 {
		command += " --detect.timeout=" + strconv.Itoa(p.BlackduckTimeout)
	}

	// RAPID,STATELESS,INTELLIGENT
	if p.BlackduckScanMode != "" {
		switch p.BlackduckScanMode {
		case "RAPID", "STATELESS", "INTELLIGENT":
			command += " --detect.blackduck.scan.mode=" + p.BlackduckScanMode
		default:
			log.Printf("Unexpected BlackduckScanMode: %s \n Scan mode can be RAPID, STATELESS, INTELLIGENT.", p.BlackduckScanMode)
		}
	}

	if p.BLackduckProperties != "" {
		command += " " + p.BLackduckProperties
	}

	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Running command: %s\n", command)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute Black Duck scan command: %w", err)
	}

	return nil
}
