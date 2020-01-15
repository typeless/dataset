package dataset

import "reflect"

type Series struct {
}

type Parser interface {
	Parse(idx int, s string) error
}

type xSeries struct {
	Name  string
	eltsz int
	n     int
	ty    reflect.Type
}

type stringSeries struct {
	xSeries
	Data []string
}

func (s *stringSeries) Parse(i int, s string) error {

}

type intSeries struct {
	xSeries
	Data []int
}

func NewSeries(name string, xs ...interface{}) {
	switch xs.(type) {
	case string:
		data
	}
	//buf := make([]byte, len(xs) * )
}

type Dataset struct {
}
