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
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	if t.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < t.NumField(); i++ {
		tf := t.Field(i)

		tag, ok := tf.Tag.Lookup("protobuf")
		if !ok {
			continue
		}
		_ = tag // todo parse me

		vf := v.Field(i)
		fv := pbfield{
			goname: tf.Name,
		}
		if tf.Type.Kind() == reflect.Struct {
			c := parseType(vf.Interface())
			fv.obj = &c
		}
	}
	return
}
