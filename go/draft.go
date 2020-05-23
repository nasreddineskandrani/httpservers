/*
package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "bytes"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

type VersionData struct {
    Version string
    Test string
}

func readJSONFromUrl(url string) (*VersionData, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var v *VersionData
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respByte := buf.Bytes()
	if err := json.Unmarshal(respByte, &v); err != nil {
		return nil, err
	}

	return v, nil
}

func run(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/assets/version.json" {
        urlCheckVersion := r.URL.Query().Get("urlCheck")
        log.Print(urlCheckVersion)
        enableCors(&w)
        var data *VersionData = &VersionData{}
        data, err := readJSONFromUrl(urlCheckVersion)
        log.Print(data.Test, err)
        fmt.Fprintf(w, "{\"version\": \"" + data.Version + "\"}")
        return
    }


    http.FileServer(http.Dir("./../dist"))
}

func main() {
        http.HandleFunc("/", run)
        http.ListenAndServe(":1222", nil)
}
*/

/* THE GOOD ONE */

package main

import (
	"log"
	"net/http"
	spa "github.com/nasreddineskandrani/spa-server"
)

func run(w http.ResponseWriter, r *http.Request) {
	log.Print("123")
}

func main() {
	http.HandleFunc("/version", run)
	http.Handle("/", spa.SpaHandler("../dist/test", "index.html"))
    // other handlers
    http.ListenAndServe(":3333", nil)
}