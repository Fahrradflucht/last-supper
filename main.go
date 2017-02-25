package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Fahrradflucht/last-supper/hex2rgba"
	"github.com/Fahrradflucht/last-supper/image"
	"github.com/Fahrradflucht/last-supper/label"
	"github.com/gorilla/mux"
)

func main() {
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
				Text:  vars["text"],
				Color: textcolor})

			image.Encode(w, img, vars["format"])
		})

	http.Handle("/", r)

	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
