package interact

import (
        "fmt"
	"io/ioutil"
	"log"

        "golang.org/x/crypto/ssh"
)

var remoteUser, remoteHost string

// initial connection
func Init(){
     //connect to remote host
     connection, session := connect()
     // run main.go without any command line arguments
     ExecuteCommand("/usr/local/go/bin/go run main.go")
     defer connection.Close()
     defer session.Close()
}
func ExecuteCommand(cmd string) string{
        //connect to remote host
        connection, session := connect()
        // execute go program on remote host and get its combined standard output and standard error
        out, _ := session.CombinedOutput(cmd)
        defer connection.Close()
        defer session.Close()
	return string(out)
}
func connect() (*ssh.Client, *ssh.Session) {
        var port =  "22"
        if remoteUser == "" {
                fmt.Print("remoteUser: ")
                fmt.Scan(&remoteUser)
                fmt.Print("remoteHost: ")
                fmt.Scan(&remoteHost)
        }
        // get key
        key, err := ioutil.ReadFile("./ec2.pem")
	if err != nil {
 	   log.Fatalf("unable to read key: %v", err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
    	    log.Fatalf("unable to parse key: %v", err)
	}
	
        // configure authentication
        sshConfig := &ssh.ClientConfig{
                User:            remoteUser,
                Auth:            []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
        }

        // start a client connection to SSH server
        connection, err := ssh.Dial("tcp", remoteHost+":"+port, sshConfig)
        if err != nil {
                connection.Close()
                panic(err)
        }
        // create session
        session, err := connection.NewSession()
        if err != nil {
                session.Close()
                panic(err)
        }

        return connection, session
}
