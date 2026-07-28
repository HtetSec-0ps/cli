package shared

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cli/cli/v2/api"
)

const (
	AuthenticationKey = "authentication"
	SigningKey        = "signing"
)

type sshKey struct {
	ID        int64
	Key       string
	Title     string
	Type      string
	CreatedAt time.Time `json:"created_at"`
}

func UserKeys(httpClient *http.Client, host, userHandle string) ([]sshKey, error) {
	resource := "user/keys"
	if userHandle != "" {
		resource = fmt.Sprintf("users/%s/keys", userHandle)
	}
	path := fmt.Sprintf("%s?per_page=%d", resource, 100)

	keys, err := getUserKeys(httpClient, host, path)

	if err != nil {
		return nil, err
	}

	for i := 0; i < len(keys); i++ {
		keys[i].Type = AuthenticationKey
	}

	return keys, nil
}

func UserSigningKeys(httpClient *http.Client, host, userHandle string) ([]sshKey, error) {
	resource := "user/ssh_signing_keys"
	if userHandle != "" {
		resource = fmt.Sprintf("users/%s/ssh_signing_keys", userHandle)
	}
	path := fmt.Sprintf("%s?per_page=%d", resource, 100)

	keys, err := getUserKeys(httpClient, host, path)

	if err != nil {
		return nil, err
	}

	for i := 0; i < len(keys); i++ {
		keys[i].Type = SigningKey
	}

	return keys, nil
}

func getUserKeys(httpClient *http.Client, hostname, path string) ([]sshKey, error) {
	var keys []sshKey
	// TODO(api-client-rollout)
	// This line of code is part of a mechanical roll out of the api client.
	// As a follow up, consider whether the api client can be injected to this call site, rather than constructed
	err := api.NewClientFromHTTP(httpClient).REST(hostname, http.MethodGet, path, nil, &keys)
	if err != nil {
		return nil, err
	}

	return keys, nil
}
