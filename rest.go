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

// Util functions for Users
// TODO can I do something like this, and pass some annonymous function?
/*
func FindUser(values []User, u User, f func(u1, u2 User) bool) User {
    for _, v := range values {
        if f(u, v) {
            return v
        }
    }
    return nil
}
*/

func FindUser(values []User, id int) User {
    for _, v := range values {
        if v.Id == id {
            return v
        }
    }
    return User{Id: 0, FirstName: "", LastName: ""}
}

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
	userId, _ := strconv.Atoi(r.URL.Query().Get("id"))
	if userId != 0 {
		getUserEndpoint(w, r, userId)
		return
	}
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

func getUserEndpoint(w http.ResponseWriter, r *http.Request, userId int) {
	if r.Method == "GET" {
		// id, _ := strconv.Atoi(r.URL.Path[1:])
		// u := FindUser(users, id)
		u := FindUser(users, userId)

		userJson, err := json.Marshal(u)
	 
	    if err != nil {
	        panic(err)
	    }
	 
	     w.Write(userJson)
	} else {
		http.Error(w, "Invalid request method.", 405)
	}
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
    setUpUsersEndpoint()
    setUpStaticPages()
    http.ListenAndServe(":8080", nil)
}
