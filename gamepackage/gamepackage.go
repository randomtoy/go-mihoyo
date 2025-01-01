package gamepackage

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Launcher struct {
	Url        string
	GameID     string
	LauncherID string
}

func Get(launcher Launcher, target any) error {

	if launcher.Url == "" {
		return fmt.Errorf("launcher URL is empty")
	}
	if launcher.LauncherID == "" {
		return fmt.Errorf("launcher ID is empty")
	}

	stringUrl := fmt.Sprintf("%s?launcher_id=%s", launcher.Url, launcher.LauncherID)
	if launcher.GameID != "" {
		stringUrl = fmt.Sprintf("%s&game_ids[]=%s", stringUrl, launcher.GameID)
	}

	response, err := http.Get(stringUrl)
	if err != nil {
		return fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return nil
}
