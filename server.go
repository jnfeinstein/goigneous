package main

import (
	"fmt"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"github.com/martini-contrib/render"
	"goigneous/app"
	. "goigneous/app/models"
	"goigneous/config"
	"net/http"
	"strconv"
)

func checkError(err error) bool {
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return true
	}
	return false
}

func main() {
	m := martini.Classic()

	config.Initialize(m)

	m.Use(render.Renderer())
	m.Use(auth.Basic("igneous", "joel"))

	db, err := app.MakeDb()
	if checkError(err) {
		return
	}

	m.Get("/", func(r render.Render, req *http.Request) {
		r.HTML(200, "index", req.Host)
	})

	m.Post("/documents/new", binding.Bind(Document{}), func(w http.ResponseWriter, req *http.Request, doc Document) {
		err := db.Add(&doc)
		if checkError(err) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "http://%s/documents/%d", req.Host, doc.Id)
	})

	m.Get("/documents/:id", func(w http.ResponseWriter, p martini.Params) {
		id, err := strconv.Atoi(p["id"])
		if checkError(err) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		doc, err := db.Get(id)
		if checkError(err) || doc == nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "%s", doc.Content)
	})

	m.Put("/documents/:id", binding.Bind(Document{}), func(w http.ResponseWriter, req *http.Request, p martini.Params, doc Document) {
		id, err := strconv.Atoi(p["id"])
		if checkError(err) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		doc.Id = id // Links the new content to the old db entry
		count, err := db.Update(&doc)
		if count != 1 || err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "http://%s/documents/%d", req.Host, id)
	})

	m.Delete("/documents/:id", func(w http.ResponseWriter, p martini.Params) {
		id, err := strconv.Atoi(p["id"])
		if checkError(err) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = db.Remove(id)
		if checkError(err) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	})

	m.Run()
}
