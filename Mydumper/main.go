package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Backup (B) or Restore (R): ")
	MydumperOrMyloader, _ := reader.ReadString('\n')
	MydumperOrMyloader = strings.TrimSpace(MydumperOrMyloader)

	if MydumperOrMyloader == "" {
		log.Fatal("Please specify Mydumper or Myloader")
	}

	if MydumperOrMyloader == "B" {
		fmt.Print("Enter the hostname (default: localhost): ")
		host, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("Enter the user (default: root): ")
		user, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("Enter the password: ")
		pass, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("Enter the database you want to backup: ")
		db, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("Specify the path to store backups (default: /home/mydbops/storeBackups): ")
		outputDir, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("Enter the verbose level (default: 3): ")
		verbose, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}

		host = strings.TrimSpace(host)
		user = strings.TrimSpace(user)
		pass = strings.TrimSpace(pass)
		db = strings.TrimSpace(db)
		outputDir = strings.TrimSpace(outputDir)
		verbose = strings.TrimSpace(verbose)

		if host == "" {
			host = "localhost"
		}
		if user == "" {
			user = "root"
		}
		if outputDir == "" {
			outputDir = "/home/mydbops/storeBackups"
		}
		if verbose == "" {
			verbose = "3"
		}

		cmd := exec.Command(
			"mydumper",
			"--host="+host,
			"--user="+user,
			"--password="+pass,
			"--database="+db,
			"--outputdir="+outputDir,
			"--verbose="+verbose,
			"--compress",
		)

		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal("Backup error: ", err)
		}
		fmt.Println("Backup Result:\n", string(output))

		


	} else if MydumperOrMyloader == "R" {
		fmt.Print("Enter the hostname (default: localhost): ")
		host, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("Enter the user (default: root): ")
		user, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("Enter the password: ")
		pass, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("Enter the database you want to restore: ")
		db, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("Specify the path to restore (default: /home/mydbops/storeBackups): ")
		restorePath, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("Enter the verbose level (default: 3): ")
		verbose, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}

		host = strings.TrimSpace(host)
		user = strings.TrimSpace(user)
		pass = strings.TrimSpace(pass)
		db = strings.TrimSpace(db)
		restorePath = strings.TrimSpace(restorePath)
		verbose = strings.TrimSpace(verbose)

		if host == "" {
			host = "localhost"
		}
		if user == "" {
			user = "root"
		}
		if restorePath == "" {
			restorePath = "/home/mydbops/storeBackups"
		}
		if verbose == "" {
			verbose = "3"
		}

		fmt.Println("Restoring with myloader...")
		fmt.Println("Host:", host)
		fmt.Println("User:", user)
		fmt.Println("DB:", db)
		fmt.Println("Path:", restorePath)
		fmt.Println("Verbose:", verbose)

		cmd := exec.Command(
			"myloader",
			"--host="+host,
			"--user="+user,
			"--password="+pass,
			"--database="+db,
			"--directory="+restorePath,
			"--verbose="+verbose,
		)

		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Restore Result:\n", string(output))

	} else {
		log.Fatal("Invalid command!!!")
	}
}
