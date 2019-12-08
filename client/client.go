package client

import (
	"encoding/json"
	"github.com/RocketLauncherFF/rocketlauncherff/core"
	"net/http"
)

type ffclient interface {
	Enabled(string) (bool, error)
	EnabledWithDefault(string, bool) bool
}

type FFClient struct {
	serverName string
	httpClient *http.Client
}

func NewFFClient(serverName string) *FFClient {
	return &FFClient{serverName: serverName, httpClient: &http.Client{}}
}

func (client *FFClient) Enabled(ff string) (bool, error) {
	var responseFF core.FeatureFlag
	resp, err := client.httpClient.Get(client.serverName + "/" + ff)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&responseFF)
	return responseFF.Enabled, err
}

func (client *FFClient) EnabledWithDefault(ff string, def bool) bool {
	enabled, err := client.Enabled(ff)
	if err != nil {
		return def
	}
	return enabled
}
