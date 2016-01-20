package main_test

import (
	. "gowiki"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	//"io/ioutil"
)

var _ = Describe("Gowiki", func() {

	p := &Page{Title: "WikiTitle", Body: []byte("WikiBody")}

	Context("Checks if true equals true", func() {
		It("True equal true?", func() {
			Expect(true).To(Equal(true))
		})
	})

	Context("Test Save a page", func() {
		It("Title should be WikiTitle", func() {
			Expect(p.Title).To(Equal("WikiTitle"))
		})

		It("Save should save a file called WikiTitle.txt", func() {
			p.Save()
			//filename := p.Title + ".txt"
			//ioutil.WriteFile(filename, p.Body, 0600)
			Ω("WikiTitle.txt").Should(BeAnExistingFile())
		})
	})

	Context("", func() {
		It("", func() {
			Expect().To(Equal(""))
		})
	})

	Context("", func() {
		It("", func() {
			Expect().To(Equal(""))
		})
	})

	Context("", func() {
		It("", func() {
			Expect().To(Equal(""))
		})
	})

})
