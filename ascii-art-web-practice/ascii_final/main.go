package main

import (
    "fmt"
    "net/http"
    "os"
    "strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("file not found")
    }

    if len(data) == 0 {
        return nil, fmt.Errorf("file is empty")
    }

    if len(data) < 855 {
        return nil, fmt.Errorf("error: ")
    }

    lines := strings.Split(string(data), "\n")
    listMap := make(map[rune][]string)
    for i := 0; i < 95; i++ {
        ch := rune(i + 32)
        start := i*9 + 1
        w := lines[start : start+8]
        line := make([]string, 8)
        copy(line, w)
        listMap[ch] = line
    }
    return listMap, nil
}
func formHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprintf(w, `<form method="POST"  action="/ascii">

        <label for="fname">Enter Word1</label>
        <textarea id="fname" name="fname"></textarea><br>
        <br>
        <button>Submit</button><br>

        <br>
        <select id="banner" name="banner"> 
            <option value="standard">Standard</option>
            <option value="shadow">Shadow</option>
            <option value="thinkertoy">Thinkertoy</option>
        </select>

    </form>`)
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")

    if r.Method == "POST" {
        b := r.FormValue("banner")
        banner, err := LoadBanner(b + ".txt")
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "Something went wrong...")
            return
        }
        text := r.FormValue("fname")
        if text == "" {
            w.WriteHeader(http.StatusBadRequest)
            fmt.Fprintf(w, "Bad Request")
            return 
        }
        lines := strings.Split(text, "\n")
        fmt.Fprintf(w, "<pre>")
        for _, line := range lines {
            line = strings.TrimRight(line, "\r")
            if line == "" {
                continue
            }
            for i := 0; i < 8; i++ {
                for _, ch := range line {
                    fmt.Fprintf(w, banner[ch][i])
                }
                fmt.Fprintf(w, "\n")
            }
        }
        fmt.Fprintf(w, "</pre>")
        fmt.Fprintf(w, "<a href=`/`>Back</a>")

    } else {
        w.WriteHeader(http.StatusMethodNotAllowed)
        fmt.Fprintf(w, "Something went wrong...")
    }
}
func main() {
    http.HandleFunc("/", formHandler)
    http.HandleFunc("/ascii", asciiHandler)
    http.ListenAndServe(":3000", nil)

}



