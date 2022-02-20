package main

import (
	"bytes"
	"encoding/json"
	"image/png"
	"log"
	"net/http"
	"strconv"
)

//func main() {
//	dirs, err := ioutil.ReadDir(".")
//	if err != nil {
//		panic(err)
//	}
//	for _, d := range dirs {
//		dirname := d.Name()
//		files, err := ioutil.ReadDir("./" + dirname)
//		if err != nil {
//			panic(err)
//		}
//
//		ff := make([]string, len(files), len(files))
//		for i, f := range files {
//			ff[i] = f.Name()
//			fmt.Println(f.Name())
//		}
//
//		for i := range ff {
//			tmp := strings.Join(strings.Split(fmt.Sprintf("%04s", strconv.FormatInt(int64(i), 2)), ""), "_")
//			name := dirname + "/" + tmp + ".png"
//			err := os.Rename(dirname+"/"+ff[i], name)
//			if err != nil {
//				panic(err)
//			}
//		}
//
//		for i := len(ff); i < 16; i++ {
//
//			tmp := strings.Join(strings.Split(fmt.Sprintf("%04s", strconv.FormatInt(int64(i), 2)), ""), "_")
//			name2 := dirname + "/" + tmp + ".png"
//
//			tmp = strings.Join(strings.Split(fmt.Sprintf("%04s", strconv.FormatInt(int64(i%len(ff)), 2)), ""), "_")
//			name1 := dirname + "/" + tmp + ".png"
//			// Read all content of src to data
//			data, err := ioutil.ReadFile(name1)
//			if err != nil {
//				panic(err)
//			}
//			// Write data to dst
//			err = ioutil.WriteFile(name2, data, 0644)
//			if err != nil {
//				panic(err)
//			}
//		}
//	}
//}
var m *Maker

func init() {
	var err error
	m, err = NewMaker("./static/data")
	if err != nil {
		panic(err)
	}
}

func writeImage(w http.ResponseWriter, g *Genome) {
	img, err := m.Make(g)
	if err != nil {
		log.Printf("unable to write image %s/n", err.Error())
		return
	}

	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, img); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

func blueHandler(w http.ResponseWriter, r *http.Request) {
	var g Genome

	err := json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeImage(w, &g)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/img", blueHandler)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
