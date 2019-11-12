package main

import (
        "fmt"

        "golang.org/x/crypto/ssh"
)

var remoteUser, remoteHost, pw string

func main() {
        var out []byte
        out = executeCommand("/usr/local/go/bin/go run main.go")
        fmt.Println("output from remote go program:", string(out))
}
func executeCommand(cmd string) []byte {
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
                fmt.Print("password: ")
                fmt.Scan(&pw)
        }
        // configure authentication
        sshConfig := &ssh.ClientConfig{
                User:            remoteUser,
                Auth:            []ssh.AuthMethod{ssh.Password(pw)},
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

