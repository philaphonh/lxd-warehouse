package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/ArtyDev57/lxd-warehouse/api/model"
	"github.com/ArtyDev57/lxd-warehouse/api/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// CreateUser function
func CreateUser(c echo.Context) error {
	username := c.FormValue("username")
	firstname := c.FormValue("firstname")
	lastname := c.FormValue("lastname")
	email := c.FormValue("email")
	password := c.FormValue("password")
	bytepassword := []byte(password)
	hashedpassword, err := bcrypt.GenerateFromPassword(bytepassword, 10)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error Hashing Password")
	}
	role, err := strconv.Atoi(c.FormValue("role"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid Input")
	}
	file, err := c.FormFile("image")
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "File Not Found")
	}

	src, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Error Opening File")
	}
	defer src.Close()

	timestamp := time.Now().Format("2006-01-02T15:04:05")

	filename := "api/upload/" + timestamp + "_" + file.Filename

	dst, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error Creating File")
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error Copying File")
	}

	hostname := c.Request().Host
	if protocol := c.Request().Proto; protocol != "HTTP/2" {
		filename = "http://" + hostname + "/upload/" + timestamp + "_" + file.Filename
	} else {
		filename = "https://" + hostname + "/upload/" + timestamp + "_" + file.Filename
	}

	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Conection")
	}
	defer db.Close()

	stmt := "INSERT INTO tbl_users VALUES(DEFAULT, ?, ?, ?, ?, ?, ?, ?)"

	result, err := db.Exec(stmt, username, firstname, lastname, email, hashedpassword, role, filename)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Execution")
	}

	id, err := result.LastInsertId()
	fmt.Println("Inserted ID: ", id)

	return c.String(http.StatusCreated, "Successful")
}

// UpdateUser function
func UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid Input")
	}
	firstname := c.FormValue("firstname")
	lastname := c.FormValue("lastname")
	email := c.FormValue("email")
	password := c.FormValue("password")
	bytepassword := []byte(password)
	hashedpassword, err := bcrypt.GenerateFromPassword(bytepassword, 10)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error Hashing Password")
	}
	role, err := strconv.Atoi(c.FormValue("role"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid Input")
	}

	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Conection")
	}
	defer db.Close()

	file, err := c.FormFile("image")
	if err != nil {
		fmt.Println("No File Uploaded")
	} else {
		src, err := file.Open()
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "Error Opening File")
		}
		defer src.Close()

		timestamp := time.Now().Format("2006-01-02T15:04:05")

		filename := "api/upload/" + timestamp + "_" + file.Filename

		dst, err := os.Create(filename)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "Error Creating File")
		}
		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "Error Copying File")
		}

		hostname := c.Request().Host
		if protocol := c.Request().Proto; protocol != "HTTP/2" {
			filename = "http://" + hostname + "/upload/" + timestamp + "_" + file.Filename
		} else {
			filename = "https://" + hostname + "/upload/" + timestamp + "_" + file.Filename
		}

		stmt := "UPDATE tbl_users SET user_image_url = ? WHERE id = ?"
		result, err := db.Exec(stmt, filename, id)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "Error DB Execution")
		}

		rows, err := result.RowsAffected()
		fmt.Println("Rows affected: ", rows)

		if rows == 0 {
			fmt.Println("No File Updated")
		}
	}

	if password != "" {
		stmt := `UPDATE tbl_users SET 
	user_firstname = ?, 
	user_lastname = ?, 
	user_email = ?,
	user_password = ?,
	user_role = ?
	WHERE id = ?
	`
		result, err := db.Exec(stmt, firstname, lastname, email, hashedpassword, role, id)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "Error DB Execution")
		}

		rows, err := result.RowsAffected()
		fmt.Println("Rows affected: ", rows)

		if rows == 0 {
			return c.String(http.StatusOK, "Nothing Updated")
		}
	} else {
		stmt := `UPDATE tbl_users SET 
	user_firstname = ?, 
	user_lastname = ?, 
	user_email = ?, 
	user_role = ?
	WHERE id = ?
	`
		result, err := db.Exec(stmt, firstname, lastname, email, role, id)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "Error DB Execution")
		}

		rows, err := result.RowsAffected()
		fmt.Println("Rows affected: ", rows)

		if rows == 0 {
			return c.String(http.StatusOK, "Nothing Updated")
		}
	}

	return c.String(http.StatusOK, "Successful")
}

