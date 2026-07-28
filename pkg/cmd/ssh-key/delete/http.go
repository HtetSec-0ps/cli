package delete

import (
	"fmt"
	"net/http"

	"github.com/cli/cli/v2/api"
)

type sshKey struct {
	Title string
}

func deleteSSHKey(httpClient *http.Client, host string, keyID string) error {
	path := fmt.Sprintf("user/keys/%s", keyID)
	// TODO(api-client-rollout)
	// This line of code is part of a mechanical roll out of the api client.
	// As a follow up, consider whether the api client can be injected to this call site, rather than constructed
	return api.NewClientFromHTTP(httpClient).REST(host, http.MethodDelete, path, nil, nil)
}

func getSSHKey(httpClient *http.Client, host string, keyID string) (*sshKey, error) {
	var key sshKey
	path := fmt.Sprintf("user/keys/%s", keyID)
	// TODO(api-client-rollout)
	// This line of code is part of a mechanical roll out of the api client.
	// As a follow up, consider whether the api client can be injected to this call site, rather than constructed
	err := api.NewClientFromHTTP(httpClient).REST(host, http.MethodGet, path, nil, &key)
	if err != nil {
		return nil, err
	}

	return &key, nil
}
