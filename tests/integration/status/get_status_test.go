package clone_tabnews_test

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	integration_setup "github.com/gabriel-jm/clone-tabnews/tests/integration/setup"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tkuchiki/faketime"
)

func getBody(res *http.Response) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	body := strings.TrimSuffix(buf.String(), "\n")

	return body
}

var _ = Describe("/api/v1/status", Ordered, func() {
	var url string
	var testServer *httptest.Server
	f := faketime.NewFaketime(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

	BeforeAll(func() {
		testServer = integration_setup.RunTestServer()
		url = fmt.Sprintf("%s/api/v1/status", testServer.URL)
		f.Do()
	})

	AfterAll(func() {
		testServer.Close()
		f.Undo()
	})

	It("Should return 200 when server is health", func(ctx SpecContext) {
		resp, err := http.Get(url)

		if err != nil {
			log.Fatalf("Error: %v", err)
			ctx.Err()
		}

		Expect(resp.StatusCode).To(Equal(200))

		responseBody := getBody(resp)

		Expect(responseBody).To(Equal(`{"updatedAt":"2009-11-10T23:00:00Z"}`))
	})
})
