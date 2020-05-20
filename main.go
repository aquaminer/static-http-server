package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)
const DIR_PATH string = "/var/www";

func FileServer(fs http.FileSystem) http.Handler {
	fileServer := http.FileServer(fs)
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("static"))))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fs.Open(path.Clean(r.URL.Path))
		if os.IsNotExist(err) {
			w.WriteHeader(200)
			var file, _ = ioutil.ReadFile(path.Clean(DIR_PATH + "/index.html"))
			w.Write(file);

			return
		}
		fileServer.ServeHTTP(w, r)
	})
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", FileServer(http.Dir(DIR_PATH))))
}
