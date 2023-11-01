package awssdk

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/ruelala/aws-mfa/pkg/creds"
	"github.com/ruelala/aws-mfa/pkg/utils"
)

func GetSessionCreds(perm_profile, session_profile, mfa_serial, mfa_token string) {
	client := STSClient(perm_profile)
	input := &sts.GetSessionTokenInput{
		DurationSeconds: aws.Int32(int32(129600)),
		SerialNumber:    aws.String(mfa_serial),
		TokenCode:       aws.String(mfa_token),
	}
	resp, err := client.GetSessionToken(context.Background(), input)
	utils.Panic(err)

	creds.WriteSessionCreds(session_profile, resp.Credentials)
	utils.Log(
		"INFO",
		fmt.Sprintf(
			"successfully refreshed temporary credentials for %s profile (expires: %s)\n",
			session_profile,
			aws.ToTime(resp.Credentials.Expiration).Local().Format(time.RFC3339),
		),
	)
}
