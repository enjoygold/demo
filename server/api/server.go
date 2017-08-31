package main
import(
	"encoding/json"
//	"io"
	"fmt"
	"net/http"
	"os"
	//"strconv"
	"strings"
	"github.com/gorilla/mux"
)

var chaincodeName string

type response struct{
        GpCoin string      `json:"gpcoin,omitempty"`
        USD    string       `json:"usd,omitempty"`
}


func writeHead(w http.ResponseWriter) http.ResponseWriter {
        w.Header().Set("Access-Control-Allow-Origin", "*")             //
        w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //
        w.Header().Set("content-type", "application/json")
        return w
}

func topupHandle(w http.ResponseWriter, req *http.Request) {
	writeHead(w)
	req.ParseForm()
	amount, found1 := req.Form["Amount"]
	user, found2   := req.Form["User"]

	if !(found1 && found2) {
		fmt.Fprintf(w, "Wrong arguments")
		return
	}

	err := topup(chaincodeName, amount[0], user[0])
	if err != nil {
		fmt.Fprintf(w, "faild")
	}

	fmt.Fprintf(w, "ok")
}


func investHandle(w http.ResponseWriter, req *http.Request) {
	writeHead(w)
	req.ParseForm()
	amount, found1 := req.Form["Amount"]
	user, found2   := req.Form["User"]

	if !(found1 && found2) {
		fmt.Fprintf(w, "Wrong arguments")
		return
	}

	err := invest(chaincodeName, amount[0], user[0])
	if err != nil {
		fmt.Fprintf(w, "faild")
	}

	fmt.Fprintf(w, "ok")
}


func cashoutHandle(w http.ResponseWriter, req *http.Request) {
	writeHead(w)
	req.ParseForm()
	amount, found1 := req.Form["Amount"]
	user, found2   := req.Form["User"]

	if !(found1 && found2) {
		fmt.Fprintf(w, "Wrong arguments")
		return
	}

	err := cashout(chaincodeName, amount[0], user[0])
	if err != nil {
		fmt.Fprintf(w, "faild")
	}

	fmt.Fprintf(w, "ok")
}

func transferHandle(w http.ResponseWriter, req *http.Request) {
	writeHead(w)
	req.ParseForm()
	amount, found1 := req.Form["Amount"]
	from, found2   := req.Form["From"]
	to, found3     := req.Form["To"]

	if !(found1 && found2 && found3) {
		fmt.Fprintf(w, "Wrong arguments")
	}

	err := transfer(chaincodeName, amount[0], from[0], to[0])
	if err != nil {
		fmt.Fprintf(w, "faild")
	}

	fmt.Fprintf(w, "ok")
}

func queryHandle(w http.ResponseWriter, req *http.Request) {
	var r response

	writeHead(w)
	req.ParseForm()
	user, found := req.Form["User"]
	if !found {
		fmt.Fprintf(w, "wrong arguments")
		return
	}

	res, err := CheckUser(chaincodeName, user[0])
	if err != nil {
		fmt.Fprintf(w, "failed")
	}


	if res == "Null"{
		r = response{
                GpCoin  : "0",
                USD     : "0",
		}
	}else {

		amounts := strings.Split(res, ",")
		fmt.Println(res)
		r = response{
		GpCoin  : amounts[0],
		USD	: amounts[1],
		}
	}

	b, err := json.Marshal(&r)
        if err == nil{
                fmt.Fprint(w, string(b))
        }else{
                fmt.Fprint(w, err)
                return
	}
}

var router = mux.NewRouter()

func main() {
	var err error
	if err = initNVP(); err != nil {
	appLogger.Debugf("Failed initiliazing NVP [%s]", err)
                os.Exit(-1)
	}

	chaincodeName, err = deploy()
	if err != nil {
		appLogger.Debugf("Failed with initiliazing")
		os.Exit(-1)
	}

	//http.HandleFunc("/login", login)
	router.HandleFunc("/topup", topupHandle).Methods("POST")
	router.HandleFunc("/invest", investHandle).Methods("POST")
	router.HandleFunc("/cashout", cashoutHandle).Methods("POST")
	router.HandleFunc("/transfer", transferHandle).Methods("POST")
	router.HandleFunc("/query", queryHandle).Methods("POST")
	http.Handle("/", router)


	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}


