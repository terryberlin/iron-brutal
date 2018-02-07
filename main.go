package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

//GetStore : GetStore is a structure to hold store demographic data
type GetStore struct {
	FieldName  *string `db:"FieldName" json:"FieldName"`
	FieldValue *string `db:"FieldValue" json:"FieldValue"`
	SortValue  *string `db:"SortValue" json:"SortValue"`
}

type page struct {
	FieldName string
	GetStores []GetStore
}

func main() {
	http.HandleFunc("/getstores", GetStores)
	http.ListenAndServe(":8000", nil)
}

//GetStores : GetStores is a function to request store data.
func GetStores(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content Type", "text/html")
	templates := template.New("template")
	templates.New("Body").Parse(doc)
	templates.New("List").Parse(docList)

	copyfromstore := r.URL.Query().Get("copyfromstore")

	sql := `Exec getStoreDemographics $1`

	getstores := []GetStore{}
	err := DB().Select(&getstores, sql, copyfromstore)
	if err != nil {
		log.Println(err)
	}

	page := page{FieldName: "Get Stores", GetStores: getstores}
	templates.Lookup("Body").Execute(w, page)
}

//DB : DB is a function that connects to SQL server.
func DB() *sqlx.DB {
	serv := os.Getenv("DB_SERVER")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	database := os.Getenv("DB_DATABASE")

	db, err := sqlx.Connect("mssql", fmt.Sprintf(`server=%s;user id=%s;password=%s;database=%s;log64;encrypt=disable`, serv, user, pass, database))

	if err != nil {
		log.Println(err)
	}
	return db
}

const doc = `
 <!DOCTYPE html>
 <html>
     <head><title>{{.FieldName}}</title></head>
     <body>
         <h1>Store Edits</h1>
         {{template "List" .GetStores}}
     </body>
 </html>
 `
const docList = `
<ul >
    {{range .}}
	<li>{{.FieldName}}: <input value={{ .FieldValue}}></input></li>
    {{end}}
</ul>
<button (click)="processFBC(newFBC.value)" class="btn btn-primary btn-sm">Update It</button>
`
