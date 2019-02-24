package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/orville-wright/eplhack/eplstuff"
	"github.com/orville-wright/eplhack/mylogger"
)

/*
func keepLines(s string, n int) string {
	result := strings.Join(strings.Split(s, "\n")[:n], "\n")
	return strings.Replace(result, "\r", "", -1)
}
*/

/*
// move this out to its own package when I'm ready to mess wiht goquery
func processElement(index int, element *goquery.Selection) {
	// See if the href attribute exists on the element
	href, exists := element.Attr("title")
	if exists {
		fmt.Println(href)
	}
}
*/

// Hack #1

func hack1() {
	// XXXX
	log.Print("===================================")
	log.Print("*** #1.0 starting...")
	login_name := "badusername"
	password := "badpassword"
	loginURL := "https://ois-orinda-ca.schoolloop.com/portal/login?etarget=login_form"
	//loginURL := "https://ois-orinda-ca.schoolloop.com/portal/login"
	urlData := url.Values{}
	urlData.Set("login_name", login_name)
	urlData.Set("password", password)
	log.Print("*** #1.1 : URL: ", loginURL)
	log.Printf("*** #1.2 : username / password: %s / %s", login_name, password)
	log.Print("*** #1.3 : POSTform now")
	resp4, err := http.PostForm(loginURL, urlData)
	//resp4, err := http.Post(loginURL, "text/html", urlData)

	mylogger.Info.Printf("*** #1.4 POSTform - using URL: %s", resp4.Request.URL)
	mylogger.Info.Printf("*** #1.4 POSTform - using URLdata: %v", urlData)
	mylogger.Info.Println("*** #1.5 POSTform - Status...", resp4.Status)
	mylogger.Info.Println("*** #1.5 POSTform - req form values...", resp4.Request.Form)

	if err != nil {
		log.Fatal(err)
	} else {
		mylogger.Info.Println("*** #1.6 POSTform ERR Status resp4/auto-req ...", err)
	}

	fmt.Println("*** XXXX HACK #1 Resp status : ", resp4.Status)
	//fmt.Println("Header: ", resp4.Header)
	//fmt.Println("Body: ", resp4.Body)
	fmt.Println("*** XXXX HACK #1 Resp Header values")
	i := 1
	for key, value := range resp4.Header {
		fmt.Println(i, "-", key, ":", value)
		i++
	}

	log.Print("*** XXXX HACK #1 Resp Cookies...")
	for ckey, cookie := range resp4.Cookies() {
		fmt.Println(ckey, ":", "Cookie:", cookie.Name, " ", cookie.Value)
	}

	log.Print("*** XXXX HACK #1 stop...")
	defer resp4.Body.Close()
}

// end hack1

func hack2() {
	log.Printf("\n===================================")
	log.Print("*** HACK #2 starting...")

	log.Print("*** #2.0 init empty GET client/Req...")
	resp0, _ := http.Get("https://ois-orinda-ca.schoolloop.com/")
	log.Printf("*** #2.1 Do GET on URL %s", resp0.Request.URL)
	log.Printf("*** #2.2 Status: %s", resp0.Status)
	//Info.Println("*** XXX HACK #2 GET Headers...", resp0.Header)

	fmt.Println("*** #2.3 Resp Header values")
	i := 1
	for key, value := range resp0.Header {
		fmt.Println(i, "-", key, ":", value)
		i++
	}

	log.Print("*** #2.4 Resp Cookies...")
	for ckey, cookie := range resp0.Cookies() {
		fmt.Println(ckey, ":", "Cookie:", cookie.Name, " ", cookie.Value)
	}

	/*
	   // move this out to its own package, when I'm ready to mess arround with goquery
	   	log.Print("*** HACK #2 >>>GO-QUERY dump on resp.body doc<<<")
	   	document, _ := goquery.NewDocumentFromReader(resp0.Body)
	   	document.Each(processElement)

	   	defer resp0.Body.Close()
	   	log.Print("*** XXXX HACK #2 stop...")
	*/

}

// end hack2

