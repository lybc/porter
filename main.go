package main

import (
    "github.com/gocolly/colly"
    "github.com/urfave/cli"
    "log"
    "os"
    "time"
    "zwdu-doawloader-go/drivers"
)

func main() {
    app := &cli.App{
        Name:   "porter",
        Usage:  "我不生产资源，我只是资源的搬运工",
        Author: "lybc",
        Email:  "yibocheng.li@gmail.com",
        Commands: []cli.Command{
            {
                Name:  "ebook",
                Usage: "下载小说",
                Flags: []cli.Flag{
                    cli.IntFlag{
                        Name:     "concurrent",
                        Usage:    "并发数量",
                        Required: false,
                        Value:    2,
                    },
                    cli.IntFlag{
                        Name:     "delay",
                        Usage:    "随机延迟秒数",
                        Required: false,
                        Value:    2,
                    },
                },
                Action: func(c *cli.Context) error {
                    url := c.Args().Get(0)
                    output := c.Args().Get(1)
                    concurrent := c.Int("concurrent")
                    delay := c.Int("delay")

                    ebook := drivers.NewEbook()
                    ebook.Collector.Limit(&colly.LimitRule{
                        DomainGlob:  "*www.zwdu.com/*",
                        Parallelism: concurrent,
                        RandomDelay: time.Duration(delay) * time.Second,
                    })

                    ebook.Download(url, output)
                    return nil
                },
            },
            {
                Name: "music",
                Usage: "下载音乐",
                Action: func(c *cli.Context) error {
                    url := c.Args().Get(0)
                    output := c.Args().Get(1)
                    drivers.NewNeteaseMusic(url).Download(url, output)
                    return nil
                },
            },
        },
    }

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}
