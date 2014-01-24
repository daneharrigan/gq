package gq

type Condition struct {
	Column string
	Operator string
	Value interface{}
}

func Equal(column string, value interface{}) *Condition {
	return &Condition{Column: column, Value: value, Operator: "="}
}


func NotEqual(column string, value interface{}) *Condition {
	return &Condition{Column: column, Value: value, Operator: "!="}
}

func Like(column string, value interface{}) *Condition {
	return &Condition{Column: column, Value: value, Operator: "LIKE"}
}

func IsNull(column string) *Condition {
	return &Condition{Column: column, Value: nil, Operator: "IS"}
}
