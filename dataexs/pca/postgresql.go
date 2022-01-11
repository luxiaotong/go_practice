package pca

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/lib/pq"
	"github.com/pkg/errors"
)

//Impl 是DBInterface抽象的PostgreSQL具体实现
type Impl struct {
	DB *sql.DB
}

// TableList 是用于获取指定数据库的Table列表
// 	param:
// 	return: []TableInfo, 表信息数组
// 	return: error, 错误信息
func (m *Impl) TableList() ([]*TableInfo, error) {
	if m.DB == nil {
		return nil, ErrDBNotConnect
	}
	q := "SELECT table_name, table_type FROM information_schema.tables WHERE table_schema=ANY(current_schemas(false))"
	rows, err := m.DB.Query(q)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	var tbName, tbType string
	var t TableType
	var result []*TableInfo

	for rows.Next() {
		_ = rows.Scan(&tbName, &tbType)
		switch tbType {
		case "BASE TABLE":
			t = TableType_TABLE
		case "VIEW":
			t = TableType_VIEW
		}
		info, err := m.tableFields(tbName)
		if err != nil {
			return nil, errors.Wrapf(err, "table: %s", tbName)
		}
		result = append(result, &TableInfo{Name: tbName, Type: t, FieldCount: uint32(len(info.Fields))})
	}

	return result, nil
}

// TableDetail 是用于获取指定表的字段列表
// 	param: string tbName, 表名
// 	return: TableDetail
// 	return: error, 错误信息
func (m *Impl) TableDetail(tbName string) (*TableFieldsInfo, error) {
	info, err := m.tableFields(tbName)
	if err != nil {
		return nil, err
	}

	tbName = pq.QuoteIdentifier(tbName)
	c, err := m.getRecordSize(tbName)
	if err != nil {
		return nil, err
	}
	info.RecordSize = c
	return info, nil
}

// tableFields 返回表字段信息
//
//  param: string tbName
//  return: *TableFieldsInfo
//  return: error
func (m *Impl) tableFields(tbName string) (*TableFieldsInfo, error) {
	if m.DB == nil {
		return nil, ErrDBNotConnect
	}
	q := "SELECT a.column_name, a.data_type, a.character_maximum_length, a.numeric_precision, " +
		" a.numeric_scale, a.datetime_precision, tco.constraint_type" +
		" FROM information_schema.columns a" +
		" LEFT JOIN information_schema.key_column_usage kcu ON kcu.column_name = a.column_name" +
		" AND kcu.table_name = a.table_name" +
		" LEFT JOIN information_schema.table_constraints tco ON kcu.constraint_name = tco.constraint_name" +
		" AND kcu.constraint_schema = tco.constraint_schema AND tco.constraint_type = 'PRIMARY KEY'" +
		" WHERE a.table_name = $1"
	rows, err := m.DB.Query(q, tbName)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	var (
		colName, colType                   string
		constraintType                     sql.NullString
		charMaxLen, numPrecision, numScale sql.NullInt64
		dtPrecistion                       sql.NullInt64
		list                               []*FieldInfo
	)

	for rows.Next() {
		_ = rows.Scan(&colName, &colType, &charMaxLen, &numPrecision, &numScale, &dtPrecistion, &constraintType)

		sizeText := ""
		switch colType {
		case "bit", "bit varying", "character varying":
			if charMaxLen.Int64 > 0 {
				sizeText = fmt.Sprintf("(%d)", charMaxLen.Int64)
			}
		case "character":
			if charMaxLen.Int64 > 0 {
				sizeText = fmt.Sprintf("(%d)", charMaxLen.Int64)
			} else if charMaxLen.Int64 == 0 {
				colType = "character varying"
			}
		case "numeric":
			// fmt.Println(numPrecision, numScale)
			if numPrecision.Int64 > 0 && numScale.Int64 > 0 {
				sizeText = fmt.Sprintf("(%d,%d)", numPrecision.Int64, numScale.Int64)
			} else if numPrecision.Int64 > 0 {
				sizeText = fmt.Sprintf("(%d)", numPrecision.Int64)
			}
		case "USER-DEFINED":
			colType = "character varying"
		}
		colType += sizeText

		list = append(list, &FieldInfo{Name: colName, Type: colType, IsPrimary: constraintType.Valid})
	}

	return &TableFieldsInfo{Fields: list}, nil
}

