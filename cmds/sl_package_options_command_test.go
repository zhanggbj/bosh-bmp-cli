package cmds_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	cmds "github.com/maximilien/bosh-bmp-cli/cmds"
)

var _ = Describe("sl-package-options command", func() {

	var (
		args    []string
		options cmds.Options
		cmd     cmds.Command
	)

	BeforeEach(func() {
		args = []string{"bmp", "sl-package-options"}
		options = cmds.Options{
			Verbose: false,
		}

		cmd = cmds.NewSlPackageOptionsCommand(options)
	})

	Describe("NewSlPackageOptionsCommand", func() {
		It("create new SlPackageOptionsCommand", func() {
			Expect(cmd).ToNot(BeNil())

			cmd2 := cmds.NewSlPackageOptionsCommand(options)
			Expect(cmd2).ToNot(BeNil())
			Expect(cmd2).To(Equal(cmd))
		})
	})

	Describe("#Name", func() {
		It("returns the name of a SlPackageOptionsCommand", func() {
			Expect(cmd.Name()).To(Equal("sl-package-options"))
		})
	})

	Describe("#Description", func() {
		It("returns the description of a SlPackageOptionsCommand", func() {
			Expect(cmd.Description()).To(Equal("List all options of Softlayer package"))
		})
	})

	Describe("#Usage", func() {
		It("returns the usage text of a SlPackageOptionsCommand", func() {
			Expect(cmd.Usage()).To(Equal("bmp sl-package-options"))
		})
	})

	Describe("#Validate", func() {
		It("validates a good SlPackageOptionsCommand", func() {
			validate, err := cmd.Validate()
			Expect(validate).To(BeTrue())
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe("#Execute", func() {
		It("executes a good SlPackageOptionsCommand", func() {
			rc, err := cmd.Execute(args)
			Expect(rc).To(Equal(0))
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
