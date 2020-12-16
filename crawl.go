package main

import (
  "database/sql"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "net/url"
  "regexp"
  "strings"
  "time"

  "./materials"
  _ "./materials/linux"
  "golang.org/x/net/html"
  _ "github.com/go-sql-driver/mysql"
)

var (
  input         string
  recmp       = regexp.MustCompile("\\<([\\w\\W])+?\\>")
  recmpscript = regexp.MustCompile("\\<script(|[\\w\\W])+?>(|[\\w\\W])+?</script\\>")
  recmpstyle  = regexp.MustCompile("\\<style(|[\\w\\W])+?>(|[\\w\\W])+?</style\\>")
  inpdeletion = strings.NewReplacer(
    ",","", ".","", "?","", "_","",
    "<","", ">","", "/","", "\\","",
    "}","", "*","", "+","", "]","", ":","", ";","", "`","",
    "!","", "\"","", "#","", "$","", "%","", "@","",
    "&","", "'","", "(","", ")","", "=","", "~","", "|","",
    "^","", "-","", "{","", "　","", "\t","", "\n","",
  ) 
)

// net/http start
var (

  tsp = &http.Transport{
    DisableCompression: true,
    IdleConnTimeout:    80 * time.Second,
    DialContext:        http.DefaltTransport.DialContext,
    ForceAttemptHTTP2:  true,
    Proxy: http.ProxyFromEnvironment,
  }

  client = &http.Client{
    Transport: tsp,
    Timeout:   150 * time.Second,
    CheckRedirect: checkredirectfunc,
  }

)

func checkredirectfunc(req *http.Request, via []*http.Request) error{
  if len(via) >= 16 {
    return http.ErrUseLastResponse
  }
  return nil
}
// net/http end

type acquisitionlist struct{
  Titles  string
  Hosts   string
  Bodys   string
  Result1 string
  Result2 string
  Result3 string
}

func checkerror(errmese string, err error){
  if err != nil {
    fmt.Printf("%s\n", errmese)
    log.Fatal(err)
    main()
  }
}

func main(){
  fmt.Printf("クローリングしたいURLを入力してください\n")
  fmt.Printf("警告メッセージ\n")
  fmt.Println("何が起きても開発者である私にはは責任を持ちません")
  fmt.Scan(&input)
  geturlfunc(input)
}

func geturlfunc(url string){
  resp := getresponsefunc(url)
  defer resp.Body.Close()
  bodys, err := ioutil.ReadAll(resp.Body)
  checkerror("",err)

  ahtmltoken := html.NewTokenizer(strings.NewReader(string(bodys)))
  for {
    switch ahtmltoken.Next() {
    case html.ErrorToken:
      return
    case html.StartTagToken:
      t := ahtmltoken.Token()
      if t.Data == "a"{
        for _, v := range t.Attr {
          if v.Key == "href"{
            autocheckurlfunc(v.Val, url)
          }
        }
      }
    }
  }
}

func autocheckurlfunc(oksurl,urls string){
  u, err := url.Parse(oksurl)
  checkerror("", err)
  rel, err := u.Parse(urls)
  checkerror("", err)
  cse := getresponsefunc(rel.String())
  acquisitionfunc(cse, rel.String())  
}

func getresponsefunc(s string)(r *http.Response){
  time.Sleep(2 * time.Second)
  requt, err := http.NewRequest("GET", s, nil)
  checkerror("URLがありませんもしくはURL無効です",err)
  requt.Header.Add("User-Agent", "Super Mozilla/8.6 && Super Computers/15.6 && Quantum Computers/5.8 (Search Engine Bots)")
  resp, err := client.Do(requt)
  checkerror("", err)
  return resp
}

func acquisitionfunc(resp *http.Response, geturl string){
  defer resp.Body.Close()
  if resp.StatusCode == 404{
    return
  }
  bodys, err := ioutil.ReadAll(resp.Body)
  checkerror("なんか解析失敗した",err)

  htmltoken        := html.NewTokenizer(strings.NewReader(string(bodys)))
  acquisitiontitle := titlegetfunc(htmltoken)
  acquisitionbody  := documentfunc(htmltoken)

  deletebodys      := recmpscript.ReplaceAllString(string(bodys),"")
  deletebodys2     := inpdeletion.Replace(recmp.ReplaceAllString(deletebodys,""))
  rsl1,rsl2,rsl3   := materials.Hashtags(deletebodys2)
  
  databaselist := acquisitionlist{
    Titles:  acquisitiontitle,
    Hosts:   geturl,
    Bodys:   acquisitionbody,
    Result1: rsl1,
    Result2: rsl2,
    Result3: rsl3,
  }
  databaselist.databaseinsertstructfunc()
}

func titlegetfunc(ztitle *html.Tokenizer)(titilereturn string){
  for{
    tt := ztitle.Next()
    switch tt{
    case html.ErrorToken:
      return
    case html.StartTagToken:
      t := ztitle.Token()
      if t.Data == "title"{
        ztitle.Next()
        i := ztitle.Token()
        return i.Data
      }
    }
  }
}

func documentfunc(bodytoken *html.Tokenizer)(retubod string){
  for {
    tt := bodytoken.Next()
    switch tt {
    case html.ErrorToken:
      return
    case html.StartTagToken:
      t := bodytoken.Token()
      if t.Data == "meta"{
        for _, v := range t.Attr{
          if v.Key == "name" && v.Val == "description"{
            for _, v := range t.Attr{
              if v.Key == "content"{
                return v.Val
              }
            }
          }
        }
      }
      
      if t.Data == "p"{
        bodytoken.Next()
        i := bodytoken.Token()
        return i.Data
      }
    }
  }
}

func (accept acquisitionlist)databaseinsertstructfunc(){
  db, err := sql.Open("mysql", "root:xzAinagithub@tcp(127.0.0.1:3306)/search")
  checkerror("mysql起動失敗しました",err)
  defer db.Close()
  informationinsert, err := db.Prepare("INSERT INTO search(title,url,setu,result1,result2,result3) VALUES(?,?,?,?,?,?)")
  checkerror("完璧な列がありません",err)
  informationinsert.Exec(accept.Titles, accept.Hosts, accept.Bodys, accept.Result1, accept.Result2, accept.Result3)
}