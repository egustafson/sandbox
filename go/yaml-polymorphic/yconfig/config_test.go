package yconfig_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	"polyyaml/yconfig"
)

func TestUnmarshalMasterConfig_NullDecryptor(t *testing.T) {
	data := `
id: master0
decryptor:
  type: null-decryptor
`

	var cfg yconfig.MasterConfig
	err := yaml.Unmarshal([]byte(data), &cfg)
	assert.NoError(t, err)
	assert.Equal(t, "master0", cfg.ID)

	nullDecryptor, ok := cfg.Decryptor.Decryptor.(*yconfig.NullDecryptor)
	assert.True(t, ok)
	_ = nullDecryptor // just to avoid unused variable warning
}

func TestUnmarshalMasterConfig_PasswordDecryptor(t *testing.T) {
	data := `
id: master1
decryptor:
  type: password-decryptor
  password: secret123
`

	var cfg yconfig.MasterConfig
	err := yaml.Unmarshal([]byte(data), &cfg)
	assert.NoError(t, err)
	assert.Equal(t, "master1", cfg.ID)

	pwdDecryptor, ok := cfg.Decryptor.Decryptor.(*yconfig.PasswordDecryptor)
	assert.True(t, ok)
	assert.Equal(t, "secret123", pwdDecryptor.Password)
}

func TestUnmarshalMasterConfig_YubiKeyDecryptor(t *testing.T) {
	data := `
id: master2
decryptor:
  type: yubikey-decryptor
  slot: 2
  pin: 123456
`

	var cfg yconfig.MasterConfig
	err := yaml.Unmarshal([]byte(data), &cfg)
	assert.NoError(t, err)
	assert.Equal(t, "master2", cfg.ID)

	ykDecryptor, ok := cfg.Decryptor.Decryptor.(*yconfig.YubiKeyDecryptor)
	assert.True(t, ok)
	assert.Equal(t, 2, ykDecryptor.Slot)
	assert.Equal(t, "123456", ykDecryptor.PIN)
}

func TestUnmarshalMasterConfig_UnknownDecryptorType(t *testing.T) {
	data := `
id: master3
decryptor:
  type: nonexistent-type
`

	var cfg yconfig.MasterConfig
	err := yaml.Unmarshal([]byte(data), &cfg)
	assert.ErrorIs(t, err, yconfig.ErrDecryptorType)
}
