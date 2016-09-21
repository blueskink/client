// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

// +build darwin

package client

import (
	"strings"

	"github.com/keybase/client/go/install"
)

func (c *CmdStatus) osSpecific(status *fstatus) error {
	serviceStatus := install.KeybaseServiceStatus(c.G(), "service", 0, c.G().Log)
	kbfsStatus := install.KeybaseServiceStatus(c.G(), "kbfs", 0, c.G().Log)

	productVersion, buildVersion, err := c.osVersionAndBuild()
	if err != nil {
		c.G().Log.Warning("Error determining OS version: %s", err)
	}
	status.OSVersion = strings.Join([]string{productVersion, buildVersion}, "-")

	if len(serviceStatus.Pid) > 0 {
		status.Service.Running = true
		status.Service.Pid = serviceStatus.Pid
	}

	if len(kbfsStatus.Pid) > 0 {
		status.KBFS.Running = true
		status.KBFS.Pid = kbfsStatus.Pid
	}

	return nil
}
