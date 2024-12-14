// Makesure the server is kept running. Folder:= webServer
// strings.Builder := It is a package. It is used to effeciently build a string using Write method.
// The content we receive is in byte format. That's why we have to convert it into string readable format for us to understand.

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb")
	// PeformGetRequest()
	// PeformPostJsonRequest()
	PeformPostFormRequest()
}

func PeformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length:", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	// fmt.Println(content)
	// fmt.Println(string(content))
	
	fmt.Println("ByteCount is:", byteCount)
	fmt.Println(responseString.String())
	
}

func PeformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	// Fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price": 0,
			"platform":"learnCodeOnline.in"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PeformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	// formdata
	data := url.Values{}
	data.Add("firstname", "mayur")
	data.Add("lastname", "patil")
	data.Add("email", "patilmayur.2745@gmail.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
