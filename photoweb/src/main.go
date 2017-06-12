// web project main.go
package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func viewHandle(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if b := isExists(imagePath); !b {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func isExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	} /* else {
		return false
	}*/
	return os.IsExist(err) //错误更加详细

}

func listHandle(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir(UPLOAD_DIR)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//	var listHtml string
	locals := make(map[string]interface{})
	images := []string{}
	for _, file := range files {
		imageId := file.Name()
		//		listHtml += "<li><a href=\"/view?id=" + imageId + "\">" + imageId + "</a></li>"
		images = append(images, imageId)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8") //保证客户端显示为html
	//	io.WriteString(w, "<ol>"+listHtml+"</ol>")
	locals["images"] = images
	t, er := template.ParseFiles("html/list.html")
	if er != nil {
		http.Error(w, er.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, locals)
}

func uploadHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8") //保证客户端显示为html
		//		io.WriteString(w, "<form method=\"POST\" action=\"/upload\"  enctype=\"multipart/form-data\">Choose an image to upload: <input name=\"image\" type=\"file\" /><input type=\"submit\" value=\"Upload\" /></form>")
		t, err := template.ParseFiles("html/upload.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
		return
	}
	if r.Method == "POST" {
		f, h, e := r.FormFile("image")
		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		filename := h.Filename
		t, er := os.Create(UPLOAD_DIR + "/" + filename)
		if er != nil {
			http.Error(w, er.Error(), http.StatusInternalServerError)
			return
		}
		defer t.Close()
		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

const (
	UPLOAD_DIR = "uploads"
)

func main() {
	http.HandleFunc("/", listHandle)
	http.HandleFunc("/upload", uploadHandle)
	http.HandleFunc("/view", viewHandle)
	log.Println("启动....")
	log.Println(32 << 20)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("web服务启动失败", err.Error())
	}
	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()
}
