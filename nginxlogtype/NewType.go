package nginxlogtype

func NewType(typeString string) Variant {
	return RangesMap[typeString]
}
