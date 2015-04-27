package driver

import (
	"docker-extension/dockeropts"
)

type UbuntuSystemdDriver struct {
	ubuntuBaseDriver
	systemdBaseDriver
}

func (u UbuntuSystemdDriver) BaseOpts() []string {
	return []string{"--daemon", "-H=fd://"}
}

func (u UbuntuSystemdDriver) ChangeOpts(args string) error {
	const cfg = "/lib/systemd/system/docker.service"
	e := dockeropts.SystemdUnitEditor{}
	return rewriteOpts(e, cfg, args)
}
