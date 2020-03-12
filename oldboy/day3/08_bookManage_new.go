package main

/*
@Time : 2020-03-12 22:44
@Author : audiRStony
@File : 08_bookManage_new.go
@Software: GoLand
*/
import (
    "fmt"
    "os"
)

type book struct {
    title   string
    author  string
    price   float32
    publish bool
}

// 指定用户输出的信息
func userInputInfo() *book {
    var (
        title   string
        author  string
        price   float32
        publish bool
    )

    fmt.Println("请输出书籍信息")
    fmt.Print("书名:")
    fmt.Scan(&title)
    fmt.Print("作者:")
    fmt.Scan(&author)
    fmt.Print("价格:")
    fmt.Scan(&price)
    fmt.Print("是否上架[ture|false]")
    fmt.Scan(&publish)
    fmt.Println(title, author, price, publish)
    book := newBook(title, author, price, publish)
    return book

}

// book的指针切片，用来存放所有书籍
var allbooks = make([]*book, 0, 100)
// 定义新书的构造函数
func newBook(title, author string, price float32, publish bool) *book {
    return &book{
        title:   title,
        author:  author,
        price:   price,
        publish: publish,
    }
}

// 菜单栏
func showMenu() {
    fmt.Println("欢迎使用书籍管理系统BMS！")
    fmt.Println("1.添加书籍")
    fmt.Println("2.修改书籍信息")
    fmt.Println("3.展示书籍信息")
    fmt.Println("4.退出")
    fmt.Println()
}

// 添加书籍
func addBook() {
    book := userInputInfo()
    for _, b := range allbooks {
        if b.title == book.title {
            fmt.Printf("<%s> The book existed\n", b.title)
            return
        }
    }

    allbooks = append(allbooks, book) // 把添加的书信息存入切片
    fmt.Println("add book successfully")
}

// 修改书籍
func updateBook() {

    book := userInputInfo()
    for index, b := range allbooks {
        if b.title == book.title { // 判断书是否存在，不存在则直接返回
            allbooks[index] = book
            fmt.Printf("<%s> modified successful\n", b.title)
            return
        }
    }
    fmt.Printf("<%s> not existed\n", book.title)
}

// 展示信息
func showBooks() {

    if len(allbooks) == 0 { // 判断allbooks是否有书
        fmt.Println("穷B，啥也没有！")
        return
    }

    for _, b := range allbooks {
        fmt.Printf("<%s> 作者: %s,价格: %.2f,是否销售: %t\n", b.title, b.author, b.price, b.publish)
    }
}

func main() {
    for {
        showMenu()
        var option int
        fmt.Scanln(&option)
        switch option {
        case 1:
            addBook()
        case 2:
            updateBook()
        case 3:
            showBooks()
        case 4:
            fmt.Println("exit")
            os.Exit(0)
        }
    }
}
