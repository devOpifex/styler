package main

func isShadow(token string) bool {
	return token == "shadow"
}

func makeShadow(token, next string) (string, bool) {
	if !isShadow(token) {
		return "", false
	}

	if next == "sm" {
		return `-webkit-box-shadow: 0px 10px 5px 0px rgba(0,0,0,0.1);-moz-box-shadow: 0px 10px 5px 0px rgba(0,0,0,0.1);box-shadow: 0px 10px 5px 0px rgba(0,0,0,0.1)`, true
	}

	if next == "md" {
		return `-webkit-box-shadow: 0px 13px 12px -3px rgba(0,0,0,0.23);-moz-box-shadow: 0px 13px 12px -3px rgba(0,0,0,0.23);box-shadow: 0px 13px 12px -3px rgba(0,0,0,0.23)`, true
	}

	if next == "lg" {
		return `-webkit-box-shadow: 0px 13px 12px 0px rgba(0,0,0,0.32);-moz-box-shadow: 0px 13px 12px 0px rgba(0,0,0,0.32);box-shadow: 0px 13px 12px 0px rgba(0,0,0,0.32)`, true
	}

	return "", false
}
