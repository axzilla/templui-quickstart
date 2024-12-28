// main.go
package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/a-h/templ"
	"github.com/axzilla/templui-quickstart/assets"
	"github.com/axzilla/templui-quickstart/store"
	"github.com/axzilla/templui-quickstart/ui/pages"
	"github.com/joho/godotenv"
)

func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	users := store.Store.List()
	err := pages.UserListView(users).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	if err := store.Store.Delete(id); err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	InitDotEnv()

	mux := http.NewServeMux()

	SetupAssetsRoutes(mux)
	mux.Handle("GET /", templ.Handler(pages.Landing()))
	mux.HandleFunc("GET /api/users", handleGetUsers)
	mux.HandleFunc("DELETE /api/users/{id}", handleDeleteUser)

	fmt.Println("Server is running on http://localhost:8090")
	http.ListenAndServe(":8090", mux)
}

func InitDotEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func SetupAssetsRoutes(mux *http.ServeMux) {
	var isDevelopment = os.Getenv("GO_ENV") != "production"
	assetHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isDevelopment {
			w.Header().Set("Cache-Control", "no-store")
		}
		var fs http.Handler
		if isDevelopment {
			fs = http.FileServer(http.Dir("./assets"))
		} else {
			fs = http.FileServer(http.FS(assets.Assets))
		}
		fs.ServeHTTP(w, r)
	})
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", assetHandler))
}
