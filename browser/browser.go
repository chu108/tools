package browser

import (
	"os/exec"
	"runtime"
)

type Browser struct {
}

func NewBrowser() *Browser {
	return &Browser{}
}

//打开网页
func (*Browser) OpenUrl(url string) (err error) {
	switch runtime.GOOS {
	case "windows": //windows
		err = exec.Command(`cmd`, `/c`, `start`, url).Start()
	case "darwin": //Linux
		err = exec.Command(`open`, url).Start()
	default: //Mac
		err = exec.Command(`xdg-open`, url).Start()
	}
	return
}
