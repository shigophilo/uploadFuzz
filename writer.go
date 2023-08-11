package main

import (
	"fmt"
	"os"

	"github.com/qax-os/excelize"
)

func writer(results []string) {
	// 使用os.Stat函数获取文件信息
	file := Outfile + ".xlsx"
	_, err := os.Stat(file)

	if err == nil {
		f, err := excelize.OpenFile(file)
		if err != nil {
			fmt.Println(err)
		}
		//	fmt.Println("写入数据")
		f.SetSheetRow("上传详情", "A2", &results)
		//插入空白行
		//fmt.Println("插入空白行")
		err1 := f.InsertRows("上传详情", 2, 1)
		if err1 != nil {
			fmt.Println(err1)
		}
		if err := f.SaveAs(file); err != nil {
			fmt.Println(err)
		}
	} else if os.IsNotExist(err) {
		//如果文件不存在,创建,并写入第一行
		f := excelize.NewFile()
		index, _ := f.NewSheet("上传详情")
		f.SetActiveSheet(index)
		f.SetCellValue("上传详情", "A1", "ID")
		f.SetCellValue("上传详情", "B1", "状态码")
		f.SetCellValue("上传详情", "C1", "请求包")
		f.SetCellValue("上传详情", "D1", "返回包")
		f.SetCellValue("上传详情", "E1", "特征匹配")
		if err := f.SaveAs(file); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("打开保存文件发生错误,结果可能无法保存:", err)
	}
}
