package prefix_test

import (
	"fmt"
	"io"

	. "github.com/mokiat/preftime/prefix"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Writer", func() {
	var writer io.Writer
	var out *gbytes.Buffer

	BeforeEach(func() {
		out = gbytes.NewBuffer()
		writer = NewPrefixWriter(out, NewIndexPrefixFunction())
	})

	It("should prefix whole lines", func() {
		line := "a whole line\n"
		count, err := fmt.Fprint(writer, line)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(count).Should(Equal(len(line)))
		Ω(out).Should(gbytes.Say("\\[1\\] a whole line\n"))
	})

	It("should prefix partial lines", func() {
		line := "a partial line"
		count, err := fmt.Fprint(writer, line)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(count).Should(Equal(len(line)))
		Ω(out).Should(gbytes.Say("\\[1\\] a partial line"))
	})

	It("should prefix multiple lines", func() {
		line := "a whole line\nand "
		firstCount, err := fmt.Fprint(writer, line)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(firstCount).Should(Equal(len(line)))
		line = "a partial one"
		secondCount, err := fmt.Fprint(writer, line)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(secondCount).Should(Equal(len(line)))

		Ω(out).Should(gbytes.Say("\\[1\\] a whole line\n\\[2\\] and a partial one"))
	})
})
