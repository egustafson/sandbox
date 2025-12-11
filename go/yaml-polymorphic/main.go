package main

import (
	"embed"
	"fmt"
	"io/fs"
	"polyyaml/yconfig"

	"gopkg.in/yaml.v3"
)

//go:embed data/*
var cfgs embed.FS

func main() {

	cfgFileList, err := fs.Glob(cfgs, "data/*.yaml")
	if err != nil {
		fmt.Println("Error globbing files:", err)
		return
	}

	for _, cfgFile := range cfgFileList {

		data, err := cfgs.ReadFile(cfgFile)
		if err != nil {
			fmt.Println("Error reading file:", err)
			continue
		}

		var cfgvar yconfig.MasterConfig
		err = yaml.Unmarshal(data, &cfgvar)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		switch cfgvar.Decryptor.Decryptor.GetType() {
		case yconfig.NullDecryptorType:
			nullDecryptor, ok := cfgvar.Decryptor.Decryptor.(*yconfig.NullDecryptor)
			if !ok {
				fmt.Println("Decryptor is not of type NullDecryptor")
				return
			}
			printNullInfo(nullDecryptor)
		case yconfig.PasswordDecryptorType:
			pwdDecryptor, ok := cfgvar.Decryptor.Decryptor.(*yconfig.PasswordDecryptor)
			if !ok {
				fmt.Println("Decryptor is not of type PasswordDecryptor")
				return
			}
			printPasswordInfo(pwdDecryptor)
		case yconfig.YubiKeyDecryptorType:
			yubiDecryptor, ok := cfgvar.Decryptor.Decryptor.(*yconfig.YubiKeyDecryptor)
			if !ok {
				fmt.Println("Decryptor is not of type YubiKeyDecryptor")
				return
			}
			printYubiKeyInfo(yubiDecryptor)
		default:
			fmt.Println("Unknown decryptor type")
		}
		fmt.Println("-----")
	}
}

func printNullInfo(nullDecryptor *yconfig.NullDecryptor) {
	fmt.Println("Type: ", nullDecryptor.GetType())
}

func printPasswordInfo(pwdDecryptor *yconfig.PasswordDecryptor) {
	fmt.Println("Type:     ", pwdDecryptor.GetType())
	fmt.Println("Password: ", pwdDecryptor.Password)
}

func printYubiKeyInfo(ykDecryptor *yconfig.YubiKeyDecryptor) {
	fmt.Println("Type: ", ykDecryptor.GetType())
	fmt.Println("Slot: ", ykDecryptor.Slot)
	fmt.Println("PIN:  ", ykDecryptor.PIN)
}
