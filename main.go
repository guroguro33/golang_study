package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
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

// 毎回renderTemplateでhtml読み込まなくていいようにキャッシュする
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// // template.ParseFilesでテンプレートオブジェクトを作成する
	// t, _ := template.ParseFiles(tmpl + ".html")
	// // TemplateオブジェクトのExecuteメソッドで値を埋め込む
	// t.Execute(w, p) // 第１引数：出力先、第２引数：埋め込むデータ

	// キャッシュしたテンプレート名をExecuteTemplateでページデータpと共に渡す
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	// // URLは/view/test
	// // 「/view/」の６文字分削除したURL⇨titleを取得
	// title := r.URL.Path[len("/view/"):]
	// title（ファイル名）を指定し、textファイルの中身をpとする
	p, err := loadPage(title)
	if err != nil {
		// textファイルからページが生成できなかった時はeditページにリダイレクト
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	// // URLは/edit/test
	// title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/save/"):]
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

// titleを取得するための正規表現を作成
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// 各handlerにtitleを渡すためのハンドラーを準備
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// FindStringSubmatchで一致したURLをスライスで格納して返す
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2]) // m[0]はパスの最後全部、m[1]からスライスしたものが入っている
	}
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
