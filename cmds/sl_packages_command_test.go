package cmds_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	cmds "github.com/maximilien/bosh-bmp-cli/cmds"
)

var _ = Describe("sl-packages command", func() {

	var (
		args    []string
		options cmds.Options
		cmd     cmds.Command
	)

	BeforeEach(func() {
		args = []string{"bmp", "sl-packages"}
		options = cmds.Options{
			Verbose: false,
		}

		cmd = cmds.NewSlPackagesCommand(options)
	})

	Describe("NewSlPackagesCommand", func() {
		It("create new SlPackagesCommand", func() {
			Expect(cmd).ToNot(BeNil())

			cmd2 := cmds.NewSlPackagesCommand(options)
			Expect(cmd2).ToNot(BeNil())
			Expect(cmd2).To(Equal(cmd))
		})
	})

	Describe("#Name", func() {
		It("returns the name of a SlPackagesCommand", func() {
			Expect(cmd.Name()).To(Equal("sl-packages"))
		})
	})

	Describe("#Description", func() {
		It("returns the description of a SlPackagesCommand", func() {
			Expect(cmd.Description()).To(Equal("List all Softlayer packages"))
		})
	})

	Describe("#Usage", func() {
		It("returns the usage text of a SlPackagesCommand", func() {
			Expect(cmd.Usage()).To(Equal("bmp sl-packages"))
		})
	})

	Describe("#Validate", func() {
		It("validates a good SlPackagesCommand", func() {
			validate, err := cmd.Validate()
			Expect(validate).To(BeTrue())
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe("#Execute", func() {
		It("executes a good SlPackagesCommand", func() {
			rc, err := cmd.Execute(args)
			Expect(rc).To(Equal(0))
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
