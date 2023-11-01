# aws-mfa

Generates or refreshes temporary aws credentials via STS and stores them to
support tools that don't behave nicely when mfa is required. To do this, we have
the idea of "permanent" credentials and temporary credentials. To support
existing scripts/tooling, the tool looks for a permanent profile using a suffix
rather than generating a temporary profile with one.

## Features

- Stores generated credentials in your credentials file for reuse
- Expiration is stored in the credentials file to prevent unnecessary refreshes
  (can be overridden with `-f/-force`)
- Stores your mfa serial in the credentials file
- Customizable suffix for the "permanent" credentials
- Customizable duration (within the limits of STS)
- Replace existing permanent key pair with new pair

## Install

Head over to [releases](https://github.com/RueLaLa/aws-mfa/releases) and
download the latest version for your OS/Architecture, and place the downloaded
binary in your PATH.

### MacOS

You may run into an issue with MacOS gatekeeper blocking you from running the
application because it is "unverified". If you run into that, you can add the
application to the approved list by running this in your terminal of choice

```
sudo spctl --add /path/to/aws-mfa
```

Additionally, MacOS may complain about this being malicious software:

<img width="263" alt="malicious-error" src="https://user-images.githubusercontent.com/8377014/190272917-06a11f3e-9419-41ed-89d7-961c4218899a.png">

Fix this by going into System Preferences -> Security & Privacy, and allowing
this app.

## Usage

```
$ ./aws-mfa -h
aws-mfa

  Flags: 
       --version        Displays the program version string.
    -h --help           Displays help with available flag, subcommand, and positional value parameters.
    -p --profile        profile to create MFA creds with (default: default)
    -f --force          force MFA recreation regardless of existing tokens
    -r --refresh-keys   force refresh your existing permanent IAM keys
    -s --suffix         suffix to match to find static credentials file (default: permanent)
```

## Example

Basic example with an `mfa_serial` defined in the credentials file and a
`default-permanent` section

```
# ~/.aws/credentials

[default-permanent]
aws_access_key_id     = <YOUR_ACCESS_KEY_ID>
aws_secret_access_key = <YOUR_SECRET_ACCESS_KEY>
mfa_serial            = arn:aws:iam::<ACCOUNT_ID>:mfa/<DEVICE>
```

Run `aws-mfa` and follow the prompt to provide your mfa token. After refreshing
the tokens, your credentials file will contain a new `default` section.

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

If you don't provide a profile with the `--profile` flag, it will use the value
`default` and look for a `default-permanent` profile to use.

```
$ ./aws-mfa --profile <my-other-profile>
```
