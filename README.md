# Prohibition of commercial use(商用禁止)

# scraping search engine(スクレイピング検索エンジン)
I'm not sure in every way.There's no scraping going on, and I'm not taking responsibility for that.（私は性能に自信がありませんスクレイピングを行います責任を持ちません）

# What you need. (あなたに必要なもの)
mecab shared library for (linux or windows)((windows or linux)のmecab共有ライブラリ),
Go 1.15,
Mysql 8.0.21,
gcc 8.1.0,
Mecab 0.996 UTF-8 dynamiclibrary(動的ライブラリ),

Required package(golang)(必要パッケージ golang)
--- golang ---
```
go get golang.org/x/net/html
go get github.com/go-sql-driver/mysql
```
Install as described above.（上記のようにインストールしてください）
--- mysql ---
Create a database with the database name search（search名のデータベースを作ってください）|
Please make your table as follows.　（以下のようにテーブルを作ってくださいませ。)       　|
```
> DESCRIBE search;
+---------+---------------+------+-----+---------+-------+
| Field   | Type          | Null | Key | Default | Extra |
+---------+---------------+------+-----+---------+-------+
| title   | varchar(45)   | YES  |     | NULL    |       |
| url     | varchar(2045) | YES  |     | NULL    |       |
| setu    | varchar(300)  | YES  |     | NULL    |       |
| result1 | varchar(12)   | YES  |     | NULL    |       |
| result2 | varchar(12)   | YES  |     | NULL    |       |
| result3 | varchar(12)   | YES  |     | NULL    |       |
+---------+---------------+------+-----+---------+-------+
```
# How to use  (使い方)
1. Download the above file as a zip file and unzip it(上記のファイルをZIPファイルとしてダウンロードし、解凍してください。) OR  ``` git clone https://github.com/kiri139/scrapingsearchengine-golang.git ```
2. switch (mecab shared library) {
   case linux:
     Place "libmecab.so" in "materials\linux\" ( "materials/linux/" に "libmecab.so" を配置します。)
     Library linking of <stdlib.h> with gcc(gccで<stdlib.h>をライブラリリンク)
   
   case windows:
   Place "libmecab.dll" in "materials" ( "materials/" に "libmecab.dll" を配置します。)
   
}
2. ``` go run crawl.go ```
3. ``` go run main.go ```
4. ``` http://localhost:8080/home/ ``` Access to (にアクセス！)
That's it! (以上です！)

# Contact　(連絡先)
discord: kiri#3492
I'd love it if you could point out improvements to my programming code, etc.!(私のプログラミングコードの改善をなどを指摘してくれると嬉しいです！)
Anyone who wants to get involved with me, by all means!(私に関わりたい人是非とも！)

END... (終わり...)
