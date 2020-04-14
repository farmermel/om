package commands

import (
	"fmt"
	"github.com/pivotal-cf/jhanda"
	"github.com/pivotal-cf/om/api"
	"github.com/pivotal-cf/om/interpolate"
	"gopkg.in/yaml.v2"
	"sort"
	"strings"
)

type ConfigureOpsman struct {
	service     updateSSLCertificateService
	logger      logger
	environFunc func() []string
	Options     struct {
		ConfigFile string   `long:"config"    short:"c"         description:"path to yml file containing all config fields (see docs/configure-director/README.md for format)" required:"true"`
		VarsFile   []string `long:"vars-file" short:"l"         description:"load variables from a YAML file"`
		VarsEnv    []string `long:"vars-env"  env:"OM_VARS_ENV" description:"load variables from environment variables (e.g.: 'MY' to load MY_var=value)"`
		Vars       []string `long:"var"       short:"v"         description:"load variable from the command line. Format: VAR=VAL"`
		OpsFile    []string `long:"ops-file"                    description:"YAML operations file"`
	}
}

type opsmanConfig struct {
	SSLCertificate *struct {
		Certificate string `yaml:"certificate"`
		PrivateKey  string `yaml:"private_key"`
	} `yaml:"ssl-certificate"`
	Field map[string]interface{} `yaml:",inline"`
}

//counterfeiter:generate -o ./fakes/configure_opsman_service.go --fake-name ConfigureOpsmanService . configureOpsmanService
type configureOpsmanService interface {
	UpdateSSLCertificate(api.SSLCertificateInput) error
}

func NewConfigureOpsman(environFunc func() []string, service configureOpsmanService, logger logger) ConfigureOpsman {
	return ConfigureOpsman{
		environFunc: environFunc,
		service:     service,
		logger:      logger,
	}
}

func (c ConfigureOpsman) Execute(args []string) error {
	if _, err := jhanda.Parse(&c.Options, args); err != nil {
		return fmt.Errorf("could not parse configure-opsman flags: %s", err)
	}

	config, err := c.interpolateConfig()
	if err != nil {
		return err
	}

	err = c.validate(config)
	if err != nil {
		return err
	}

	if config.SSLCertificate != nil {
		c.logger.Printf("Updating SSL Certificate...\n")
		err = c.updateSSLCertificate(config)
		if err != nil {
			return err
		}
		c.logger.Printf("Successfully applied custom SSL Certificate.\n")
		c.logger.Printf("Please allow about 1 min for the new certificate to take effect.\n")
	}
	return nil
}

func (c ConfigureOpsman) updateSSLCertificate(config *opsmanConfig) error {
	return c.service.UpdateSSLCertificate(api.SSLCertificateInput{
		CertPem:       config.SSLCertificate.Certificate,
		PrivateKeyPem: config.SSLCertificate.PrivateKey,
	})
}

func (c ConfigureOpsman) interpolateConfig() (*opsmanConfig, error) {
	configContents, err := interpolate.Execute(interpolate.Options{
		TemplateFile:  c.Options.ConfigFile,
		VarsFiles:     c.Options.VarsFile,
		EnvironFunc:   c.environFunc,
		Vars:          c.Options.Vars,
		VarsEnvs:      c.Options.VarsEnv,
		OpsFiles:      c.Options.OpsFile,
		ExpectAllKeys: true,
	})
	if err != nil {
		return nil, err
	}

	var config opsmanConfig
	err = yaml.UnmarshalStrict(configContents, &config)
	if err != nil {
		return nil, fmt.Errorf("could not be parsed as valid configuration: %s: %s", c.Options.ConfigFile, err)
	}
	return &config, nil
}

func (c ConfigureOpsman) validate(config *opsmanConfig) error {
	invalidFields := []string{}
	if len(config.Field) > 0 {
		for key, _ := range config.Field {
			if key == "opsman-configuration" {
				continue
			}
			invalidFields = append(invalidFields, key)
		}
	}
	if len(invalidFields) > 0 {
		sort.Strings(invalidFields)
		return fmt.Errorf("unrecognized top level key(s) in config file:\n%s", strings.Join(invalidFields, "\n"))
	}
	return nil
}

func (c ConfigureOpsman) Usage() jhanda.Usage {
	return jhanda.Usage{
		Description:      "This authenticated command configures settings available on the \"Settings\" page in the Ops Manager UI",
		ShortDescription: "configures values present on the Ops Manager settings page",
		Flags:            c.Options,
	}
}
