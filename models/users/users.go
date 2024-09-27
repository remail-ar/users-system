// remail-app users
package main

type User struct {
	ID           uint64 `db:"id"`
	Username     string `db:"username"`
	Password     string `db:"password"`
	hashFunction string `db:"hash_function"`
	FirstName    string `db:"FirstName"`
	LastName     string `db:"LastName"`
}
