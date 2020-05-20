package commands

import (
    "github.com/urfave/cli"
)

var RootCmd = &cli.App{
    Name:   "porter",
    Usage:  "我们不生产资源，我只是资源的搬运工",
    Author: "lybc",
    Email:  "yibocheng.li@gmail.com",
}