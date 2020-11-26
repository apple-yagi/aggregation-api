package pkg

import (
	"aggregation-mod/internal/apps/aggregation-api/com"
	"encoding/csv"
	"io"
	"log"
	"os"
)

func generateData(label string) (data com.Data) {
	d := com.Data{}
	d.Label = label
	return d
}

func FormatToData(csv_path string) (dl com.DataList) {
	f, err := os.Open(csv_path)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	r := csv.NewReader(f)
	record, err := r.Read()

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	var datalist com.DataList
	record_len := len(record)

	for i := 0; i < record_len; i++ {
		datalist = append(datalist, generateData(record[i]))
	}

	for {
		reco, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		for i := 0; i < record_len; i++ {
			datalist[i].Value = append(datalist[i].Value, reco[i])
		}
	}

	return datalist
}
