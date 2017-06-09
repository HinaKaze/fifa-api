package api

func GetURLs() map[string]string {
	platformHosts := make(map[string]string)
	platformHosts["pc"] = "https://utas.external.s2.fut.ea.com:443"
	platformHosts["ps3"] = "https://utas.external.s2.fut.ea.com:443"
	platformHosts["ps4"] = "https://utas.external.s2.fut.ea.com:443"
	platformHosts["xbox"] = "https://utas.external.s3.fut.ea.com:443"
	platformHosts["xbox360"] = "https://utas.external.s3.fut.ea.com:443"
	platformHosts["ios"] = "https://utas.external.fut.ea.com:443"
	platformHosts["and"] = "https://utas.external.fut.ea.com:443"

	urls := make(map[string]string)
	urls["main_site"] = "https://www.easports.com"
	urls["futweb"] = "https://www.easports.com/iframe/fut17/?baseShowoffUrl=https%3A%2F%2Fwww.easports.com%2Ffifa%2Fultimate-team%2Fweb-app%2Fshow-off&guest_app_uri=http%3A%2F%2Fwww.easports.com%2Ffifa%2Fultimate-team%2Fweb-app&locale=en_US"
	urls["fut_config"] = "https://www.easports.com/iframe/fut17/bundles/futweb/web/flash/xml/site_config.xml"
	urls["fut_home"] = "https://www.easports.com/fifa/ultimate-team/web-app"
	urls["fut"] = ""
	urls["fut_question"] = "https://www.easports.com/iframe/fut17/p/ut/game/fifa17/phishing/question"
	urls["fut_validate"] = "https://www.easports.com/iframe/fut17/p/ut/game/fifa17/phishing/validate"
	urls["fut_auth"] = "https://www.easports.com/iframe/fut17/p/ut/auth"
	urls["fut_captcha_img"] = "https://www.easports.com/iframe/fut17/p/ut/captcha/img"
	urls["fut_captcha_validate"] = "https://www.easports.com/iframe/fut17/p/ut/captcha/img"
	urls["fut_host"] = platformHosts["ps4"]
	urls["shards"] = "https://www.easports.com/iframe/fut17/p/ut/shards/v2"
	urls["acc_info"] = "https://www.easports.com/iframe/fut17/p/ut/game/fifa17/user/accountinfo"
	urls["card_info"] = "https://fifa17.content.easports.com/fifa/fltOnlineAssets/CC8267B6-0817-4842-BB6A-A20F88B05418/2017/fut/items/web/"
	urls["messages"] = "https://www.easports.com/iframe/fut17/bundles/futweb/web/flash/xml/localization/messages.en_US.xml"
	urls["mass_info"] = "/ut/game/fifa17/usermassinfo"
	return urls
}

//TODO update url
