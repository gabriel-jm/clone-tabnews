package clone_tabnews_test

import (
	"log"
	"testing"

	"github.com/gabriel-jm/clone-tabnews/src/db"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIntegration(t *testing.T) {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatalf("Loading env error: %v", err)
	}

	db.Init()

	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Setup")
}
