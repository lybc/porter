package commands

import (
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/urfave/cli"
	"porter/utils"
)

var qiniuCmd = cli.Command{
	Name:  "qiniu",
	Usage: "下载七牛云对象存储中的文件",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "ak",
			Usage:    "accessKey",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "sk",
			Usage:    "secretKey",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "bucket",
			Usage:    "bucket",
			Required: true,
		},
		&cli.StringFlag{
			Name:  "output",
			Usage: "output",
			Value: "./",
		},
		&cli.IntFlag{
			Name:  "concurrent",
			Usage: "同时下载数量",
			Value: 8,
		},
	},
	Action: download,
}

func init() {
	RootCmd.Commands = append(RootCmd.Commands, qiniuCmd)
}

func download(ctx *cli.Context) {
	ak := ctx.String("ak")
	sk := ctx.String("sk")
	bucket := ctx.String("bucket")
	output := ctx.String("output")
	concurrent := ctx.Int("concurrent")

	mac := qbox.NewMac(ak, sk)
	cfg := storage.Config{
		UseHTTPS: false,
	}
	limit := 1000
	marker := ""
	bucketManager := storage.NewBucketManager(mac, &cfg)
	downloader := utils.NewDownloader(output)
	downloader.Concurrent = concurrent
	for {
		entries, _, nextMarker, hasNext, err := bucketManager.ListFiles(bucket, "", "", marker, limit)
		if err != nil {
			panic(err)
		}
		domain, err := bucketManager.ListBucketDomains(bucket)
		if err != nil {
			panic(err)
		}
		for _, entry := range entries {
			url := storage.MakePublicURL(domain[0].Domain, entry.Key)
			downloader.AppendResource(entry.Key, "http://"+url)

		}
		if hasNext {
			marker = nextMarker
		} else {
			break
		}
	}
	downloader.Start()
}
