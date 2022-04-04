package schema

type UpdateOnlyAnnotation struct{}

func (UpdateOnlyAnnotation) Name() string {
	return "UpdateOnly"
}
