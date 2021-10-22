/*
   Created by wangw at 2021/9/2 10:34
   Copyright (c) 2013-present, Xiamen Dianchu Technology Co.,Ltd.
*/

package excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"strconv"
	"time"
)

type Int64 interface {
	Value(interface{}) (int64, error)
}

type Float64 interface {
	Value(interface{}) (float64, error)
}

type Str interface {
	Value(interface{}) (string, error)
}

type Date interface {
	Value(interface{}) (string, error)
}

type Time interface {
	Value(interface{}) (string, error)
}

type DateTime interface {
	Value(interface{}) (string, error)
}

type Enum interface {
	Value(interface{}) (interface{}, error)
}

// Dict 数据字典 比如活动类型等
type Dict interface {
	Value(interface{}) (interface{}, error)
}

// Cell 单元格定义
type Cell struct {
	Row   int
	Col   int
	Value interface{}
}

func (c *Cell) CellName() string {
	name, _ := excelize.CoordinatesToCellName(c.Col+1, c.Row+1)
	return name
}

func (c *Cell) Int64Val(i Int64) (value int64, err error) {
	return i.Value(c.Value)
}

func (c *Cell) Float64Val(f Float64) (value float64, err error) {
	return f.Value(c.Value)
}

func (c *Cell) StrVal(s Str) (value string, err error) {
	return s.Value(c.Value)
}

func (c *Cell) DateVal(d Date) (value string, err error) {
	return d.Value(c.Value)
}

func (c *Cell) TimeVal(t Time) (value string, err error) {
	return t.Value(c.Value)
}

func (c *Cell) DateTimeVal(dt DateTime) (value string, err error) {
	return dt.Value(c.Value)
}

func (c *Cell) EnumVal(e Enum) (value interface{}, err error) {
	return e.Value(c.Value)
}

func (c *Cell) DictVal(d Dict) (value interface{}, err error) {
	return d.Value(c.Value)
}

// Block 数据块定义
type Block struct {
	Name string
	Row  int
	Col  int
	Data [][]*Cell
}

func (b *Block) GetCell(row, col int) (*Cell, error) {
	return getCell(b.Data, row, col)
}

// Sheet Excel 页签定义
type Sheet struct {
	Loc  int
	Name string
	Row  int
	Col  int
	Data [][]*Cell
}

func (s *Sheet) GetCell(row, col int) (*Cell, error) {
	return getCell(s.Data, row, col)
}

func getCell(data [][]*Cell, row, col int) (*Cell, error) {
	cellName, err := excelize.CoordinatesToCellName(col+1, row+1)
	if err != nil {
		return nil, err
	}
	if !(len(data) > row) {
		return nil, fmt.Errorf("单元格%v值为空", cellName)
	}
	if !(len(data[row]) > col) {
		return nil, fmt.Errorf("单元格%v值为空", cellName)
	}
	return data[row][col], nil
}

// GetBlock 通过块名称 读取对应块数据
func (s *Sheet) GetBlock(blockName string) (*Block, error) {
	data := make([][]*Cell, 0)
	start := false

	for _, row := range s.Data {
		if len(row) == 0 {
			continue
		}
		v, _ := row[0].StrVal(StrCellValue)
		if v == blockName {
			start = true
			continue
		}
		if start && v == "" {
			break
		}
		if start {
			data = append(data, row)
		}
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("%v数据块未找到", blockName)
	}

	return &Block{Name: blockName, Data: data}, nil
}

// Excel 定义
type Excel struct {
	Name   string
	Sheets []*Sheet
}

type ExcelReader interface {
	OpenReader(fileName string, r io.Reader) (*Excel, error)
}

var (
	StrCellValue      = &StrCell{}
	Int64CellValue    = &Int64Cell{}
	Float64CellValue  = &Float64Cell{}
	DateCellValue     = &DateCell{}
	TimeCellValue     = &TimeCell{}
	DateTimeCellValue = &DateTimeCell{}
)

type StrCell struct {
}

func (sc *StrCell) Value(v interface{}) (string, error) {
	return v.(string), nil
}

type Int64Cell struct {
}

func (ic *Int64Cell) Value(v interface{}) (int64, error) {
	return strconv.ParseInt(v.(string), 10, 64)
}

type Float64Cell struct {
}

func (fc *Float64Cell) Value(v interface{}) (float64, error) {
	return strconv.ParseFloat(v.(string), 64)
}

type DateCell struct {
}

func (dc *DateCell) Value(v interface{}) (string, error) {
	_, err := time.Parse("2006-01-02", v.(string))
	if err != nil {
		return "", err
	}
	return v.(string), nil
}

type TimeCell struct {
}

func (tc *TimeCell) Value(v interface{}) (string, error) {
	_, err := time.Parse("15:04:05", v.(string))
	if err != nil {
		return "", err
	}
	return v.(string), nil
}

type DateTimeCell struct {
}

func (dtc *DateTimeCell) Value(v interface{}) (string, error) {
	_, err := time.Parse("2006-01-02 15:04:05", v.(string))
	if err != nil {
		return "", err
	}
	return v.(string), nil
}
