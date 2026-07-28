package delete

import (
	"fmt"
	"net/http"

	"github.com/cli/cli/v2/api"
)

type gpgKey struct {
	ID    int64
	KeyID string `json:"key_id"`
}

func deleteGPGKey(httpClient *http.Client, host, id string) error {
	path := fmt.Sprintf("user/gpg_keys/%s", id)
	// TODO(api-client-rollout)
	// This line of code is part of a mechanical roll out of the api client.
	// As a follow up, consider whether the api client can be injected to this call site, rather than constructed
	return api.NewClientFromHTTP(httpClient).REST(host, "DELETE", path, nil, nil)
}

func getGPGKeys(httpClient *http.Client, host string) ([]gpgKey, error) {
	var keys []gpgKey
	// TODO(api-client-rollout)
	// This line of code is part of a mechanical roll out of the api client.
	// As a follow up, consider whether the api client can be injected to this call site, rather than constructed
	err := api.NewClientFromHTTP(httpClient).REST(host, "GET", "user/gpg_keys?per_page=100", nil, &keys)
	if err != nil {
		return nil, err
	}
	return keys, nil
}
