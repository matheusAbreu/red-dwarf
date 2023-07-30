package config

import (
	"os"
	"strings"
)

func GetEnvString(variableName string) string{
	return os.Getenv(variableName)
}

func GetEnvBool(variableName string) bool{
	envVar := GetEnvString(variableName)
	return strings.ToLower(envVar) == "true"
}
