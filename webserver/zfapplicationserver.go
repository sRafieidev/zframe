package webserver

import (
	"ZGOPROJ/core/ServiceManager"
	"encoding/asn1"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func engineRequestManager(w http.ResponseWriter, r *http.Request) {
	//if (r.RequestURI.)
	//dt := time.Now()
	//color.Yellow(" Request:Time(%s) URI:(%s)", dt.String(), r.RequestURI)
	//color.Yellow(r.RequestURI)
	var addressforprocess = r.RequestURI
	isreq, isexist, content := requestisfile(addressforprocess)
	if isreq {
		if isexist {
			w.Write(content)
		} else {
			errorHandler(w, r, 404)
		}
	} else if !isreq {

		exist, serviceitem := checkExistFunction(addressforprocess)
		if exist {
			ServiceManager.ExecuteService(*serviceitem, w, r)
			//c := &serviceitem.FunctionQuery
			//w.Write([]byte(*c))
			// parse input data ZGJSON
		} else {
			w.Write([]byte("function not exist!\n"))
		}
	}
	//mux.Handle("/", http.FileServer(http.Dir("./sys/http/public")))
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "ZFrame Application Engine Request Not Found ")
	}
}

func requestisfile(requestcontent string) (isreq bool, isexit bool, filedata []byte) {
	if strings.Contains(requestcontent, ".") {
		fileaddress := "./sys/http/public" + requestcontent
		data, err := ioutil.ReadFile(fileaddress)
		if err != nil {
			return true, false, asn1.NullBytes
		} else {
			return true, true, data
		}
	} else {
		return false, false, asn1.NullBytes
	}
}

func checkExistFunction(functionname_in string) (isexit bool, Item *ServiceManager.SysRestFunction) {
	valindexofquestionmark := strings.Index(functionname_in, "?")
	var functionname string = functionname_in

	if valindexofquestionmark > 0 {
		functionname = functionname_in[0:valindexofquestionmark]
	}
	for i := 0; i < len(ServiceManager.ServiceListStruct); i++ {
		if ServiceManager.ServiceListStruct[i].ProcessFunctionName == functionname {
			return true, &ServiceManager.ServiceListStruct[i]
		}
	}
	//nullobject := ServiceManager.Sys_rest_function{}
	return false, nil
}

func StartWebApplicationServerZFGO(inc chan int) {
	color.Green("Start Application Listener")
	var rt int = 0
	//get the value of the ADDR environment variable
	addr := os.Getenv("ADDR")
	//if it's blank, default to ":80", which means
	//listen port 80 for requests addressed to any host
	if len(addr) == 0 {
		addr = ":9090"
	}
	//create a new mux (router)
	//the mux calls different functions for
	//different resource paths
	mux := http.NewServeMux()
	//tell it to call the HelloHandler() function
	//when someone requests the resource path `/hello`
	mux.HandleFunc("/", engineRequestManager)
	//mux.Handle("/", http.FileServer(http.Dir("./sys/http/public")))
	//start the web server using the mux as the root handler,
	//and report any errors that occur.
	//the ListenAndServe() function will block so
	//this program will continue to run until killed
	//ZLog.Printf("server is listening at %s...", addr)
	color.Green("server is listening at %s...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
	inc <- rt
	close(inc)
}
