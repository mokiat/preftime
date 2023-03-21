package prefix_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"

	"github.com/mokiat/preftime/prefix"
)

var _ = Describe("Func", func() {
	var (
		prefixFunc prefix.Func
		out        *gbytes.Buffer
	)

	BeforeEach(func() {
		out = gbytes.NewBuffer()
	})

	Describe("IndexFunc", func() {
		BeforeEach(func() {
			prefixFunc = prefix.IndexFunc()
		})

		It("should write incrementing number", func() {
			_, err := prefixFunc(out)
			Expect(err).ToNot(HaveOccurred())
			Expect(out).To(gbytes.Say("\\[1\\] "))

			_, err = prefixFunc(out)
			Expect(err).ToNot(HaveOccurred())
			Expect(out).To(gbytes.Say("\\[2\\] "))
		})

		When("a write error occurs", func() {
			BeforeEach(func() {
				out.Close()
			})

			It("should report an error", func() {
				_, err := prefixFunc(out)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("TimestampFunc", func() {
		BeforeEach(func() {
			prefixFunc = prefix.TimestampFunc()
		})

		It("should write timestamps", func() {
			_, err := prefixFunc(out)
			Expect(err).ToNot(HaveOccurred())
			Expect(out).To(gbytes.Say("\\[\\d\\d\\d\\d-\\d\\d-\\d\\d \\d\\d:\\d\\d:\\d\\d.\\d\\d\\d\\] "))
		})

		When("a write error occurs", func() {
			BeforeEach(func() {
				out.Close()
			})

			It("should report an error", func() {
				_, err := prefixFunc(out)
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
