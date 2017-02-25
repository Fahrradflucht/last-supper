package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/Fahrradflucht/last-supper/hex2rgba"
	"github.com/Fahrradflucht/last-supper/image"
	"github.com/Fahrradflucht/last-supper/label"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	port := flag.Int("port", 8080, "Port on which to run the server.")
	fontPath := flag.String("font", "", "Path to a ttf file to be used as image font")
	flag.Parse()

	var fontBytes []byte
	if *fontPath != "" {
		var err error
		fontBytes, err = ioutil.ReadFile(*fontPath)
		if err != nil {
			log.Fatal(err)
		}
	}

	r := mux.NewRouter()

	r.HandleFunc(
		"/{width:[0-9]+}x{height:[0-9]+}/{backcolor}/{textcolor}/{text}.{format}",
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)

			// Save since the regex guards against non-numeral chars
			// TODO: Limit max size
			width, _ := strconv.Atoi(vars["width"])
			height, _ := strconv.Atoi(vars["height"])

			bgcolor, err := hex2rgba.Convert(vars["backcolor"], 255)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			textcolor, err := hex2rgba.Convert(vars["textcolor"], 255)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			img := image.New(width, height, bgcolor, label.ImageLabel{
				Text:      vars["text"],
				Color:     textcolor,
				FontBytes: fontBytes})

			image.Encode(w, img, vars["format"])
		})

	h := handlers.CompressHandler(r)
	http.Handle("/", h)

	log.Printf("Listening on port %d...", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
