package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/v33/github"
)

// RepositoriesFunc represents a function that retrieves GitHub repositories using the given query and options.
type RepositoriesFunc = func(context context.Context, query string, options *github.SearchOptions) (*github.RepositoriesSearchResult, *github.Response, error)

// Track tracks public GitHub repositories using the given function, continuously updating according to the given
// interval.
//
// The given interval must be greater than zero.
func Track(fun RepositoriesFunc, interval time.Duration) error {
	for ; ; <-time.Tick(interval) {
		con := context.Background()
		listOptions := github.ListOptions{PerPage: 3}
		searchOptions := &github.SearchOptions{ListOptions: listOptions, Sort: "updated"}
		result, _, err := fun(con, "is:public", searchOptions)
		if err != nil {
			return err
		}
		for _, repository := range result.Repositories {
			fmt.Println(*repository.Name)
		}
	}
}
