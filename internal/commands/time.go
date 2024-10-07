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
			location, _ := time.LoadLocation("Asia/Taipei")
			var err error
			space := strings.Count(calculateTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			}
			if space == 1 {
				layout = "2006-01-02 15:04:05"
			}
			currentTimer, err = time.ParseInLocation(layout, calculateTime, location)
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
