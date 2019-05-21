package templates

import (
	"os"
	"text/template"
)

const tax = 10.00 / 100.00

type Product struct {
	Name  string
	Price float32
}

func (p Product) PriceWithTax() float32 {
	return p.Price * (1 + tax)
}

const templateString = `
{{- "Item Information" }}
Name: {{ .Name }}
Price: {{ printf "$%.2f" .Price }}
Price (inc GST): {{ calctax .Price | printf "$%.2f" }}
`

func PipelineDemo1() {
	p := Product{
		Name:  "Chocolate WPI",
		Price: 35.0,
	}

	fm := template.FuncMap{}
	fm["calctax"] = func(price float32) float32 {
		return price * (1 + tax)
	}
	t := template.Must(template.New("").Funcs(fm).Parse(templateString))
	t.Execute(os.Stdout, p)
}
