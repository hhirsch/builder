package helpers

func GetBannerText() string {
	var bannerText string = `
 ____        _ _     _           
| __ ) _   _(_) | __| | ___ _ __ 
|  _ \| | | | | |/ _' |/ _ \ '__|
| |_) | |_| | | | (_| |  __/ |   
|____/ \__,_|_|_|\__,_|\___|_|   
`
	return bannerText
}

func GetHelpText() string {
	var bannerText string = `
Builder - server maintenance tool

builder [command]

help      show this help
init      initialize project in directory
update    run pending migrations scripts
migration re-run specific migration
script    run specific script
`
	return bannerText
}
