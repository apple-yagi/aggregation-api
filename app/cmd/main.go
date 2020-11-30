package main

import (
	"aggregation-mod/pkg/external"
	"aggregation-mod/pkg/external/pg"
)

func main() {
	defer pg.CloseConn()

	if err := external.Router.Run(":3030"); err != nil {
		return
	}
}
