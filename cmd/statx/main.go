package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/mdp/go-statx"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "apikey",
			Value:  "",
			Usage:  "API key for Statx",
			EnvVar: "API_KEY",
		},
		cli.StringFlag{
			Name:   "authtoken",
			Value:  "",
			Usage:  "Auth Token for Statx",
			EnvVar: "AUTH_TOKEN",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "login",
			Usage: "login to StatX",
			Action: func(c *cli.Context) error {
				return login(c.Args().First(), c.String("clientname"))
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "clientname",
					Value: "mdp/go-statx",
					Usage: "the client name",
				},
			},
		},
		{
			Name:  "list",
			Usage: "list a resource",
			Action: func(c *cli.Context) error {
				return list(c)
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "group",
					Value: "",
					Usage: "the group id of the resource",
				},
				cli.StringFlag{
					Name:  "stat",
					Value: "",
					Usage: "the stat id of the resource",
				},
			},
		},
		{
			Name:  "update",
			Usage: "update a stat",
			Action: func(c *cli.Context) error {
				return update(c)
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "group",
					Value: "",
					Usage: "the group id of the resource",
				},
				cli.StringFlag{
					Name:  "stat",
					Value: "",
					Usage: "the stat id of the resource",
				},
				cli.StringFlag{
					Name:  "value",
					Value: "",
					Usage: "the value to update",
				},
			},
		},
	}

	app.Run(os.Args)
}

func login(phoneNumber, clientName string) error {
	if len(phoneNumber) <= 0 {
		return cli.NewExitError("must specify a phone number", 1)
	}
	client := statx.NewClient(nil)
	authResponse, _, err := client.Auth.Login(phoneNumber, clientName)
	if err != nil {
		fmt.Printf("Error: %+v", err)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter verification code: ")
	code, _ := reader.ReadString('\n')
	code = strings.TrimSpace(code)
	credentials, _, err := client.Auth.Verify(code, authResponse)
	if err != nil {
		fmt.Printf("Error: %+v", err)
		return err
	}
	fmt.Printf("Credentials\nAPIKey:%s\nAuthToken:%s\n", *credentials.APIKey, *credentials.AuthToken)
	return nil
}

func print(j interface{}) {
	json, err := json.MarshalIndent(j, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(json))
}

func update(c *cli.Context) error {
	client := statx.NewAuthenticatedClient(nil, c.GlobalString("apikey"), c.GlobalString("authtoken"))
	stat := &statx.Stat{Value: c.String("value")}
	updated, _, err := client.Stats.Update(c.String("group"), c.String("stat"), stat)
	print(updated)
	return err
}

func list(c *cli.Context) error {
	client := statx.NewAuthenticatedClient(nil, c.GlobalString("apikey"), c.GlobalString("authtoken"))
	if len(c.String("group")) > 0 {
		if len(c.String("stat")) > 0 {
			stat, _, err := client.Stats.Get(c.String("group"), c.String("stat"))
			print(stat)
			return err
		}
		statList, _, err := client.Stats.List(c.String("group"))
		print(statList)
		return err
	}
	groupList, _, err := client.Groups.List()
	print(groupList)
	return err
}
