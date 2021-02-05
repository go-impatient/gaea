package schema

import "github.com/facebook/ent"

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return nil
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return nil
}
