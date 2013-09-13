package orm

import(
	"fmt"
)

var (
	sqlMapper SqlMapper = SqlMapper{StmtMap: make(map[string]func(Query)*Statement)}
)

const (
	SQL_KEY_FORMAT = `%s.%s`
	SQL_INSERT_KEY = `sql_insert_key`
	SQL_SELECT_KEY = `sql_select_key`
	SQL_GET_KEY = `sql_get_key`
	SQL_UPDATE_KEY = `sql_update_key`
	SQL_DELETE_KEY = `sql_delete_key`
)

type SqlMapper struct{
	StmtMap map[string]func(Query)*Statement
}

func AddStmt(namespace, sqlKey string, stmtFunc func(Query)*Statement) {
	sqlMapper.StmtMap[fmt.Sprintf(SQL_KEY_FORMAT, namespace, sqlKey)] = stmtFunc
}

func AddSaveStmt(namespace string, stmtFunc func(Query)*Statement) {
	sqlMapper.StmtMap[fmt.Sprintf(SQL_KEY_FORMAT, namespace, SQL_INSERT_KEY)] = stmtFunc
}

func AddListStmt(namespace string, stmtFunc func(Query)*Statement) {
	sqlMapper.StmtMap[fmt.Sprintf(SQL_KEY_FORMAT, namespace, SQL_SELECT_KEY)] = stmtFunc
}

func AddGetStmt(namespace string, stmtFunc func(Query)*Statement) {
	sqlMapper.StmtMap[fmt.Sprintf(SQL_KEY_FORMAT, namespace, SQL_GET_KEY)] = stmtFunc
}

func AddUpdateStmt(namespace string, stmtFunc func(Query)*Statement) {
	sqlMapper.StmtMap[fmt.Sprintf(SQL_KEY_FORMAT, namespace, SQL_UPDATE_KEY)] = stmtFunc
}

func AddDeleteStmt(namespace string, stmtFunc func(Query)*Statement) {
	sqlMapper.StmtMap[fmt.Sprintf(SQL_KEY_FORMAT, namespace, SQL_DELETE_KEY)] = stmtFunc
}

func MustStmt(namespace, sqlKey string, query Query) *Statement {
	key := fmt.Sprintf(SQL_KEY_FORMAT, namespace, sqlKey)
	stmtFunc := sqlMapper.StmtMap[key]
	if stmtFunc==nil {
		panic(fmt.Sprintf("statement func not found with key[%s]", key))
	}
	stmt := stmtFunc(query)
	if stmt==nil {
		panic(fmt.Sprintf("statement not found with key[%s]", key))
	}
	fmt.Printf("sql: %s\nparams: %v\n", stmt.Sql, stmt.Params)
	return stmt
}
