// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Parity Public License
// that can be found in the LICENSE file.

package environ

import (
	"fmt"
	"strings"

	"github.com/coreos/go-semver/semver"
)

// helper function returns a list of environment variables
// that represent the semantic version.
func versions(s string) map[string]string {
	s = strings.TrimPrefix(s, "v")
	version, err := semver.NewVersion(s)
	if err != nil {
		return nil
	}
	env := map[string]string{}
	env["DRONE_SEMVER"] = version.String()
	env["DRONE_SEMVER_MAJOR"] = fmt.Sprint(version.Major)
	env["DRONE_SEMVER_MINOR"] = fmt.Sprint(version.Minor)
	env["DRONE_SEMVER_PATCH"] = fmt.Sprint(version.Patch)
	if s := string(version.PreRelease); s != "" {
		env["DRONE_SEMVER_PRE_RELEASE"] = s
	}
	if version.Metadata != "" {
		env["DRONE_SEMVER_BUILD"] = version.Metadata
	}
	return env
}
