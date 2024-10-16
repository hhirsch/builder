package helpers

import (
	_ "embed"
)

//go:embed logo.txt
var bannerText string

//go:embed help.txt
var helpText string

func GetBannerText() string {

	return bannerText
}

func GetHelpText() string {

	return helpText
}

func GetCommandNameRequiredText() string {
	return "Please provide a command name as an argument"
}
