package main

import(
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	Result string
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ascii-pract", asciiArtHandler)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return 
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	tmpl.Execute(w, PageData{Result: ""})
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		http.Error(w, "Bad request: text and banner are required", http.StatusBadRequest)
		return
	}

	bannerData, err := loadBanner(banner + ".txt")
	if err != nil {
		http.Error(w, "Bannerfile is empty", http.StatusInternalServerError)
		return
	}

	result := printArt(text, bannerData)

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	tmpl.Execute(w, PageData{Result: result})
}