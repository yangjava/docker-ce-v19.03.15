// +build linux

package daemon // import "github.com/docker/docker/daemon"

import (
	"fmt"

	aaprofile "github.com/docker/docker/profiles/apparmor"
	"github.com/opencontainers/runc/libcontainer/apparmor"
)

// Define constants for native driver
const (
	defaultApparmorProfile = "docker-default"
)

// DefaultApparmorProfile returns the name of the default apparmor profile
func DefaultApparmorProfile() string {
	if apparmor.IsEnabled() {
		return defaultApparmorProfile
	}
	return ""
}

func ensureDefaultAppArmorProfile() error {
	if apparmor.IsEnabled() {
		loaded, err := aaprofile.IsLoaded(defaultApparmorProfile)
		if err != nil {
			return fmt.Errorf("Could not check if %s AppArmor profile was loaded: %s", defaultApparmorProfile, err)
		}

		// Nothing to do.
		if loaded {
			return nil
		}

		// Load the profile.
		if err := aaprofile.InstallDefault(defaultApparmorProfile); err != nil {
			return fmt.Errorf("AppArmor enabled on system but the %s profile could not be loaded: %s", defaultApparmorProfile, err)
		}
	}

	return nil
}
