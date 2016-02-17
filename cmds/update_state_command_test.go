package cmds_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	cmds "github.com/maximilien/bosh-bmp-cli/cmds"
)

var _ = Describe("update-state command", func() {
	var (
		args    []string
		options cmds.Options
		cmd     cmds.Command
	)

	BeforeEach(func() {
		args = []string{"bmp", "update-state"}
		options = cmds.Options{
			Verbose: false,
		}

		cmd = cmds.NewUpdateStateCommand(options)
	})

	Describe("NewUpdateStateCommand", func() {
		It("create new UpdateStateCommand", func() {
			Expect(cmd).ToNot(BeNil())

			cmd2 := cmds.NewUpdateStateCommand(options)
			Expect(cmd2).ToNot(BeNil())
			Expect(cmd2).To(Equal(cmd))
		})
	})

	Describe("#Name", func() {
		It("returns the name of a UpdateStateCommand", func() {
			Expect(cmd.Name()).To(Equal("update-state"))
		})
	})

	Describe("#Description", func() {
		It("returns the description of a UpdateStateCommand", func() {
			Expect(cmd.Description()).To(Equal(`Update the server state (\"bm.state.new\", \"bm.state.using\", \"bm.state.loading\", \"bm.state.failed\", \"bm.state.deleted\")`))
		})
	})

	Describe("#Usage", func() {
		It("returns the usage text of a UpdateStateCommand", func() {
			Expect(cmd.Usage()).To(Equal("bmp update-state <bm.state.new>"))
		})
	})

	Describe("#Options", func() {
		It("returns the options of a UpdateStateCommand", func() {
			Expect(cmds.EqualOptions(cmd.Options(), options)).To(BeTrue())
		})
	})

	Describe("#Validate", func() {
		It("validates a good UpdateStateCommand", func() {
			validate, err := cmd.Validate()
			Expect(validate).To(BeTrue())
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe("#Execute", func() {
		It("executes a good UpdateStateCommand", func() {
			rc, err := cmd.Execute(args)
			Expect(rc).To(Equal(0))
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
