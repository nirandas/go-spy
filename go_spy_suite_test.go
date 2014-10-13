package spy_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoSpy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoSpy Suite")
}