func hack3() {
	log.Printf("\n===================================")
	log.Print("***  HACK 3.0 starting...")
	log.Print("*** #3.1 init empty GET client/Req...")

	client2 := http.Client{}
	//request2, err := http.NewRequest("POST", "https://ois-orinda-ca.schoolloop.com/portal/login", nil)
	request2, err := http.NewRequest("POST", "https://ois-orinda-ca.schoolloop.com/portal/login?etarget=login_form", nil)
	request2.SetBasicAuth("badusername", "badpassowrd")

	resp2, err := client2.Do(request2) //POST
	mylogger.Info.Printf("*** #3.2 do manual POST - using URL: %s", resp2.Request.URL)
	mylogger.Info.Println("*** #3.3 do manual POST - Status...", resp2.Status)
	if err != nil {
		log.Fatal(err)
	} else {
		mylogger.Info.Println("*** #3.4 do manual POST ERR Status resp2/req2 ...", err)
	}
	/* Get Details */
	mylogger.Info.Println("*** #3.5 craft new URL after manual POST for req2 ...")
	request2.URL, err = url.Parse("https://ois-orinda-ca.schoolloop.com/portal/parent_home")
	if err != nil {
		fmt.Printf("*** #3.6 URL Parse #FAIL Error : %s", err)
	} else {
		mylogger.Info.Println("*** #3.7 2nd URL updated to: ", request2.URL)
	}

	mylogger.Info.Println("*** #3.8 set auth creds for 2nd POST for req2 ...") //#bug this now needs to be a GET!!
	request2.SetBasicAuth("badusername", "badpassword")
	resp2, err = client2.Do(request2)
	mylogger.Info.Printf("*** #3.9 2nd GET resp2 from orig POST URL: %s", resp2.Request.URL)
	mylogger.Info.Println("*** #3.10 resp2 2nd GET Status...", resp2.Status)
	if err != nil {
		fmt.Printf("Error : %s", err)
	} else {
		mylogger.Info.Println("*** #3.11 2nd new GET from orig POST - ERR Status resp2/req2 ...", err)
	}
	defer resp2.Body.Close()

	// 2nd half...

	fmt.Println("*** #3.9 Resp2 Header values")
	i := 1
	for key, value := range resp2.Header {
		fmt.Println(i, "-", key, ":", value)
		i++
	}

	log.Print(" #3.10 Resp2 Cookies...")
	for ckey, cookie := range resp2.Cookies() {
		fmt.Println(ckey, ":", "Cookie:", cookie.Name, " ", cookie.Value)
	}
	log.Print("*** #3.11 JSON decode resp2.body...")
	var result map[string]interface{}
	json.NewDecoder(resp2.Body).Decode(&result)
	log.Println(result)

}

// end hack3

func hack4() {
	log.Printf("\n===================================")
	log.Print("*** HACK #4 starting...")
	log.Print("*** #4.0 init empty GET client/Req...")
	client4 := http.Client{}
	request1, err := http.NewRequest("GET", "https://ois-orinda-ca.schoolloop.com/portal/parent_home", nil)
	resp1, err := client4.Do(request1)
	mylogger.Info.Printf("*** #4.1 manual resp1 POST URL: %s", resp1.Request.URL)
	mylogger.Info.Println("*** #4.2 resp1 POST Status...", resp1.Status)
	if err != nil {
		log.Fatal(err)
	} else {
		mylogger.Info.Println("*** #4.3 client.do resp1/req1 ERR Status...", err)
	}
	defer resp1.Body.Close()

	mylogger.Info.Printf("*** #4.4 manual request GET URL: %s", resp1.Request.URL)
	mylogger.Info.Println("*** #4.5 GET Status...", resp1.Status)
	mylogger.Info.Println("*** #4.6 request GET postform data: ", request1.PostForm)

}

func main() {
	mylogger.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	mylogger.Info.Println("*** In main()")
	/*
		options := cookiejar.Options{
			PublicSuffixList: publicsuffix.List,
		}
		jar, err := cookiejar.New(&options)
		if err != nil {
			log.Fatal(err)
		}
	*/
	// client := http.Client{Jar: jar}
	hack1()
	hack2()
	hack3()
	hack4()

	eplstuff.Hack10()
	eplstuff.Hack20()
	eplstuff.Hack30()
	eplstuff.Hack40()

}

/*
	log.Print("Set url.Values array...")
	v := url.Values{}
	v.Set("login_name", "badusername")
	v.Add("password", "badpassword
*/

// resp1, err := client.Do(request1)
//if err != nil {
//	log.Fatal(err)
//}

// log.Print("JSON decode on NewRequest ...")

//var result map[string]interface{}
//json.NewDecoder(resp1.Body).Decode(&result)
//log.Println(result)

//request1.ParseForm()

// request1.PostForm = url.Values{"login_name": {"badusername"}, "password": {"badpassword"}}

// data, err := ioutil.ReadAll(resp1.Body)

// log.Print("Here comes the response data page...")

//log.Println("ParseForm.Form: ", request2.Form)             // Print server side info
//log.Println("ParseForm.URL.Path: ", request2.URL.Path)     // Print server side info
//log.Println("ParseForm.URL.scheme: ", request2.URL.Scheme) // Print server side info
//log.Println("ParseForm.method req1: ", request2.Method)    // Print server side info

// Info.Println("Resp0 - Body...", keepLines(string(body), 3))

//Info.Println("Resp1 manual http.Request GET - Status...", resp1.Status)
//log.Println(resp1.Status) // Print the response Status
//Info.Println("Resp1 manual http.Request GET - postform...", request1.PostForm)
// Info.Println("Resp2 - Headers...", resp2.Header) // Print the response Status
//Info.Println("Resp2 simple GET - Status...", resp2.Status)
// body, _ := ioutil.ReadAll(resp2.Body)
// log.Println(string(body)) // print whole html of user profile data
