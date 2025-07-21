package cmd_test

import (
	"CLI-Calculator-2/cmd"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("add command", func() {
	It("subs 10 - 3 and gets 7", func() {
		output, err := cmd.ExecCmd(cmd.SubCmd(), "10", "3")
		Expect(err).To(BeNil())
		Expect(output).To(Equal("Result: 7\n"))
	})
})
