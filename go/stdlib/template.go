package main

import (
	"bytes"
	"fmt"
	"text/template"
)

var templateText = `{{ if ge .Balance 100 -}} {{ .Name }}, добрый день! Ваш баланс - {{ .Balance }}₽. Все в порядке. ` +
	`{{- else if ge .Balance 1 -}} {{ .Name }}, добрый день! Ваш баланс - {{ .Balance }}₽. Пора пополнить. ` + 
	`{{- else -}} {{ .Name }}, добрый день! Ваш баланс - {{ .Balance }}₽. Доступ заблокирован. {{- end }}`

type User struct {
	Name    string
	Balance int
}

func renderToString(tpl *template.Template, data any) string {
	var buf bytes.Buffer
	tpl.Execute(&buf, data)
	return buf.String()
}

func main() {
	tpl := template.New("message")
	tpl = template.Must(tpl.Parse(templateText))

	user := User{"Алиса", 111}
	got := renderToString(tpl, user)

	const want = "Алиса, добрый день! Ваш баланс - 500₽. Все в порядке."
	if got != want {
		fmt.Printf("%v: got '%v'", user, got)
	}
}
