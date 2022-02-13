package main

import (
  "fmt"
	"log"
	"net/http"
  "strings"

  "github.com/eknkc/amber"
)


func pageHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

    compiler := amber.New()
	
		location := strings.ToLower(r.URL.Path[1:])

    var entry string

		if location == "" || len(location) == 0 {
			entry = INDEX_PAGE
		} else {
			entry = location
		}

		file := fmt.Sprintf("%s/%s.amber", ROOT_DIR, entry)
			
		err := compiler.ParseFile(file)
	
		if err != nil {
				
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
	
		} else {

		  template, err := compiler.Compile()
	
		  if err != nil {
				
			  log.Println(err)
			  w.WriteHeader(http.StatusInternalServerError)
			  return
	
		  }
	
		  template.Execute(w, nil)

    }

	
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

} // pageHandler

