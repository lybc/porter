package main

import (
    "fmt"
    "github.com/bmaupin/go-epub"
    "github.com/gocolly/colly"
    "github.com/urfave/cli"
    "golang.org/x/text/encoding/simplifiedchinese"
    "log"
    "net"
    "net/http"
    "os"
    "regexp"
    "sort"
    "strconv"
    "strings"
    "time"
)

type Chapter struct {
    Title string
    Body  string
    index int
}

type Ebook struct {
    Name       string
    Cover      string
    Author     string
    Status     string
    LastUpdate string
    Desc       string
    Chapters   []Chapter
}

func ChineseFormat(raw string) string {
    utf8string, _ := simplifiedchinese.GBK.NewDecoder().String(raw)
    return strings.ReplaceAll(utf8string, "聽", "")
}

func (book *Ebook) Len() int {
    return len(book.Chapters)
}

func (book *Ebook) Swap(i, j int) {
    book.Chapters[i], book.Chapters[j] = book.Chapters[j], book.Chapters[i]
}

func (book *Ebook) Less(i, j int) bool {
    return book.Chapters[i].index < book.Chapters[j].index
}

func (book *Ebook) PrintMeta() {
    fmt.Println(book.Name)
    fmt.Println(book.Author)
    fmt.Println(book.Status)
    fmt.Println(book.LastUpdate)
    fmt.Println(book.Desc)
}

func (book *Ebook) Crawl(url string) {
    c := colly.NewCollector(
        colly.MaxDepth(2),
        colly.Async(true),
    )

    c.WithTransport(&http.Transport{
        Proxy: http.ProxyFromEnvironment,
        DialContext: (&net.Dialer{
            Timeout:   30 * time.Second,
            KeepAlive: 30 * time.Second,
        }).DialContext,
        MaxIdleConns:          100,
        IdleConnTimeout:       90 * time.Second,
        TLSHandshakeTimeout:   10 * time.Second,
        ExpectContinueTimeout: 1 * time.Second,
    })

    c.Limit(&colly.LimitRule{
        DomainGlob:  "*www.zwdu.com/*",
        Parallelism: 5,
        RandomDelay: 2 * time.Second,
    })

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL)
    })

    c.OnError(func(_ *colly.Response, err error) {
        log.Println("Something went wrong:", err)
    })

    c.OnHTML("#maininfo", func(e *colly.HTMLElement) {
        book.Name = ChineseFormat(e.DOM.Find("h1").Text())
        book.Cover, _ = e.DOM.Find("#fmimg > img").Attr("src")
        book.Author = ChineseFormat(e.DOM.Find("#info > p:nth-child(2)").Text())
        book.Status = ChineseFormat(e.DOM.Find("#info > p:nth-child(3)").Text())
        book.LastUpdate = ChineseFormat(e.DOM.Find("#info > p:nth-child(4)").Text())
        book.Desc = ChineseFormat(e.DOM.Find("#intro > p:nth-child(1)").Text())
    })

    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        //e.Request.Visit(e.Attr("href"))
        href := e.Attr("href")
        if match, _ := regexp.Match(`/book/\d+/\d+.html`, []byte(href)); match {
            fmt.Printf("开始抓取: %s\n", href)
            fullUrl := e.Request.AbsoluteURL(href)
            c.Visit(fullUrl)
        }
    })

    c.OnHTML("div.content_read", func(e *colly.HTMLElement) {
        var chapter Chapter
        chapter.Title = ChineseFormat(e.DOM.Find("h1").Text())
        rawBody, _ := e.DOM.Find("#content").Html()
        chapter.Body = ChineseFormat(rawBody)
        urlSuffix := regexp.MustCompile(`\d+.html`).Find([]byte(e.Request.URL.Path))
        chapter.index, _ = strconv.Atoi(strings.ReplaceAll(string(urlSuffix), ".html", ""))
        book.Chapters = append(book.Chapters, chapter)
    })

    c.Visit(url)
    c.Wait()
}

func (book *Ebook) CreateEpub(path string) {
    sort.Sort(book)
    e := epub.NewEpub(book.Name)
    e.SetAuthor(book.Author)
    e.SetDescription(book.Desc)
    e.SetTitle(book.Name)
    e.SetCover(book.Cover, "")
    for _, element := range book.Chapters {
        body := "<h1>" + element.Title + "</h1>" + "<br/><br/>" + element.Body
        //fmt.Println(body)
        e.AddSection(body, element.Title, "", "")
    }
    filename := path + "/" + book.Name + ".epub"
    e.Write(filename)
}

func main() {
    app := &cli.App{
        Name:  "zwdu",
        Usage: "抓取八一中文网小说转换成epub",
        Flags: []cli.Flag{
            cli.StringFlag{
                Name:     "start",
                Usage:    "抓取起始页面",
                Required: true,
            },
            cli.StringFlag{
                Name:     "output",
                Usage:    "文件输出路径",
                Required: true,
            },
        },
        Action: func(c *cli.Context) error {
            output := c.String("output")
            start := c.String("start")

            ebook := &Ebook{}
            ebook.Crawl(start)
            ebook.PrintMeta()
            ebook.CreateEpub(output)
            return nil
        },
    }

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}
