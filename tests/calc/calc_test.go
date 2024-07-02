package clone_tabnews_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/gabriel-jm/clone-tabnews/src/calc"
)

var _ = Describe("Calc", func() {
	It("Should return the correct value from a sum", func() {
		value := Calc(1, 2)
		Expect(value).To(Equal(3))
	})
})
