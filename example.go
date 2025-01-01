package main

import (
	"fmt"

	"github.com/randomtoy/mihoyosdk/gamepackage"
)

func main() {

	var response gamepackage.Response

	launcher := gamepackage.Launcher{
		Url:        "https://sg-hyp-api.hoyoverse.com/hyp/hyp-connect/api/getGamePackages",
		LauncherID: "VYTpXlbWo8",
		// GameID:     "4ziysqXOQ8",
	}

	err := gamepackage.Get(launcher, &response)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Parsed response: %+v\n", response.Data.GamePackages)
	fmt.Println("Done")
}
