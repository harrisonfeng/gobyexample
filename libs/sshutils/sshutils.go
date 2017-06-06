package sshutils

// This package includes utilities implemented by Golang.
// @author Harrison Feng <feng.harrison@gmail.com

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
		// log.Fatal("Failed to run: " + err.Error())
		log.Fatal("Errors: " + b.String())
		return b.String(), err
	}

	o := b.String()

	return o, nil
}

func Service(params map[string]string, name string, action string, su bool) error {

	var cmd string

	conf := GetSshConfig(params["user"], params["password"])

	if su == true {
		// This requires the user has the same password as root.
		cmd = fmt.Sprintf("echo '%s' | su -c 'service %s %s' root",
			params["password"],
			name,
			action)
	} else {
		cmd = fmt.Sprintf("service %s %s", name, action)
	}

	_, err := ExecSshCmd(cmd, params["hostname"], "22", conf)
	if err != nil {
		return err
	}

	return nil

}

func DownloadWithCurl(params map[string]string, url string, targetDir string, su bool) error {

	var cmd string

	if su == true {
		// This requires the user has the same password as root.
		cmd = fmt.Sprint("echo '%s' | su -c 'cd %s && curl -O %s' root",
			params["password"],
			targetDir,
			url)
	} else {
		cmd = fmt.Sprint("cd %s && curl -O %s", targetDir, url)
	}

	conf := GetSshConfig(params["user"], params["password"])

	_, err := ExecSshCmd(cmd, params["hostname"], "22", conf)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil

}
