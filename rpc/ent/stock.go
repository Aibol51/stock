// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-core/rpc/ent/stock"
)

// Stock is the model entity for the Stock schema.
type Stock struct {
	config `json:"-"`
	// ID of the ent.
	// UUID
	ID uuid.UUID `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Status 1: normal 2: ban | 状态 1 正常 2 禁用
	Status uint8 `json:"status,omitempty"`
	// Stock name
	StockName string `json:"stock_name,omitempty"`
	// Stock code
	StockCode int32 `json:"stock_code,omitempty"`
	// Stock code
	IsRecommend  bool `json:"is_recommend,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Stock) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case stock.FieldIsRecommend:
			values[i] = new(sql.NullBool)
		case stock.FieldStatus, stock.FieldStockCode:
			values[i] = new(sql.NullInt64)
		case stock.FieldStockName:
			values[i] = new(sql.NullString)
		case stock.FieldCreatedAt, stock.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case stock.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Stock fields.
func (s *Stock) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case stock.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case stock.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case stock.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case stock.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				s.Status = uint8(value.Int64)
			}
		case stock.FieldStockName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field stock_name", values[i])
			} else if value.Valid {
				s.StockName = value.String
			}
		case stock.FieldStockCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field stock_code", values[i])
			} else if value.Valid {
				s.StockCode = int32(value.Int64)
			}
		case stock.FieldIsRecommend:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_recommend", values[i])
			} else if value.Valid {
				s.IsRecommend = value.Bool
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Stock.
// This includes values selected through modifiers, order, etc.
func (s *Stock) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// Update returns a builder for updating this Stock.
// Note that you need to call Stock.Unwrap() before calling this method if this Stock
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Stock) Update() *StockUpdateOne {
	return NewStockClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Stock entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Stock) Unwrap() *Stock {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Stock is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Stock) String() string {
	var builder strings.Builder
	builder.WriteString("Stock(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", s.Status))
	builder.WriteString(", ")
	builder.WriteString("stock_name=")
	builder.WriteString(s.StockName)
	builder.WriteString(", ")
	builder.WriteString("stock_code=")
	builder.WriteString(fmt.Sprintf("%v", s.StockCode))
	builder.WriteString(", ")
	builder.WriteString("is_recommend=")
	builder.WriteString(fmt.Sprintf("%v", s.IsRecommend))
	builder.WriteByte(')')
	return builder.String()
}

// Stocks is a parsable slice of Stock.
type Stocks []*Stock