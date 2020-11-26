package pkg

import (
	"aggregation-mod/internal/apps/aggregation-api/com"
	"fmt"
)

func PrintData(dl com.DataList) {
	for _, d := range dl {
		fmt.Println(d)
	}
}
