package _sys

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/chu108/tools/_str"
	"os/exec"
)

func getCmd(commName string, arg ...string) *exec.Cmd {
	cmdPath, err := exec.LookPath(commName)
	if err != nil {
		panic(err)
	}
	return exec.Command(cmdPath, arg...)
}

//执行命令并返回结果
func Exec(commName string, arg ...string) (string, error) {
	output, err := getCmd(commName, arg...).CombinedOutput()
	outputStr := _str.NewBytes().ToStr(output)

	return outputStr, err
}

//执行命令并直接输出结果
func ExecPipe(commName string, arg ...string) error {
	// 命令的错误输出和标准输出都连接到同一个管道
	cmd := getCmd(commName, arg...)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Println(string(tmp))
		if err != nil {
			break
		}
	}
	return cmd.Wait()
}

func ExecGrep(commName string, arg ...string) (string, error) {
	var out, stderr bytes.Buffer
	cmd := getCmd(commName, arg...)
	cmd.Stderr = &stderr
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return out.String(), errors.New(err.Error() + ":" + stderr.String())
	}
	return out.String(), nil
}

//解析命令行字符串参数
func FlagString(name, value, usage string) string {
	val := flag.String(name, value, usage)
	flag.Parse()
	return *val
}

//解析命令行int参数
func FlagInt(name string, value int, usage string) int {
	val := flag.Int(name, value, usage)
	flag.Parse()
	return *val
}

//解析命令行int64参数
func FlagInt64(name string, value int64, usage string) int64 {
	val := flag.Int64(name, value, usage)
	flag.Parse()
	return *val
}

//解析命令行int64参数
func FlagFloat64(name string, value int64, usage string) float64 {
	val := flag.Float64(name, 0, usage)
	flag.Parse()
	return *val
}

//解析命令行bool参数
func FlagBool(name string, value bool, usage string) bool {
	val := flag.Bool(name, value, usage)
	flag.Parse()
	return *val
}