func quoteJoin(fields []string) string {
	ff := make([]string, len(fields))
	for i, field := range fields {
		ff[i] = pq.QuoteIdentifier(field)
	}
	return strings.Join(ff, ",")
}

func realData(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	// log.Debug("type is %s : %v", reflect.TypeOf(v).String(), v)
	switch v := v.(type) {
	case []uint8:
		// log.Debug("it's string %s", v)
		return string(v)
	default:
		return v
	}
}

func (m *Impl) DataList(tbName string, fields, orders []string, pageNo, pageSize uint32) ([]map[string]interface{}, uint32, error) {
	return m.DataListPos(tbName, fields, orders, pageNo*pageSize, pageSize)
}

func (m *Impl) DataListPos(tbName string, fields, orders []string, offset, limit uint32) ([]map[string]interface{}, uint32, error) {
	if m.DB == nil {
		return nil, 0, ErrDBNotConnect
	}
	tbName = pq.QuoteIdentifier(tbName)
	fff := schemasToFields(fields)
	ooo := schemasToFields(orders)
	selects := quoteJoin(fff)
	if len(selects) == 0 {
		selects = "*"
	}
	order := quoteJoin(ooo)
	if len(order) > 0 {
		order = "ORDER BY " + order
	}
	q := fmt.Sprintf("SELECT %s FROM %s %s LIMIT $1 OFFSET $2", selects, tbName, order)
	// log.Debug("query is %s", q)
	rows, err := m.DB.Query(q, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer func() {
		_ = rows.Close()
	}()

	cols, err := rows.Columns()
	if err != nil {
		return nil, 0, err
	}

	var result []map[string]interface{}
	vals := make([]interface{}, len(cols))

	for rows.Next() {
		for i := range cols {
			vals[i] = &vals[i]
		}

		err = rows.Scan(vals...)
		if err != nil {
			return nil, 0, err
		}

		dict := make(map[string]interface{}, len(cols))
		for i, raw := range vals {
			dict[cols[i]] = realData(raw)
		}
		result = append(result, dict)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	if offset > 0 {
		return result, 0, nil
	}

	c, err := m.getRecordSize(tbName)
	if err != nil {
		return nil, 0, err
	}

	return result, c, nil
}

func (m *Impl) getRecordSize(table string) (uint32, error) {
	var c uint32
	q := fmt.Sprintf("SELECT count(*) FROM %s", table)
	row := m.DB.QueryRow(q)
	err := row.Scan(&c)
	return c, err
}

// LastModified 是用于获取指定数据库的Log Sequence Number
//  param:
//  return: string, 返回log sequence number, 如果没有则返回空字符串
//  return: error
func (m *Impl) LastModified() (string, error) {
	if m.DB == nil {
		return "", ErrDBNotConnect
	}
	var l string

	row := m.DB.QueryRow("SELECT pg_current_wal_flush_lsn();")
	err := row.Scan(&l)
	if err != nil {
		return "", err
	}

	return l, err
}

func (m *Impl) Close() {
	if m.DB == nil {
		return
	}
	_ = m.DB.Close()
	m.DB = nil
}

func (m *Impl) ToPostgresType(fields []*FieldInfo) []*FieldInfo {
	return fields
}

func schemaToField(s string) string {
	const schema = "schema://"
	if strings.HasPrefix(s, schema) {
		return s[len(schema):]
	}
	return s
}

func schemasToFields(fields []string) []string {
	fff := make([]string, len(fields))
	for i, field := range fields {
		fff[i] = schemaToField(field)
	}
	return fff
}

// CreateTable 用于创建表
// 	param: string tbName 数据表名
// 	param: TableDetail tbDetail 数据表的信息, 包括字段信息, 主键
// 	return: error, 错误信息, 成功时返回 nil
func (m *Impl) CreateTable(tbName string, tbDetail *TableFieldsInfo) error {
	if m.DB == nil {
		return ErrDBNotConnect
	}

	var cols []string
	var pks []string
	for _, field := range tbDetail.Fields {
		n := pq.QuoteIdentifier(schemaToField(field.Name))
		col := n + " " + field.Type
		col += " NULL"
		cols = append(cols, col)
		if field.IsPrimary {
			pks = append(pks, n)
		}
	}

	if len(pks) > 0 {
		pri := fmt.Sprintf("Primary Key(%s)", strings.Join(pks, ","))
		cols = append(cols, pri)
	}

	colStr := strings.Join(cols, ",")

	q := fmt.Sprintf("CREATE TABLE %s (%s)", pq.QuoteIdentifier(tbName), colStr)

	_, err := m.DB.Exec(q)
	if err != nil {
		return errors.Wrap(err, "create table: "+q)
	}
	return nil
}

func (m *Impl) DropTable(tbName string) error {
	if m.DB == nil {
		return ErrDBNotConnect
	}

	// if exists drop it
	q := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", pq.QuoteIdentifier(tbName))

	_, err := m.DB.Exec(q)
	if err != nil {
		return err
	}
	return nil
}

func (m *Impl) DropAll(user string) error {
	if m.DB == nil {
		return ErrDBNotConnect
	}

	// if exists drop it
	q := fmt.Sprintf("DROP OWNED BY %s", pq.QuoteIdentifier(user))

	_, err := m.DB.Exec(q)
	if err != nil {
		return err
	}
	return nil
}

// CreateIndex 用于创建表索引
// 	param: string tbName 表名
// 	param: string idxName 索引名
// 	param: []string idx 表的索引
// 	return: error, 错误信息, 成功时返回 nil
func (m *Impl) CreateIndex(tbName, idxName string, idx []string) error {
	if m.DB == nil {
		return ErrDBNotConnect
	}

	idxStr := strings.Join(idx, ",")
	q := fmt.Sprintf("CREATE INDEX %s ON %s USING btree (%s)",
		pq.QuoteIdentifier(idxName), pq.QuoteIdentifier(tbName), idxStr)

	_, err := m.DB.Exec(q)
	if err != nil {
		return err
	}
	return nil
}

// Insert 用于批量插入数据
// 	param: string tbName 表名
// 	param: []map[string]interface{} data 数据
// 	return: error, 错误信息, 成功时返回 nil
func (m *Impl) Insert(tbName string, data []map[string]interface{}) error {
	if m.DB == nil {
		return ErrDBNotConnect
	}

	if len(data) == 0 {
		return nil
	}

	txn, e := m.DB.Begin()
	if e != nil {
		return e
	}

	var cols []string
	for k := range data[0] {
		cols = append(cols, k)
	}
	ncols := schemasToFields(cols)
	stmt, e := txn.Prepare(pq.CopyIn(tbName, ncols...))
	if e != nil {
		return e
	}

	defer func() {
		_ = stmt.Close()
	}()

	for _, d := range data {
		var vals []interface{}
		for _, v := range cols {
			vals = append(vals, d[v])
		}
		_, e = stmt.Exec(vals...)
		if e != nil {
			_ = txn.Rollback()
			return e
		}

	}

	_, err := stmt.Exec()
	if err != nil {
		_ = txn.Rollback()
		return err
	}

	err = txn.Commit()
	if err != nil {
		// log.Error("error: %v", err)
		return err
	}

	return nil
}
