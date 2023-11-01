package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/ruelala/aws-mfa/pkg/awssdk"
	"github.com/ruelala/aws-mfa/pkg/creds"
	"github.com/ruelala/aws-mfa/pkg/utils"
)

func main() {
	args := utils.ParseArgs()

	// construct permanent profile name and check that it exists
	perm_profile := fmt.Sprintf("%s-%s", args.Profile, args.Suffix)
	if !creds.IniSectionExists(perm_profile) {
		utils.Panic(fmt.Errorf("couldn't find %s profile in aws credentials file", perm_profile))
	}

	// check for existing mfa credentials
	if creds.IniSectionExists(args.Profile) {
		expire_time, err := creds.GetIniVal(args.Profile, "expires").Time()
		utils.Panic(err)
		if !args.Force && time.Now().Before(expire_time) {
			utils.Log("INFO", "creds haven't expired yet, use -f/-force to force renewal\n")
			os.Exit(0)
		}
	}

	// load and validate existing mfa device arn
	utils.Log("INFO", fmt.Sprintf("refreshing temporary credentials for %s profile\n", args.Profile))
	mfa_serial := creds.GetIniVal(perm_profile, "mfa_serial").String()
	_, err := arn.Parse(mfa_serial)
	if err != nil {
		utils.Panic(fmt.Errorf("mfa_serial ARN in %s aws credentials - %s", perm_profile, err))
	}

	// get mfa token and generate session credentials
	scanner := bufio.NewScanner(os.Stdin)
	utils.Log("INFO", "enter your MFA token code: ")
	scanner.Scan()
	mfa_token := scanner.Text()
	awssdk.GetSessionCreds(perm_profile, args.Profile, mfa_serial, mfa_token)

	// replace existing permanent IAM key-pair with new pair in credentials file
	if args.RefreshKeys {
		awssdk.RotateKeys(perm_profile, args.Profile)
	}
}
