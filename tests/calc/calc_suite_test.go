package clone_tabnews_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCalc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tests/Calc/Calc")
}
