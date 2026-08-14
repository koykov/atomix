package pbvector

type pbtyp struct {
	wire   wire
	num    uint
	opt    bool
	goname string
	pbname string
	ver    ver
}

type wire uint8

type ver uint8

func parseType(x any) pbtyp {
	// todo implement me
	return pbtyp{}
}
