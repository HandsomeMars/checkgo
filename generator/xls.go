package generator

import (
	"bytes"
	"checkgo/conf"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"net/http"
	"time"
)

func init() {
	http.HandleFunc("/generator/xls", GeneratorXlsHandler)
	http.HandleFunc("/generator/import", ImportXlsHandler)
}

// GeneratorXlsHandler  导出xls模板
func GeneratorXlsHandler(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	xls, err := generatorXls()
	if err != nil {
		return
	}

	err = xls.Write(buf)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	disposition := fmt.Sprintf("attachment; filename=\"%s-%s.xlsx\"", "模板", time.Now().Format("2006-01-02 15:04:05"))
	w.Header().Set("Content-Disposition", disposition)
	w.WriteHeader(http.StatusOK)
	xls.WriteTo(w)
}

// ImportXlsHandler  导入xls模板
func ImportXlsHandler(w http.ResponseWriter, r *http.Request) {

}

func generatorXls() (*excelize.File, error) {

	activitys, err := conf.GetChecker().GetAllActivity()
	if err != nil {
		log.Fatal(err)
	}

	xlsx := excelize.NewFile()
	for _, activity := range activitys {
		// Create a new sheet.
		index := xlsx.NewSheet(activity.Name)
		row, col := 0, 1
		xlsx.SetActiveSheet(index)
		// Set value of a cell.
		for _, comment := range activity.Comments {
			if _, row, err = writeComment(xlsx, activity.Name, col, row+1, comment); err != nil {
				return xlsx, err
			}
		}
	}

	// Save xlsx file by the given path.
	if err := xlsx.SaveAs("./Book1.xlsx"); err != nil {
		return xlsx, err
	}

	return xlsx, nil
}

func write(xlsx *excelize.File, sheet string, col, row int, data interface{}) error {
	log.Println(sheet, data, row, col)
	index, err := excelize.CoordinatesToCellName(col, row)
	if err != nil {
		return err
	}
	xlsx.SetCellValue(sheet, index, data)
	return nil
}

func writeComment(xlsx *excelize.File, sheet string, col, row int, comment *conf.Comment) (int, int, error) {

	//写表头
	err := write(xlsx, sheet, col, row, comment.Name)
	if err != nil {
		return col, row, err
	}

	//组件起始位置
	//indexRow := row + 1
	//indexCol := col
	currentRow := row + 1
	currentCol := col
	maxRow := currentRow
	maxCol := currentCol

	r, l := 0, 0
	if comment.Layout == conf.Vertical {
		r = 1
	} else {
		l = 1
	}

	//写字段
	for _, filed := range comment.Filed {
		err := write(xlsx, sheet, currentCol, currentRow, filed.Name)
		if err != nil {
			return currentCol, currentRow, err
		}

		var rows []interface{}
		if filed.Example != nil {
			rows = filed.Example.([]interface{})
		}

		for ri, cols := range rows {
			if filed.Type == conf.Array {
				array := cols.([]interface{})
				for li, data := range array {
					if currentCol+li+l > maxCol {
						maxCol = currentCol + li
					}
					err := write(xlsx, sheet, currentCol+li+l, currentRow+ri+r, data)
					if err != nil {
						return currentCol, currentRow, err
					}
				}
			} else {
				err := write(xlsx, sheet, currentCol+l, currentRow+ri+r, cols)
				if err != nil {
					return currentCol, currentRow, err
				}
			}
		}

		currentCol += r
		if currentCol > maxCol {
			maxCol = currentCol
		}

		currentRow += l
		if currentRow > maxRow {
			maxRow = currentRow
		}

		//if comment.Layout == conf.Vertical {
		//	//水平迁移
		//	currentCol += 1
		//	if currentCol > maxCol {
		//		maxCol = currentCol
		//	}
		//} else {
		//	//垂直迁移
		//	currentRow += 1
		//	if currentRow > maxRow {
		//		maxRow = currentRow
		//	}
		//}
	}

	////指针归位
	//if comment.Layout == conf.Vertical {
	//	currentCol = indexCol
	//	currentRow += 1
	//} else {
	//	currentRow = indexRow
	//	currentCol += 1
	//}

	////写子模板数据
	//for _, filed := range comment.Filed {
	//	var rows []interface{}
	//	if filed.Example != nil {
	//		rows = filed.Example.([]interface{})
	//	}
	//
	//	for r, cols := range rows {
	//		if filed.Type == conf.Array {
	//			array := cols.([]interface{})
	//			for l, data := range array {
	//				if currentCol+l > maxCol {
	//					maxCol = currentCol + l
	//				}
	//				err := write(xlsx, sheet, currentCol+l, currentRow+r, data)
	//				if err != nil {
	//					return currentCol, currentRow, err
	//				}
	//			}
	//		} else {
	//			err := write(xlsx, sheet, currentCol, currentRow+r, cols)
	//			if err != nil {
	//				return currentCol, currentRow, err
	//			}
	//		}
	//	}
	//
	//	if comment.Layout == conf.Vertical {
	//		//水平迁移
	//		currentCol += 1
	//		if currentCol > maxCol {
	//			maxCol = currentCol
	//		}
	//	} else {
	//		//垂直迁移
	//		currentRow += 1
	//		if currentRow > maxRow {
	//			maxRow = currentRow
	//		}
	//	}
	//}

	return maxCol, maxRow, nil
}
