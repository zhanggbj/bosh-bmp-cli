package common_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	common "github.com/maximilien/bosh-bmp-cli/common"
	fakes "github.com/maximilien/bosh-bmp-cli/common/fakes"
)

var _ = Describe("DefaultPrinter", func() {
	var (
		printer common.Printer
		fakeUi  *fakes.FakeUi
	)

	Describe("when verbose is true", func() {
		BeforeEach(func() {
			printer = common.NewDefaultPrinter(fakeUi, true)
		})

		It("#Printf", func() {
			Expect(true).To(Equal(false))
		})

		It("#Println", func() {
			Expect(true).To(Equal(false))
		})
	})

	Describe("when verbose is false", func() {
		BeforeEach(func() {
			printer = common.NewDefaultPrinter(fakeUi, false)
		})

		It("#Printf", func() {
			Expect(true).To(Equal(false))
		})

		It("#Println", func() {
			Expect(true).To(Equal(false))
		})
	})
})
