package util

//func ReadExcel(filepath string) () {
//	// 首先读excel
//	xlsx, err := excelize.OpenFile(filepath)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	// Get all the rows in the Sheet2.
//	rows := xlsx.GetRows("Sheet2")
//	var datas []Data
//	for i, row := range rows {
//		// 去掉第一行，第一行是表头
//		if i == 0 {
//			continue
//		}
//		var data Data
//		for j, colCell := range row {
//			// 去掉前后空格（自己封装的方法，去掉字符串前后的特殊字符）
//			colCell = tools.TrimPrefixSuffix(colCell, " ")
//			// 排除第一列为Null
//			if j == 0 && colCell == "Null" {
//				continue
//			}
//			// 第一列即是一级
//			if j == 0 && colCell != "Null" {
//				data.FirstClass = colCell
//			}
//			// 第二列即是二级
//			if j == 1 {
//				data.SecondClass = colCell
//			}
//			// 三级
//			if j == 2 {
//				data.ThirdClass = colCell
//			}
//
//		}
//		fmt.Println(util.StringifyJson(data))
//		datas = append(datas, data)
//	}
//}
