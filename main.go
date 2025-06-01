package main

import (
	"os"
	"html/template"
	"log"
	"net/http"

	"github.com/imrkgofficial/gtts"
)

func main() {
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/generate", generateHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server running at http://localhost:9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	text := r.FormValue("text")
	if text == "" {
		http.Error(w, "Text cannot be empty", http.StatusBadRequest)
		return
	}

	// Create TTS client
	tts := gtts.New()

	// Generate mp3 bytes
	mp3data, err := tts.Speech(text, gtts.LangEnglish)
	if err != nil {
		http.Error(w, "Failed to generate MP3: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Save mp3 to file (optional) or directly stream
	// Let's save to static/output.mp3
	err = saveMP3("static/output.mp3", mp3data)
	if err != nil {
		http.Error(w, "Failed to save MP3: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect user to the mp3 file
	http.Redirect(w, r, "/static/output.mp3", http.StatusSeeOther)
}

func saveMP3(path string, data []byte) error {
	f, err := openFile(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	return err
}

func openFile(path string) (*os.File, error) {
	return os.Create(path)
}
