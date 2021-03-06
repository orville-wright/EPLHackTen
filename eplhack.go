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

	"github.com/orville-wright/EPLHackTen/mylogger"
)

// Globaldbug Global var for debug status testing
var Globaldbug bool // default zero value = false

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
	// hack1 basic http.POSTform with credentials passed in body
	mylogger.Info.Printf("\n================ 1.0 ===================")
	mylogger.Info.Println("basic http.PostForm...")
	loginName := u //username
	password := p  //password
	loginURL := "https://ois-orinda-ca.schoolloop.com/portal/login?etarget=login_form"
	//loginURL := "https://ois-orinda-ca.schoolloop.com/portal/login"
	urlData := url.Values{}
	urlData.Set("login_name", loginName)
	urlData.Set("password", password)
	mylogger.Info.Println("#1.1 : URL: ", loginURL)
	mylogger.Info.Printf("#1.2 : username / password: %s / %s", loginName, password)
	if Globaldbug == true {
		mylogger.Info.Println("#1.3 : Exec POSTform now")
	}
	resp4, err := http.PostForm(loginURL, urlData)
	//resp4, err := http.Post(loginURL, "text/html", urlData)
	if Globaldbug == true {
		mylogger.Info.Printf("\n==== #1.3.1 POSTform data ====")
		mylogger.Info.Println("#1.4 POSTform request URL:", resp4.Request.URL)
		mylogger.Info.Println("#1.4.1 POSTform URLdata:", urlData)
		mylogger.Info.Println("#1.5 POSTform status:", resp4.Status)
		mylogger.Info.Println("#1.5.1 POSTform data:", resp4.Request.Form)
	}

	if err != nil {
		log.Fatal(err)
	} else {
		if Globaldbug == true {
			mylogger.Info.Println("#1.6 POSTform ERR Status:", err)
			mylogger.Info.Println("#1.7 Resp status: ", resp4.Status)
		}
	}

	//fmt.Println("Header: ", resp4.Header)
	//fmt.Println("Body: ", resp4.Body)
	if Globaldbug == true {
		mylogger.Info.Printf("\n==== #1.8 resp headers ====")
		i := 1
		// cycle through header data fields
		for key, value := range resp4.Header {
			// split out the individual cookies
			if key == "Set-Cookie" {
				mylogger.Info.Println("#1.9 Resp Cookies...")
				for ckey, cookie := range resp4.Cookies() {
					//for ckey, cookie := range resp4.Request.Cookies() {
					fmt.Printf("C%v - %v - Cookie: %v\n", ckey, cookie.Name, cookie.Value)
				} // for :: cookie splitter
			} // if :: test for cookie section
			fmt.Printf("H%v - %v - %v\n", i, key, value)
			i++
		} // for :: main key,value loop
	} // test :: Globaldbug

	defer resp4.Body.Close()
} // end func decl

