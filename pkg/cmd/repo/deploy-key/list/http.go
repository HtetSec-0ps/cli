package list

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cli/cli/v2/api"
	"github.com/cli/cli/v2/internal/ghrepo"
)

type deployKey struct {
	ID        int64     `json:"id"`
	Key       string    `json:"key"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	ReadOnly  bool      `json:"read_only"`
}

func repoKeys(httpClient *http.Client, repo ghrepo.Interface) ([]deployKey, error) {
	path := fmt.Sprintf("repos/%s/%s/keys?per_page=100", repo.RepoOwner(), repo.RepoName())

	var keys []deployKey
	// TODO(api-client-rollout)
	// This line of code is part of a mechanical roll out of the api client.
	// As a follow up, consider whether the api client can be injected to this call site, rather than constructed
	err := api.NewClientFromHTTP(httpClient).REST(repo.RepoHost(), "GET", path, nil, &keys)
	if err != nil {
		return nil, err
	}

	return keys, nil
}
