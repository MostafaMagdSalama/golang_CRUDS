package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	//create table if not exists

	_, err = db.Exec("create table if not exists users(id serial primary key , name text , email text)")
	if err != nil {
		log.Fatal(err)
		return
	}
	//   handle router
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers(db)).Methods("GET")
	router.HandleFunc("/user/{id}", getuserById(db)).Methods("GET")		
	router.HandleFunc("/user", createUser(db)).Methods("POST")
	router.HandleFunc("/user/{id}",deleteUserBuId(db)).Methods("DELETE")
	// router.HandleFunc("/users/{id}", updateuser(db)).Methods("PUT")

	// start server
	log.Fatal(http.ListenAndServe(":8000", jsonContentTypeMiddleware(router)))
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func createUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, e *http.Request) {
		var user User
		json.NewDecoder(e.Body).Decode(&user)
		// if err != nil {
		// 	log.Fatal(err)
		// 	return
		// }
		fmt.Println(user)
		_, err := db.Exec("insert into  users (name,email) values ($1, $2)", user.Name, user.Email)
		if err != nil {
			log.Fatal(err)
			return

		}
		json.NewEncoder(w).Encode(user)

	}
}

func getUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("select * from users")

		if err != nil {
			log.Fatal("error in query")
			return
		}
		defer rows.Close()
		users := []User{}
		for rows.Next() {
			var u User
			if err := rows.Scan(&u.Name, &u.ID, &u.Email); err != nil {
				log.Fatal(err)
			}
			users = append(users, u)

		}
		json.NewEncoder(w).Encode(users)

	}
}


func getuserById(db *sql.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
    var user User
params:= mux.Vars(r);
id:= params["id"];
		err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
log.Fatal(err)
return
		}


		json.NewEncoder(w).Encode(user)
	}
}


func deleteUserBuId(db *sql.DB )http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
params := mux.Vars(r)
id := params["id"];
_,err:= db.Exec("delete from users where id = $1 ", id);
if err !=nil {
	log.Fatal(err);
	return ;
}
json.NewEncoder(w).Encode("User Deleted")
	}
}

func updateUser(db *sql.DB)http.HandlerFunc{
return func(w http.ResponseWriter , r *http.Request){
params :=mux.Vars(r);
id:=params["id"];
var user User;
if err:= json.NewDecoder(r.Body).Decode(&user);err!=nil{
	log.Fatal(err)
	return ;
}
_,err:= db.Exec("update users set name = $1 , email = $2 where id = $3",user.Name , user.Email , id )
if err !=nil {
	log.Fatal(err)
	return ;
}
json.NewEncoder(w).Encode(user)	
}
}



















