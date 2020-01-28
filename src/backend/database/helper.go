package database

import "fmt"

func TableName(t Table) string {
	return t.TableName()
}

func InnerJoin(tableName string) string {
	return fmt.Sprintf("INNER JOIN %s ", tableName)
}

func On(formatter string, args ...interface{}) string {
	return "ON " + fmt.Sprintf(formatter, args...)
}

type c struct {
	col string
}

func Col(table, col string) c {
	return c{col: fmt.Sprintf("%s.%s", table, col)}
}

func (r c) Like(matchStr string) string {
	return r.col + fmt.Sprintf(" LIKE '%s'", matchStr)
}
