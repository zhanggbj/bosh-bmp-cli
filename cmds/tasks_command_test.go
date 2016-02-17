package cmds_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	cmds "github.com/maximilien/bosh-bmp-cli/cmds"
)

var _ = Describe("tasks command", func() {
	var (
		args    []string
		options cmds.Options
		cmd     cmds.Command
	)

	BeforeEach(func() {
		args = []string{"bmp", "tasks"}
		options = cmds.Options{
			Verbose: false,
		}

		cmd = cmds.NewTasksCommand(options)
	})

	Describe("NewTasksCommand", func() {
		It("create new TasksCommand", func() {
			Expect(cmd).ToNot(BeNil())

			cmd2 := cmds.NewTasksCommand(options)
			Expect(cmd2).ToNot(BeNil())
			Expect(cmd2).To(Equal(cmd))
		})
	})

	Describe("#Name", func() {
		It("returns the name of a TasksCommand", func() {
			Expect(cmd.Name()).To(Equal("tasks"))
		})
	})

	Describe("#Description", func() {
		It("returns the description of a TasksCommand", func() {
			Expect(cmd.Description()).To(Equal(`List tasks: \"option --latest number\", \"The number of latest tasks, default is 50\"`))
		})
	})

	Describe("#Usage", func() {
		It("returns the usage text of a TasksCommand", func() {
			Expect(cmd.Usage()).To(Equal("bmp tasks"))
		})
	})

	Describe("#Options", func() {
		It("returns the options of a TasksCommand", func() {
			Expect(cmds.EqualOptions(cmd.Options(), options)).To(BeTrue())
		})
	})

	Describe("#Validate", func() {
		It("validates a good TasksCommand", func() {
			validate, err := cmd.Validate()
			Expect(validate).To(BeTrue())
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe("#Execute", func() {
		It("executes a good TasksCommand", func() {
			rc, err := cmd.Execute(args)
			Expect(rc).To(Equal(0))
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
