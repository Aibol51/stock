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
	"github.com/suyuan32/simple-admin-core/rpc/ent/account"
	"github.com/suyuan32/simple-admin-core/rpc/ent/comment"
	"github.com/suyuan32/simple-admin-core/rpc/ent/like"
	"github.com/suyuan32/simple-admin-core/rpc/ent/post"
	"github.com/suyuan32/simple-admin-core/rpc/ent/view"
)

// PostCreate is the builder for creating a Post entity.
type PostCreate struct {
	config
	mutation *PostMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (pc *PostCreate) SetCreatedAt(t time.Time) *PostCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PostCreate) SetNillableCreatedAt(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PostCreate) SetUpdatedAt(t time.Time) *PostCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PostCreate) SetNillableUpdatedAt(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetStatus sets the "status" field.
func (pc *PostCreate) SetStatus(u uint8) *PostCreate {
	pc.mutation.SetStatus(u)
	return pc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pc *PostCreate) SetNillableStatus(u *uint8) *PostCreate {
	if u != nil {
		pc.SetStatus(*u)
	}
	return pc
}

// SetDeletedAt sets the "deleted_at" field.
func (pc *PostCreate) SetDeletedAt(t time.Time) *PostCreate {
	pc.mutation.SetDeletedAt(t)
	return pc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pc *PostCreate) SetNillableDeletedAt(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetDeletedAt(*t)
	}
	return pc
}

// SetTitle sets the "title" field.
func (pc *PostCreate) SetTitle(s string) *PostCreate {
	pc.mutation.SetTitle(s)
	return pc
}

// SetContent sets the "content" field.
func (pc *PostCreate) SetContent(s string) *PostCreate {
	pc.mutation.SetContent(s)
	return pc
}

// SetAuthor sets the "author" field.
func (pc *PostCreate) SetAuthor(s string) *PostCreate {
	pc.mutation.SetAuthor(s)
	return pc
}

// SetCategory sets the "category" field.
func (pc *PostCreate) SetCategory(s string) *PostCreate {
	pc.mutation.SetCategory(s)
	return pc
}

// SetTags sets the "tags" field.
func (pc *PostCreate) SetTags(s string) *PostCreate {
	pc.mutation.SetTags(s)
	return pc
}

// SetSummary sets the "summary" field.
func (pc *PostCreate) SetSummary(s string) *PostCreate {
	pc.mutation.SetSummary(s)
	return pc
}

// SetCover sets the "cover" field.
func (pc *PostCreate) SetCover(s string) *PostCreate {
	pc.mutation.SetCover(s)
	return pc
}

// SetID sets the "id" field.
func (pc *PostCreate) SetID(u uuid.UUID) *PostCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PostCreate) SetNillableID(u *uuid.UUID) *PostCreate {
	if u != nil {
		pc.SetID(*u)
	}
	return pc
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (pc *PostCreate) SetAccountID(id uuid.UUID) *PostCreate {
	pc.mutation.SetAccountID(id)
	return pc
}

// SetNillableAccountID sets the "account" edge to the Account entity by ID if the given value is not nil.
func (pc *PostCreate) SetNillableAccountID(id *uuid.UUID) *PostCreate {
	if id != nil {
		pc = pc.SetAccountID(*id)
	}
	return pc
}

// SetAccount sets the "account" edge to the Account entity.
func (pc *PostCreate) SetAccount(a *Account) *PostCreate {
	return pc.SetAccountID(a.ID)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (pc *PostCreate) AddCommentIDs(ids ...uuid.UUID) *PostCreate {
	pc.mutation.AddCommentIDs(ids...)
	return pc
}

// AddComments adds the "comments" edges to the Comment entity.
func (pc *PostCreate) AddComments(c ...*Comment) *PostCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddCommentIDs(ids...)
}

// AddLikeIDs adds the "likes" edge to the Like entity by IDs.
func (pc *PostCreate) AddLikeIDs(ids ...uuid.UUID) *PostCreate {
	pc.mutation.AddLikeIDs(ids...)
	return pc
}

// AddLikes adds the "likes" edges to the Like entity.
func (pc *PostCreate) AddLikes(l ...*Like) *PostCreate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return pc.AddLikeIDs(ids...)
}

