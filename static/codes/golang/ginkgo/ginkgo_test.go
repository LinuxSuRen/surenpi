package test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Books Suite")
}

var _ = Describe("Book", func() {
	var (
		longBook  Book
		shortBook Book
	)

	BeforeEach(func() {
		longBook = Book{
			Name: "guide",
		}
		shortBook = Book{
			Name: "script",
		}
	})

	Describe("Different names", func() {
		It("success", func() {
			Expect(longBook.Name).To(Equal("guide"))
			Expect(shortBook.Name).To(Equal("script"))
			Expect(shortBook.Name).Should(Equal("script"))
			Expect(shortBook.Published).To(BeFalse())
			Expect(shortBook.Published).NotTo(BeTrue())
			Expect(shortBook.Published).ToNot(BeTrue())

			_, err := foo()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(foo()).Should(Succeed())
		})

		Context("page number", func() {
			JustBeforeEach(func() {
				longBook.Pages = 10
			})

			It("pages", func() {
				Expect(longBook.Pages).To(Equal(10))
			})
		})
	})
})
