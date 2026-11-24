package textTemplate

import (
	"fmt"
	"text/template"
	"os"
)

func GoTextTemplate(){
	fmt.Println("Go text template is running..")
}
// For a short message, fmt.Sprintf may be enough. 
// Templates become useful when the output has multiple sections and repeated items
func TextTemplateCaller(){
   t1 := template.New("t1")
   t1, err := t1.Parse("Value is {{.}}\n")
   if err != nil{
	panic(err)
   }
//    fmt.Println(t1)
   t1.Execute(os.Stdout, "Kaleb")
   t1.Execute(os.Stdout, 9)
   t1.Execute(os.Stdout, []string{"awel", "mamush"})

   create := func(name, t string)*template.Template{
	return template.Must(template.New(name).Parse(t))
   }
   t2 := create("t2", "Value of t2 {{.Name}}\n")
   t2.Execute(os.Stdout, struct {
	Name string
   }{"Mamo"})

   t2.Execute(os.Stdout, map[string]string{
	"Name": "Kemal",
   })

   t3 := create("t3", "{{if . -}}yes {{else -}}no {{end}}\n")
   t3.Execute(os.Stdout, "")
   t3.Execute(os.Stdout, true)

   t4 := create("t4",
    "Range: {{range.}}{{.}} {{end}} \n",
   )

   t4.Execute(os.Stdout, []string{
	"one",
	"two",
	"three",
	"four",
   })


   customerReciept()
}

// Imagine your online store needs a plain-text receipt for every order. 
// The layout stays the same, but the customer name, order number, and purchased items change.
func customerReciept(){
	type reciept struct{
		Customer string
		OrderId string
		Items []string
	}

	create := func(name, t string)*template.Template{
		return template.Must(template.New(name).Parse(t))
	}

	t1 := create("t1", 
     "Reciept for {{.Customer}}\nOrder ID: {{.OrderId}}\n Items: \n{{range.Items}}-{{.}}\n{{end}}",
    )
	t1.Execute(os.Stdout, reciept{
		Customer: "Girma",
		OrderId: "001",
		Items: []string{"PC", "Charger", "Cable"},
	})
}