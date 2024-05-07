package app
import (
	
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicacao de linha de comando"
	app.Usage = "Busca IPs e Nomes de Servidores na internet"
	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPs na internet",
			Flags:[]cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				
				},
			},
			Action: buscarIps,
		},
		
	}
	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal()
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}