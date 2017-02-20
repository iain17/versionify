package mux_versionify

import (
	"github.com/Pocketbrain/versionify"
	"fmt"
)

func Registrator(version *versionify.Version, methods versionify.Methods) {
	fmt.Println(version.String())
	fmt.Println(methods)
}