package prefix_test

import (
	. "github.com/mokiat/preftime/prefix"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Function", func() {
	var prefixFunc PrefixFunction
	var out *gbytes.Buffer

	BeforeEach(func() {
		out = gbytes.NewBuffer()
	})

	Describe("IndexPrefixFunction", func() {
		BeforeEach(func() {
			prefixFunc = NewIndexPrefixFunction()
		})

		It("should write incrementing number", func() {
			_, err := prefixFunc(out)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(out).Should(gbytes.Say("\\[1\\] "))

			_, err = prefixFunc(out)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(out).Should(gbytes.Say("\\[2\\] "))
		})

		Context("when a write error occurs", func() {
			BeforeEach(func() {
				out.Close()
			})

			It("should report an error", func() {
				_, err := prefixFunc(out)
				Ω(err).Should(HaveOccurred())
			})
		})
	})

	Describe("TimePrefixFunction", func() {
		BeforeEach(func() {
			prefixFunc = NewTimePrefixFunction()
		})

		It("should write timestamps", func() {
			_, err := prefixFunc(out)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(out).Should(gbytes.Say("\\[\\d\\d\\d\\d-\\d\\d-\\d\\d \\d\\d:\\d\\d:\\d\\d.\\d\\d\\d\\] "))
		})

		Context("when a write error occurs", func() {
			BeforeEach(func() {
				out.Close()
			})

			It("should report an error", func() {
				_, err := prefixFunc(out)
				Ω(err).Should(HaveOccurred())
			})
		})
	})
})
