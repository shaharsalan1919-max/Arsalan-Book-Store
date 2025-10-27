package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"
    "strconv"
    "github.com/gorilla/mux"
)

type Book struct {
    ID int `json:"id"`
    Title string `json:"title"`
    Author string `json:"author"`
    Price float64 `json:"price"`
    Available bool `json:"available"`
}

var catalogFile = "catalog.json"

func load() ([]Book, error) {
    f, err := os.ReadFile(catalogFile)
    if err != nil {
        return []Book{}, nil
    }
    var books []Book
    if err := json.Unmarshal(f, &books); err != nil {
        return nil, err
    }
    return books, nil
}

func save(books []Book) error {
    b, _ := json.MarshalIndent(books, "", "  ")
    return os.WriteFile(catalogFile, b, 0644)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
    books, _ := load()
    json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])
    books, _ := load()
    for _, b := range books {
        if b.ID == id {
            json.NewEncoder(w).Encode(b)
            return
        }
    }
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(map[string]string{"error":"not found"})
}

func createBook(w http.ResponseWriter, r *http.Request) {
    var book Book
    json.NewDecoder(r.Body).Decode(&book)
    books, _ := load()
    id := 1
    if len(books) > 0 {
        id = books[len(books)-1].ID + 1
    }
    book.ID = id
    books = append(books, book)
    save(books)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])
    var updated Book
    json.NewDecoder(r.Body).Decode(&updated)
    books, _ := load()
    for i, b := range books {
        if b.ID == id {
            updated.ID = id
            books[i] = updated
            save(books)
            json.NewEncoder(w).Encode(updated)
            return
        }
    }
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(map[string]string{"error":"not found"})
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])
    books, _ := load()
    for i, b := range books {
        if b.ID == id {
            books = append(books[:i], books[i+1:]...)
            save(books)
            json.NewEncoder(w).Encode(map[string]string{"deleted":"ok"})
            return
        }
    }
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(map[string]string{"error":"not found"})
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/books", getBooks).Methods("GET")
    r.HandleFunc("/books/{id}", getBook).Methods("GET")
    r.HandleFunc("/books", createBook).Methods("POST")
    r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
    r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
    log.Println("Catalog Service running on :4000")
    http.ListenAndServe(":4000", r)
}
