package svc_db

import "github.com/walleframe/walle/util"

const (
	Columns   = "__COLUMNS__"
	TableName = "__TABLE__"
)

type StringCondition[T any] struct {
	w    *T
	buf  *util.Builder
	name string
}

func NewStringCondition[T any](w *T, buf *util.Builder, name string) *StringCondition[T] {
	return &StringCondition[T]{
		w:    w,
		buf:  buf,
		name: "id",
	}
}

func (c *StringCondition[T]) Equal(v string) *T {
	//if c.buf.Len() > 0 {
	//	c.buf.Write([]byte(" and `"))
	//} else {
	//	c.buf.Write([]byte(" where `"))
	//}
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` = '"))
	c.buf.WriteString(v)
	c.buf.WriteByte('\'')
	return c.w
}

func (c *StringCondition[T]) NotEqual(v string) *T {
	//if c.buf.Len() > 0 {
	//	c.buf.Write([]byte(" and `"))
	//} else {
	//	c.buf.Write([]byte(" where `"))
	//}
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` != '"))
	c.buf.WriteString(v)
	c.buf.WriteByte('\'')
	return c.w
}

func (c *StringCondition[T]) Like(v string) *T {
	//if c.buf.Len() > 0 {
	//	c.buf.Write([]byte(" and `"))
	//} else {
	//	c.buf.Write([]byte(" where `"))
	//}
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` like '"))
	c.buf.WriteString(v)
	c.buf.WriteByte('\'')
	return c.w
}

func (c *StringCondition[T]) In(vals ...string) *T {
	//if c.buf.Len() > 0 {
	//	c.buf.Write([]byte(" and `"))
	//} else {
	//	c.buf.Write([]byte(" where `"))
	//}
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` in ("))
	for k := 0; k < len(vals); k++ {
		if k > 0 {
			c.buf.Write([]byte(",'"))
		} else {
			c.buf.WriteByte('\'')
		}
		c.buf.WriteString(vals[k])
		c.buf.WriteByte('\'')
	}
	c.buf.Write([]byte(")"))
	return c.w
}

type IntSignedCondition[T any, V int | int8 | int16 | int32 | int64] struct {
	w    *T
	buf  *util.Builder
	name string
}

func NewIntSignedCondition[T any, V int | int8 | int16 | int32 | int64](w *T, buf *util.Builder, name string) *IntSignedCondition[T, V] {
	return &IntSignedCondition[T, V]{
		w:    w,
		buf:  buf,
		name: "id",
	}
}

func (c *IntSignedCondition[T, V]) LessThen(v V) *T {
	//if c.buf.Len() > 0 {
	//	c.buf.Write([]byte(" and `"))
	//} else {
	//	c.buf.Write([]byte(" where `"))
	//}
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` < "))
	c.buf.WriteInt64(int64(v))
	return c.w
}

func (c *IntSignedCondition[T, V]) LessEqual(v V) *T {
	//if c.buf.Len() > 0 {
	//	c.buf.Write([]byte(" and `"))
	//} else {
	//	c.buf.Write([]byte(" where `"))
	//}
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` <= "))
	c.buf.WriteInt64(int64(v))
	return c.w
}

func (c *IntSignedCondition[T, V]) GreaterThen(v V) *T {
	//if c.buf.Len() > 0 {
	//	c.buf.Write([]byte(" and `"))
	//} else {
	//	c.buf.Write([]byte(" where `"))
	//}
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` > "))
	c.buf.WriteInt64(int64(v))
	return c.w
}

func (c *IntSignedCondition[T, V]) GreaterEqual(v V) *T {
	//if c.buf.Len() > 0 {
	//	c.buf.Write([]byte(" and `"))
	//} else {
	//	c.buf.Write([]byte(" where `"))
	//}
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` >= "))
	c.buf.WriteInt64(int64(v))
	return c.w
}

func (c *IntSignedCondition[T, V]) Equal(v V) *T {
	//if c.buf.Len() > 0 {
	//	c.buf.Write([]byte(" and `"))
	//} else {
	//	c.buf.Write([]byte(" where `"))
	//}
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` = "))
	c.buf.WriteInt64(int64(v))
	return c.w
}

func (c *IntSignedCondition[T, V]) NotEqual(v V) *T {
	//if c.buf.Len() > 0 {
	//	c.buf.Write([]byte(" and `"))
	//} else {
	//	c.buf.Write([]byte(" where `"))
	//}
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` != "))
	c.buf.WriteInt64(int64(v))
	return c.w
}

func (c *IntSignedCondition[T, V]) Between(min, max V) *T {
	if min > max {
		min, max = max, min
	}
	//if c.buf.Len() > 0 {
	//	c.buf.Write([]byte(" and `"))
	//} else {
	//	c.buf.Write([]byte(" where `"))
	//}
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` between "))
	c.buf.WriteInt64(int64(min))
	c.buf.Write([]byte(" and "))
	c.buf.WriteInt64(int64(max))
	return c.w
}

func (c *IntSignedCondition[T, V]) In(vals ...V) *T {
	//if c.buf.Len() > 0 {
	//	c.buf.Write([]byte(" and `"))
	//} else {
	//	c.buf.Write([]byte(" where `"))
	//}
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` in ("))
	for k := 0; k < len(vals); k++ {
		if k > 0 {
			c.buf.WriteByte(',')
		}
		c.buf.WriteInt64(int64(vals[k]))
	}
	c.buf.Write([]byte(")"))
	return c.w
}

type IntUnSignedCondition[T any, V uint | uint8 | uint16 | uint32 | uint64] struct {
	w    *T
	buf  *util.Builder
	name string
}

func NewIntUnSignedCondition[T any, V uint | uint8 | uint16 | uint32 | uint64](w *T, buf *util.Builder, name string) *IntUnSignedCondition[T, V] {
	return &IntUnSignedCondition[T, V]{
		w:    w,
		buf:  buf,
		name: "id",
	}
}

func (c *IntUnSignedCondition[T, V]) LessThen(v V) *T {
	//if c.buf.Len() > 0 {
	//	c.buf.Write([]byte(" and `"))
	//} else {
	//	c.buf.Write([]byte(" where `"))
	//}
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` < "))
	c.buf.WriteUint64(uint64(v))
	return c.w
}

func (c *IntUnSignedCondition[T, V]) LessEqual(v V) *T {
	//if c.buf.Len() > 0 {
	//	c.buf.Write([]byte(" and `"))
	//} else {
	//	c.buf.Write([]byte(" where `"))
	//}
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` <= "))
	c.buf.WriteUint64(uint64(v))
	return c.w
}

func (c *IntUnSignedCondition[T, V]) GreaterThen(v V) *T {
	//if c.buf.Len() > 0 {
	//	c.buf.Write([]byte(" and `"))
	//} else {
	//	c.buf.Write([]byte(" where `"))
	//}
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` > "))
	c.buf.WriteUint64(uint64(v))
	return c.w
}

func (c *IntUnSignedCondition[T, V]) GreaterEqual(v V) *T {
	//if c.buf.Len() > 0 {
	//	c.buf.Write([]byte(" and `"))
	//} else {
	//	c.buf.Write([]byte(" where `"))
	//}
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` >= "))
	c.buf.WriteUint64(uint64(v))
	return c.w
}

func (c *IntUnSignedCondition[T, V]) Equal(v V) *T {
	//if c.buf.Len() > 0 {
	//	c.buf.Write([]byte(" and `"))
	//} else {
	//	c.buf.Write([]byte(" where `"))
	//}
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` = "))
	c.buf.WriteUint64(uint64(v))
	return c.w
}

func (c *IntUnSignedCondition[T, V]) NotEqual(v V) *T {
	//if c.buf.Len() > 0 {
	//	c.buf.Write([]byte(" and `"))
	//} else {
	//	c.buf.Write([]byte(" where `"))
	//}
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` != "))
	c.buf.WriteUint64(uint64(v))
	return c.w
}

func (c *IntUnSignedCondition[T, V]) Between(min, max V) *T {
	if min > max {
		min, max = max, min
	}
	//if c.buf.Len() > 0 {
	//	c.buf.Write([]byte(" and `"))
	//} else {
	//	c.buf.Write([]byte(" where `"))
	//}
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` between "))
	c.buf.WriteUint64(uint64(min))
	c.buf.Write([]byte(" and "))
	c.buf.WriteUint64(uint64(max))
	return c.w
}

func (c *IntUnSignedCondition[T, V]) In(vals ...V) *T {
	//if c.buf.Len() > 0 {
	//	c.buf.Write([]byte(" and `"))
	//} else {
	//	c.buf.Write([]byte(" where `"))
	//}
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` in ("))
	for k := 0; k < len(vals); k++ {
		if k > 0 {
			c.buf.WriteByte(',')
		}
		c.buf.WriteUint64(uint64(vals[k]))
	}
	c.buf.Write([]byte(")"))
	return c.w
}
