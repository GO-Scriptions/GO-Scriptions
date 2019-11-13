package interact

import (
        "fmt"
	"io/ioutil"
	"log"

        "golang.org/x/crypto/ssh"
)

var remoteUser, remoteHost string

func ExecuteCommand(cmd string) []byte {
        //connect to remote host
        connection, session := connect()
        // execute go program on remote host and get its combined standard output and standard error
        out, _ := session.CombinedOutput(cmd)

        defer connection.Close()
        defer session.Close()
        return out
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
