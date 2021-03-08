package internal_test

import (
	"context"
	"errors"
	"testing"

	"github.com/elimity-com/backend-intern-exercise/internal"
	"github.com/google/go-github/v33/github"
)

var errMock = errors.New("mock")

func TestTrack(t *testing.T) {
	mock := mock{t: t}
	if err := internal.Track(mock.repositories, 1); err != errMock {
		t.Fatalf("got invalid error: %v", err)
	}
}

type mock struct {
	t *testing.T
}

func (m mock) repositories(_ context.Context, query string, options *github.SearchOptions) (*github.RepositoriesSearchResult, *github.Response, error) {
	t := m.t
	if query != "is:public" {
		t.Fatalf("got invalid query: %s", query)
	}
	if sort := options.Sort; sort != "updated" {
		t.Fatalf("got invalid sort: %s", sort)
	}
	return nil, nil, errMock
}
