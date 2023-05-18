package gohttp_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	goformat "goroutine/helpers/request"
	"goroutine/helpers/slices"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func Test_getRequest(t *testing.T) {

	resp, err := http.Get(
		"https://jsonplaceholder.typicode.com/posts/1",
	)
	goerror(err)

	// method 1

	// read the response body
	// body, err := ioutil.ReadAll(resp.Body)
	// goerror(err)
	// defer resp.Body.Close()
	// var data map[string]interface{}
	// err = json.Unmarshal(body, &data)
	// goerror(err)
	// fmt.Println("response body", data)
	// return

	//method 2
	var p post
	err = json.NewDecoder(resp.Body).Decode(&p)
	goerror(err)
	resp.Body.Close()
	log.Println(p) // call:  type Stringer interface {  String() string }
}

func Test_postRequest(t *testing.T) {

	py := payload{
		Name:  "test",
		Email: "test@example",
	}

	body, err := json.Marshal(py)
	goerror(err)

	resp, err := http.Post(
		"https://jsonplaceholder.typicode.com/posts",
		"application/json",
		bytes.NewBuffer(body), // write request body
	)

	resolveResponseBody(resp)

}

func Test_toString(t *testing.T) {

	var i int32 = 10
	var u uint8 = 10
	fmt.Println(goformat.FormatString(nil))
	fmt.Println(goformat.FormatString(12.24))
	fmt.Println(goformat.FormatString(u))
fmt.Println(goformat.FormatString(i))

}

func resolveResponseBody(resp *http.Response) (respBody map[string]interface{}) {
	var err error = nil

	// read the response body and close
	body, err := ioutil.ReadAll(resp.Body)
	goerror(err)
	defer resp.Body.Close()

	switch contentType := resp.Header.Get("Content-Type"); true {

	case strings.HasPrefix(contentType, "application/json"):

		// resolve body
		contentType := resp.Header.Get("Content-Type")
		if !strings.HasPrefix(contentType, "application/json") {
			panic(fmt.Errorf("expect Content-Type json, get %s", resp.Header.Get("Content-Type")))
		}

		err = json.Unmarshal(body, &respBody)
		goerror(err)

		log.Printf("response body = %+v\n", respBody)

	case strings.HasPrefix(contentType, "application/x-www-form-urlencoded"):

		re, err := url.ParseQuery(string(body))
		for k, v := range re {
			respBody[k] = v
		}
		goerror(err)

		// ParseQuery returns : type Values map[string][]string
		log.Printf("response body = %+v\n", respBody)

	default:
		panic(fmt.Errorf("unsupport reponse content-type: %s", contentType))
	}
	return
}

func goerror(err error) {
	if err != nil {
		panic(err)
	}
}

type post struct {
	UserID *int    `json:"user_id,omitempty"`
	ID     *int    `json:"id,omitempty"`
	Title  *string `json:"title,omitempty"`
	Body   *string `json:"body,omitempty"`
	Fake   *string `json:",omitempty"` // json key is Fake
}

func (p post) String() string {
	s, _ := json.MarshalIndent(p, "", "  ")
	return fmt.Sprintf("%T: %s\n", p, s)
}

type payload struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (p payload) String() string {
	s, _ := json.MarshalIndent(p, "", "  ")
	return fmt.Sprintf("%T: %s\n", p, s)
}

func Test_panic(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			log.Println("first layer e = ", e)

			e := recover()
			log.Println("first layer take recover again, e= ", e)

			if e := recover(); e != nil {
				log.Println("second layer e = ", e)
			}
		}

	}()

	panic(fmt.Errorf("this is an test of recover "))

}
