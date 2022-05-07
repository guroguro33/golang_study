package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// template.ParseFilesでテンプレートオブジェクトを作成する
	t, _ := template.ParseFiles(tmpl + ".html")
	// TemplateオブジェクトのExecuteメソッドで値を埋め込む
	t.Execute(w, p) // 第１引数：出力先、第２引数：埋め込むデータ
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// URLは/view/test
	// 「/view/」の６文字分削除したURLを取得
	title := r.URL.Path[len("/view/"):]
	// title（ファイル名）を指定し、textファイルの中身をpとする
	p, _ := loadPage(title)
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	// URLは/edit/test
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
