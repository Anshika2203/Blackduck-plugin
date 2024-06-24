// Copyright 2024 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

type Plugin struct {
	BlackduckURL        string `envconfig:"PLUGIN_BLACKDUCK_URL"`
	BlackduckToken      string `envconfig:"PLUGIN_BLACKDUCK_TOKEN"`
	BlackduckProject    string `envconfig:"PLUGIN_BLACKDUCK_PROJECT"`
	BLackduckProperties string `envconfig:"PLUGIN_BLACKDUCK_PROPERTIES"`
	LogLevel            string `envconfig:"PLUGIN_LOG_LEVEL"`
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
	moreProperties := p.BLackduckProperties

	if bdURL == "" || bdToken == "" {
		return fmt.Errorf("BLACKDUCK_URL and BLACKDUCK_TOKEN environment variables must be set")
	}

	command := fmt.Sprintf("bash <(curl -s -L https://detect.synopsys.com/detect9.sh) --blackduck.url=\"%s\" --blackduck.api.token=\"%s\" --detect.project.name=\"%s\" --blackduck.trust.cert=true", bdURL, bdToken, bdProject)

	if moreProperties != "" {
		command += " " + moreProperties
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
