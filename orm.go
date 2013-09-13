package orm

import(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db = initDb()
)

func initDb() (db *sql.DB){
	db, err := sql.Open("sqlite3", "./big_moment.db")
	if err != nil {
		panic(err)
	}
	return db
}

func Close(){
	db.Close()
}

type Query interface {
}

type RowScanner interface {
	Scan(dest ...interface{}) error
}

type Statement struct {
	Sql string
	Scan func(RowScanner) (interface{}, error)
	Params []interface{}
}

func Delete(namespace string, query Query) (count int64, err error) {
	return ExecUpdate(namespace, SQL_DELETE_KEY, query)
}

func Update(namespace string, query Query) (count int64, err error) {
	return ExecUpdate(namespace, SQL_UPDATE_KEY, query)
}
/**
 * save query to database, execute insert into table command
 */
func ExecUpdate(namespace, sqlKey string, query Query) (count int64, err error) {
	stmt := MustStmt(namespace, sqlKey, query)
	pstmt, err := db.Prepare(stmt.Sql)
	if err!= nil {
		return
	}
	result, err := pstmt.Exec(stmt.Params...)
	if err!= nil {
		return
	}
	return result.RowsAffected()
}

func Save(namespace string, query Query) (id int64, err error) {
	return ExecSave(namespace, SQL_INSERT_KEY, query)
}
/**
 * save query to database, execute insert into table command
 */
func ExecSave(namespace, sqlKey string, query Query) (id int64, err error) {
	stmt := MustStmt(namespace, sqlKey, query)
	pstmt, err := db.Prepare(stmt.Sql)
	if err!= nil {
		return
	}
	result, err := pstmt.Exec(stmt.Params...)
	if err!= nil {
		return
	}
	return result.LastInsertId()
}

func List(namespace string, query Query) (entities []interface{}, err error) {
	return ExecList(namespace, SQL_SELECT_KEY, query)
}
/**
 * list entities from database
 */
func ExecList(namespace, sqlKey string, query Query) (entities []interface{}, err error) {
	stmt := MustStmt(namespace, sqlKey, query)
	pstmt, err := db.Prepare(stmt.Sql)
	if err!= nil {
		return
	}
	rows, err := pstmt.Query(stmt.Params...)
	if err!= nil {
		return
	}
	defer rows.Close()
	entities = make([]interface{}, 0)
	for rows.Next() {
		entity, err := stmt.Scan(rows)
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}
	err = rows.Err()
	return
}

func Get(namespace string, query Query) (entity interface{}, err error) {
	return ExecGet(namespace, SQL_GET_KEY, query)
}

/**
 * get one entity from database
 */
func ExecGet(namespace, sqlKey string, query Query) (entity interface{}, err error) {
	stmt := MustStmt(namespace, sqlKey, query)
	pstmt, err := db.Prepare(stmt.Sql)
	if err!= nil {
		return
	}
	row := pstmt.QueryRow(stmt.Params...)
	entity, err = stmt.Scan(row)
	if err != nil {
		return nil, err
	}
	return
}
