package printer

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

// read more: http://www.termsys.demon.co.uk/vtansi.htm

// ScrGoto 光标移动到(x,y)
func ScrGoto(x, y int) {
	fmt.Print("\033[" + strconv.Itoa(x) + ";" + strconv.Itoa(y) + "H")
}

// ScrCleanLine 清除光标所在行
func ScrCleanLine() {
	fmt.Print("\033[2K\r")
}

// ScrClean 清除屏幕
//   i == 0 清除从光标到屏幕结尾
//   i == 1 清除光标到屏幕开始
//   i == 2 清除整个屏幕
func ScrClean(i int) {
	fmt.Print("\033[" + strconv.Itoa(i) + "J")
}

// ScrCleanEntire 清除整个屏幕并且移动光标到屏幕开始
func ScrCleanEntire() {
	ScrClean(2)
	ScrGoto(1, 1)
}

// ScrCursorUp 将光标向上移动lines行
func ScrCursorUp(lines int) {
	fmt.Print("\033[" + strconv.Itoa(lines) + "A")
}

// ScrCursorDown 将光标向下移动lines行
func ScrCursorDown(lines int) {
	fmt.Print("\033[" + strconv.Itoa(lines) + "B")
}

// ScrCursorLeft 将光标向左移动lines行
func ScrCursorLeft(cols int) {
	fmt.Print("\033[" + strconv.Itoa(cols) + "D")
}

// ScrCursorRight 将光标向右移动lines行
func ScrCursorRight(cols int) {
	fmt.Print("\033[" + strconv.Itoa(cols) + "C")
}

// ScrCursorSave 记录当前光标位置
func ScrCursorSave() {
	fmt.Print("\033[s")
}

// ScrCursorRestore 移动光标到上次存储的位置
func ScrCursorRestore() {
	fmt.Print("\033[u")
}

// ScrColor 输出颜色设置
// f 前景色 0-15
// b 背景色 0-7
func ScrColor(f int, b int) {
	var tmp string
	var colors = [8]int{0, 4, 2, 6, 1, 5, 3, 7}

	fg := colors[int(math.Mod(float64(f), 8))] + 30
	bg := colors[int(math.Mod(float64(b), 8))] + 40

	if b > 7 {
		tmp += "5;"
	} else {
		tmp += "0;"
	}

	if f > 7 {
		tmp += "1;"
	}

	fmt.Printf("\033[%s%d;%dm", tmp, fg, bg)
}

// ColorReset 重置颜色设置
func ScrColorReset() {
	fmt.Print("\033[0m")
}

// Wait 显示等待图表
func Wait(stop chan bool) {
	for {
		select {
		case <-stop:
			return
		default:
		}
		fmt.Print("|")
		time.Sleep(50 * time.Millisecond)
		fmt.Print("\010/")
		time.Sleep(50 * time.Millisecond)
		fmt.Print("\010-")
		time.Sleep(50 * time.Millisecond)
		fmt.Print("\010\\")
		time.Sleep(50 * time.Millisecond)
		fmt.Print("\010")
	}
}
