package main

import (
	"log"
	"net/http"
	"flag"
	"fmt"
)

type Env struct {
	kw Keywords
}

func main() {
	log.Println("Industry Recommendation service started...")

	// tcp listener
	var port = flag.String("port", "3000", "The server port")
	flag.Parse()

	// Env
	env :=&Env{}
	env.kw = *env.kw.New()

	fmt.Println(env.kw.Data["licensing"])

	log.Println(":"+*port+"/recs started...")
	http.HandleFunc("/recs", env.getIndustriesFromKeyword)
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatal(err)
	}

}


func (env *Env) getIndustriesFromKeyword(w http.ResponseWriter, r *http.Request) {

	fmt.Println(env.kw.Data["licensing"])
	fmt.Println("method call")
	val, err := env.kw.Get("licensing")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)

}