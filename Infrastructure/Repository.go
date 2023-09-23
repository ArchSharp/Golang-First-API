package Infrastructure

import (
	"fmt"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func AddToDB(tableName string, object map[string]interface{}) (map[string]interface{}, error) {
	columnNames := make([]string, 0, len(object))
	parameterPlaceholders := make([]string, 0, len(object))
	parameterValues := make([]interface{}, 0, len(object))

	i := 1
	for key, value := range object {
		columnNames = append(columnNames, key)
		parameterPlaceholders = append(parameterPlaceholders, fmt.Sprintf("$%d", i))
		parameterValues = append(parameterValues, value)
		i++
	}

	insertQuery := "" //fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING *", tableName, pq.QuoteIdentifierArray(columnNames), parameterPlaceholders)

	var insertedData map[string]interface{}
	err := db.QueryRow(insertQuery, parameterValues...).Scan(pq.Map(&insertedData))
	if err != nil {
		return nil, err
	}

	return insertedData, nil
}
