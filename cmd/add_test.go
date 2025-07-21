package cmd_test

import (
	"CLI-Calculator-2/cmd"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("add command", func() {
	It("adds 5 + 3 and gets 8", func() {
		output, err := cmd.ExecCmd(cmd.AddCmd(), "5", "3")
		Expect(err).To(BeNil())
		Expect(output).To(Equal("Result: 8\n"))
	})
})
