package main_test

import (
	"fmt"
	. "gowiki"
	"io/ioutil"
	"log"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	//"io/ioutil"

	//"net/http/httptest"
)

var _ = Describe("Gowiki", func() {

	BeforeEach(func() {
		//go StartServer()
	})

	p := &Page{Title: "WikiTitle", Body: []byte("WikiBody")}

	Context("Checks if true equals true", func() {
		It("True equal true?", func() {
			Expect(true).To(Equal(true))
		})
	})

	Context("Test View a page", func() {
		server := ghttp.NewServer()
		server.AppendHandlers(ViewHandler)
		fmt.Println(server.Addr())
		res, err := http.Get("http://" + server.Addr())
		if err != nil {
			log.Fatal(err)
		}
		greeting, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", greeting)
		//It("Title should be WikiTitle", func() {
		//	Expect(p.Title).To(Equal("WikiTitle"))
		//})

		//It("Save should save a file called WikiTitle.txt", func() {
		//	p.Save()
		//filename := p.Title + ".txt"
		//ioutil.WriteFile(filename, p.Body, 0600)
		//	Ω("WikiTitle.txt").Should(BeAnExistingFile())
		//})
	})

	Context("Test Load and View a page", func() {
		It("Load should load a file called WikiTitle.txt", func() {
			LoadPage(p.Title)
			// How do we check if Page was loaded and viewable?
		})
	})

	//Describe("The Gowiki Client", func() {
	//var server *ghttp.Server
	//var client *gowiki.Client

	//BeforeEach(func() {
	//server = ghttp.NewServer()
	//client = gowiki.NewClient(server.URL())
	//})

	//AfterEach(func() {
	//shut down the server between tests
	//server.Close()
	//})

	//})

})
