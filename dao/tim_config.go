package dao

import (
	"reflect"

	"github.com/donnie4w/gdao"
)

type im_config_Valuestr struct {
	gdao.Field
	fieldName  string
	FieldValue *string
}

func (c *im_config_Valuestr) Name() string {
	return c.fieldName
}

func (c *im_config_Valuestr) Value() interface{} {
	return c.FieldValue
}

type im_config_Createtime struct {
	gdao.Field
	fieldName  string
	FieldValue *string
}

func (c *im_config_Createtime) Name() string {
	return c.fieldName
}

func (c *im_config_Createtime) Value() interface{} {
	return c.FieldValue
}

type im_config_Remark struct {
	gdao.Field
	fieldName  string
	FieldValue *string
}

func (c *im_config_Remark) Name() string {
	return c.fieldName
}

func (c *im_config_Remark) Value() interface{} {
	return c.FieldValue
}

type im_config_Id struct {
	gdao.Field
	fieldName  string
	FieldValue *int32
}

func (c *im_config_Id) Name() string {
	return c.fieldName
}

func (c *im_config_Id) Value() interface{} {
	return c.FieldValue
}

type im_config_Keyword struct {
	gdao.Field
	fieldName  string
	FieldValue *string
}

func (c *im_config_Keyword) Name() string {
	return c.fieldName
}

func (c *im_config_Keyword) Value() interface{} {
	return c.FieldValue
}

type Im_config struct {
	gdao.Table
	Id         *im_config_Id
	Keyword    *im_config_Keyword
	Valuestr   *im_config_Valuestr
	Createtime *im_config_Createtime
	Remark     *im_config_Remark
}

func (u *Im_config) GetId() int32 {
	return *u.Id.FieldValue
}

func (u *Im_config) SetId(arg int64) {
	u.Table.ModifyMap[u.Id.fieldName] = arg
	v := int32(arg)
	u.Id.FieldValue = &v
}

func (u *Im_config) GetKeyword() string {
	return *u.Keyword.FieldValue
}

func (u *Im_config) SetKeyword(arg string) {
	u.Table.ModifyMap[u.Keyword.fieldName] = arg
	v := string(arg)
	u.Keyword.FieldValue = &v
}

func (u *Im_config) GetValuestr() string {
	return *u.Valuestr.FieldValue
}

func (u *Im_config) SetValuestr(arg string) {
	u.Table.ModifyMap[u.Valuestr.fieldName] = arg
	v := string(arg)
	u.Valuestr.FieldValue = &v
}

func (u *Im_config) GetCreatetime() string {
	return *u.Createtime.FieldValue
}

func (u *Im_config) SetCreatetime(arg string) {
	u.Table.ModifyMap[u.Createtime.fieldName] = arg
	v := string(arg)
	u.Createtime.FieldValue = &v
}

func (u *Im_config) GetRemark() string {
	return *u.Remark.FieldValue
}

func (u *Im_config) SetRemark(arg string) {
	u.Table.ModifyMap[u.Remark.fieldName] = arg
	v := string(arg)
	u.Remark.FieldValue = &v
}

func (t *Im_config) Query(columns ...gdao.Column) ([]Im_config, error) {
	if columns == nil {
		columns = []gdao.Column{t.Createtime, t.Remark, t.Id, t.Keyword, t.Valuestr}
	}
	rs, err := t.Table.Query(columns...)
	if rs == nil || err != nil {
		return nil, err
	}
	ts := make([]Im_config, 0, len(rs))
	c := make(chan int16, len(rs))
	for _, rows := range rs {
		t := NewIm_config()
		go copyIm_config(c, rows, t, columns)
		<-c
		ts = append(ts, *t)
	}
	return ts, nil
}

func copyIm_config(channle chan int16, rows []interface{}, t *Im_config, columns []gdao.Column) {
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

func (t *Im_config) QuerySingle(columns ...gdao.Column) (*Im_config, error) {
	if columns == nil {
		columns = []gdao.Column{t.Createtime, t.Remark, t.Id, t.Keyword, t.Valuestr}
	}
	rs, err := t.Table.QuerySingle(columns...)
	if rs == nil || err != nil {
		return nil, err
	}
	rt := NewIm_config()
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

func (t *Im_config) Select(columns ...gdao.Column) (*Im_config, error) {
	if columns == nil {
		columns = []gdao.Column{t.Createtime, t.Remark, t.Id, t.Keyword, t.Valuestr}
	}
	rows, err := t.Table.Selects(columns...)
	defer rows.Close()
	if err != nil || rows == nil {
		return nil, err
	}
	buff := make([]interface{}, len(columns))
	if rows.Next() {
		n := NewIm_config()
		cpIm_config(buff, n, columns)
		row_err := rows.Scan(buff...)
		if row_err != nil {
			return nil, row_err
		}
		return n, nil
	}
	return nil, nil
}

func (t *Im_config) Selects(columns ...gdao.Column) ([]*Im_config, error) {
	if columns == nil {
		columns = []gdao.Column{t.Createtime, t.Remark, t.Id, t.Keyword, t.Valuestr}
	}
	rows, err := t.Table.Selects(columns...)
	defer rows.Close()
	if err != nil || rows == nil {
		return nil, err
	}
	ns := make([]*Im_config, 0)
	buff := make([]interface{}, len(columns))
	for rows.Next() {
		n := NewIm_config()
		cpIm_config(buff, n, columns)
		row_err := rows.Scan(buff...)
		if row_err != nil {
			return nil, row_err
		}
		ns = append(ns, n)
	}
	return ns, nil
}

func cpIm_config(buff []interface{}, t *Im_config, columns []gdao.Column) {
	for i, column := range columns {
		field := column.Name()
		switch field {
		case "id":
			buff[i] = &t.Id.FieldValue
		case "keyword":
			buff[i] = &t.Keyword.FieldValue
		case "valuestr":
			buff[i] = &t.Valuestr.FieldValue
		case "createtime":
			buff[i] = &t.Createtime.FieldValue
		case "remark":
			buff[i] = &t.Remark.FieldValue
		}
	}
}

func NewIm_config(tableName ...string) *Im_config {
	createtime := &im_config_Createtime{fieldName: "createtime"}
	createtime.Field.FieldName = "createtime"
	remark := &im_config_Remark{fieldName: "remark"}
	remark.Field.FieldName = "remark"
	id := &im_config_Id{fieldName: "id"}
	id.Field.FieldName = "id"
	keyword := &im_config_Keyword{fieldName: "keyword"}
	keyword.Field.FieldName = "keyword"
	valuestr := &im_config_Valuestr{fieldName: "valuestr"}
	valuestr.Field.FieldName = "valuestr"
	table := &Im_config{Keyword: keyword, Valuestr: valuestr, Createtime: createtime, Remark: remark, Id: id}
	table.Table.ModifyMap = make(map[string]interface{})
	if len(tableName) == 1 {
		table.Table.TableName = tableName[0]
	} else {
		table.Table.TableName = "im_config"
	}
	return table
}
