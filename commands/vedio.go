package commands

import "github.com/urfave/cli"

var videoCmd = cli.Command{
   Name:  "video",
   Usage: "下载视频",
    Action: func(c *cli.Context) error {
        url := c.Args().Get(0)
        output := c.Args().Get(1)
        downloadBilibili(url, output)
        return nil
    },
}

func init() {
   RootCmd.Commands = append(RootCmd.Commands, videoCmd)
}

func downloadBilibili(resourceUrl string, output string) {

}
