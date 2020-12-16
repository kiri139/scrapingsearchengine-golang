package main

import (
  "log"
  "net/http"
  "html/template"
  "html"
  "strings"
  "database/sql"
  "fmt"

  _ "github.com/go-sql-driver/mysql"
)

var mysqlcheckset = strings.NewReplacer(`"`, `\"`, `'`,`\'`,`%`,`\%`,`_`,`\_`)//

type Information struct {
  StructTitle       string
  StructHost        string
  StructExplanation string
}

func CheckError(err error) {
  if err != nil {
	  log.Fatal(err)
  }
}

func main() {
  http.HandleFunc("/home", topfunc)
  http.HandleFunc("/home/search", resultfunc)
  http.ListenAndServe(":8080", nil)
}

func topfunc(w http.ResponseWriter, r *http.Request) {
  indexFile, _ := template.ParseFiles("index.html")
  indexFile.Execute(w, "index.html")
}

func resultfunc(w http.ResponseWriter, r *http.Request) {
  searchFile, _ := template.ParseFiles("search.html")
  getinp := r.FormValue("keyword")
  checkinp := html.EscapeString(getinp)
  myresult := mysqlopenfunc(getinp)
  htmlinsert := struct {
    Mese            string
    InformationSets []Information
  }{
    Mese: checkinp,
    InformationSets: myresult,
  }



  searchFile.ExecuteTemplate(w, "search.html", htmlinsert)
}

func mysqlopenfunc(getinp string)[]Information{
  mysqlspecialoks := mysqlcheckset.Replace(getinp)
  db, err := sql.Open("mysql", "root:xzAinagithub@tcp(127.0.0.1:3306)/search")
  CheckError(err)
  defer db.Close()
  //I could only come up with a slightly cumbersome way to write it (LOL
  mysqlstatementLIKE := fmt.Sprintf("SELECT * FROM search WHERE CONCAT(title, url, setu) LIKE'%%%s%%';",mysqlspecialoks)
  mysqlsearch, err := db.Query(mysqlstatementLIKE) //Cognitively, a mysql statement!
  CheckError(err)
  defer mysqlsearch.Close()
  var (
    dbtitle     string
    dburl       string
    dbsetu      string
    dbresult1   string
    dbresult2   string
    dbresult3   string
    slicestruct []Information
  )
  for mysqlsearch.Next(){
    err := mysqlsearch.Scan(&dbtitle, &dburl, &dbsetu, &dbresult1, &dbresult2, &dbresult3)
    CheckError(err)
    slicestruct = append(slicestruct, Information{
      StructTitle: html.EscapeString(dbtitle),
      StructHost:  html.EscapeString(dburl),
      StructExplanation: html.EscapeString(dbsetu),
    })
  }
  err = mysqlsearch.Err()
  CheckError(err)
  return slicestruct
}