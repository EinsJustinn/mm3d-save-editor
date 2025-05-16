package main

import (
	"flag"
	"fmt"
	"mm3d-save-editor/save"
	"os"
)

func main() {
	file := flag.String("file", "", "Path to the save file")
	money := flag.Int("money", -1, "Set the money amount (use -money)")
	bank := flag.Int("bank", -1, "Set the bank account balance (use -bank)")
	name := flag.String("name", "", "Change the player name (use --name)")

	flag.Parse()

	if *file == "" {
		fmt.Println("Error: Please specify a save file using --file.")
		os.Exit(1)
	}

	moneySet := *money != -1
	bankSet := *bank != -1
	nameSet := *name != ""

	if !moneySet && !bankSet && !nameSet {
		fmt.Println("Error: You must provide at least one of the following flags: --money, --bank, or --name.")
		os.Exit(1)
	}

	saveFile, err := save.Load(*file)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Save file: %s\n", *file)
	if moneySet {
		err := saveFile.SetMoney(*money)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
	}
	if bankSet {
		err := saveFile.SetBank(*bank)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
	}
	if nameSet {
		err := saveFile.SetName(*name)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
	}

	err = saveFile.SaveToFile(*file)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Save file updated")
}
