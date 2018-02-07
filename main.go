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

type edit struct {
	FieldName  string
	FieldValue string
}

//GetStore : GetStore is a structure to hold store demographic data
type GetStore struct {
	FieldName  *string `db:"FieldName" json:"FieldName"`
	FieldValue *string `db:"FieldValue" json:"FieldValue"`
}

type page struct {
	FieldName  string
	StoreEdits []edit
}

type page2 struct {
	FieldName string
	GetStores []GetStore
}

func main() {
	http.HandleFunc("/getstores", GetStores)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content Type", "text/html")

		templates := template.New("template")
		templates.New("Body").Parse(doc)
		templates.New("List").Parse(docList)

		storeEdits := []edit{{FieldName: "Store", FieldValue: "1701"}, {FieldName: "Organization", FieldValue: "TJJohnTest"}, {FieldName: "Inventory_Group", FieldValue: "TJMMFA"}}

		page := page{FieldName: "Store Edits", StoreEdits: storeEdits}
		templates.Lookup("Body").Execute(w, page)

	})

	http.ListenAndServe(":8000", nil)
}

// const docList = `
// <ul >
//     {{range .}}
// 	<li>{{.FieldName}}: <input value={{ .FieldValue}}></input></li>
// 	<!--input ngModel name="inputStore" id="inputStore" #inputStore="ngModel" (change)="onClick()" [(ngModel)]="tstore" (keyup.enter)="onClick()" (keydown.tab)="onClick()"/-->
//     {{end}}
// </ul>
// <button (click)="processFBC(newFBC.value)" class="btn btn-primary btn-sm">Update It</button>
// `

//GetStores : GetStores is a function to request store data.
func GetStores(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content Type", "text/html")
	templates := template.New("template")
	templates.New("Body").Parse(doc2)
	templates.New("List").Parse(docList2)

	//copyfromstore := r.URL.Query().Get("copyfromstore")
	//copytostore := r.URL.Query().Get("copytostore")

	//sql := `Exec get_store_demographics $1, $2`
	sql := `
			select 'Store' as FieldName, '1702' as FieldValue
		union select 'Organization','TJJohnTestMe'
		union select 'Inventory_Group','TJMMFAMe'
	`
	getstores := []GetStore{}
	err := DB().Select(&getstores, sql)
	if err != nil {
		log.Println(err)
	}
	//json, err := json.Marshal(getstores)
	if err != nil {
		log.Println(err)
	}
	//fmt.Fprintf(w, string(json))

	page2 := page2{FieldName: "Get Stores", GetStores: getstores}
	templates.Lookup("Body").Execute(w, page2)

}

//DB : DB is a function that connects to SQL server.
func DB() *sqlx.DB {
	serv := os.Getenv("DB_SERVER")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	database := os.Getenv("DB_DATABASE")

	db, err := sqlx.Connect("mssql", fmt.Sprintf(`server=%s;user id=%s;password=%s;database=%s;log64;encrypt=disable`, serv, user, pass, database))

	//db, err := sqlx.Connect("mssql", "server=localhost;user id=QUIKSERVE/terryb;password=;database=quikserve_dev;log64;encrypt=disable")

	if err != nil {
		log.Println(err)
	}
	return db
}

//<!--Link className="btn btn-primary" to={ `/units/${this.props.match.params.id}/edit` } -->

// const docList = `
// <div>
//     <Row className="justify-content-between mb-4" >
//       <Col xs="4">
//         <h2>{ this.state.unit.name }</h2>
//       </Col>
//       <Col xs="2">
// 		<Link className="btn btn-primary">
//           <IconMenuItem icon="edit" text="Edit" />
//         </Link>
//       </Col>
//     </Row>
//     <Row>
//       <Col xs="6">
//         <Card>
//           <CardBlock>
//             <CardTitle>Details</CardTitle>
//             <KeyVal k="Name" val={ this.state.unit.name } />
//             <KeyVal k="Identifier" val={ this.state.unit.identifier } />
//             <KeyVal k="Created" val={ moment.unix(this.state.unit.created_at).fromNow() } />
//             <KeyVal k="Last Updated" val={ moment.unix(this.state.unit.updated_at).fromNow() } />
//           </CardBlock>
//         </Card>
//       </Col>
//       <Col xs="6">
//         <Card>
//           <CardBlock>
//             <CardTitle>Contact Information</CardTitle>
//             <KeyVal k="Phone" val={ formatPhoneNumber(parsePhoneNumber(this.state.unit.phone), 'National') } />
//             <KeyVal k="Address" val={ this.state.unit.address.street_1 } />
//             <KeyVal k="" val={ this.state.unit.address.street_2 } />
//             <KeyVal k="State" val={ this.state.unit.address.state } />
//             <KeyVal k="Zip" val={ this.state.unit.address.zip } />
//             <KeyVal k="Country" val={ this.state.unit.address.country } />
//           </CardBlock>
//         </Card>
//       </Col>
//     </Row>
//   </div>
// `

const doc = `
 <!DOCTYPE html>
 <html>
     <head><title>{{.FieldName}}</title></head>
     <body>
         <h1>Store Edits</h1>
         {{template "List" .StoreEdits}}
     </body>
 </html>
 `

const docList = `
<ul >
    {{range .}}
	<li>{{.FieldName}}: <input value={{ .FieldValue}}></input></li>
	<!--input ngModel name="inputStore" id="inputStore" #inputStore="ngModel" (change)="onClick()" [(ngModel)]="tstore" (keyup.enter)="onClick()" (keydown.tab)="onClick()"/-->
    {{end}}
</ul>
<button (click)="processFBC(newFBC.value)" class="btn btn-primary btn-sm">Update It</button>
`

const doc2 = `
 <!DOCTYPE html>
 <html>
     <head><title>{{.FieldName}}</title></head>
     <body>
         <h1>Store Edits</h1>
         {{template "List" .GetStores}}
     </body>
 </html>
 `

const docList2 = `
<ul >
    {{range .}}
	<li>{{.FieldName}}: <input value={{ .FieldValue}}></input></li>
	<!--input ngModel name="inputStore" id="inputStore" #inputStore="ngModel" (change)="onClick()" [(ngModel)]="tstore" (keyup.enter)="onClick()" (keydown.tab)="onClick()"/-->
    {{end}}
</ul>
<button (click)="processFBC(newFBC.value)" class="btn btn-primary btn-sm">Update It</button>
`
