package helpers

import (
	_ "embed"
)

//go:embed logo.txt
var bannerText string

func GetBannerText() string {

	return bannerText
}

func GetCommandNameRequiredText() string {
	return "Please provide a command name as an argument."
}
