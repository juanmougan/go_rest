package main
 
import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
)

// Data structures
type Message struct {
    Text string	`json:"text"`
}

type User struct {
    Id        int   `json:"id,omitempty"`
    FirstName string   `json:"first_name,omitempty"`
    LastName  string   `json:"last_name,omitempty"`
}

var users []User

// Endpoints
func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome, %s!", r.URL.Path[1:])
}

func about (w http.ResponseWriter, r *http.Request) {
 
    m := Message{"Sample API written by Juan Manuel Mougan"}
    b, err := json.Marshal(m)
 
    if err != nil {
        panic(err)
    }
 
     w.Write(b)
}

// Fancy (and useful) function from StackOverflow ;)
func getJson(url string, target interface{}) error {
    r, err := http.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}

func getUsersEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		usersJson, err := json.Marshal(users)
	 
	    if err != nil {
	        panic(err)
	    }
	 
	     w.Write(usersJson)
	} else {
		http.Error(w, "Invalid request method.", 405)
	}
}

func GetUserEndpoint(w http.ResponseWriter, r *http.Request) {
	// id := r.URL.Path[1:]
	// TODO filter by id
}

func GetUserEndpointByLastName(w http.ResponseWriter, r *http.Request) {
	lastName := r.URL.Query()["last_name"]
	if lastName != nil {
		// TODO filter list
	}
}

func createUserEndpoint(w http.ResponseWriter, r *http.Request) {
	var u User
	id, _ := strconv.Atoi(r.URL.Path[1:])
	json.NewDecoder(r.Body).Decode(u)
	
	// Create the user
	u.Id = id
	users = append(users, u)
	
	// Return list with new user
	b, err := json.Marshal(users)
 
    if err != nil {
        panic(err)
    }
 
     w.Write(b)
}

func DeleteUserEndpoint() {

}

func updateUserEndpoint() {

}

// Set up endpoints
func setUpStaticPages() {
	http.HandleFunc("/", index)
    http.HandleFunc("/about/", about)
}

func setUpUsersEndpoint() {
	http.HandleFunc("/users", getUsersEndpoint)
//    http.HandleFunc("/users/{id}", GetUserEndpoint).Methods("GET")
//    http.HandleFunc("/users/{id}", createUserEndpoint).Methods("POST")
//    http.HandleFunc("/users/{id}", updateUserEndpoint).Methods("PUT")
//    http.HandleFunc("/users/{id}", DeleteUserEndpoint).Methods("DELETE")
}

func mockData() {
	users = append(users, User{Id: 1, FirstName: "Donald", LastName: "Trump"})
	users = append(users, User{Id: 2, FirstName: "Hillary", LastName: "Clinton"})
}
 
func main() {
    mockData()
    setUpStaticPages()
    setUpUsersEndpoint()
    http.ListenAndServe(":8080", nil)
}
