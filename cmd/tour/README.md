# tour

字串轉換工具以及時間處理工具

## 主要指令邏輯

```golang
package commands

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
}

```

## word 指令邏輯

```golang
package commands

import (
	"log"
	"strings"

	"github.com/leetcode-golang-classroom/golang-cli-sample/internal/word"
	"github.com/spf13/cobra"
)

var str string
var mode int8
var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "單字格式轉換",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToLower(str)
		case ModeLower:
			content = word.ToUpper(str)
		case ModeUnderscopeToLowerCamelCase:
			content = word.UnderscopeToLowerCamelCase(str)
		case ModeUnderscopeToUpperCamelCase:
			content = word.UnderscopeToUpperCamelCase(str)
		case ModeCamelCaseToUnderscope:
			content = word.CamelCaseToUnderscope(str)
		default:
			log.Fatalf("暫不支援該轉換模式，請執行 help word 檢視說明文件")
		}
		log.Printf("輸出結果: %s\n", content)
	},
}

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscopeToUpperCamelCase
	ModeUnderscopeToLowerCamelCase
	ModeCamelCaseToUnderscope
)

var desc = strings.Join(
	[]string{
		"該子指令支援各種單字格式轉換，模式如下",
		"1:全部單字轉為大寫",
		"2:全部單字轉為小寫",
		"3:底線單字轉為大寫駝峰單字",
		"4:底線單字轉為小寫駝峰單字",
		"5:駝峰單字轉為底線單字",
	}, "\n",
)

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "請輸入單字內容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "請輸入單字轉換的模式")
}
```

## 時間主要指令

```golang
package commands

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/leetcode-golang-classroom/golang-cli-sample/internal/timer"
	"github.com/spf13/cobra"
)

var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "時間格式處理",
	Long:  "時間格式處理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "取得現在時間",
	Long:  "取得現在時間",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("輸出結果: %s, %d\n", nowTime.Format(time.RFC3339), nowTime.Unix())
	},
}

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "計算所需時間",
	Long:  "計算所需時間",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			space := strings.Count(calculateTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			}
			if space == 1 {
				layout = "2006-01-02 15:04:05"
			}
			currentTimer, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, err := strconv.Atoi(calculateTime)
				if err != nil {
					log.Fatalf("strconv calculateTime err: %v", err)
				}
				currentTimer = time.Unix(int64(t), 0)
			}
		}
		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timeer.GetCalculateTime err: %v", err)
		}
		log.Printf("輸出結果: %s, %d\n", t.Format(layout), t.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)
	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", `需要計算的時間，有效時間單位為時間戳記或是格式化之後的時間`)
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "0s", `持續時間，有效時間單位為"ns","us","ms", "s", "m", "h"`)
}
```

特別可以注意到的是 timeCmd 包含著 nowTimeCmd 與　calculateTimeCmd　兩個子指令。這兩個子指令必須要在 timeCmd 作 AddCommand 才能夠生效。

## 時區問題

在初始化 Local 時，預設會從　/etc/localtime 取出時區。
假設想要指定時區，可以特別使用　LoadLocation 讀出時區格式，然後設定到　timer 上
特別注意的是想要　Parse 時間字串為某個時區的格式，必須使用 ParseInLocation 帶入設定的時區
