package schema

import "entgo.io/ent"

// Reply holds the schema definition for the Reply entity.
type Reply struct {
	ent.Schema
}

// Fields of the Reply.
func (Reply) Fields() []ent.Field {
	return nil
}

// Edges of the Reply.
func (Reply) Edges() []ent.Edge {
	return nil
}
