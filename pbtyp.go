package pbvector

import "reflect"

type pbtyp struct {
	fields []pbfield
}

type pbfield struct {
	wire   wire
	num    uint
	opt    bool
	goname string
	pbname string
	ver    ver

	obj *pbtyp
}

type wire uint8

const (
	wireVarint     wire = 0
	wireFixed32    wire = 5
	wireFixed64    wire = 1
	wireBytes      wire = 2
	wireStartGroup wire = 3
	wireEndGroup   wire = 4
)

type ver uint8

func parseType(x any) (r pbtyp) {
	v := reflect.ValueOf(x)
	if v.Type().Kind() != reflect.Struct {
		return
	}
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		fv := pbfield{
			// todo fill me
		}
		if f.Type().Kind() == reflect.Struct {
			c := parseType(f.Interface())
			fv.obj = &c
		}
	}
	return
}
