package main

import (
	"github.com/UnknwoonUser/Crypto/aes_impl/src/action"
	"github.com/urfave/cli/v2"
	"os"
	"sort"
)

func main() {
	var app = &cli.App{
		Name:                 "应用密码学实践-2019141440070-罗鉴",
		Usage:                "AES加密与解密",
		EnableBashCompletion: true,
		Commands: cli.Commands{
			{
				Name:   "encrypt",
				Usage:  "AES加密",
				Action: action.EncryptAction,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "mode",
						Usage:    "指定加密的操作模式，有ECB、CBC、CFB、OFB四种",
						Aliases:  []string{"m"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "plainfile",
						Usage:    "指定明文件的位置和名称",
						Aliases:  []string{"p"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "keyfile",
						Usage:    "指定密钥文件的位置和名称",
						Aliases:  []string{"k"},
						Required: true,
					},
					&cli.StringFlag{
						Name:    "vifile",
						Usage:   "指定初始化向量文件的位置和名称",
						Aliases: []string{"v"},
					},
					&cli.StringFlag{
						Name:     "cipherfile",
						Usage:    "指定密文文件的位置和名称",
						Aliases:  []string{"c"},
						Required: true,
					},
				},
			},
			{
				Name:   "decrypt",
				Usage:  "AES解密",
				Action: action.DecryptAction,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "mode",
						Usage:    "指定解密的操作模式，有ECB、CBC、CFB、OFB四种",
						Aliases:  []string{"m"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "plainfile",
						Usage:    "指定明文件的位置和名称",
						Aliases:  []string{"p"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "keyfile",
						Usage:    "指定密钥文件的位置和名称",
						Aliases:  []string{"k"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "vifile",
						Usage:    "指定初始化向量文件的位置和名称",
						Aliases:  []string{"v"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "cipherfile",
						Usage:    "指定密文文件的位置和名称",
						Aliases:  []string{"c"},
						Required: true,
					},
				},
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			_ = ctx.App.Command("help").Action(ctx)
			_ = action.AfterAction(ctx)
			return
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	app.Run(os.Args)

}
