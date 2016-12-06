package adept

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/nu7hatch/gouuid"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func createSessionId() {

}

// getSession reads session values and
func getSession(r *http.Request) User {
	u := User{}
	session, _ := store.Get(r, "adept")
	val := session.Values["user_id"]
	fmt.Println(session.Values)
	if val != nil && val != 0 {
		u.ID = val.(int)

		u.Retrieve()
		u.Authenticated = true
	}

	return u
}

func setSession(u User) {
	fmt.Println(u.SessionID)
	DB.Exec("INSERT INTO sessions SET session_id=?, user_id=?, session_start_time=?, session_last_time=? ON DUPLICATE KEY UPDATE session_id=?, session_last_time=?", u.SessionID, u.ID, Now(), Now(), u.SessionID, Now())
}

func deleteSession(u User) {
	DB.Exec("DELETE FROM sessions WHERE session_id=? OR user_id=?", u.SessionID, u.ID)
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func makeHash(salt string, p string) (string, string) {
	if salt == "" {
		salt = RandStringBytes(16)
	}

	key := p + salt
	hash := sha1.New()
	hash.Write([]byte(key))
	hashOut := base64.URLEncoding.EncodeToString(hash.Sum(nil))
	return salt, string(hashOut)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	Templates.ExecuteTemplate(w, "register.html", nil)
}

func registerProcess(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	email := r.FormValue("email")
	pass := r.FormValue("password")

	salt, hash := makeHash("", pass)
	u := User{Username: username, Password: hash, Salt: salt, Email: email}
	u.Create()
	fmt.Println(hash, username, email, pass)
}

func authProcess(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	pass := r.FormValue("password")

	u := User{}
	u.Username = username
	u.Password = pass

	u.Retrieve()
	_, hashed := makeHash(u.Salt, pass)
	if u.Password == hashed {
		session, err := store.Get(r, "adept")
		if err != nil {

		}
		sn, err := uuid.NewV4()

		if err != nil {
			//
		}
		session.Values["session_id"] = sn.String()
		session.Values["user_id"] = u.ID
		u.SessionID = sn.String()
		session.Save(r, w)
		setSession(u)
		http.Redirect(w, r, "/dashboard", 302)
	} else {
		fmt.Println("BASD LOGIN")
	}
	fmt.Println(u)
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	LoadTemplates()
	Templates.ExecuteTemplate(w, "login.html", nil)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	u := getSession(r)
	deleteSession(u)
	session, err := store.Get(r, "adept")
	if err != nil {

	}
	var ud User
	ud.Authenticated = false
	ud.ID = 0
	ud.SessionID = ""
	session.Values["session_id"] = ""
	session.Values["user_id"] = 0

	session.Save(r, w)

	http.Redirect(w, r, "/", 302)
}
