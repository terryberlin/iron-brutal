package main

import (
	"net/http"
	"text/template"
)

type edit struct {
	FieldName  string
	FieldValue string
}

type page struct {
	FieldName  string
	StoreEdits []edit
}

func main() {
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

const docList = `
<ul >
    {{range .}}
	<li>{{.FieldName}}: <input value={{ .FieldValue}}></input></li>
	<!--input ngModel name="inputStore" id="inputStore" #inputStore="ngModel" (change)="onClick()" [(ngModel)]="tstore" (keyup.enter)="onClick()" (keydown.tab)="onClick()"/-->
    {{end}}
</ul>
<button (click)="processFBC(newFBC.value)" class="btn btn-primary btn-sm">Update It</button>
`

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
