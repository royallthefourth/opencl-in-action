package ch2

import (
	"fmt"
	"github.com/jgillich/go-opencl/cl"
	"github.com/urfave/cli"
	"log"
)

// See section 2.2.3 for details.
func GetExtensions(_ *cli.Context) {
	platforms, err := cl.GetPlatforms()
	// TODO show extensions
	for _, plat := range platforms {
		fmt.Println(plat.Name())
	}

	if err != nil {
		log.Fatal(err)
	}
}
