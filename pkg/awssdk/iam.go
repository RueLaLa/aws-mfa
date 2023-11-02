package awssdk

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/ruelala/aws-mfa/pkg/creds"
	"github.com/ruelala/aws-mfa/pkg/utils"
)

func RotateKeys(perm_profile, profile string) {
	utils.Log("INFO", "refreshing permanent IAM keys, this will generate a new set of AWS keys and REPLACE/DELETE THE EXISTING KEYS\n")
	original_key := creds.GetIniVal(perm_profile, "aws_access_key_id").String()
	utils.Log("INFO", fmt.Sprintf("original key-pair: %s\n", original_key))

	client := IAMClient(profile)
	new_key, err := client.CreateAccessKey(context.Background(), &iam.CreateAccessKeyInput{})
	utils.Panic(err)

	_, err = client.DeleteAccessKey(
		context.Background(),
		&iam.DeleteAccessKeyInput{AccessKeyId: aws.String(original_key)},
	)
	utils.Panic(err)

	creds.WritePermanentCreds(perm_profile, new_key)
	utils.Log("INFO", "new keys written to credentials file\n")
}
