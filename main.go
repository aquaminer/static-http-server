package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

func getEnv(env string, defVal string) string {
	val, exists := os.LookupEnv(env);
	if(exists){
		return val;
	}else{
		return defVal;
	}
}
var dirPath = getEnv("DIR_PATH", "/var/www");
var fallbackFilePath = path.Clean(getEnv("FALLBACK_FILE", dirPath + "/index.html"));
var fallbackFile, _ = ioutil.ReadFile(fallbackFilePath)

func FileServer(fs http.FileSystem) http.Handler {
	fileServer := http.FileServer(fs)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fs.Open(path.Clean(r.URL.Path))
		if os.IsNotExist(err) {
			w.WriteHeader(200)
			w.Write(fallbackFile)

			return
		}
		fileServer.ServeHTTP(w, r)
	})
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", FileServer(http.Dir(dirPath))))
}
