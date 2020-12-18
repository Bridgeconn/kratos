package internal

import (
	"context"
	"testing"

	"github.com/ory/x/configx"
	"github.com/ory/x/dbal"
	"github.com/ory/x/stringsx"

	"github.com/stretchr/testify/require"

	"github.com/ory/x/logrusx"

	"github.com/ory/kratos/driver"
	"github.com/ory/kratos/driver/config"
	"github.com/ory/kratos/selfservice/hook"
	"github.com/ory/kratos/x"
)

const UnsetDefaultIdentitySchema = "file://not-set.schema.json"

func NewConfigurationWithDefaults() *config.Config {
	c := config.MustNew(logrusx.New("", ""),
		configx.WithValues(map[string]interface{}{
			"log.level":                                      "trace",
			config.ViperKeyDSN:                               dbal.InMemoryDSN,
			config.ViperKeyHasherArgon2ConfigMemory:          16384,
			config.ViperKeyHasherArgon2ConfigIterations:      1,
			config.ViperKeyHasherArgon2ConfigParallelism:     1,
			config.ViperKeyHasherArgon2ConfigSaltLength:      16,
			config.ViperKeyHasherArgon2ConfigKeyLength:       16,
			config.ViperKeyCourierSMTPURL:                    "smtp://foo:bar@baz.com/",
			config.ViperKeySelfServiceBrowserDefaultReturnTo: "https://www.ory.sh/redirect-not-set",
			config.ViperKeyDefaultIdentitySchemaURL:          UnsetDefaultIdentitySchema,
		}),
		configx.SkipValidation(),
	)
	return c
}

// NewFastRegistryWithMocks returns a registry with several mocks and an SQLite in memory database that make testing
// easier and way faster. This suite does not work for e2e or advanced integration tests.
func NewFastRegistryWithMocks(t *testing.T) (*config.Config, *driver.RegistryDefault) {
	conf, reg := NewRegistryDefaultWithDSN(t, "")
	reg.WithCSRFTokenGenerator(x.FakeCSRFTokenGenerator)
	reg.WithCSRFHandler(x.NewFakeCSRFHandler(""))
	reg.WithHooks(map[string]func(config.SelfServiceHook) interface{}{
		"err": func(c config.SelfServiceHook) interface{} {
			return &hook.Error{Config: c.Config}
		},
	})

	require.NoError(t, reg.Persister().MigrateUp(context.Background()))
	return conf, reg
}

// NewRegistryDefaultWithDSN returns a more standard registry without mocks. Good for e2e and advanced integration testing!
func NewRegistryDefaultWithDSN(t *testing.T, dsn string) (*config.Config, *driver.RegistryDefault) {
	c := NewConfigurationWithDefaults()
	c.MustSet(config.ViperKeyDSN, stringsx.Coalesce(dsn, dbal.InMemoryDSN))

	reg, err := driver.NewRegistryFromDSN(c, logrusx.New("", ""))
	require.NoError(t, err)
	reg.Config().MustSet("dev", true)
	require.NoError(t, reg.Init())
	return c, reg.(*driver.RegistryDefault)
}
