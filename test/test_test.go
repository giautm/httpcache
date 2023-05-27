package test_test

import (
	"testing"

	"github.com/giautm/httpcache"
	"github.com/giautm/httpcache/test"
)

func TestMemoryCache(t *testing.T) {
	test.Cache(t, httpcache.NewMemoryCache())
}
