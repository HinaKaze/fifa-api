package api

func GetChromeHeader() map[string]string {
	chromeHeader := make(map[string]string)
	chromeHeader["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"
	chromeHeader["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"
	chromeHeader["Accept-Encoding"] = "gzip,deflate,sdch, br"
	chromeHeader["Accept-Language"] = "en-US,en;q=0.8"
	chromeHeader["Connection"] = "keep-alive"
	chromeHeader["DNT"] = "1"
	return chromeHeader
}
