package feature

import (
	"os"
	"testing"

	"github.com/DATA-DOG/godog"
)

func TestMain(m *testing.M) {
	status := godog.RunWithOptions("REST API", func(s *godog.Suite) {
		FeatureContext(s)
	}, godog.Options{
		Format: "progress",
		Paths:  []string{"."},
		Output: os.Stdout,
	})

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}
