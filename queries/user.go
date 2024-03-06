package queries

import (
	"database/sql"
	"log"
  "crypto/rand"
  "encoding/hex"

	"github.com/chamanbravo/upstat/database"
	"github.com/chamanbravo/upstat/models"
	"github.com/chamanbravo/upstat/serializers"

  "golang.org/x/crypto/argon2"
)

type params struct {
    memory      uint32
    iterations  uint32
    parallelism uint8
    saltLength  uint32
    keyLength   uint32
}

// Establish the parameters to use for Argon2.
var p = &params{
    memory:      64 * 1024,
    iterations:  3,
    parallelism: 2,
    saltLength:  16,
    keyLength:   32,
}

func generateFromPassword(password string, p *params) (hash string, err error) {
    // Generate a cryptographically secure random salt.
    //salt, err := generateRandomBytes(p.saltLength)
    salt, err := []byte("NiPlOyV|db=YOw&027s%!l,.+Y*v1c"), nil
    if err != nil {
        return "", err
    }

    // Pass the plaintext password, salt and parameters to the argon2.IDKey
    // function. This will generate a hash of the password using the Argon2id
    // variant.
    hash = hex.EncodeToString(argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength))

    return hash, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
    b := make([]byte, n)
    _, err := rand.Read(b)
    if err != nil {
        return nil, err
    }

    return b, nil
}

func SaveUser(u *serializers.UserSignUp) error {
	stmt, err := database.DB.Prepare("INSERT INTO users(username, email, password) VALUES($1, $2, $3)")
	if err != nil {
		log.Println("Error when trying to prepare statement")
		log.Println(err)
		return err
	}
	defer stmt.Close()

  // Pass the plaintext password and parameters to our generateFromPassword
  // helper function.
  hash, err := generateFromPassword(u.Password, p)
  if err != nil {
    log.Fatal(err)
  }

	_, err = stmt.Exec(u.Username, u.Email, hash)
	if err != nil {
		log.Println("Error when trying to save user")
		log.Println(err)
		return err
	}

	return nil
}

func FindUserByUsernameAndEmail(u *serializers.UserSignUp) (*models.User, error) {
	stmt, err := database.DB.Prepare("SELECT id, username, email FROM users WHERE username = $1 OR email = $2")
	if err != nil {
		log.Println("Error when trying to prepare statement")
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	user := new(models.User)
	result := stmt.QueryRow(u.Username, u.Email).Scan(&user.ID, &user.Username, &user.Email)
	if result != nil {
		if result == sql.ErrNoRows {
			return nil, nil
		}
		log.Println("Error when trying to find user")
		log.Println(err)
		return nil, err
	}
	return user, nil
}

func FindUserByUsernameAndPassword(username, password string) (*models.User, error) {
	stmt, err := database.DB.Prepare("SELECT id, username, email, firstname, lastname FROM users WHERE username = $1 AND password = $2")
	if err != nil {
		log.Println("Error when trying to prepare statement")
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

  hash, err := generateFromPassword(password, p)
  if err != nil {
    log.Fatal(err)
  }

	user := new(models.User)
	result := stmt.QueryRow(username, hash).Scan(&user.ID, &user.Username, &user.Email, &user.Firstname, &user.Lastname)
	if result != nil {
		if result == sql.ErrNoRows {
			return nil, nil
		}
		log.Println("Error when trying to find user")
		log.Println(result)
		return nil, result
	}
	return user, nil
}

func FindUserByUsername(username string) (*models.User, error) {
	stmt, err := database.DB.Prepare("SELECT id, username, email, firstname, lastname FROM users WHERE username = $1")
	if err != nil {
		log.Println("Error when trying to prepare statement")
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	user := new(models.User)
	result := stmt.QueryRow(username).Scan(&user.ID, &user.Username, &user.Email, &user.Firstname, &user.Lastname)
	if result != nil {
		if result == sql.ErrNoRows {
			return nil, nil
		}
		log.Println("Error when trying to find user")
		log.Println(err)
		return nil, err
	}
	return user, nil
}

func UsersCount() (int, error) {
	stmt, err := database.DB.Prepare("SELECT COUNT(*) FROM users")
	if err != nil {
		log.Println("Error when trying to prepare statement")
		log.Println(err)
		return -1, err
	}
	defer stmt.Close()

	var count int
	err = stmt.QueryRow().Scan(&count)
	if err != nil {
		log.Println("Error when trying to retrieve user count")
		log.Println(err)
		return -1, err
	}

	return count, nil
}

func UpdatePassword(username string, u *serializers.UpdatePasswordIn) error {
	stmt, err := database.DB.Prepare("UPDATE users SET password = $2 WHERE username = $1 AND password = $3")
	if err != nil {
		log.Println("Error when trying to prepare statement")
		log.Println(err)
		return err
	}
	defer stmt.Close()

  current_hash, err := generateFromPassword(u.CurrentPassword, p)
  if err != nil {
    log.Fatal(err)
  }

  new_hash, err := generateFromPassword(u.NewPassword, p)
  if err != nil {
    log.Fatal(err)
  }

	_, err = stmt.Exec(username, new_hash, current_hash)
	if err != nil {
		log.Println("Error when trying to update password")
		log.Println(err)
		return err
	}

	return nil
}

func UpdateAccount(username string, u *serializers.UpdateAccountIn) error {
	stmt, err := database.DB.Prepare("UPDATE users SET firstname = $1, lastname = $2 WHERE username = $3")
	if err != nil {
		log.Println("Error when trying to prepare statement")
		log.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Firstname, u.Lastname, username)
	if err != nil {
		log.Println("Error when trying to update password")
		log.Println(err)
		return err
	}

	return nil
}
