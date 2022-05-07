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
	// 「/view/」の６文字分削除したURL⇨titleを取得
	title := r.URL.Path[len("/view/"):]
	// title（ファイル名）を指定し、textファイルの中身をpとする
	p, err := loadPage(title)
	if err != nil {
		// textファイルからページが生成できなかった時はeditページにリダイレクト
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	}
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

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	// textareaのnameがbodyなので指定してvalueを取得する
	body := r.FormValue("body")
	// titleとbodyを入れたページを生成
	p := &Page{Title: title, Body: []byte(body)}
	// txtファイルを生成し、エラーが発生した場合はerrが入る
	err := p.save()
	if err != nil {
		// Errorメソッドでエラーを返す
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// リダイレクトし、302を返す
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
