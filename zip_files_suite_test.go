package zip_files_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestZipFiles(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ZipFiles Suite")
}
