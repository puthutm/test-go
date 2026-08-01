package handlers

import (
	"data-referensi/helpers"
	"fmt"

	"github.com/xuri/excelize/v2"
)

/* Export Data From Model */
func ModelExport(
	fileSaveAs string,
	headers []string,
	data []map[string]interface{},
) error {
	file := excelize.NewFile()
	sheetName := "Sheet1"
	file.NewSheet(sheetName)

	// Menulis header ke dalam file Excel
	for i, header := range headers {
		col := string(rune('A' + i))
		cell := fmt.Sprintf("%s1", col)
		file.SetCellValue(sheetName, cell, header)
	}

	for rowIndex, row := range data {
		rowNum := rowIndex + 2
		for colIndex, header := range headers {
			col := string(rune('A' + colIndex))
			cell := fmt.Sprintf("%s%d", col, rowNum)
			value := row[header]
			file.SetCellValue(sheetName, cell, value)
		}
	}

	for i := range headers {
		col := string(rune('A' + i))
		helpers.ExcelAutoSizeColumn(file, sheetName, col, len(data))
	}

	if err := file.SaveAs(fileSaveAs); err != nil {
		return fmt.Errorf("failed to save XLSX file: %v", err)
	}

	return nil
}

func ModelImport(
	filePath string,
	headers []string,
	processRow func(row map[string]string) error,
) error {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to open excel file: %v", err)
	}

	sheetName := "Sheet1"
	if sheetName == "" {
		return fmt.Errorf("sheet not found in file")
	}

	rows, err := file.GetRows(sheetName)
	if err != nil {
		return fmt.Errorf("failed to get rows: %v", err)
	}

	if len(rows) < 1 {
		return fmt.Errorf("file is empty or missing header row")
	}

	fileHeaders := rows[0]
	if len(fileHeaders) != len(headers) {
		return fmt.Errorf("header mismatch: expected %d headers but got %d", len(headers), len(fileHeaders))
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}

		rowData := make(map[string]string)
		for j, cell := range row {
			if j < len(headers) {
				rowData[headers[j]] = cell
			}
		}
		if err := processRow(rowData); err != nil {
			return fmt.Errorf("error processing row %d: %v", i+1, err)
		}
	}

	return nil
}
