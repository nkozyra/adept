package adept

type User struct {
	ID            int
	Username      string
	Password      string
	Email         string
	Salt          string
	Authenticated bool
	HasSession    bool
	SessionID     string
}

type UserSection struct {
	User User
}

func CountUsers() int {
	var count int
	DB.QueryRow("SELECT count(*) FROM users").Scan(&count)
	return count
}

func (u User) Create() {
	DB.Exec("INSERT INTO users SET user_name=?, user_email=?, user_salt=?, user_password=?", u.Username, u.Email, u.Salt, u.Password)
}

func (u *User) Retrieve() {

	if u.ID == 0 {
		DB.QueryRow("SELECT user_id, user_name, user_salt, user_password FROM users WHERE user_name=?", u.Username).Scan(&u.ID, &u.Username, &u.Salt, &u.Password)
	}

	if u.ID != 0 {
		DB.QueryRow("SELECT user_id, user_name, user_salt, user_password FROM users WHERE user_id=?", u.ID).Scan(&u.ID, &u.Username, &u.Salt, &u.Password)
	}

}
