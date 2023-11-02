package creds

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/sts/types"
	"github.com/go-ini/ini"
	"github.com/ruelala/aws-mfa/pkg/utils"
)

func GetIniVal(profile, key string) *ini.Key {
	cfg, err := ini.Load(config.DefaultSharedCredentialsFilename())
	utils.Panic(err)
	val := cfg.Section(profile).Key(key)
	return val
}

func IniSectionExists(profile string) bool {
	cfg, err := ini.Load(config.DefaultSharedCredentialsFilename())
	utils.Panic(err)
	section := cfg.Section(profile)
	if len(section.Keys()) == 0 {
		return false
	} else {
		return true
	}
}

func WriteSessionCreds(profile string, creds *types.Credentials) {
	cfg, err := ini.Load(config.DefaultSharedCredentialsFilename())
	utils.Panic(err)
	cfg.Section(profile).Key("aws_access_key_id").SetValue(aws.ToString(creds.AccessKeyId))
	cfg.Section(profile).Key("aws_secret_access_key").SetValue(aws.ToString(creds.SecretAccessKey))
	cfg.Section(profile).Key("aws_session_token").SetValue(aws.ToString(creds.SessionToken))
	cfg.Section(profile).Key("expires").SetValue(aws.ToTime(creds.Expiration).Local().Format(time.RFC3339))
	cfg.SaveTo(config.DefaultSharedCredentialsFilename())
}

func WritePermanentCreds(perm_profile string, new_key *iam.CreateAccessKeyOutput) {
	cfg, _ := ini.Load(config.DefaultSharedCredentialsFilename())
	cfg.Section(perm_profile).Key("aws_access_key_id").SetValue(aws.ToString(new_key.AccessKey.AccessKeyId))
	cfg.Section(perm_profile).Key("aws_secret_access_key").SetValue(aws.ToString(new_key.AccessKey.SecretAccessKey))
	cfg.SaveTo(config.DefaultSharedCredentialsFilename())
}
