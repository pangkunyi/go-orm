package orm

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

type Page struct {
    PageSize int
    StartIdx int
    EndIdx int
    PageNum int
}

func (this *Page) SetPage(page int){
    if page < 1 {
        this.PageNum = 1
    }
    if this.PageSize < 1 {
        this.PageSize = 10
    }
    this.StartIdx = (this.PageNum - 1) * this.PageSize 
    this.EndIdx = this.StartIdx + this.PageSize - 1
}
