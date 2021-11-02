package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
	"timeChangeTools/internal"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

var nowTimeCmd = &cobra.Command{
	Use:  "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := internal.GetnowTime()
		log.Printf("输出结果: %s,%d",nowTime.Format("2006-01-02 15:04:05"),nowTime.Unix())
	},
}

var calculationTime string
var duration string

var calcualteTimeCmd = &cobra.Command{
	Use: "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if calculationTime == ""{
			currentTimer = internal.GetnowTime()
		}else{
			var err error
			space := strings.Count(calculationTime,":")
			if space == 0 {
				layout = "2006-01-02"
			}
			if space == 1 {
				layout = "2006-01-02 15:04"
			}
			currentTimer,err = time.Parse(layout,calculationTime)
			if err != nil {
				t,_ := strconv.Atoi(calculationTime)
				currentTimer = time.Unix(int64(t),0)
			}
			t, err := internal.GetCalculateTime(currentTimer, duration)
			if err != nil {
				log.Fatalf("timer.GetCalculateTime err: %v", err)
			}
			log.Printf("输出结果: %s, %d", t.Format(layout), t.Unix())

		}
	},
}

func init(){
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calcualteTimeCmd)

	calcualteTimeCmd.Flags().StringVarP(&calculationTime,"calculate","c","","需要单位时间戳或者已经格式化后的时间")
	calcualteTimeCmd.Flags().StringVarP(&duration,"duration","d","",`持续时间，有效时间单位为"ns", "us" (or "µs"), "ms", "s", "m", "h`)
}