// AddViewIDs adds the "views" edge to the View entity by IDs.
func (pc *PostCreate) AddViewIDs(ids ...uuid.UUID) *PostCreate {
	pc.mutation.AddViewIDs(ids...)
	return pc
}

// AddViews adds the "views" edges to the View entity.
func (pc *PostCreate) AddViews(v ...*View) *PostCreate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return pc.AddViewIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (pc *PostCreate) Mutation() *PostMutation {
	return pc.mutation
}

// Save creates the Post in the database.
func (pc *PostCreate) Save(ctx context.Context) (*Post, error) {
	if err := pc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PostCreate) SaveX(ctx context.Context) *Post {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PostCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PostCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PostCreate) defaults() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		if post.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized post.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := post.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		if post.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized post.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := post.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.Status(); !ok {
		v := post.DefaultStatus
		pc.mutation.SetStatus(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		if post.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized post.DefaultID (forgotten import ent/runtime?)")
		}
		v := post.DefaultID()
		pc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pc *PostCreate) check() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Post.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Post.updated_at"`)}
	}
	if _, ok := pc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Post.title"`)}
	}
	if _, ok := pc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Post.content"`)}
	}
	if _, ok := pc.mutation.Author(); !ok {
		return &ValidationError{Name: "author", err: errors.New(`ent: missing required field "Post.author"`)}
	}
	if _, ok := pc.mutation.Category(); !ok {
		return &ValidationError{Name: "category", err: errors.New(`ent: missing required field "Post.category"`)}
	}
	if _, ok := pc.mutation.Tags(); !ok {
		return &ValidationError{Name: "tags", err: errors.New(`ent: missing required field "Post.tags"`)}
	}
	if _, ok := pc.mutation.Summary(); !ok {
		return &ValidationError{Name: "summary", err: errors.New(`ent: missing required field "Post.summary"`)}
	}
	if _, ok := pc.mutation.Cover(); !ok {
		return &ValidationError{Name: "cover", err: errors.New(`ent: missing required field "Post.cover"`)}
	}
	return nil
}

func (pc *PostCreate) sqlSave(ctx context.Context) (*Post, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
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
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PostCreate) createSpec() (*Post, *sqlgraph.CreateSpec) {
	var (
		_node = &Post{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(post.Table, sqlgraph.NewFieldSpec(post.FieldID, field.TypeUUID))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(post.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(post.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.Status(); ok {
		_spec.SetField(post.FieldStatus, field.TypeUint8, value)
		_node.Status = value
	}
	if value, ok := pc.mutation.DeletedAt(); ok {
		_spec.SetField(post.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := pc.mutation.Title(); ok {
		_spec.SetField(post.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := pc.mutation.Content(); ok {
		_spec.SetField(post.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := pc.mutation.Author(); ok {
		_spec.SetField(post.FieldAuthor, field.TypeString, value)
		_node.Author = value
	}
	if value, ok := pc.mutation.Category(); ok {
		_spec.SetField(post.FieldCategory, field.TypeString, value)
		_node.Category = value
	}
	if value, ok := pc.mutation.Tags(); ok {
		_spec.SetField(post.FieldTags, field.TypeString, value)
		_node.Tags = value
	}
	if value, ok := pc.mutation.Summary(); ok {
		_spec.SetField(post.FieldSummary, field.TypeString, value)
		_node.Summary = value
	}
	if value, ok := pc.mutation.Cover(); ok {
		_spec.SetField(post.FieldCover, field.TypeString, value)
		_node.Cover = value
	}
	if nodes := pc.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.AccountTable,
			Columns: []string{post.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.account_posts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: post.CommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.LikesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   post.LikesTable,
			Columns: post.LikesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(like.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ViewsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   post.ViewsTable,
			Columns: post.ViewsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(view.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PostCreateBulk is the builder for creating many Post entities in bulk.
type PostCreateBulk struct {
	config
	err      error
	builders []*PostCreate
}

// Save creates the Post entities in the database.
func (pcb *PostCreateBulk) Save(ctx context.Context) ([]*Post, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Post, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PostMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PostCreateBulk) SaveX(ctx context.Context) []*Post {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PostCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PostCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
