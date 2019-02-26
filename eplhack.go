package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/orville-wright/EPLHackTen/eplstuff"
	"github.com/orville-wright/EPLHackTen/mylogger"
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

func hack1(u string, p string) {
	// XXXX
	log.Print("===================================")
	log.Print("*** #1.0 starting...")
	loginName := u //username
	password := p  //password
	loginURL := "https://ois-orinda-ca.schoolloop.com/portal/login?etarget=login_form"
	//loginURL := "https://ois-orinda-ca.schoolloop.com/portal/login"
	urlData := url.Values{}
	urlData.Set("login_name", loginName)
	urlData.Set("password", password)
	log.Print("*** #1.1 : URL: ", loginURL)
	log.Printf("*** #1.2 : username / password: %s / %s", loginName, password)
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
	// no username/password required. Just a GET request going on here...
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

func hack3(u string, p string) {
	// Required:  username/passowrd
	log.Printf("\n===================================")
	log.Print("***  HACK 3.0 starting...")
	log.Print("*** #3.1 init empty GET client/Req...")

	client2 := http.Client{}
	//request2, err := http.NewRequest("POST", "https://ois-orinda-ca.schoolloop.com/portal/login", nil)
	request2, err := http.NewRequest("POST", "https://ois-orinda-ca.schoolloop.com/portal/login?etarget=login_form", nil)
	request2.SetBasicAuth(u, p)

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
	// no username/password used here. Just a GET happening...
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
	myastate := 99
	myValue := &myastate
	myptr := myValue

	mylogger.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	mylogger.Info.Println("*** In main()")
	log.Printf("\n===================================")
	log.Print("*** #0.0 Debug logging initalized ...")
	// CMD Line args processing
	usernamePtr := flag.String("username", "no_username", "Username to log in as")
	passwordPtr := flag.String("password", "no_password", "Password credentials")
	debugPtr := flag.Bool("debugon", false, "Enable **INFO level debug output")
	//numbPtr := flag.Int("numb", 42, "NUMB an int")
	//var svar string
	//flag.StringVar(&svar, "svar", "bar", "SVAR a string var")
	flag.Parse()

	log.Printf("\n===================================")
	log.Print("*** #0.1 CMD Line args[]...")
	fmt.Println("Username:", *usernamePtr)
	fmt.Println("Password:", *passwordPtr)
	fmt.Println("Debug status:", *debugPtr)
	fmt.Println("Args - raw string passed:", os.Args[1:])
	fmt.Println("tail:", flag.Args())
	// t1 := *usernamePtr
	// t2 := *passwordPtr

	var argvarray [2]string
	argvarray[0] = *usernamePtr
	argvarray[1] = *passwordPtr

	//*myValue = x
	for x, argvloop := range argvarray {
		/* debug var checking...
		log.Printf("Top of for loop...")
		log.Printf("myptr: %v", myptr)
		log.Printf("*myptr: %v", *myptr)
		log.Printf("&myptr: %v", &myptr)
		log.Printf("myValue: %v", myValue)
		log.Printf("&myValue: %v", &myValue)
		log.Printf("*myValue: %v", *myValue)
		log.Printf("myastate: %v", myastate)
		log.Printf("&myastate: %v", &myastate)
		*/
		log.Printf("Looping: %v arg: %s", x, argvarray[x])
		switch argvloop {
		case "", " ":
			log.Printf("Argv %x looks bad (space or empty string): %s...", x, argvloop)
			*myValue = 2 // bad argv state
		case "no_username":
			log.Printf("Argv %x looks bad (no username): %s...", x, argvloop)
			*myValue = 3 // bad argv state
		case "no_password":
			log.Printf("Argv %x looks bad (no password): %s...", x, argvloop)
			*myValue = 4 // bad argv state
		default: // good-ish
			log.Printf("In DEFAULT: %v...", x)
			if x == 1 { // only go inside if for-loop has completed
				if *myValue == 99 { // clean argv state
					log.Printf("In DEFAULT & everything is good...")
					// exec code here...
					log.Printf("state ptr: %v state value: %v - username/password credentials provided. Executing...", *myptr, myastate)
					hack1(*usernamePtr, *passwordPtr)
					hack2()
					hack3(*usernamePtr, *passwordPtr)
					hack4()
					eplstuff.Hack10(*usernamePtr, *passwordPtr)
					eplstuff.Hack20()
					eplstuff.Hack30(*usernamePtr, *passwordPtr)
					eplstuff.Hack40()
				} else {
					log.Print("One of the Args was bad. NOT execing...")
				}
			} else {
				log.Print("In DEFAULT but not finished looping yet...")
			}
		} // switch tests
	} // outer FOR loop of switch tests
} // end main()

//fmt.Println("svar:", svar)

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
