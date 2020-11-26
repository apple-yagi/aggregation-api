package pkg

import (
	"aggregation-mod/internal/apps/aggregation-api/com"
	"encoding/json"
	"fmt"
	"log"
)

func DataToJson(dl com.DataList) {
	for _, d := range dl {
		d_json, err := json.Marshal(d)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		fmt.Println(d_json)
	}
}
