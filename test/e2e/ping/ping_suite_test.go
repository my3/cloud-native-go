package ping

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ping Suite")
}
