package svc_db

import "github.com/walleframe/walle/util"

type StringConditionV2[T any] struct {
	w    *T
	buf  *util.Builder
	name string
}

func NewStringConditionV2[T any](w *T, buf *util.Builder, name string) *StringConditionV2[T] {
	return &StringConditionV2[T]{
		w:    w,
		buf:  buf,
		name: name,
	}
}

func (c *StringConditionV2[T]) Equal() *T {
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` = ?"))
	return c.w
}

func (c *StringConditionV2[T]) NotEqual() *T {
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` != ?"))
	return c.w
}

func (c *StringConditionV2[T]) Like() *T {
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` like ?"))
	return c.w
}

type NumberConditionV2[T any] struct {
	w    *T
	buf  *util.Builder
	name string
}

func NewNumberConditionV2[T any](w *T, buf *util.Builder, name string) *NumberConditionV2[T] {
	return &NumberConditionV2[T]{
		w:    w,
		buf:  buf,
		name: name,
	}
}

func (c *NumberConditionV2[T]) LessThen() *T {
	//if c.buf.Len() > 0 {
	//	c.buf.Write([]byte(" and `"))
	//} else {
	//	c.buf.Write([]byte(" where `"))
	//}
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` < ?"))
	return c.w
}

func (c *NumberConditionV2[T]) LessEqual() *T {
	//if c.buf.Len() > 0 {
	//	c.buf.Write([]byte(" and `"))
	//} else {
	//	c.buf.Write([]byte(" where `"))
	//}
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` <= ?"))
	return c.w
}

func (c *NumberConditionV2[T]) GreaterThen() *T {
	//if c.buf.Len() > 0 {
	//	c.buf.Write([]byte(" and `"))
	//} else {
	//	c.buf.Write([]byte(" where `"))
	//}
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` > ?"))
	return c.w
}

func (c *NumberConditionV2[T]) GreaterEqual() *T {
	//if c.buf.Len() > 0 {
	//	c.buf.Write([]byte(" and `"))
	//} else {
	//	c.buf.Write([]byte(" where `"))
	//}
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` >= ?"))
	return c.w
}

func (c *NumberConditionV2[T]) Equal() *T {
	//if c.buf.Len() > 0 {
	//	c.buf.Write([]byte(" and `"))
	//} else {
	//	c.buf.Write([]byte(" where `"))
	//}
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` = ?"))
	return c.w
}

func (c *NumberConditionV2[T]) NotEqual() *T {
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` != ?"))
	return c.w
}

func (c *NumberConditionV2[T]) Between() *T {
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` between ? and ?"))
	return c.w
}

func (c *NumberConditionV2[T]) In(vals ...int64) *T {
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

func (c *NumberConditionV2[T]) InInt32(vals ...int32) *T {
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

type ExecOrder[T any] struct {
	w    *T
	buf  *util.Builder
	name string
}

func NewExecOrder[T any](w *T, buf *util.Builder, name string) *ExecOrder[T] {
	return &ExecOrder[T]{
		w:    w,
		buf:  buf,
		name: name,
	}
}

func (c *ExecOrder[T]) Asc() *T {
	c.buf.WriteString(" order by ")
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` asc"))
	return c.w
}

func (c *ExecOrder[T]) Desc() *T {
	c.buf.WriteString(" order by ")
	c.buf.WriteByte('`')
	c.buf.WriteString(c.name)
	c.buf.Write([]byte("` desc"))
	return c.w
}
