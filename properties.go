package main

import (
	"fmt"
	flags "github.com/jessevdk/go-flags"
	"github.com/magiconair/properties"
	"os"
)

var options struct {
	Get struct {
		File string `short:"f" long:"file" description:"Path to properties file" value-name:"FILE" required:"true"`
		Key  string `short:"k" long:"key" description:"Key" required:"true"`
	} `command:"get" description:"Reads property value by passed 'key' from specified file and prints it. If property is not found then empty string is printed"`

	Put struct {
		// 	File string `short:"f" long:"file" description:"Path to properties file" value-name:"FILE" required:"true"`
		// 	Key string `short:"k" long:"key" description:"Key" required:"true"`
		// 	Value string `short:"v" long:"value" description:"Value"`
	} `command:"put" description:"Not implemented yet (UNDER CONSTRUCTION)"`
}

func main() {
	parser := flags.NewParser(&options, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		flagsErr, ok := err.(*flags.Error)
		if ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else if ok && flagsErr.Type == flags.ErrCommandRequired {
			parser.WriteHelp(os.Stdout)
			os.Exit(1)
		} else {
			os.Exit(1)
		}
	}

	commandName := parser.Command.Active.Name
	if commandName == "get" {
		p := properties.MustLoadFile(options.Get.File, properties.UTF8)
		value := p.GetString(options.Get.Key, "")

		fmt.Println(value)
	} else {
		err := fmt.Errorf("command '%s' is not implemented yet", commandName)
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
