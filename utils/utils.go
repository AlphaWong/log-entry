package utils

import "github.com/google/uuid"

func GetUuid() string {
	var u, _ = uuid.NewRandom()
	return u.String()
}
