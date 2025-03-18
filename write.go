package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func writeTo() {
	// 创建一个新的 Excel 文件
	f := excelize.NewFile()
	// 获取默认工作表
	sheetName := f.GetSheetName(f.GetActiveSheetIndex())

	// 定义表头字段
	headers := []string{
		"学年", "学期", "课程号", "课程名称", "开课学院", "学分", "课程性质", "开课类型", "年级", "专业",
		"成绩录入老师", "职称", "周日", "周一", "周二", "周三", "周四", "周五", "周六",
	}

	// 写入表头
	for colIndex, header := range headers {
		// 写入表头数据
		cell, err := excelize.CoordinatesToCellName(colIndex+1, 1)
		if err != nil {
			fmt.Printf("坐标转换出错: %v\n", err)
			return
		}
		if err := f.SetCellValue(sheetName, cell, header); err != nil {
			fmt.Printf("写入单元格数据出错: %v\n", err)
			return
		}
	}

	// 保存 Excel 文件
	if err := f.SaveAs("courses.xlsx"); err != nil {
		fmt.Printf("保存文件出错: %v\n", err)
		return
	}

	fmt.Println("Excel 文件创建成功，已保存为 courses.xlsx")
}
