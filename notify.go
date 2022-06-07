package nogisched

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Notify(msg string) (string, error) {
	token := os.Getenv("NOGISCHED_NOTIFY_TOKEN")
	if token == "" {
		return "", fmt.Errorf("the environment variable NOGISCHED_NOTIFY_TOKEN is NOT set")
	}
	client := &http.Client{}
	param := url.Values{}
	param.Add("message", msg)
	req, err := http.NewRequest(http.MethodPost,
		"https://notify-api.line.me/api/notify",
		strings.NewReader(param.Encode()))
	if err != nil {
		return "", fmt.Errorf("request create error: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request send error: %w", err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("response read error: %w", err)
	}
	return string(bodyText), nil
}
