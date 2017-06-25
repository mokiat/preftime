package prefix_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPrefix(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Prefix Suite")
}
