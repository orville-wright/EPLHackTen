package eplstuff

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/orville-wright/eplhack/mylogger"
)

// Hack10 does some EPL API tests
func Hack10() {
	// Hack10
	mylogger.Info.Printf("\n===================================")
	mylogger.Info.Print("*** HACK #10.0 starting...")
	username := "badusername"
	password := "badpassword"
	loginURL := "https://ois-orinda-ca.schoolloop.com/portal/login?etarget=login_form"
	urlData := url.Values{}
	urlData.Set("login_name", username)
	urlData.Set("password", badpassword)
	mylogger.Info.Print("*** #10.1 : URL: ", loginURL)
	mylogger.Info.Printf("*** #10.2 : username / password: %s / %s", username, password)
	mylogger.Info.Print("*** #10.3 : POST now")
	resp4, err := http.PostForm(loginURL, urlData)

	if err != nil {
		fmt.Println("Fatal error on http.Post - ERR: ", err)
	}

	mylogger.Info.Println("*** #10.4 Resp status : ", resp4.Status)
	//fmt.Println("Header: ", resp4.Header)
	//fmt.Println("Body: ", resp4.Body)
	mylogger.Info.Println("*** #10.5 Resp Header values")
	i := 1
	for key, value := range resp4.Header {
		fmt.Println(i, "-", key, ":", value)
		i++
	}

	mylogger.Info.Print("*** #10.6 Resp Cookies...")
	for ckey, cookie := range resp4.Cookies() {
		fmt.Println(ckey, ":", "Cookie:", cookie.Name, " ", cookie.Value)
	}

	mylogger.Info.Print("*** #10.7 stop...")
	defer resp4.Body.Close()
}

// end hack1

// Hack20 does some more EPL API tests
func Hack20() {
	mylogger.Info.Printf("\n===================================")
	mylogger.Info.Print("*** HACK #20.0 starting...")

	mylogger.Info.Print("*** #20.1 init empty GET client/Req...")
	resp0, _ := http.Get("https://ois-orinda-ca.schoolloop.com/")
	mylogger.Info.Printf("*** #20.2 Do GET on URL %s", resp0.Request.URL)
	mylogger.Info.Printf("*** #20.3 Status: %s", resp0.Status)
	//Info.Println("*** XXX HACK #2 GET Headers...", resp0.Header)

	fmt.Println("*** #20.4 Resp Header values")
	i := 1
	for key, value := range resp0.Header {
		fmt.Println(i, "-", key, ":", value)
		i++
	}

	mylogger.Info.Print("*** #20.5 Resp Cookies...")
	for ckey, cookie := range resp0.Cookies() {
		fmt.Println(ckey, ":", "Cookie:", cookie.Name, " ", cookie.Value)
	}

	defer resp0.Body.Close()
	mylogger.Info.Print("*** #20.6 stop...")
}

// end hack2

// Hack30 does even more EPL API testing
func Hack30() {
	mylogger.Info.Printf("\n===================================")
	mylogger.Info.Print("*** HACK #30.0 starting...")
	mylogger.Info.Print("*** #30.1 init empty GET client/Req...")

	client2 := http.Client{}
	//request2, err := http.NewRequest("POST", "https://ois-orinda-ca.schoolloop.com/portal/login", nil)
	request2, err := http.NewRequest("POST", "https://ois-orinda-ca.schoolloop.com/portal/login?etarget=login_form", nil)
	request2.SetBasicAuth("badusername", "badpassword")

	resp2, err := client2.Do(request2) //POST
	//Info.Printf("*** #3.2 do manual POST - using URL: %s", resp2.Request.URL)
	mylogger.Info.Printf("*** #30.2 do manual POST - using URL: %s", resp2.Request.URL)
	//Info.Println("*** #3.3 do manual POST - Status...", resp2.Status)
	mylogger.Info.Println("*** #30.3 do manual POST - Status...", resp2.Status)
	if err != nil {
		log.Fatal(err)
	} else {
		mylogger.Info.Println("*** #30.4 do manual POST ERR Status resp2/req2 ...", err)
	}
	/* Get Details */
	mylogger.Info.Println("*** #30.5 craft new URL after manual POST for req2 ...")
	request2.URL, err = url.Parse("https://ois-orinda-ca.schoolloop.com/portal/parent_home")
	if err != nil {
		fmt.Printf("*** #30.6 URL Parse #FAIL Error : %s", err)
	} else {
		mylogger.Info.Println("*** #30.7 2nd URL updated to: ", request2.URL)
	}

	mylogger.Info.Println("*** #30.8 set auth creds for 2nd POST for req2 ...") //#bug this now needs to be a GET!!
	request2.SetBasicAuth("badusername", "badpassword")
	resp2, err = client2.Do(request2)
	mylogger.Info.Printf("*** #30.9 2nd GET resp2 from orig POST URL: %s", resp2.Request.URL)
	mylogger.Info.Println("*** #30.10 resp2 2nd GET Status...", resp2.Status)
	if err != nil {
		fmt.Printf("Error : %s", err)
	} else {
		mylogger.Info.Println("*** #30.11 2nd new GET from orig POST - ERR Status resp2/req2 ...", err)
	}
	defer resp2.Body.Close()

	// 2nd half...

	fmt.Println("*** #30.12 Resp2 Header values")
	i := 1
	for key, value := range resp2.Header {
		fmt.Println(i, "-", key, ":", value)
		i++
	}

	mylogger.Info.Print("*** #3.13 Resp2 Cookies...")
	for ckey, cookie := range resp2.Cookies() {
		fmt.Println(ckey, ":", "Cookie:", cookie.Name, " ", cookie.Value)
	}
	mylogger.Info.Print("*** #30.14 JSON decode resp2.body...")
	var result map[string]interface{}
	json.NewDecoder(resp2.Body).Decode(&result)
	mylogger.Info.Println(result)

}

// end hack3

// Hack40 does even more crazy EPL API testing
func Hack40() {
	mylogger.Info.Printf("\n===================================")
	mylogger.Info.Print("*** HACK #40.0 starting...")
	mylogger.Info.Print("*** #40.1 init empty GET client/Req...")
	client4 := http.Client{}
	request1, err := http.NewRequest("GET", "https://ois-orinda-ca.schoolloop.com/portal/parent_home", nil)
	resp1, err := client4.Do(request1)
	mylogger.Info.Printf("*** #40.2 manual resp1 POST URL: %s", resp1.Request.URL)
	mylogger.Info.Println("*** #40.3 resp1 POST Status...", resp1.Status)
	if err != nil {
		log.Fatal(err)
	} else {
		mylogger.Info.Println("*** #40.4 client.do resp1/req1 ERR Status...", err)
	}
	defer resp1.Body.Close()

	mylogger.Info.Printf("*** #40.5 manual request GET URL: %s", resp1.Request.URL)
	mylogger.Info.Println("*** #40.6 GET Status...", resp1.Status)
	mylogger.Info.Println("*** #40.7 request GET postform data: ", request1.PostForm)

}
