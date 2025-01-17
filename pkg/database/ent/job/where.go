// Code generated by entc, DO NOT EDIT.

package job

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
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
func IDNotIn(ids ...string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
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
func IDGT(id string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Cronjobid applies equality check predicate on the "cronjobid" field. It's identical to CronjobidEQ.
func Cronjobid(v int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCronjobid), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Func applies equality check predicate on the "func" field. It's identical to FuncEQ.
func Func(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFunc), v))
	})
}

// Cron applies equality check predicate on the "cron" field. It's identical to CronEQ.
func Cron(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCron), v))
	})
}

// Mode applies equality check predicate on the "mode" field. It's identical to ModeEQ.
func Mode(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMode), v))
	})
}

// ResourceIds applies equality check predicate on the "resourceIds" field. It's identical to ResourceIdsEQ.
func ResourceIds(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResourceIds), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// Metadata applies equality check predicate on the "metadata" field. It's identical to MetadataEQ.
func Metadata(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMetadata), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// CronjobidEQ applies the EQ predicate on the "cronjobid" field.
func CronjobidEQ(v int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCronjobid), v))
	})
}

// CronjobidNEQ applies the NEQ predicate on the "cronjobid" field.
func CronjobidNEQ(v int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCronjobid), v))
	})
}

// CronjobidIn applies the In predicate on the "cronjobid" field.
func CronjobidIn(vs ...int) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCronjobid), v...))
	})
}

// CronjobidNotIn applies the NotIn predicate on the "cronjobid" field.
func CronjobidNotIn(vs ...int) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCronjobid), v...))
	})
}

// CronjobidGT applies the GT predicate on the "cronjobid" field.
func CronjobidGT(v int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCronjobid), v))
	})
}

// CronjobidGTE applies the GTE predicate on the "cronjobid" field.
func CronjobidGTE(v int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCronjobid), v))
	})
}

// CronjobidLT applies the LT predicate on the "cronjobid" field.
func CronjobidLT(v int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCronjobid), v))
	})
}

// CronjobidLTE applies the LTE predicate on the "cronjobid" field.
func CronjobidLTE(v int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCronjobid), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// FuncEQ applies the EQ predicate on the "func" field.
func FuncEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFunc), v))
	})
}

// FuncNEQ applies the NEQ predicate on the "func" field.
func FuncNEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFunc), v))
	})
}

// FuncIn applies the In predicate on the "func" field.
func FuncIn(vs ...string) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFunc), v...))
	})
}

// FuncNotIn applies the NotIn predicate on the "func" field.
func FuncNotIn(vs ...string) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFunc), v...))
	})
}

// FuncGT applies the GT predicate on the "func" field.
func FuncGT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFunc), v))
	})
}

// FuncGTE applies the GTE predicate on the "func" field.
func FuncGTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFunc), v))
	})
}

// FuncLT applies the LT predicate on the "func" field.
func FuncLT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFunc), v))
	})
}

// FuncLTE applies the LTE predicate on the "func" field.
func FuncLTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFunc), v))
	})
}

// FuncContains applies the Contains predicate on the "func" field.
func FuncContains(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFunc), v))
	})
}

// FuncHasPrefix applies the HasPrefix predicate on the "func" field.
func FuncHasPrefix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFunc), v))
	})
}

// FuncHasSuffix applies the HasSuffix predicate on the "func" field.
func FuncHasSuffix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFunc), v))
	})
}

// FuncEqualFold applies the EqualFold predicate on the "func" field.
func FuncEqualFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFunc), v))
	})
}

// FuncContainsFold applies the ContainsFold predicate on the "func" field.
func FuncContainsFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFunc), v))
	})
}

// CronEQ applies the EQ predicate on the "cron" field.
func CronEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCron), v))
	})
}

// CronNEQ applies the NEQ predicate on the "cron" field.
func CronNEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCron), v))
	})
}

// CronIn applies the In predicate on the "cron" field.
func CronIn(vs ...string) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCron), v...))
	})
}

// CronNotIn applies the NotIn predicate on the "cron" field.
func CronNotIn(vs ...string) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCron), v...))
	})
}

// CronGT applies the GT predicate on the "cron" field.
func CronGT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCron), v))
	})
}

// CronGTE applies the GTE predicate on the "cron" field.
func CronGTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCron), v))
	})
}

// CronLT applies the LT predicate on the "cron" field.
func CronLT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCron), v))
	})
}

// CronLTE applies the LTE predicate on the "cron" field.
func CronLTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCron), v))
	})
}

// CronContains applies the Contains predicate on the "cron" field.
func CronContains(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCron), v))
	})
}

// CronHasPrefix applies the HasPrefix predicate on the "cron" field.
func CronHasPrefix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCron), v))
	})
}

// CronHasSuffix applies the HasSuffix predicate on the "cron" field.
func CronHasSuffix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCron), v))
	})
}

// CronEqualFold applies the EqualFold predicate on the "cron" field.
func CronEqualFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCron), v))
	})
}

// CronContainsFold applies the ContainsFold predicate on the "cron" field.
func CronContainsFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCron), v))
	})
}

// ModeEQ applies the EQ predicate on the "mode" field.
func ModeEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMode), v))
	})
}

// ModeNEQ applies the NEQ predicate on the "mode" field.
func ModeNEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMode), v))
	})
}

