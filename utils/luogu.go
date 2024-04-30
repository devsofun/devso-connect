package utils

import (
	"io"
	"net/http"
	"strings"
)

func CheckParameterInURL(url string, parameter string) (bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	userAgent := `Devso Connect`
	resp.Header.Set("User-Agent", userAgent)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	content := string(body)
	println(content)

	if strings.Contains(content, parameter) {
		return true, nil
	}

	return false, nil
}
