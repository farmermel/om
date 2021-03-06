<!--- This file is autogenerated from the files in docsgenerator/templates/interpolate --->
&larr; [back to Commands](../README.md)

# `om interpolate`

The `interpolate` command allows you to test template interpolation in
isolation.

For example if you are modifying a product configuration template you can use
`om interpolate` to verify that the generated config looks correct before
running `om configure-product`.

## Command Usage
```

interpolates variables into a manifest

Usage:
  om [options] interpolate [<args>]

Flags:
  --config, -c             string             path for file to be interpolated
  --ops-file, -o           string (variadic)  YAML operations files
  --path                   string             extract specified value out of the interpolated file (e.g.: /private_key). The rest of the file will not be printed.
  --skip-missing, -s       bool               allow skipping missing params
  --var, -v                string (variadic)  load variable from the command line. Format: VAR=VAL
  --vars-env, OM_VARS_ENV  string (variadic)  load variables from environment variables (e.g.: 'MY' to load MY_var=value)
  --vars-file, -l          string (variadic)  load variables from a YAML file

Global Flags:
  --ca-cert, OM_CA_CERT                                  string  OpsManager CA certificate path or value
  --client-id, -c, OM_CLIENT_ID                          string  Client ID for the Ops Manager VM (not required for unauthenticated commands)
  --client-secret, -s, OM_CLIENT_SECRET                  string  Client Secret for the Ops Manager VM (not required for unauthenticated commands)
  --connect-timeout, -o, OM_CONNECT_TIMEOUT              int     timeout in seconds to make TCP connections (default: 10)
  --decryption-passphrase, -d, OM_DECRYPTION_PASSPHRASE  string  Passphrase to decrypt the installation if the Ops Manager VM has been rebooted (optional for most commands)
  --env, -e                                              string  env file with login credentials
  --help, -h                                             bool    prints this usage information (default: false)
  --password, -p, OM_PASSWORD                            string  admin password for the Ops Manager VM (not required for unauthenticated commands)
  --request-timeout, -r, OM_REQUEST_TIMEOUT              int     timeout in seconds for HTTP requests to Ops Manager (default: 1800)
  --skip-ssl-validation, -k, OM_SKIP_SSL_VALIDATION      bool    skip ssl certificate validation during http requests (default: false)
  --target, -t, OM_TARGET                                string  location of the Ops Manager VM
  --trace, -tr, OM_TRACE                                 bool    prints HTTP requests and response payloads
  --username, -u, OM_USERNAME                            string  admin username for the Ops Manager VM (not required for unauthenticated commands)
  --version, -v                                          bool    prints the om release version (default: false)
  OM_VARS_ENV                                            string  load vars from environment variables by specifying a prefix (e.g.: 'MY' to load MY_var=value)

```

## Interpolation

Given a template file with a variable reference:

```yaml
# config.yml
key: ((variable_name))
```

Values can be provided from a separate variables yaml file (`--vars-file`) or from environment variables (`--vars-env`).

To load variables from a file use the `--vars-file` flag.

```yaml
# vars.yml
variable_name: some_value
```

```
om interpolate \
  --config config.yml \
  --vars-file vars.yml
```

To load variables from a set of environment variables, specify the common
environment variable prefix with the `--vars-env` flag.

```
OM_VAR_variable_name=some_value om interpolate \
  --config config.yml \
  --vars-env OM_VAR
```

The interpolation support is inspired by similar features in BOSH. You can
[refer to the BOSH documentation](https://bosh.io/docs/cli-int/) for details on how interpolation
is performed.