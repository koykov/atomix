package pbvector

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

func parseType(x any) pbtyp {
	// todo implement me
	return pbtyp{}
}
