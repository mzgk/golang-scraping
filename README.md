# Golangでスクレイピング

## 参考
- 2019.4.14 技術書典6
- 株式会社メディアドゥ Tech Do Book vol.01
- 第3章 Goによるはじめてのスクレイピング

## 改修
- goqueryの使用方法を最新化
  - NewDocument(Deprecate) -> NewDocumentFromReader
  - https://github.com/PuerkitoBio/goqueryを参照

## 実行方法
```bash
# Webサーバの起動（スクレイピング用サイト）
$ cd techdo-book_web
$ go run main.go

# スクレイピング
$ cd ../
$ go run scraping.go
```
