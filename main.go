package main

import (
	"belajar-routing/connection"
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var Data = map[string]interface{}{
	"Title":   "Personal Web",
	"isLogin": true,
}

// Struct Type
type Blog struct {
	Id          int
	Title       string
	Author      string
	Content     string
	Image       string
	Post_Date   time.Time
	Format_Date string
}

// ARRAY OF OBJ
// var Blogs = []Blog{
// 	{
// 		Title:     "Pasar Coding Indonesia dinilai Masih Cukup Menjanjikan",
// 		Post_Date: time.Now().String(),
// 		Author:    "Jordi El Nino",
// 		Content:   "Lorem ipsum dolor sit, amet consectetur adipisicing elit. Odit sit excepturi reiciendis.",
// 	},
// }

func main() {

	// deklarsi NewRouter
	router := mux.NewRouter()

	connection.DatabaseConnect()
	// Creating static Folder
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("./images/"))))

	// Deklarasi Handling
	router.HandleFunc("/", testing).Methods("GET")
	router.HandleFunc("/Home", Home).Methods("GET")
	router.HandleFunc("/Contact", contact).Methods("GET")
	router.HandleFunc("/AddProject", Add).Methods("GET")
	router.HandleFunc("/Add-blog", AddBlog).Methods("POST")
	router.HandleFunc("/Delete/{id}", Deleteblog).Methods("GET")
	router.HandleFunc("/Blog-detail/{id}", blog).Methods("GET")

	http.ListenAndServe("localhost:3636", router)
}

// DEKLARASI FUNGSI
func testing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("It's Work"))

	// fmt.Println("Testing 1")
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Text/html; Charset=UTF-8")

	//PARSEFILES template index.html
	var tmpl, err = template.ParseFiles("index.html")

	//ERR HANDLING
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	resp := map[string]interface{}{
		"Data": Data,
		// "Blogs": Blogs,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, resp)
}

// ADD-BLOG
func AddBlog(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)

	}

	// Title := r.PostForm.Get("title")
	// Content := r.PostForm.Get("content")

	// var newBlog = Blog{
	// 	Title:     Title,
	// 	Post_Date: time.Now().String(),
	// 	Author:    "Jordi El Nino",
	// 	Content:   Content,
	// }

	// Blogs = append(Blogs, newBlog)
	http.Redirect(w, r, "/AddProject", http.StatusMovedPermanently)

}

func Deleteblog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Text/html; Charset=UTF-8")
	// id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// Blogs = append(Blogs[:id], Blogs[id+1:]...)
	http.Redirect(w, r, "/AddProject", http.StatusMovedPermanently)
}

// HANDLING FUNC WITH QUERY STR
func blog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; Charset=UTF-8")

	// id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// PARSEFILES template miniBlog
	var tmpl, err = template.ParseFiles("Blog-detail.html")

	// ERR HANDLING
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	blogDet := Blog{}
	// for i, data := range Blogs {
	// 	if i == id {
	// 		blogDet = Blog{
	// 			Title:     data.Title,
	// 			Post_Date: data.Post_Date,
	// 			Author:    data.Author,
	// 			Content:   data.Content,
	// 		}
	// 	}

	// }

	resp := map[string]interface{}{
		"Data": Data,
		"Blog": blogDet,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, resp)

}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;CHARSET=UTF-8")

	var tmpl, err = template.ParseFiles("contactMe.html")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpl.Execute(w, Data)
	w.WriteHeader(http.StatusOK)
	fmt.Println("Still work")
}

func Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;CHARSET-UTF8")

	var tmpl, err = template.ParseFiles("addP.html")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
	}
	//
	rows, _ := connection.Conn.Query(context.Background(), "SELECT id, title, image, content, post_date FROM public.tb_blog;")

	var result []Blog
	for rows.Next() {
		var each = Blog{}

		var err = rows.Scan(&each.Id, &each.Title, &each.Image, &each.Content, &each.Post_Date)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		each.Author = "Teguh Fauzi"
		each.Format_Date = each.Post_Date.Format("17 February 2000")

		result = append(result, each)
	}

	resp := map[string]interface{}{
		"Data":  Data,
		"Blogs": result,
	}

	tmpl.Execute(w, resp)
	w.WriteHeader(http.StatusOK)

}