// GetAllUser function
func GetAllUser(c echo.Context) error {
	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Connection")
	}
	defer db.Close()

	var (
		id        int
		username  string
		firstName string
		lastName  string
		email     string
		role      string
		image     string
	)

	query := `
	SELECT tbl_users.id,
       tbl_users.username,
       tbl_users.user_firstname,
       tbl_users.user_lastname,
       tbl_users.user_email,
       tbl_user_roles.role_name,
       tbl_users.user_image_url FROM tbl_users
           INNER JOIN tbl_user_roles on tbl_users.user_role = tbl_user_roles.id
	`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Execution")
	}
	defer rows.Close()

	result := make([]model.UserClean, 0)

	for rows.Next() {
		if err := rows.Scan(&id, &username, &firstName, &lastName, &email, &role, &image); err != nil {
			fmt.Println(err)
			break
		}
		result = append(result, model.UserClean{
			ID:        id,
			Username:  username,
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			Role:      role,
			Image:     image,
		})
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Rows Error")
	}

	return c.JSON(http.StatusOK, result)
}

// GetUser function
func GetUser(c echo.Context) error {
	user := c.Param("username")
	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Connection")
	}
	defer db.Close()

	var (
		id        int
		username  string
		firstName string
		lastName  string
		email     string
		role      int
		image     string
	)

	query := `
	SELECT tbl_users.id,
       tbl_users.username,
       tbl_users.user_firstname,
       tbl_users.user_lastname,
       tbl_users.user_email,
       tbl_users.user_role,
	   tbl_users.user_image_url FROM tbl_users
	   WHERE username = ?
	`

	row := db.QueryRow(query, user)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Execution")
	}
	if err := row.Scan(&id, &username, &firstName, &lastName, &email, &role, &image); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error Row Scanning")
	}

	result := &model.UserCleanRaw{ID: id, Username: username, FirstName: firstName, LastName: lastName, Email: email, Role: role, Image: image}

	return c.JSON(http.StatusOK, result)
}

// GetAllUserRole function
func GetAllUserRole(c echo.Context) error {
	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Connection")
	}
	defer db.Close()

	var (
		id       int
		roleName string
	)

	query := `SELECT id, role_name FROM tbl_user_roles`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Execution")
	}

	result := make([]model.UserRole, 0)

	for rows.Next() {
		if err := rows.Scan(&id, &roleName); err != nil {
			fmt.Println(err)
			break
		}
		result = append(result, model.UserRole{ID: id, RoleName: roleName})
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Rows Error")
	}

	return c.JSON(http.StatusOK, result)
}

// Login function
func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	bytepassword := []byte(password)

	// Info retrieved from database
	var (
		id    int
		user  string
		pass  []byte
		role  int
		image string
	)

	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Connection")
	}
	defer db.Close()

	query := "SELECT id, username, user_password, user_role, user_image_url FROM tbl_users WHERE username = ?"

	row := db.QueryRow(query, username)
	row.Scan(&id, &user, &pass, &role, &image)
	if user != username {
		return c.String(http.StatusUnauthorized, "Not Found")
	}
	if err := bcrypt.CompareHashAndPassword(pass, bytepassword); err != nil {
		fmt.Println(err)
		return c.String(http.StatusUnauthorized, "Not Authorized")
	}

	claim := &model.AuthClaim{
		user,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte("nothing_is_secret"))
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, "Error Tokenizing Credential")
	}

	return c.JSON(http.StatusOK, &model.Auth{UserID: id, Username: user, Role: role, ImageURL: image, Token: t})
}
