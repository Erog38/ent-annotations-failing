package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

// Fields of the User
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique(),
		field.Time("update_time").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(UnsettableAnnotation{}).
			Immutable(),
		field.Time("created_time").
			Default(time.Now).
			Annotations(UnsettableAnnotation{}).
			Immutable(),
		field.String("updated_by").
			Default("unknown user").
			Annotations(UpdateOnlyAnnotation{}),
		field.String("created_by").
			Immutable().
			NotEmpty(),
	}
}
