package db

import "log"

type ClipboardValue struct {
	Id    int
	Key   string
	Value string
}

func InsertClipboardValue(key string, value string) {
	insertSQL := `INSERT INTO clipboard (Key, Value) VALUES (?, ?)`
	_, err := database.Exec(insertSQL, key, value)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteClipboardValue(id int) {
	deleteSQL := `DELETE FROM clipboard WHERE Id = ?`
	_, err := database.Exec(deleteSQL, id)
	if err != nil {
		log.Println("Test")
		log.Fatal(err)
	}
}

func GetAllClipboardValues() []*ClipboardValue {
	getSql := `SELECT id, key, value FROM clipboard ORDER BY created_at DESC`
	rows, err := database.Query(getSql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	results := []*ClipboardValue{}
	for rows.Next() {
		var id int
		var key string
		var value string

		if err := rows.Scan(&id, &key, &value); err != nil {
			log.Fatal(err)
		}

		results = append(results, &ClipboardValue{
			Id:    id,
			Key:   key,
			Value: value,
		})
	}

	return results
}
