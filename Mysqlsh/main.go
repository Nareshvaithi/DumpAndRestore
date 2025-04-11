package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main(){
	enableCmd := exec.Command("mysql", "-u", "root", "-p", "-e", "SET GLOBAL local_infile = 1;")
	_, err := enableCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to enable local_infile: %s", err)
	}

	
	fmt.Print("Press (B) Backup or Press (R) Restore: ")
	reader := bufio.NewReader(os.Stdin)
	BackupOrRestore,_ :=reader.ReadString('\n')
	BackupOrRestore = strings.TrimSpace(BackupOrRestore)



	if BackupOrRestore == "B"{
		fmt.Print("Enter user (default: root): ")
		user,err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("Enter password : ")
		password,err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("Enter host (default: localhost): ")
		host,err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("Enter db you want to backup: ")
		db,err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("Enter path to backup stores (default: home/mydbops/mysqlshBackup): ")
		path,err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		user = strings.TrimSpace(user)
		password = strings.TrimSpace(password)
		host = strings.TrimSpace(host)
		db = strings.TrimSpace(db)
		path = strings.TrimSpace(path)

		if user == "" {
			user = "root"
		}

		if host == "" {
			host = "localhost"
		}

		if path == "" {
			path = "/home/mydbops/mysqlshBackup"
		}

		executeBackup := fmt.Sprintf("util.dumpSchemas(['%s'],'%s',{ocimds:false})",db,path)
		cmd := exec.Command(
			"mysqlsh",
			"--js",
			"--user="+user,
			"--password="+password,
			"--host="+host,
			"--execute="+executeBackup,
		)

		output,err := cmd.CombinedOutput()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(output))

	}else if BackupOrRestore == "R" {

		fmt.Print("Enter user (default: root): ")
		user,err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("Enter password: ")
		password,err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("Enter host (default: localhost): ")
		host,err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("Enter db you want to restore: ")
		db,err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("Enter path to restore files stores (default: home/mydbops/mysqlshBackup): ")
		path,err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		user = strings.TrimSpace(user)
		password = strings.TrimSpace(password)
		host = strings.TrimSpace(host)
		db = strings.TrimSpace(db)
		path = strings.TrimSpace(path)

		if user == "" {
			user = "root"
		}

		if host == "" {
			host = "localhost"
		}

		if path == "" {
			path = "/home/mydbops/mysqlshBackup"
		}

		executeRestore := fmt.Sprintf("util.loadDump('%s', {schema: '%s'})",path,db)

		cmd := exec.Command(
			"mysqlsh",
			"--js",
			"--user="+user,
			"--password="+password,
			"--host="+host,
			"--execute="+executeRestore,
		)


		output,err := cmd.CombinedOutput()
		
		fmt.Println("restored: \n",string(output))

		if err != nil {
			log.Fatal(err)
		}


	}else {
		log.Fatal("B for Backup and R for Restore")
	}
}