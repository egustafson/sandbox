// Package yconfig is a YAML example of a polymorphic configuration structure.
package yconfig

import "errors"

const (
	NullDecryptorType     = "null-decryptor"
	PasswordDecryptorType = "password-decryptor"
	YubiKeyDecryptorType  = "yubikey-decryptor"
)

type MasterConfig struct {
	ID        string        `yaml:"id"`
	Decryptor DecryptConfig `yaml:"decryptor"`
}

var ErrDecryptorType = errors.New("unknown decryptor type")

type DecryptConfig struct {
	Decryptor DecryptorType `yaml:",inline"`
}

type DecryptorType interface {
	GetType() string
}

type NullDecryptor struct{}

func (dc *NullDecryptor) GetType() string {
	return NullDecryptorType
}

type PasswordDecryptor struct {
	Password string `yaml:"password"`
}

func (dc *PasswordDecryptor) GetType() string {
	return PasswordDecryptorType
}

type YubiKeyDecryptor struct {
	Slot int    `yaml:"slot"`
	PIN  string `yaml:"pin"`
}

func (dc *YubiKeyDecryptor) GetType() string {
	return YubiKeyDecryptorType
}

func (dc *DecryptConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var aux struct {
		Type string `yaml:"type"`
	}

	if err := unmarshal(&aux); err != nil {
		return err
	}

	switch aux.Type {
	case NullDecryptorType:
		var nullDecryptor NullDecryptor
		if err := unmarshal(&nullDecryptor); err != nil {
			return err
		}
		dc.Decryptor = &nullDecryptor
	case PasswordDecryptorType:
		var pwdDecryptor PasswordDecryptor
		if err := unmarshal(&pwdDecryptor); err != nil {
			return err
		}
		dc.Decryptor = &pwdDecryptor
	case YubiKeyDecryptorType:
		var ykDecryptor YubiKeyDecryptor
		if err := unmarshal(&ykDecryptor); err != nil {
			return err
		}
		dc.Decryptor = &ykDecryptor
	default:
		return ErrDecryptorType
	}
	return nil
}
