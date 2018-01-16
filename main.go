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
	<li>{{.FieldName}}: <input>{{ .FieldValue}}</input></li>
	<!--input ngModel name="inputStore" id="inputStore" #inputStore="ngModel" (change)="onClick()" [(ngModel)]="tstore" (keyup.enter)="onClick()" (keydown.tab)="onClick()"/-->
    {{end}}
</ul>
<button (click)="processFBC(newFBC.value)" class="btn btn-primary btn-sm">Update</button>
`

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
