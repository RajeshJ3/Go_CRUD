package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// User is a struct that represents a user in our application
type User struct {
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Post is a struct that represents a single post
type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author User   `json:"author"`
}

// `posts` is a list of all our posts
var posts []Post = []Post{}

// `main` function starts from here
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/post/add/", addPost).Methods("POST")
	router.HandleFunc("/post/all", fetchPosts).Methods("GET")
	router.HandleFunc("/post/{id}", fetchPost).Methods("GET")
	router.HandleFunc("/post/{id}", patchPost).Methods("PATCH")
	router.HandleFunc("/post/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/post/{id}", deletePost).Methods("DELETE")

	http.ListenAndServe(":5000", router) // We want our server to run at port `5000`

}

// Adding a new post
func addPost(w http.ResponseWriter, r *http.Request) {
	// This represents, we are working with JSON data
	w.Header().Set("Content-Type", "application/json")

	// creating a new instance
	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost)

	// Adding new post to the existing list
	posts = append(posts, newPost)

	// This response will be returned
	json.NewEncoder(w).Encode(posts)
}

// Fetching all posts
func fetchPosts(w http.ResponseWriter, r *http.Request) {
	// This represents, we are working with JSON data
	w.Header().Set("Content-Type", "application/json")

	// This response will be returned
	json.NewEncoder(w).Encode(posts)
}

// Fetching single a post
func fetchPost(w http.ResponseWriter, r *http.Request) {
	// This represents, we are working with JSON data
	w.Header().Set("Content-Type", "application/json")

	// Extracting post id from URL parameter
	var idParam string = mux.Vars(r)["id"]
	// Converting string to int
	id, err := strconv.Atoi(idParam)

	// Error checking
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	// Error checking
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}

	// Searching for post in the posts list
	post := posts[id]

	// This response will be returned
	json.NewEncoder(w).Encode(post)
}

// Update post function
func updatePost(w http.ResponseWriter, r *http.Request) {
	// This represents, we are working with JSON data
	w.Header().Set("Content-Type", "application/json")

	// Extracting post id from URL parameter
	var idParam string = mux.Vars(r)["id"]
	// Converting string to int
	id, err := strconv.Atoi(idParam)

	// Error checking
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	// Error checking
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}

	// Creating new post
	var updatedPost Post
	json.NewDecoder(r.Body).Decode(&updatedPost)

	// Replacing the existing one with the new one
	posts[id] = updatedPost

	// This response will be returned
	json.NewEncoder(w).Encode(updatedPost)
}

// Patch post function
func patchPost(w http.ResponseWriter, r *http.Request) {
	// This represents, we are working with JSON data
	w.Header().Set("Content-Type", "application/json")

	// Extracting post id from URL parameter
	var idParam string = mux.Vars(r)["id"]
	// Converting string to int
	id, err := strconv.Atoi(idParam)

	// Error checking
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	// Error checking
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}

	// Updating the fields of existing post
	post := &posts[id]
	json.NewDecoder(r.Body).Decode(&post)

	// This response will be returned
	json.NewEncoder(w).Encode(post)
}

// Delete post function
func deletePost(w http.ResponseWriter, r *http.Request) {
	// This represents, we are working with JSON data
	w.Header().Set("Content-Type", "application/json")

	// Extracting post id from URL parameter
	var idParam string = mux.Vars(r)["id"]
	// Converting string to int
	id, err := strconv.Atoi(idParam)

	// Error checking
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	// Error checking
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}

	// Removing the specified post, by shifting back all by a unit
	posts = append(posts[:id], posts[id+1:]...)

	// This response will be returned
	w.WriteHeader(200)
}
