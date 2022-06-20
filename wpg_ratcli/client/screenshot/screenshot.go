package screenshot

import (
	"fmt"
	"image/png"
	"io/ioutil"
	"os"
	"os/user"
	"strings"

	"github.com/kbinani/screenshot"
)

func getTempDir() string {
	usr, _ := user.Current()
	app := usr.HomeDir + "\\AppData\\Local\\Temp\\"
	dir := strings.Replace(app, "\\", "/", -1)
	return dir
}
func GetScreen() []byte {
	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fileName := fmt.Sprintf("%stmp.log", getTempDir())
	fmt.Println(fileName)
	file, _ := os.Create(fileName)
	defer file.Close()
	png.Encode(file, img)
	filebyte, _ := ioutil.ReadFile(fileName)
	return filebyte
}
