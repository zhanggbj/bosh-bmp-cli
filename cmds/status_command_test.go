package cmds_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	cmds "github.com/maximilien/bosh-bmp-cli/cmds"
)

var _ = Describe("status command", func() {

	var (
		args    []string
		options cmds.Options
		cmd     cmds.Command
	)

	BeforeEach(func() {
		args = []string{"bmp", "status"}
		options = cmds.Options{
			Verbose: false,
		}

		cmd = cmds.NewStatusCommand(options)
	})

	Describe("NewStatusCommand", func() {
		It("create new StatusCommand", func() {
			Expect(cmd).ToNot(BeNil())

			cmd2 := cmds.NewStatusCommand(options)
			Expect(cmd2).ToNot(BeNil())
			Expect(cmd2).To(Equal(cmd))
		})
	})

	Describe("#Name", func() {
		It("returns the name of a StatusCommand", func() {
			Expect(cmd.Name()).To(Equal("status"))
		})
	})

	Describe("#Description", func() {
		It("returns the description of a StatusCommand", func() {
			Expect(cmd.Description()).To(Equal("show bmp status"))
		})
	})

	Describe("#Usage", func() {
		It("returns the usage text of a StatusCommand", func() {
			Expect(cmd.Usage()).To(Equal("bmp status"))
		})
	})

	Describe("#Validate", func() {
		It("validates a good StatusCommand", func() {
			validate, err := cmd.Validate()
			Expect(validate).To(BeTrue())
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe("#Execute", func() {
		It("executes a good StatusCommand", func() {
			rc, err := cmd.Execute(args)
			Expect(rc).To(Equal(0))
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
