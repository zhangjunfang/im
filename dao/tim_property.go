package dao

import (
	"reflect"

	"github.com/donnie4w/gdao"
)

type im_property_Id struct {
	gdao.Field
	fieldName  string
	FieldValue *int32
}

func (c *im_property_Id) Name() string {
	return c.fieldName
}

func (c *im_property_Id) Value() interface{} {
	return c.FieldValue
}

type tim_property_Keyword struct {
	gdao.Field
	fieldName  string
	FieldValue *string
}

func (c *tim_property_Keyword) Name() string {
	return c.fieldName
}

func (c *tim_property_Keyword) Value() interface{} {
	return c.FieldValue
}

type tim_property_Valueint struct {
	gdao.Field
	fieldName  string
	FieldValue *int32
}

func (c *tim_property_Valueint) Name() string {
	return c.fieldName
}

func (c *tim_property_Valueint) Value() interface{} {
	return c.FieldValue
}

type tim_property_Valuestr struct {
	gdao.Field
	fieldName  string
	FieldValue *string
}

func (c *tim_property_Valuestr) Name() string {
	return c.fieldName
}

func (c *tim_property_Valuestr) Value() interface{} {
	return c.FieldValue
}

type tim_property_Remark struct {
	gdao.Field
	fieldName  string
	FieldValue *string
}

func (c *tim_property_Remark) Name() string {
	return c.fieldName
}

func (c *tim_property_Remark) Value() interface{} {
	return c.FieldValue
}

type Im_property struct {
	gdao.Table
	Valueint *tim_property_Valueint
	Valuestr *tim_property_Valuestr
	Remark   *tim_property_Remark
	Id       *im_property_Id
	Keyword  *tim_property_Keyword
}

func (u *Im_property) GetValuestr() string {
	return *u.Valuestr.FieldValue
}

func (u *Im_property) SetValuestr(arg string) {
	u.Table.ModifyMap[u.Valuestr.fieldName] = arg
	v := string(arg)
	u.Valuestr.FieldValue = &v
}

func (u *Im_property) GetRemark() string {
	return *u.Remark.FieldValue
}

func (u *Im_property) SetRemark(arg string) {
	u.Table.ModifyMap[u.Remark.fieldName] = arg
	v := string(arg)
	u.Remark.FieldValue = &v
}

func (u *Im_property) GetId() int32 {
	return *u.Id.FieldValue
}

func (u *Im_property) SetId(arg int64) {
	u.Table.ModifyMap[u.Id.fieldName] = arg
	v := int32(arg)
	u.Id.FieldValue = &v
}

func (u *Im_property) GetKeyword() string {
	return *u.Keyword.FieldValue
}

func (u *Im_property) SetKeyword(arg string) {
	u.Table.ModifyMap[u.Keyword.fieldName] = arg
	v := string(arg)
	u.Keyword.FieldValue = &v
}

func (u *Im_property) GetValueint() int32 {
	return *u.Valueint.FieldValue
}

func (u *Im_property) SetValueint(arg int64) {
	u.Table.ModifyMap[u.Valueint.fieldName] = arg
	v := int32(arg)
	u.Valueint.FieldValue = &v
}

func (t *Im_property) Query(columns ...gdao.Column) ([]Im_property, error) {
	if columns == nil {
		columns = []gdao.Column{t.Id, t.Keyword, t.Valueint, t.Valuestr, t.Remark}
	}
	rs, err := t.Table.Query(columns...)
	if rs == nil || err != nil {
		return nil, err
	}
	ts := make([]Im_property, 0, len(rs))
	c := make(chan int16, len(rs))
	for _, rows := range rs {
		t := NewIm_property()
		go copyIm_property(c, rows, t, columns)
		<-c
		ts = append(ts, *t)
	}
	return ts, nil
}

