package entity

type ParseableTag struct {
	Tag string			`json:"tag"`
	Class string		`json:"class"`
	Repeatable bool		`json:"repeatable"`
}
