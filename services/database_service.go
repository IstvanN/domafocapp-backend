package services

import (
	"strings"
)

// GetPostgresDBSpecsFromURL splits the given postgres URL and returns with the specs separated
func GetPostgresDBSpecsFromURL(dbURL string) (dbHost, dbUser, dbPort, dbPw, dbName string) {
	userAndPw := strings.Replace(strings.Split(dbURL, "@")[0], "postgres://", "", -1)
	hostAndPortAndName := strings.Split(dbURL, "@")[1]

	dbUser = strings.Split(userAndPw, ":")[0]
	dbPw = strings.Split(userAndPw, ":")[1]
	dbHost = strings.Split(hostAndPortAndName, ":")[0]
	dbPort = strings.Split(strings.Split(hostAndPortAndName, ":")[1], "/")[0]
	dbName = strings.Split(strings.Split(hostAndPortAndName, ":")[1], "/")[1]
	return
}
