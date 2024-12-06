package zip_files_test

import (
	"github.com/metadiv-io/zip_files"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Zipper", func() {
	It("should zip files and unzip them", func() {
		zipper := zip_files.New()
		Expect(zipper.CountFiles()).To(Equal(0))

		zipper.AddFile("test.txt", []byte("test"))
		zipper.AddFile("test2.txt", []byte("test2"))
		Expect(zipper.CountFiles()).To(Equal(2))
		Expect(zipper.Files()).To(Equal([]string{"test.txt", "test2.txt"}))

		zip, err := zipper.Zip()
		Expect(err).NotTo(HaveOccurred())
		Expect(len(zip)).To(BeNumerically(">", 0))

		files, err := zip_files.Unzip(zip)
		Expect(err).NotTo(HaveOccurred())
		Expect(files).To(Equal(map[string][]byte{"test.txt": []byte("test"), "test2.txt": []byte("test2")}))
	})
})
