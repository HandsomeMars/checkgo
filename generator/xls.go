package generator

import (
	"bytes"
	"checkgo/conf"
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"net/http"
	"time"
)

func init() {
	http.HandleFunc("/generator/xls", GeneratorXlsHandler)
}

type sheet struct {
	sheet string
	row   int
	col   int
	drow  int
	dcol  int
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

func generatorXls() (*excelize.File, error) {

	activitys, err := conf.GetChecker().GetAllActivity()
	if err != nil {
		log.Fatal(err)
	}

	xlsx := excelize.NewFile()
	for _, activity := range activitys {
		sheet := &sheet{
			sheet: activity.Name,
			row:   1,
			col:   1,
			drow:  1,
			dcol:  1,
		}
		// Create a new sheet.
		index := xlsx.NewSheet(activity.Name)
		xlsx.SetActiveSheet(index)
		// Set value of a cell.
		for _, comment := range activity.Comments {
			if handleComment(xlsx, sheet, comment); err != nil {
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

//handleComment 处理组件
func handleComment(xlsx *excelize.File, sheet *sheet, comment *conf.Comment) error {
	switch comment.Type {
	case conf.Table:
		if err := handleTableComment(xlsx, sheet, comment); err != nil {
			return err
		}
	case conf.Form:
		if err := handleFormComment(xlsx, sheet, comment); err != nil {
			return err
		}
	default:
		return errors.New("未知类型" + string(comment.Type))
	}
	return nil
}

func handleTableComment(xlsx *excelize.File, sheet *sheet, comment *conf.Comment) error {

	//组件起始位置
	currentRow := sheet.row
	currentCol := sheet.col
	//currentDRow := sheet.drow
	//currentDCol := sheet.dcol

	index, err := excelize.CoordinatesToCellName(currentCol, currentRow)
	if err != nil {
		return err
	}

	//写表头
	xlsx.SetCellValue(sheet.sheet, index, comment.Name)
	sheet.row = sheet.row + 1

	//写子组件
	for _, comment := range comment.Comments {
		//子组件表头
		index, err := excelize.CoordinatesToCellName(sheet.col, sheet.row)
		if err != nil {
			return err
		}
		xlsx.SetCellValue(sheet.sheet, index, comment.Name)
		sheet.col = sheet.col + 1
		sheet.row = sheet.row + 1
		//子组件模板数据
		if err := handleCommentValue(xlsx, sheet, comment); err != nil {
			return err
		}
		sheet.row = sheet.row - 1
	}
	//归位另起一行
	sheet.col = currentCol
	sheet.row = sheet.drow + 1

	return nil
}

func handleFormComment(xlsx *excelize.File, sheet *sheet, comment *conf.Comment) error {

	//组件起始位置
	currentRow := sheet.row
	currentCol := sheet.col

	index, err := excelize.CoordinatesToCellName(currentCol, currentRow)
	if err != nil {
		return err
	}

	//写表头
	xlsx.SetCellValue(sheet.sheet, index, comment.Name)
	sheet.row = sheet.row + 1

	//写子组件
	for _, comment := range comment.Comments {
		index, err := excelize.CoordinatesToCellName(sheet.col, sheet.row)
		if err != nil {
			return err
		}
		xlsx.SetCellValue(sheet.sheet, index, comment.Name)
		sheet.row = sheet.row + 1

		sheet.row = sheet.col + 1
		//子组件模板数据
		if err := handleCommentValue(xlsx, sheet, comment); err != nil {
			return err
		}
		sheet.row = sheet.col - 1
	}

	//归位另起一行
	sheet.col = currentCol
	sheet.row = sheet.drow + 1
	return nil
}

func handleCommentValue(xlsx *excelize.File, sheet *sheet, comment *conf.Comment) error {
	switch comment.Type {
	case conf.Array:
		if err := handleArray(xlsx, sheet, comment); err != nil {
			return err
		}
	case conf.String:
		if err := handleString(xlsx, sheet, comment); err != nil {
			return err
		}
	default:
		return errors.New("未知类型" + string(comment.Type))
	}
	return nil
}

func handleArray(xlsx *excelize.File, sheet *sheet, comment *conf.Comment) error {

	//组件起始位置
	sheet.drow = sheet.row
	sheet.dcol = sheet.col

	//currentRow := sheet.row
	currentCol := sheet.col

	//写表头
	rows := make([][]interface{}, 0)
	if comment.Example != nil {
		rows = comment.Example.([][]interface{})
	}
	for _, cols := range rows {
		sheet.drow = sheet.drow + 1
		for _, data := range cols {
			sheet.dcol = sheet.dcol + 1
			index, err := excelize.CoordinatesToCellName(sheet.dcol, sheet.drow)
			if err != nil {
				return err
			}
			xlsx.SetCellValue(sheet.sheet, index, data)
		}
		sheet.dcol = currentCol
	}

	return nil
}

func handleString(xlsx *excelize.File, sheet *sheet, comment *conf.Comment) error {

	//组件起始位置
	sheet.drow = sheet.row
	sheet.dcol = sheet.col

	//写表头
	rows := make([]interface{}, 0)
	if comment.Example != nil {
		rows = comment.Example.([]interface{})
	}
	for _, data := range rows {
		sheet.drow = sheet.drow + 1
		index, err := excelize.CoordinatesToCellName(sheet.dcol, sheet.drow)
		if err != nil {
			return err
		}
		xlsx.SetCellValue(sheet.sheet, index, data)
	}

	return nil
}
