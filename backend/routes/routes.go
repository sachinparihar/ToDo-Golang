package routes

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"htmx/model"
)

func sendTodos(w http.ResponseWriter) {

	todos, err := model.GetAllTodos()
	if err != nil {
		fmt.Println("Could not get all todos from db", err)
		return
	}

	tmpl := template.Must(template.ParseFiles("../frontend/index.html"))

	err = tmpl.ExecuteTemplate(w, "Todos", todos)
	if err != nil {
		fmt.Println("Could not execute template", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {

	todos, err := model.GetAllTodos()
	if err != nil {
		fmt.Println("Could not get all todos from db", err)
	}

	tmpl := template.Must(template.ParseFiles("../frontend/index.html"))

	err = tmpl.Execute(w, todos)
	if err != nil {
		fmt.Println("Could not execute template", err)
	}

}

func validateObjectID(hexID string) (primitive.ObjectID, error) {
	hexID = strings.TrimSpace(hexID)
	if len(hexID) != 24 {
		return primitive.NilObjectID, errors.New("Invalid ObjectID length")
	}

	validHex := regexp.MustCompile("^[0-9a-fA-F]{24}$").MatchString
	if !validHex(hexID) {
		return primitive.NilObjectID, errors.New("Invalid ObjectID format")
	}

	objID, err := primitive.ObjectIDFromHex(hexID)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return objID, nil
}

func extractID(idParam string) string {
	re := regexp.MustCompile(`[a-f0-9]{24}`)
	id := re.FindString(idParam)
	fmt.Println("Extracted ID:", id) // Add this line to log the extracted ID
	return id
}

func markTodo(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	fmt.Println("ID Param:", idParam)

	// Extract the ID from the parameter
	extractedID := extractID(idParam)

	// Validate the ObjectID
	objID, err := validateObjectID(extractedID)
	if err != nil {
		fmt.Println("Invalid ID:", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = model.MarkTodoDone(objID, true)
	if err != nil {
		fmt.Println("Could not update todo:", err)
		http.Error(w, "Could not update todo", http.StatusInternalServerError)
		return
	}

	sendTodos(w)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	fmt.Println("ID Param:", idParam)

	// Extract the ID from the parameter
	extractedID := extractID(idParam)

	// Validate the ObjectID
	objID, err := validateObjectID(extractedID)
	if err != nil {
		fmt.Println("Invalid ID:", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = model.DeleteTodoByID(objID)
	if err != nil {
		fmt.Println("Could not delete:", err)
		http.Error(w, "Could not delete", http.StatusInternalServerError)
		return
	}

	sendTodos(w)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error parsing form", err)
	}

	err = model.CreateTodo(r.FormValue("todo"))
	if err != nil {
		fmt.Println("Could not create todo", err)
	}

	sendTodos(w)
}

func SetupAndRun() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/todo/{id}", markTodo).Methods("PUT")
	mux.HandleFunc("/todo/{id}", deleteTodo).Methods("DELETE")
	mux.HandleFunc("/create", createTodo).Methods("POST")

	log.Fatal(http.ListenAndServe(":5000", mux))

}
