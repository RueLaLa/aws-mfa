aws-mfa
===

Generates or refreshes temporary aws credentials via STS and stores them to support tools that don't behave nicely when mfa is required.
To do this, we have the idea of "permanent" credentials and temporary credentials. To support existing scripts/tooling,
the tool looks for a permanent profile using a suffix rather than generating a temporary profile with one.

## Features

  * Stores generated credentials in your credentials file for reuse
  * Expiration is stored in the credentials file to prevent unnecessary refreshes (can be overridden with `-f/-force`)
  * Stores your mfa serial in the credentials file
  * Customizable suffix for the "permanent" credentials
  * Customizable duration (within the limits of STS)


## Install

Head over to [releases](https://github.com/RueLaLa/aws-mfa/releases) and download the latest version for your OS/Architecture, and place the extracted binary in your PATH.

### MacOS

You may run into an issue with MacOS gatekeeper blocking you from running the application because it is "unverified". If you run into that, you can add the application to the approved list by running this in your terminal of choice
```
sudo spctl --add /path/to/aws-mfa
```

## Usage
```
$ ./aws-mfa -h
Refreshes or generates temporary AWS credentials via STS. If you already have credentials with an
expiration that's an hour out or further, they won't be refreshed unless you use the '--force' flag.

Usage:
  aws-mfa [flags]

Flags:
  -f, --force                                      force a refresh even if unexpired credentials exist
  -h, --help                                       help for aws-mfa
  -p, --profile string                             profile that will contain the temporary credentials within the AWS shared credentials file (default "default")
  -s, --suffix string                              suffix to append to profile, used to find permanent credentials. results in <profile>-<suffix> (default "permanent")
```

## Example
Basic example with an `mfa_serial` defined in the credentials file and a `default-permanent` section

```
# ~/.aws/credentials

[default-permanent]
aws_access_key_id     = <YOUR_ACCESS_KEY_ID>
aws_secret_access_key = <YOUR_SECRET_ACCESS_KEY>
mfa_serial            = arn:aws:iam::<ACCOUNT_ID>:mfa/<DEVICE>
```

Run `aws-mfa` and follow the prompt to provide your mfa token. After refreshing the tokens, your credentials file will contain a new `default` section.

```
# ~/.aws/credentials
[default-permanent]
aws_access_key_id     = <YOUR_ACCESS_KEY_ID>
aws_secret_access_key = <YOUR_SECRET_ACCESS_KEY>
mfa_serial            = arn:aws:iam::<ACCOUNT_ID>:mfa/<DEVICE>

[default]
aws_access_key_id     = <TEMPORARY_ACCESS_KEY_ID>
aws_secret_access_key = <TEMPORARY_SECRET_ACCESS_KEY>
aws_session_token     = <SESSION_TOKEN>
expires               = 2018-05-12T03:18:07-04:00
```

### Profiles

If you don't provide a profile with the `--profile` flag, it will use the value `default` and look for a `default-permanent` profile to use.

```
$ ./aws-mfa --profile <my-other-profile>
```
