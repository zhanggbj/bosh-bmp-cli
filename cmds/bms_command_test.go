package cmds_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	cmds "github.com/maximilien/bosh-bmp-cli/cmds"
)

var _ = Describe("bms command", func() {

	var (
		args    []string
		options cmds.Options
		cmd     cmds.Command
	)

	BeforeEach(func() {
		args = []string{"bmp", "bms"}
		options = cmds.Options{
			Verbose: false,
		}

		cmd = cmds.NewBmsCommand(args, options)
	})

	Describe("NewBmsCommand", func() {
		It("create new BmsCommand", func() {
			Expect(cmd).ToNot(BeNil())

			cmd2 := cmds.NewBmsCommand(args, options)
			Expect(cmd2).ToNot(BeNil())
			Expect(cmd2).To(Equal(cmd))
		})
	})
})
