package utils

import (
	"fmt"
	"runtime"

	"github.com/integrii/flaggy"
)

// these get passed in as ldflags by goreleaser
var version string
var commit string
var date string

func print_version() string {
	go_version := runtime.Version()
	return fmt.Sprintf("aws-mfa %s built with %s on commit %s at %s", version, go_version, commit, date)
}

type Args struct {
	Profile            string
	Force, RefreshKeys bool
	Suffix             string
}

func ParseArgs() Args {
	a := Args{"default", false, false, "permanent"}
	flaggy.String(&a.Profile, "p", "profile", "profile to create MFA creds with")
	flaggy.Bool(&a.Force, "f", "force", "force MFA recreation regardless of existing tokens")
	flaggy.Bool(&a.RefreshKeys, "r", "refresh-keys", "force refresh your existing permanent IAM keys")
	flaggy.String(&a.Suffix, "s", "suffix", "suffix to match to find static credentials file")
	flaggy.SetVersion(print_version())
	flaggy.Parse()
	return a
}
