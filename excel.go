package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"time"
)

func main() {
	dt1 := time.Now()
	excel, err1 := xlsx.OpenFile("address.xlsx")

	if err1 != nil {
		fmt.Printf(err1.Error())
	}

	sheet1 := excel.Sheets[0]
	//10000行目のデータを表示
	fmt.Println(sheet1.Rows[3500].Cells[0].Value + " : " + sheet1.Rows[3500].Cells[1].Value + " : " + sheet1.Rows[3500].Cells[2].Value)

	dt2 := time.Now()
	fmt.Println(dt2.Sub(dt1).Seconds()) //Excelファイルの読込時間
}