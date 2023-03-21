package prefix_test

import (
	"fmt"
	"io"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"

	"github.com/mokiat/preftime/prefix"
)

var _ = Describe("Writer", func() {
	var (
		writer io.Writer
		out    *gbytes.Buffer
	)

	BeforeEach(func() {
		out = gbytes.NewBuffer()
		writer = prefix.NewWriter(out, prefix.IndexFunc())
	})

	It("should prefix whole lines", func() {
		const line = "a whole line\n"
		count, err := fmt.Fprint(writer, line)
		Expect(err).ToNot(HaveOccurred())
		Expect(count).To(Equal(len(line)))
		Expect(out).To(gbytes.Say("\\[1\\] a whole line\n"))
	})

	It("should prefix partial lines", func() {
		const line = "a partial line"
		count, err := fmt.Fprint(writer, line)
		Expect(err).ToNot(HaveOccurred())
		Expect(count).To(Equal(len(line)))
		Expect(out).To(gbytes.Say("\\[1\\] a partial line"))
	})

	It("should prefix multiple lines", func() {
		const line1 = "a whole line\nand "
		firstCount, err := fmt.Fprint(writer, line1)
		Expect(err).ToNot(HaveOccurred())
		Expect(firstCount).To(Equal(len(line1)))

		const line2 = "a partial one"
		secondCount, err := fmt.Fprint(writer, line2)
		Expect(err).ToNot(HaveOccurred())
		Expect(secondCount).To(Equal(len(line2)))

		Expect(out).To(gbytes.Say("\\[1\\] a whole line\n\\[2\\] and a partial one"))
	})
})
