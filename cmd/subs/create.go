package subs

import (
	"bufio"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"github.com/thewisepigeon/goo/models"
	"github.com/thewisepigeon/goo/pkg"
	"log"
	"os"
)

var CreateKeyCMD = &cobra.Command{
	Use: "create",
	Run: func(cmd *cobra.Command, args []string) {
		newKey := &models.Key{
			Key: pkg.GenerateRandomString(15),
		}
		err := newKey.Save()
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
		err = clipboard.WriteAll(newKey.Key)
		if err != nil {
			log.Println("Key created: ", newKey.Key)
			os.Exit(0)
		}
		log.Println("Key created successfuly and sent into your clipboard")
		os.Exit(0)
	},
}

var CreateActionCMD = &cobra.Command{
	Use: "create",
	Run: func(cmd *cobra.Command, args []string) {
		var name, workdir, command, recover_command string
		var err error
		interactive, _ := cmd.Flags().GetBool("i")
		if interactive {
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("> Enter the action's name")
			name, err = reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			pkg.TrimNewLineChar(&name)
			ok, message := pkg.IsValidName(name)
			if !ok {
				fmt.Println(message)
				os.Exit(0)
			}
			fmt.Println("> Enter the directory where the action should be run(use . for default)")
			workdir, err = reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			pkg.TrimNewLineChar(&workdir)
			ok = pkg.IsValidDir(workdir)
			if !ok {
				fmt.Println("Directory", workdir, "not found")
				os.Exit(0)
			}
			fmt.Println("> Enter the command to run (Use ';' to separate multiple commands)")
			command, err = reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			pkg.TrimNewLineChar(&command)
			fmt.Println("> Enter the command to run in case the previous one(s) fails (leave empty to do nothing)")
			recover_command, err = reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			pkg.TrimNewLineChar(&recover_command)
			newAction := &models.Action{
				Name:           name,
				WorkDir:        workdir,
				Command:        command,
				RecoverCommand: recover_command,
			}
			err = newAction.Save()
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			fmt.Println("New action created")
			os.Exit(0)
		}
		name, _ = cmd.Flags().GetString("name")
		ok, message := pkg.IsValidName(name)
		if !ok {
			fmt.Println(message)
			os.Exit(0)
		}
		workdir, _ = cmd.Flags().GetString("workdir")
		ok = pkg.IsValidDir(workdir)
		if !ok {
			fmt.Println("Directory", workdir, "not found")
			os.Exit(0)
		}
		command, _ = cmd.Flags().GetString("command")
		recover_command, _ = cmd.Flags().GetString("recover")
		newAction := &models.Action{
			Name:           name,
			WorkDir:        workdir,
			Command:        command,
			RecoverCommand: recover_command,
		}
		err = newAction.Save()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println("New action created")
		os.Exit(0)
	},
}

func init() {
	CreateActionCMD.Flags().Bool("i", false, "Create the command interactive mode or not")
	CreateActionCMD.Flags().String("name", "", "Name of the action to be created")
	CreateActionCMD.Flags().String("workdir", ".", "Directory where to run the action")
	CreateActionCMD.Flags().String("command", "", "Command to run")
	CreateActionCMD.Flags().String("recover", "", "Command to run if the action doe not exit with a 0 code")
}
