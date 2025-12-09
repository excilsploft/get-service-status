package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
)

// const LogPrefix = "PyPi"
// const StatusUrl = "https://status.python.org/api/v2/summary.json"

const LogPrefix = "Github"
const StatusUrl = "https://www.githubstatus.com/api/v2/summary.json"

// const LogPrefix = "NPM"
// const StatusUrl = "https://status.npmjs.org/api/v2/summary.json"

type Summary struct {
	Status `json:"status"`
}

type Status struct {
	Indicator   string `json:"indicator"`
	Description string `json:"description"`
}

func main() {

	status := Summary{}
	outLog := slog.New(slog.NewTextHandler(os.Stdout, nil))
	resp, err := http.Get(StatusUrl)
	if err != nil {
		outLog.Error(LogPrefix, "error", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	decodeErr := json.NewDecoder(resp.Body).Decode(&status)
	if decodeErr != nil {
		outLog.Error(LogPrefix, "error", "Error Decoding JSON Response")
		os.Exit(1)
	}

	switch status.Indicator {
	case "critical":
		outLog.Error(LogPrefix, "status", status.Description)
	case "major":
		outLog.Error(LogPrefix, "status", status.Description)
	case "minor":
		outLog.Error(LogPrefix, "status", status.Description)
	default:
		outLog.Info(LogPrefix, "status", status.Description)
	}

}
