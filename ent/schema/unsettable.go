package schema

type UnsettableAnnotation struct{}

func (UnsettableAnnotation) Name() string {
	return "Unsettable"
}
