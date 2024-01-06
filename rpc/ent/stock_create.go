// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-core/rpc/ent/stock"
)

// StockCreate is the builder for creating a Stock entity.
type StockCreate struct {
	config
	mutation *StockMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (sc *StockCreate) SetCreatedAt(t time.Time) *StockCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *StockCreate) SetNillableCreatedAt(t *time.Time) *StockCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *StockCreate) SetUpdatedAt(t time.Time) *StockCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *StockCreate) SetNillableUpdatedAt(t *time.Time) *StockCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetStatus sets the "status" field.
func (sc *StockCreate) SetStatus(u uint8) *StockCreate {
	sc.mutation.SetStatus(u)
	return sc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sc *StockCreate) SetNillableStatus(u *uint8) *StockCreate {
	if u != nil {
		sc.SetStatus(*u)
	}
	return sc
}

// SetStockName sets the "stock_name" field.
func (sc *StockCreate) SetStockName(s string) *StockCreate {
	sc.mutation.SetStockName(s)
	return sc
}

// SetStockCode sets the "stock_code" field.
func (sc *StockCreate) SetStockCode(s string) *StockCreate {
	sc.mutation.SetStockCode(s)
	return sc
}

// SetIsRecommend sets the "is_recommend" field.
func (sc *StockCreate) SetIsRecommend(b bool) *StockCreate {
	sc.mutation.SetIsRecommend(b)
	return sc
}

// SetNillableIsRecommend sets the "is_recommend" field if the given value is not nil.
func (sc *StockCreate) SetNillableIsRecommend(b *bool) *StockCreate {
	if b != nil {
		sc.SetIsRecommend(*b)
	}
	return sc
}

// SetStockRise sets the "stock_rise" field.
func (sc *StockCreate) SetStockRise(s string) *StockCreate {
	sc.mutation.SetStockRise(s)
	return sc
}

// SetNillableStockRise sets the "stock_rise" field if the given value is not nil.
func (sc *StockCreate) SetNillableStockRise(s *string) *StockCreate {
	if s != nil {
		sc.SetStockRise(*s)
	}
	return sc
}

// SetStockFall sets the "stock_fall" field.
func (sc *StockCreate) SetStockFall(s string) *StockCreate {
	sc.mutation.SetStockFall(s)
	return sc
}

// SetNillableStockFall sets the "stock_fall" field if the given value is not nil.
func (sc *StockCreate) SetNillableStockFall(s *string) *StockCreate {
	if s != nil {
		sc.SetStockFall(*s)
	}
	return sc
}

// SetAddTime sets the "add_time" field.
func (sc *StockCreate) SetAddTime(s string) *StockCreate {
	sc.mutation.SetAddTime(s)
	return sc
}

// SetNillableAddTime sets the "add_time" field if the given value is not nil.
func (sc *StockCreate) SetNillableAddTime(s *string) *StockCreate {
	if s != nil {
		sc.SetAddTime(*s)
	}
	return sc
}

// SetDetails sets the "details" field.
func (sc *StockCreate) SetDetails(s string) *StockCreate {
	sc.mutation.SetDetails(s)
	return sc
}

// SetNillableDetails sets the "details" field if the given value is not nil.
func (sc *StockCreate) SetNillableDetails(s *string) *StockCreate {
	if s != nil {
		sc.SetDetails(*s)
	}
	return sc
}

// SetStockTags sets the "stock_tags" field.
func (sc *StockCreate) SetStockTags(s string) *StockCreate {
	sc.mutation.SetStockTags(s)
	return sc
}

// SetNillableStockTags sets the "stock_tags" field if the given value is not nil.
func (sc *StockCreate) SetNillableStockTags(s *string) *StockCreate {
	if s != nil {
		sc.SetStockTags(*s)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *StockCreate) SetID(u uuid.UUID) *StockCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *StockCreate) SetNillableID(u *uuid.UUID) *StockCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// Mutation returns the StockMutation object of the builder.
func (sc *StockCreate) Mutation() *StockMutation {
	return sc.mutation
}

// Save creates the Stock in the database.
func (sc *StockCreate) Save(ctx context.Context) (*Stock, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StockCreate) SaveX(ctx context.Context) *Stock {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StockCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StockCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *StockCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := stock.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := stock.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.Status(); !ok {
		v := stock.DefaultStatus
		sc.mutation.SetStatus(v)
	}
	if _, ok := sc.mutation.IsRecommend(); !ok {
		v := stock.DefaultIsRecommend
		sc.mutation.SetIsRecommend(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := stock.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StockCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Stock.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Stock.updated_at"`)}
	}
	if _, ok := sc.mutation.StockName(); !ok {
		return &ValidationError{Name: "stock_name", err: errors.New(`ent: missing required field "Stock.stock_name"`)}
	}
	if _, ok := sc.mutation.StockCode(); !ok {
		return &ValidationError{Name: "stock_code", err: errors.New(`ent: missing required field "Stock.stock_code"`)}
	}
	if _, ok := sc.mutation.IsRecommend(); !ok {
		return &ValidationError{Name: "is_recommend", err: errors.New(`ent: missing required field "Stock.is_recommend"`)}
	}
	return nil
}

func (sc *StockCreate) sqlSave(ctx context.Context) (*Stock, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *StockCreate) createSpec() (*Stock, *sqlgraph.CreateSpec) {
	var (
		_node = &Stock{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(stock.Table, sqlgraph.NewFieldSpec(stock.FieldID, field.TypeUUID))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(stock.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(stock.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.Status(); ok {
		_spec.SetField(stock.FieldStatus, field.TypeUint8, value)
		_node.Status = value
	}
	if value, ok := sc.mutation.StockName(); ok {
		_spec.SetField(stock.FieldStockName, field.TypeString, value)
		_node.StockName = value
	}
	if value, ok := sc.mutation.StockCode(); ok {
		_spec.SetField(stock.FieldStockCode, field.TypeString, value)
		_node.StockCode = value
	}
	if value, ok := sc.mutation.IsRecommend(); ok {
		_spec.SetField(stock.FieldIsRecommend, field.TypeBool, value)
		_node.IsRecommend = value
	}
	if value, ok := sc.mutation.StockRise(); ok {
		_spec.SetField(stock.FieldStockRise, field.TypeString, value)
		_node.StockRise = value
	}
	if value, ok := sc.mutation.StockFall(); ok {
		_spec.SetField(stock.FieldStockFall, field.TypeString, value)
		_node.StockFall = value
	}
	if value, ok := sc.mutation.AddTime(); ok {
		_spec.SetField(stock.FieldAddTime, field.TypeString, value)
		_node.AddTime = value
	}
	if value, ok := sc.mutation.Details(); ok {
		_spec.SetField(stock.FieldDetails, field.TypeString, value)
		_node.Details = value
	}
	if value, ok := sc.mutation.StockTags(); ok {
		_spec.SetField(stock.FieldStockTags, field.TypeString, value)
		_node.StockTags = value
	}
	return _node, _spec
}

// StockCreateBulk is the builder for creating many Stock entities in bulk.
type StockCreateBulk struct {
	config
	err      error
	builders []*StockCreate
}

// Save creates the Stock entities in the database.
func (scb *StockCreateBulk) Save(ctx context.Context) ([]*Stock, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Stock, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StockMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StockCreateBulk) SaveX(ctx context.Context) []*Stock {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StockCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StockCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
