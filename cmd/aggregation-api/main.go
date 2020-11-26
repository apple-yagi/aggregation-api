package main

import (
	"aggregation-mod/internal/apps/aggregation-api/pkg"
)

func main() {
	dl := pkg.FormatToData("data/TSST_SI11_MD.csv")
	pkg.DataToJson(dl)
}
