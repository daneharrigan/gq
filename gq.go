package gq

import (
	"fmt"
	"strings"
	"database/sql"
	_ "github.com/lib/pq"
)

var connection *sql.DB

type Statement struct {
	Conditions []*Condition
	Columns []string
	TableName string
	RowLimit int
}

type Result struct {
	Error error
	SQL string
	Values []interface{}
	row *sql.Row
}

func Connect(databaseUrl string) error {
	var err error
	connection, err = sql.Open("postgres", databaseUrl)
	if err != nil {
		return err
	}

	return nil
}

func From(name string) *Statement {
	return &Statement{TableName: name}
}

func (statement *Statement) Select(columns ...string) *Statement {
	statement.Columns = append(statement.Columns, columns...)
	return statement
}

func (statement *Statement) Where(conditions ...*Condition) *Statement {
	statement.Conditions = append(statement.Conditions, conditions...)
	return statement
}

func (statement *Statement) Limit(limit int) *Statement {
	statement.RowLimit = limit
	return statement
}

func (statement *Statement) First() *Result {
	statement.Limit(1)
	return statement.prepare().execute()
}

func (statement *Statement) SQL() string {
	return statement.prepare().SQL
}

func (statement *Statement) prepare() *Result {
	result := new(Result)

	tableName := statement.TableName
	columns := strings.Join(statement.Columns, ", ")
	result.SQL = fmt.Sprintf("SELECT %s FROM %s", columns, tableName)

	if len(statement.Conditions) > 0 {
		result.SQL += " WHERE"
		var placeholder int
		for i, condition := range statement.Conditions {
			if i > 0 {
				result.SQL += " AND"
			}

			if condition.Value == nil {
				result.SQL += fmt.Sprintf(" (%s %s NULL)",
					condition.Column, condition.Operator)
			} else {
				result.Values = append(result.Values, condition.Value)
				placeholder++
				result.SQL += fmt.Sprintf(" (%s %s $%d)",
					condition.Column, condition.Operator, placeholder)
			}
		}
	}

	return result
}

func (result *Result) Apply(references ...interface{}) error {
	if result.Error != nil {
		return result.Error
	}

	if err := result.row.Scan(references...); err != nil {
		return err
	}
	return nil
}

func (result *Result) execute() *Result {
	result.row = connection.QueryRow(result.SQL, result.Values...)
	return result
}
