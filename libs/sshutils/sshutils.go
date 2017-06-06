package sshutils

// This package includes utilities implemented by Golang.

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
)

func GetSshConfig(user string, password string) *ssh.ClientConfig {

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.KeyboardInteractive(func(user, instruction string, questions []string, echos []bool) ([]string, error) {
				// Just send the password back for all questions
				answers := make([]string, len(questions))
				for i, _ := range answers {
					answers[i] = password
				}
				return answers, nil
			}),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return config

}

func ExecSshCmd(cmd string, hostname string, port string, config *ssh.ClientConfig) (string, error) {

	var b bytes.Buffer

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", hostname, port), config)
	if err != nil {
		log.Fatal("Failed to dail: ", err)
	}

	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	session.Stdout = &b

	if err := session.Run(cmd); err != nil {
		log.Fatal("Failed to run: " + err.Error())
		return b.String(), err
	}

	o := b.String()
	log.Println(o)

	return o, nil
}

func ServiceOps(params map[string]string, name string, ops string) error {

	var cmd = "service"

	conf := GetSshConfig(params["user"], params["password"])

	if name != "" && ops != "" {
		cmd = fmt.Sprintf("%s %s %s", cmd, name, ops)
	} else {
		log.Fatal("Invalid args given to name and ops")
	}

	_, err := ExecSshCmd(cmd, params["hostname"], "22", conf)
	if err != nil {
		return err
	}

	return nil

}
