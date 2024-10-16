package commands

import (
	"log"

	"github.com/leetcode-golang-classroom/golang-cli-sample/internal/sql2struct"
	"github.com/spf13/cobra"
)

var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql 轉換和處理",
	Long:  "sql 轉換和處理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql 轉換",
	Long:  "sql 轉換",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			Password: password,
			Charset:  charset,
		}
		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel connect err: %v", err)
		}
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v", err)
		}
		template := sql2struct.NewStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("template.Generate err: %v", err)
		}
	},
}

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "", "", "請輸入資料庫帳號")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "", "請輸入資料庫密碼")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "請輸入資料庫HOST")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "請輸入資料庫編碼")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "請輸入資料庫類型")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "utf8mb4", "請輸入資料庫名稱")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "", "", "請輸入資料表名稱")
}
