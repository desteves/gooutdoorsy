package rental_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRv(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rv Suite")
}
