package daoimpl

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

type ViaSSHDialer struct {
	client *ssh.Client
}

func (self *ViaSSHDialer) Dial(addr string) (net.Conn, error) {
	return self.client.Dial("tcp", addr)
}

func connectaws() (*sql.DB, *ssh.Client) {

	sshHost := "ec2-54-218-55-72.us-west-2.compute.amazonaws.com" // SSH Server Hostname/IP
	sshPort := 22                                                 // SSH Port
	sshUser := "risabh"                                           // SSH Username
	sshPass := "rpqb123"                                          // Empty string for no password
	dbUser := "onlinetestuser"                                    // DB username
	dbPass := "Rpqb_123"                                          // DB Password
	dbHost := "127.0.0.1:3306"                                    // DB Hostname/IP
	dbName := "onlinetestdb"                                      // Database name

	var agentClient agent.Agent
	// Establish a connection to the local ssh-agent
	if conn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		defer conn.Close()

		// Create a new instance of the ssh agent
		agentClient = agent.NewClient(conn)
	}

	// The client configuration with configuration option to use the ssh-agent
	sshConfig := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{},
	}

	// When the agentClient connection succeeded, add them as AuthMethod
	if agentClient != nil {
		sshConfig.Auth = append(sshConfig.Auth, ssh.PublicKeysCallback(agentClient.Signers))
	}
	// When there's a non empty password add the password AuthMethod
	if sshPass != "" {
		sshConfig.Auth = append(sshConfig.Auth, ssh.PasswordCallback(func() (string, error) {
			return sshPass, nil
		}))
	}

	// Connect to the SSH Server
	if sshcon, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", sshHost, sshPort), sshConfig); err == nil {
		//	defer sshcon.Close()

		// Now we register the ViaSSHDialer with the ssh connection as a parameter
		mysql.RegisterDial("mysql+tcp", (&ViaSSHDialer{sshcon}).Dial)

		// And now we can use our new driver with the regular mysql connection string tunneled through the SSH connection
		if db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@mysql+tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)); err == nil {

			fmt.Printf("Successfully connected to the db\n")
			return db, sshcon
			db.Close()

		} else {

			fmt.Printf("Failed to connect to the db: %s\n", err.Error())
		}
		return nil, nil
	}
	return nil, nil

}

func connection() *sql.DB {
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/onlinetestdb")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