func hack2() {
	// no username/password required. Just a GET request going on here...
	mylogger.Info.Printf("\n=============== 2 .0 ====================")
	mylogger.Info.Println("init basic vanilla GET client/Req...")
	resp0, _ := http.Get("https://ois-orinda-ca.schoolloop.com/")
	mylogger.Info.Println("#2.1 Basic GET respponse URL: ", resp0.Request.URL)
	mylogger.Info.Println("#2.2 Basic GET Status:", resp0.Status)
	//Info.Println("*** XXX HACK #2 GET Headers...", resp0.Header)
	if Globaldbug == true {
		mylogger.Info.Println("#2.3 Resp Header values")
		i := 1
		for key, value := range resp0.Header {
			fmt.Println(i, "-", key, ":", value)
			i++
		}

		log.Print("#2.4 Resp Cookies...")
		for ckey, cookie := range resp0.Cookies() {
			fmt.Println(ckey, ":", "Cookie:", cookie.Name, " ", cookie.Value)
		}
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
	mylogger.Info.Printf("\n============== 3.0 =====================")
	mylogger.Info.Print("#3.1 init manual empty POST client/Req...")

	client2 := &http.Client{}
	//request2, err := http.NewRequest("POST", "https://ois-orinda-ca.schoolloop.com/portal/login", nil)
	request2, err := http.NewRequest("POST", "https://ois-orinda-ca.schoolloop.com/portal/login?etarget=login_form", nil)
	fmt.Println("#3.1.0 manual POST orig URL:", request2.URL)
	// healthy setup
	request2.SetBasicAuth(u, p)
	request2.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.119 Safari/537.36")

	fmt.Printf("#3.1.1 : username / password: %s / %s\n", u, p)
	resp2, err := client2.Do(request2) //do the POST
	if Globaldbug == true {
		fmt.Println("#3.2 manual POST response URL:", resp2.Request.URL)
		fmt.Println("#3.3 manual POST status:", resp2.Status)
		fmt.Println("#3.3.1 manual POST orig req headers...")
		i := 0
		for key, value := range resp2.Request.Header {
			fmt.Println(i, "-", key, ":", value)
			i++
		}
		//
	}
	if err != nil {
		log.Fatal(err)
	} else {
		if Globaldbug == true {
			mylogger.Info.Println("#3.4 manual POST ERR status:", err)
		}
	}

	// debug dump data of initial login
	if Globaldbug == true {
		mylogger.Info.Printf("\n==== Data dump ====\n")
		fmt.Println("#3.4.1 Dump 1st req/resp Header values")
		i := 1
		for key, value := range resp2.Header {
			fmt.Println(i, "-", key, ":", value)
			i++
		}

		mylogger.Info.Printf("\n==== Data dump ====")
		fmt.Println("#3.4.2 1st req/resp Cookies...")
		for ckey, cookie := range resp2.Cookies() {
			fmt.Println(ckey, ":", "Cookie:", cookie.Name, " ", cookie.Value)
		}

		mylogger.Info.Printf("\n==== Data dump ====")
		fmt.Printf("\n#3.4.3 JSON decode 1st req/resp.body...")
		var result map[string]interface{}
		json.NewDecoder(resp2.Body).Decode(&result)
		fmt.Println(result)
	}
	//
	// 2nd half of action. assumes credentials post succeeded. So get new URL
	// Get Details
	fmt.Printf("\n#3.5 craft new URL after manual POST ...")
	request2.URL, err = url.Parse("https://ois-orinda-ca.schoolloop.com/portal/parent_home")
	if err != nil {
		mylogger.Info.Print("#3.6 URL Parse #FAIL Error:", err)
	} else {
		if Globaldbug == true {
			fmt.Println("#3.7 2nd URL updated to:", request2.URL)
		}
	}

	fmt.Println("#3.8 set auth creds for 2nd POST for req2 ...") //#bug this now needs to be a GET!!
	request2.SetBasicAuth(u, p)
	fmt.Printf("#3.8.1 : username / password: %s / %s\n", u, p)
	resp3, err := client2.Do(request2) // using orignal req/POST structure as setup/executed earlier
	fmt.Println("#3.9 2nd GET resp using orig POST request struct:", resp3.Request.URL)
	fmt.Println("#3.10 2nd GET Status:", resp3.Status)
	if err != nil {
		mylogger.Info.Printf("Error : %s", err)
	} else {
		if Globaldbug == true {
			fmt.Println("#3.11 2nd GET - ERR Status:", err)
		}
	}
	defer resp2.Body.Close()

	// 2nd half...
	if Globaldbug == true {
		mylogger.Info.Printf("\n==== Data dump ====")
		fmt.Println("#3.12 Dump 2nd req/resp Header values")
		i := 1
		for key, value := range resp3.Header {
			fmt.Println(i, "-", key, ":", value)
			i++
		}
		mylogger.Info.Printf("\n==== Data dump ====")
		fmt.Print("#3.13 2nd req/reso Cookies...")
		for ckey, cookie := range resp3.Cookies() {
			fmt.Printf("%v : cookie: %v name: %v\n", ckey, cookie.Name, cookie.Value)
		}

		mylogger.Info.Printf("\n#3.14 JSON decode 2nd req/res.body...")
		var result map[string]interface{}
		json.NewDecoder(resp3.Body).Decode(&result)
		fmt.Println(result)
	}
}

// end hack3

func hack4() {
	// no username/password used here. Just a GET happening...
	log.Printf("\n================ 4.0 ===================")
	log.Print("Test #4 starting...")
	log.Print("#4.0 init empty manual NewRequest basic GET client/Req...")
	client4 := http.Client{}
	request1, err := http.NewRequest("GET", "https://ois-orinda-ca.schoolloop.com/portal/parent_home", nil)
	resp1, err := client4.Do(request1)
	if err != nil {
		log.Fatal(err)
	} else {
		mylogger.Info.Println("#4.3 client.do resp1/req1 ERR Status...", err)
	}
	defer resp1.Body.Close()

	if Globaldbug == true {
		mylogger.Info.Println("#4.1 GET orignal URL:", request1.URL)
		mylogger.Info.Println("#4.1 GET response URL:", resp1.Request.URL)
		mylogger.Info.Println("#4.2 GET response Status:", resp1.Status)
		mylogger.Info.Println("#4.3 GET postform data:", request1.PostForm)
	} // debug
} // hack4()

func main() {
	myastate := 99
	myValue := &myastate
	//myptr := myValue

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

	Globaldbug = *debugPtr
	//Globaldbug := InitDebug(*debugPtr)

	log.Printf("\n===================================")
	if Globaldbug == true {
		log.Print("cmdline args[]...")
		fmt.Println("Username:", *usernamePtr)
		fmt.Println("Password:", *passwordPtr)
		fmt.Println("Debug status:", Globaldbug)
		fmt.Println("Args raw string:", os.Args[1:])
		fmt.Println("cmdline tail:", flag.Args())
	}

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
			if Globaldbug == true {
				log.Printf("Argv %x looks bad (space or empty string): %s...", x, argvloop)
			}
			*myValue = 2 // bad argv state
		case "no_username":
			if Globaldbug == true {
				log.Printf("Argv %x looks bad (no username): %s...", x, argvloop)
			}
			*myValue = 3 // bad argv state
		case "no_password":
			if Globaldbug == true {
				log.Printf("Argv %x looks bad (no password): %s...", x, argvloop)
			}
			*myValue = 4 // bad argv state
		default: // good-ish
			if Globaldbug == true {
				log.Printf("In DEFAULT loop: %v/%v...", x, *myValue)
			}
			if x == 1 { // only go inside if for-loop has completed
				if *myValue == 99 { // clean argv state
					// exec code here...
					log.Printf("All credentials provided. Executing...")
					hack1(*usernamePtr, *passwordPtr)
					hack2()
					hack3(*usernamePtr, *passwordPtr)
					hack4()
					//eplstuff.Hack10(*usernamePtr, *passwordPtr)
					//eplstuff.Hack20()
					//eplstuff.Hack30(*usernamePtr, *passwordPtr)
					//eplstuff.Hack40()
				} else {
					log.Print("One of the cmdline Args was bad. NOT execing...")
				}
			} else {
				if Globaldbug == true {
					log.Print("In DEFAULT but not finished looping yet...")
				}
			}
		} // main switch
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
