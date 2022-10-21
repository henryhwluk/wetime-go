package main

import (
	// "wetime-go/config"
	// "wetime-go/db"
	// . "wetime-go/router"
	"encoding/json"
	"fmt"

	//"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/nguyenthenguyen/docx"
	"github.com/xuri/excelize/v2"
)

func main() {

	// fmt.Println("hello world")

	// r := gin.Default()
	// r.GET("/hello", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"msg": "hello"})
	// })
	// //r.Run(":8000")

	// 配置文件
	//config.Process()

	// 链接数据库
	//db.Init()

	// 初始化路由
	//router := InitRouter()
	//router.Run(":" + config.Conf.HttpPort)

	type Building struct {
		BuildingNum1      string
		BuildingNum2      string
		BuildingNum3      string
		BuildingDeveloper string

		Buyer      string
		Address    string
		BuyerPhone string

		AgreedArea      string
		AgreedUnitPrice string
		AgreedAllPrice  string

		ActualArea string

		ErrorSize string
		ErrorArea string

		ReceiveRefund string
		ReceiveAmt    string
		FinalAllPrice string
	}

	f, err := excelize.OpenFile("交房明细表.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// cell, err := f.GetCellValue("Sheet1", "B2")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	var results []*Building

	for _, row := range rows {
		for index, colCell := range row {
			// if index <= 1 {
			// 	//fmt.Print("header不打印")
			// 	continue
			// }
			var b = &Building{}
			switch index {
			case 0:
				b.BuildingNum1 = colCell
			case 1:
				b.BuildingNum2 = colCell
			case 2:
				b.BuildingNum3 = colCell
			case 3:
				b.BuildingDeveloper = colCell
			case 4:
				b.Buyer = colCell
			case 5:
				b.Address = colCell
			case 6:
				b.BuyerPhone = colCell
			case 7:
				b.AgreedArea = colCell
			case 8:
				b.AgreedUnitPrice = colCell
			case 9:
				b.AgreedAllPrice = colCell
			case 10:
				b.ActualArea = colCell
			case 11:
				b.ErrorSize = colCell
			case 12:
				b.ErrorArea = colCell
			case 13:
				b.ReceiveRefund = colCell
			case 14:
				b.ReceiveAmt = colCell
			case 15:
				b.FinalAllPrice = colCell
			}

			results = append(results, b)

			//fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}

	json, err := json.Marshal(results)
	fmt.Print("results:", string(json))

	fmt.Println("hello world")
	return
	r, err := docx.ReadDocxFile("/Users/henryluk/Documents/gowork/src/wetime-go/old_template.docx")
	if err != nil {
		fmt.Print("sssssssss")
		panic(err)

	}
	docx1 := r.Editable()
	docx1.Replace("ND1ouQ", "new_1_1ccccccccc", -1)
	docx1.Replace("pxxsI3", "new_2_1bbbbbbbbbbb", -1)

	docx1.Replace("vOn7VR", "new31", -1)
	docx1.Replace("e6L14o", "new32", -1)

	docx1.Replace("ETWzK0", "new_4_1", -1)

	docx1.Replace("1tQFB6", "new_5_1", -1)

	docx1.Replace("rHNDIC", "new60", -1)
	docx1.Replace("tEteW4", "n61", -1)
	docx1.Replace("i5ZiRc", "n62", -1)
	docx1.Replace("PlTN5y", "n63", -1)
	docx1.Replace("6lcTHR", "n64", -1)
	docx1.Replace("LvcjjB", "n65", -1)

	//htmnmO
	docx1.Replace("htmnmO", "n71", -1)

	docx1.Replace("O1yIPp", "n81", -1)

	docx1.Replace("J6JnzO", "n91", -1)
	docx1.Replace("hvRgcH", "n92", -1)
	docx1.Replace("nurca2", "n93", -1)
	docx1.Replace("BC7ymZ", "n94", -1)
	docx1.Replace("ZHEASg", "n95", -1)
	docx1.Replace("QDDNSx", "n95", -1)

	docx1.Replace("LHxYbL", "n101", -1)
	docx1.Replace("5T8gRh", "n102", -1)
	docx1.Replace("qGVQAc", "n103", -1)
	docx1.Replace("3MLZgZ", "n104", -1)
	docx1.Replace("yNUQ6R", "n105", -1)
	docx1.Replace("5Bb2VM", "n106", -1)

	docx1.Replace("bjRvEK", "n111", -1)
	docx1.Replace("7NqZaM", "n112", -1)
	docx1.Replace("HRaHwa", "n113", -1)

	docx1.Replace("YCFP8t", "n121", -1)
	docx1.Replace("8Y6w5v", "n122", -1)
	docx1.Replace("KwDHfJ", "n123", -1)
	docx1.Replace("WUt462", "n124", -1)
	docx1.Replace("s2Ch48", "n125", -1)
	docx1.Replace("XJPJen", "n126", -1)

	docx1.WriteToFile("./new_result_1.docx")

}

// type Token struct {
// 	BuildingNum1      string
// 	BuildingNum2      string
// 	BuildingNum3      string
// 	BuildingDeveloper string

// 	Buyer      string
// 	Address    string
// 	BuyerPhone string

// 	AgreedArea      string
// 	AgreedUnitPrice string
// 	AgreedAllPrice  string

// 	ActualArea string

// 	ErrorSize string
// 	ErrorArea string

// 	ReceiveRefund string
// 	ReceiveAmt    string
// 	FinalAllPrice string
// }
