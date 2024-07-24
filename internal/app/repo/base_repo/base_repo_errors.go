package base_repo

import "fmt"

type NoRecordsToFetchError struct {
	Message string
}

func (m *NoRecordsToFetchError) Error() string {
	return m.Message
}

func (m *NoRecordsToFetchError) RaiseError(filters interface{}, tableName string) error {
	m.Message = fmt.Sprintf("no rows to fetch with filters: %+v on table: %+v", filters, tableName)
	return m
}