func copyIm_property(channle chan int16, rows []interface{}, t *Im_property, columns []gdao.Column) {
	defer func() { channle <- 1 }()
	for j, core := range rows {
		if core == nil {
			continue
		}
		field := columns[j].Name()
		setfield := "Set" + gdao.ToUpperFirstLetter(field)
		reflect.ValueOf(t).MethodByName(setfield).Call([]reflect.Value{reflect.ValueOf(gdao.GetValue(&core))})
	}
}

func (t *Im_property) QuerySingle(columns ...gdao.Column) (*Im_property, error) {
	if columns == nil {
		columns = []gdao.Column{t.Id, t.Keyword, t.Valueint, t.Valuestr, t.Remark}
	}
	rs, err := t.Table.QuerySingle(columns...)
	if rs == nil || err != nil {
		return nil, err
	}
	rt := NewIm_property()
	for j, core := range rs {
		if core == nil {
			continue
		}
		field := columns[j].Name()
		setfield := "Set" + gdao.ToUpperFirstLetter(field)
		reflect.ValueOf(rt).MethodByName(setfield).Call([]reflect.Value{reflect.ValueOf(gdao.GetValue(&core))})
	}
	return rt, nil
}

func (t *Im_property) Select(columns ...gdao.Column) (*Im_property, error) {
	if columns == nil {
		columns = []gdao.Column{t.Id, t.Keyword, t.Valueint, t.Valuestr, t.Remark}
	}
	rows, err := t.Table.Selects(columns...)
	defer rows.Close()
	if err != nil || rows == nil {
		return nil, err
	}
	buff := make([]interface{}, len(columns))
	if rows.Next() {
		n := NewIm_property()
		cpTim_property(buff, n, columns)
		row_err := rows.Scan(buff...)
		if row_err != nil {
			return nil, row_err
		}
		return n, nil
	}
	return nil, nil
}

func (t *Im_property) Selects(columns ...gdao.Column) ([]*Im_property, error) {
	if columns == nil {
		columns = []gdao.Column{t.Id, t.Keyword, t.Valueint, t.Valuestr, t.Remark}
	}
	rows, err := t.Table.Selects(columns...)
	defer rows.Close()
	if err != nil || rows == nil {
		return nil, err
	}
	ns := make([]*Im_property, 0)
	buff := make([]interface{}, len(columns))
	for rows.Next() {
		n := NewIm_property()
		cpTim_property(buff, n, columns)
		row_err := rows.Scan(buff...)
		if row_err != nil {
			return nil, row_err
		}
		ns = append(ns, n)
	}
	return ns, nil
}

func cpTim_property(buff []interface{}, t *Im_property, columns []gdao.Column) {
	for i, column := range columns {
		field := column.Name()
		switch field {
		case "keyword":
			buff[i] = &t.Keyword.FieldValue
		case "valueint":
			buff[i] = &t.Valueint.FieldValue
		case "valuestr":
			buff[i] = &t.Valuestr.FieldValue
		case "remark":
			buff[i] = &t.Remark.FieldValue
		case "id":
			buff[i] = &t.Id.FieldValue
		}
	}
}

func NewIm_property(tableName ...string) *Im_property {
	valueint := &tim_property_Valueint{fieldName: "valueint"}
	valueint.Field.FieldName = "valueint"
	valuestr := &tim_property_Valuestr{fieldName: "valuestr"}
	valuestr.Field.FieldName = "valuestr"
	remark := &tim_property_Remark{fieldName: "remark"}
	remark.Field.FieldName = "remark"
	id := &im_property_Id{fieldName: "id"}
	id.Field.FieldName = "id"
	keyword := &tim_property_Keyword{fieldName: "keyword"}
	keyword.Field.FieldName = "keyword"
	table := &Im_property{Id: id, Keyword: keyword, Valueint: valueint, Valuestr: valuestr, Remark: remark}
	table.Table.ModifyMap = make(map[string]interface{})
	if len(tableName) == 1 {
		table.Table.TableName = tableName[0]
	} else {
		table.Table.TableName = "tim_property"
	}
	return table
}
