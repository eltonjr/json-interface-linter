//go:build mutation

package main_test

import (
	"testing"

	"github.com/gtramontina/ooze"
)

func TestMutation(t *testing.T) {
	ooze.Release(
		t,
		ooze.WithRepositoryRoot("../.."),
		ooze.Parallel(),
		ooze.WithMinimumThreshold(0.8),
		ooze.IgnoreSourceFiles("testdata\\/.*"),
	)
}
