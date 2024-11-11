package views

import (
	"bytes"
	"text/template"
)

func RenderTpl(name string, data interface{}, tpl string) (*bytes.Buffer, error) {

	tmplEntity, err := template.New(name).Funcs(FuncMap).Parse(tpl)
	if err != nil {
		return nil, err
	}

	ret := bytes.NewBufferString("")
	// Run the template to verify the output.
	err = tmplEntity.Execute(ret, data)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
