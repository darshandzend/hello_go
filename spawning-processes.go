package main

import "fmt"
import "io/ioutil"
import "os/exec"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	dateCmd := exec.Command("date")

	dateOp, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOp))

	grepCmd := exec.Command("grep", "hello")

	grepIn, err := grepCmd.StdinPipe()
	check(err)
	grepOut, err := grepCmd.StdoutPipe()
	check(err)

	grepCmd.Start()

	grepIn.Write([]byte("hello grep\nGello grep\nchello grep"))
	grepIn.Close()

	grepOp, err := ioutil.ReadAll(grepOut)
	check(err)

	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepOp))

	lsCommand := exec.Command("bash", "-c", "ls -alhtS")
	lsOut, err := lsCommand.Output()
	check(err)
	fmt.Println("> ls -alhtS")
	fmt.Println(string(lsOut))

}
