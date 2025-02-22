// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package metadata

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgefield/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Metadata {
	return predicate.Metadata(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Metadata {
	return predicate.Metadata(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Metadata {
	return predicate.Metadata(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Metadata {
	return predicate.Metadata(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Metadata {
	return predicate.Metadata(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Metadata {
	return predicate.Metadata(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Metadata {
	return predicate.Metadata(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Metadata {
	return predicate.Metadata(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Metadata {
	return predicate.Metadata(sql.FieldLTE(FieldID, id))
}

// Age applies equality check predicate on the "age" field. It's identical to AgeEQ.
func Age(v int) predicate.Metadata {
	return predicate.Metadata(sql.FieldEQ(FieldAge, v))
}

// ParentID applies equality check predicate on the "parent_id" field. It's identical to ParentIDEQ.
func ParentID(v int) predicate.Metadata {
	return predicate.Metadata(sql.FieldEQ(FieldParentID, v))
}

// AgeEQ applies the EQ predicate on the "age" field.
func AgeEQ(v int) predicate.Metadata {
	return predicate.Metadata(sql.FieldEQ(FieldAge, v))
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v int) predicate.Metadata {
	return predicate.Metadata(sql.FieldNEQ(FieldAge, v))
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...int) predicate.Metadata {
	return predicate.Metadata(sql.FieldIn(FieldAge, vs...))
}

// AgeNotIn applies the NotIn predicate on the "age" field.
func AgeNotIn(vs ...int) predicate.Metadata {
	return predicate.Metadata(sql.FieldNotIn(FieldAge, vs...))
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v int) predicate.Metadata {
	return predicate.Metadata(sql.FieldGT(FieldAge, v))
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v int) predicate.Metadata {
	return predicate.Metadata(sql.FieldGTE(FieldAge, v))
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v int) predicate.Metadata {
	return predicate.Metadata(sql.FieldLT(FieldAge, v))
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v int) predicate.Metadata {
	return predicate.Metadata(sql.FieldLTE(FieldAge, v))
}

// ParentIDEQ applies the EQ predicate on the "parent_id" field.
func ParentIDEQ(v int) predicate.Metadata {
	return predicate.Metadata(sql.FieldEQ(FieldParentID, v))
}

// ParentIDNEQ applies the NEQ predicate on the "parent_id" field.
func ParentIDNEQ(v int) predicate.Metadata {
	return predicate.Metadata(sql.FieldNEQ(FieldParentID, v))
}

// ParentIDIn applies the In predicate on the "parent_id" field.
func ParentIDIn(vs ...int) predicate.Metadata {
	return predicate.Metadata(sql.FieldIn(FieldParentID, vs...))
}

// ParentIDNotIn applies the NotIn predicate on the "parent_id" field.
func ParentIDNotIn(vs ...int) predicate.Metadata {
	return predicate.Metadata(sql.FieldNotIn(FieldParentID, vs...))
}

// ParentIDIsNil applies the IsNil predicate on the "parent_id" field.
func ParentIDIsNil() predicate.Metadata {
	return predicate.Metadata(sql.FieldIsNull(FieldParentID))
}

// ParentIDNotNil applies the NotNil predicate on the "parent_id" field.
func ParentIDNotNil() predicate.Metadata {
	return predicate.Metadata(sql.FieldNotNull(FieldParentID))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Metadata {
	return predicate.Metadata(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Metadata {
	return predicate.Metadata(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.Metadata {
	return predicate.Metadata(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.Metadata) predicate.Metadata {
	return predicate.Metadata(func(s *sql.Selector) {
		step := newChildrenStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Metadata {
	return predicate.Metadata(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Metadata) predicate.Metadata {
	return predicate.Metadata(func(s *sql.Selector) {
		step := newParentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Metadata) predicate.Metadata {
	return predicate.Metadata(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Metadata) predicate.Metadata {
	return predicate.Metadata(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Metadata) predicate.Metadata {
	return predicate.Metadata(func(s *sql.Selector) {
		p(s.Not())
	})
}
