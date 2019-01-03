package templates

import (
	"bytes"
	"html/template"
	"strings"
)

func compileTemplate(value string) *template.Template {
	fns := template.FuncMap{
		"ToLower": strings.ToLower,
	}

	tmpl, err := template.New("").Funcs(fns).Parse(value)
	if err != nil {
		panic(err)
	}
	return tmpl
}

func generate(tmpl *template.Template, data interface{}) (string, error) {
	buf := bytes.NewBuffer(nil)
	if err := tmpl.Execute(buf, data); err != nil {
		return "", err
	}

	content := strings.TrimSpace(buf.String())

	return strings.Replace(content, "&lt;", "<", -1), nil
}
