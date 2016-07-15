package dependencies_test

import (
	"github.com/octoblu/go-test-the-fs-and-http/dependencies"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Install", func() {
	It("should be a thing", func() {
		Expect(dependencies.Install).NotTo(BeNil())
	})
})
