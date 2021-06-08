// Code generated by entc, DO NOT EDIT.

package accesssecurity

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Rule applies equality check predicate on the "rule" field. It's identical to RuleEQ.
func Rule(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRule), v))
	})
}

// IP applies equality check predicate on the "ip" field. It's identical to IPEQ.
func IP(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIP), v))
	})
}

// Source applies equality check predicate on the "source" field. It's identical to SourceEQ.
func Source(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSource), v))
	})
}

// Priority applies equality check predicate on the "priority" field. It's identical to PriorityEQ.
func Priority(v int64) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPriority), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.AccessSecurity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccessSecurity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.AccessSecurity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccessSecurity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.AccessSecurity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccessSecurity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.AccessSecurity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccessSecurity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// RuleEQ applies the EQ predicate on the "rule" field.
func RuleEQ(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRule), v))
	})
}

// RuleNEQ applies the NEQ predicate on the "rule" field.
func RuleNEQ(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRule), v))
	})
}

// RuleIn applies the In predicate on the "rule" field.
func RuleIn(vs ...string) predicate.AccessSecurity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccessSecurity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRule), v...))
	})
}

// RuleNotIn applies the NotIn predicate on the "rule" field.
func RuleNotIn(vs ...string) predicate.AccessSecurity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccessSecurity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRule), v...))
	})
}

// RuleGT applies the GT predicate on the "rule" field.
func RuleGT(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRule), v))
	})
}

// RuleGTE applies the GTE predicate on the "rule" field.
func RuleGTE(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRule), v))
	})
}

// RuleLT applies the LT predicate on the "rule" field.
func RuleLT(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRule), v))
	})
}

// RuleLTE applies the LTE predicate on the "rule" field.
func RuleLTE(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRule), v))
	})
}

// RuleContains applies the Contains predicate on the "rule" field.
func RuleContains(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRule), v))
	})
}

// RuleHasPrefix applies the HasPrefix predicate on the "rule" field.
func RuleHasPrefix(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRule), v))
	})
}

// RuleHasSuffix applies the HasSuffix predicate on the "rule" field.
func RuleHasSuffix(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRule), v))
	})
}

// RuleEqualFold applies the EqualFold predicate on the "rule" field.
func RuleEqualFold(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRule), v))
	})
}

// RuleContainsFold applies the ContainsFold predicate on the "rule" field.
func RuleContainsFold(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRule), v))
	})
}

// IPEQ applies the EQ predicate on the "ip" field.
func IPEQ(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIP), v))
	})
}

// IPNEQ applies the NEQ predicate on the "ip" field.
func IPNEQ(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIP), v))
	})
}

// IPIn applies the In predicate on the "ip" field.
func IPIn(vs ...string) predicate.AccessSecurity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccessSecurity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIP), v...))
	})
}

// IPNotIn applies the NotIn predicate on the "ip" field.
func IPNotIn(vs ...string) predicate.AccessSecurity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccessSecurity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIP), v...))
	})
}

// IPGT applies the GT predicate on the "ip" field.
func IPGT(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIP), v))
	})
}

// IPGTE applies the GTE predicate on the "ip" field.
func IPGTE(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIP), v))
	})
}

// IPLT applies the LT predicate on the "ip" field.
func IPLT(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIP), v))
	})
}

// IPLTE applies the LTE predicate on the "ip" field.
func IPLTE(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIP), v))
	})
}

// IPContains applies the Contains predicate on the "ip" field.
func IPContains(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIP), v))
	})
}

// IPHasPrefix applies the HasPrefix predicate on the "ip" field.
func IPHasPrefix(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIP), v))
	})
}

// IPHasSuffix applies the HasSuffix predicate on the "ip" field.
func IPHasSuffix(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIP), v))
	})
}

// IPEqualFold applies the EqualFold predicate on the "ip" field.
func IPEqualFold(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIP), v))
	})
}

// IPContainsFold applies the ContainsFold predicate on the "ip" field.
func IPContainsFold(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIP), v))
	})
}

// SourceEQ applies the EQ predicate on the "source" field.
func SourceEQ(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSource), v))
	})
}

// SourceNEQ applies the NEQ predicate on the "source" field.
func SourceNEQ(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSource), v))
	})
}

// SourceIn applies the In predicate on the "source" field.
func SourceIn(vs ...string) predicate.AccessSecurity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccessSecurity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSource), v...))
	})
}

// SourceNotIn applies the NotIn predicate on the "source" field.
func SourceNotIn(vs ...string) predicate.AccessSecurity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccessSecurity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSource), v...))
	})
}

// SourceGT applies the GT predicate on the "source" field.
func SourceGT(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSource), v))
	})
}

// SourceGTE applies the GTE predicate on the "source" field.
func SourceGTE(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSource), v))
	})
}

// SourceLT applies the LT predicate on the "source" field.
func SourceLT(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSource), v))
	})
}

// SourceLTE applies the LTE predicate on the "source" field.
func SourceLTE(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSource), v))
	})
}

// SourceContains applies the Contains predicate on the "source" field.
func SourceContains(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSource), v))
	})
}

// SourceHasPrefix applies the HasPrefix predicate on the "source" field.
func SourceHasPrefix(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSource), v))
	})
}

// SourceHasSuffix applies the HasSuffix predicate on the "source" field.
func SourceHasSuffix(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSource), v))
	})
}

// SourceEqualFold applies the EqualFold predicate on the "source" field.
func SourceEqualFold(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSource), v))
	})
}

// SourceContainsFold applies the ContainsFold predicate on the "source" field.
func SourceContainsFold(v string) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSource), v))
	})
}

// PriorityEQ applies the EQ predicate on the "priority" field.
func PriorityEQ(v int64) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPriority), v))
	})
}

// PriorityNEQ applies the NEQ predicate on the "priority" field.
func PriorityNEQ(v int64) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPriority), v))
	})
}

// PriorityIn applies the In predicate on the "priority" field.
func PriorityIn(vs ...int64) predicate.AccessSecurity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccessSecurity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPriority), v...))
	})
}

// PriorityNotIn applies the NotIn predicate on the "priority" field.
func PriorityNotIn(vs ...int64) predicate.AccessSecurity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccessSecurity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPriority), v...))
	})
}

// PriorityGT applies the GT predicate on the "priority" field.
func PriorityGT(v int64) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPriority), v))
	})
}

// PriorityGTE applies the GTE predicate on the "priority" field.
func PriorityGTE(v int64) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPriority), v))
	})
}

// PriorityLT applies the LT predicate on the "priority" field.
func PriorityLT(v int64) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPriority), v))
	})
}

// PriorityLTE applies the LTE predicate on the "priority" field.
func PriorityLTE(v int64) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPriority), v))
	})
}

// HasAssets applies the HasEdge predicate on the "assets" edge.
func HasAssets() predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AssetsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AssetsTable, AssetsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAssetsWith applies the HasEdge predicate on the "assets" edge with a given conditions (other predicates).
func HasAssetsWith(preds ...predicate.Asset) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AssetsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AssetsTable, AssetsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AccessSecurity) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AccessSecurity) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
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
func Not(p predicate.AccessSecurity) predicate.AccessSecurity {
	return predicate.AccessSecurity(func(s *sql.Selector) {
		p(s.Not())
	})
}