// ModeIn applies the In predicate on the "mode" field.
func ModeIn(vs ...string) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMode), v...))
	})
}

// ModeNotIn applies the NotIn predicate on the "mode" field.
func ModeNotIn(vs ...string) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMode), v...))
	})
}

// ModeGT applies the GT predicate on the "mode" field.
func ModeGT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMode), v))
	})
}

// ModeGTE applies the GTE predicate on the "mode" field.
func ModeGTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMode), v))
	})
}

// ModeLT applies the LT predicate on the "mode" field.
func ModeLT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMode), v))
	})
}

// ModeLTE applies the LTE predicate on the "mode" field.
func ModeLTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMode), v))
	})
}

// ModeContains applies the Contains predicate on the "mode" field.
func ModeContains(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMode), v))
	})
}

// ModeHasPrefix applies the HasPrefix predicate on the "mode" field.
func ModeHasPrefix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMode), v))
	})
}

// ModeHasSuffix applies the HasSuffix predicate on the "mode" field.
func ModeHasSuffix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMode), v))
	})
}

// ModeEqualFold applies the EqualFold predicate on the "mode" field.
func ModeEqualFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMode), v))
	})
}

// ModeContainsFold applies the ContainsFold predicate on the "mode" field.
func ModeContainsFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMode), v))
	})
}

// ResourceIdsEQ applies the EQ predicate on the "resourceIds" field.
func ResourceIdsEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResourceIds), v))
	})
}

// ResourceIdsNEQ applies the NEQ predicate on the "resourceIds" field.
func ResourceIdsNEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldResourceIds), v))
	})
}

// ResourceIdsIn applies the In predicate on the "resourceIds" field.
func ResourceIdsIn(vs ...string) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldResourceIds), v...))
	})
}

// ResourceIdsNotIn applies the NotIn predicate on the "resourceIds" field.
func ResourceIdsNotIn(vs ...string) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldResourceIds), v...))
	})
}

// ResourceIdsGT applies the GT predicate on the "resourceIds" field.
func ResourceIdsGT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldResourceIds), v))
	})
}

// ResourceIdsGTE applies the GTE predicate on the "resourceIds" field.
func ResourceIdsGTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldResourceIds), v))
	})
}

// ResourceIdsLT applies the LT predicate on the "resourceIds" field.
func ResourceIdsLT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldResourceIds), v))
	})
}

// ResourceIdsLTE applies the LTE predicate on the "resourceIds" field.
func ResourceIdsLTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldResourceIds), v))
	})
}

// ResourceIdsContains applies the Contains predicate on the "resourceIds" field.
func ResourceIdsContains(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldResourceIds), v))
	})
}

// ResourceIdsHasPrefix applies the HasPrefix predicate on the "resourceIds" field.
func ResourceIdsHasPrefix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldResourceIds), v))
	})
}

// ResourceIdsHasSuffix applies the HasSuffix predicate on the "resourceIds" field.
func ResourceIdsHasSuffix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldResourceIds), v))
	})
}

// ResourceIdsEqualFold applies the EqualFold predicate on the "resourceIds" field.
func ResourceIdsEqualFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldResourceIds), v))
	})
}

// ResourceIdsContainsFold applies the ContainsFold predicate on the "resourceIds" field.
func ResourceIdsContainsFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldResourceIds), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStatus), v))
	})
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStatus), v))
	})
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStatus), v))
	})
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStatus), v))
	})
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStatus), v))
	})
}

// MetadataEQ applies the EQ predicate on the "metadata" field.
func MetadataEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMetadata), v))
	})
}

// MetadataNEQ applies the NEQ predicate on the "metadata" field.
func MetadataNEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMetadata), v))
	})
}

// MetadataIn applies the In predicate on the "metadata" field.
func MetadataIn(vs ...string) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMetadata), v...))
	})
}

// MetadataNotIn applies the NotIn predicate on the "metadata" field.
func MetadataNotIn(vs ...string) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMetadata), v...))
	})
}

// MetadataGT applies the GT predicate on the "metadata" field.
func MetadataGT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMetadata), v))
	})
}

// MetadataGTE applies the GTE predicate on the "metadata" field.
func MetadataGTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMetadata), v))
	})
}

// MetadataLT applies the LT predicate on the "metadata" field.
func MetadataLT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMetadata), v))
	})
}

// MetadataLTE applies the LTE predicate on the "metadata" field.
func MetadataLTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMetadata), v))
	})
}

// MetadataContains applies the Contains predicate on the "metadata" field.
func MetadataContains(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMetadata), v))
	})
}

// MetadataHasPrefix applies the HasPrefix predicate on the "metadata" field.
func MetadataHasPrefix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMetadata), v))
	})
}

// MetadataHasSuffix applies the HasSuffix predicate on the "metadata" field.
func MetadataHasSuffix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMetadata), v))
	})
}

// MetadataEqualFold applies the EqualFold predicate on the "metadata" field.
func MetadataEqualFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMetadata), v))
	})
}

// MetadataContainsFold applies the ContainsFold predicate on the "metadata" field.
func MetadataContainsFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMetadata), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Job) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Job) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
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
func Not(p predicate.Job) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		p(s.Not())
	})
}
