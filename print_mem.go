package memalign

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"os"
	"reflect"
	"strconv"
)

type FieldInfo struct {
	StrID string
	FieldType reflect.Type
	FieldName string
	FieldSize int
}

type ColorFunc func(format string, a ...interface{}) string

var (
	colorList []ColorFunc
)

func init() {
	colorList = []ColorFunc{
		color.RedString,
		color.GreenString,
		color.YellowString,
		color.BlueString,
		color.MagentaString,
		color.CyanString,
	}

}



func PrintStructAlignment(q interface{}) {
	buff := make([]byte, 0)
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		typeInfo := reflect.TypeOf(q)
		fieldInfoList := make([]FieldInfo, typeInfo.NumField())


		for i := 0; i < typeInfo.NumField(); i++ {
			fieldInfoList[i].StrID = string(byte(i)+'A')
			fieldInfoList[i].FieldType = typeInfo.Field(i).Type
			fieldInfoList[i].FieldName = typeInfo.Field(i).Name
		}

		for i := 0; i < typeInfo.NumField(); i++ {
			field := typeInfo.Field(i)
			size := int(field.Type.Size())
			fieldInfoList[i].FieldSize = size

			// padding
			paddSize := int(field.Offset) - len(buff)/2
			//fmt.Println(i, ":", field.Offset, len(buff)/2)
			for j := 0; j < paddSize; j++ {
				buff = append(buff, '|', ' ')
			}
			for k := 0; k < size; k++ {
				buff = append(buff, '|', byte(i+'A'))
			}
		}
		if len(buff)%8 != 0 {
			buff = append(buff, '|')
		}

		// print
		// 1.
		fmt.Println("---- Fields in struct ----")
		PrintStructInfo(fieldInfoList)

		// 2.
		fmt.Println("---- Memory layout ----")
		ColorFormatPrint(buff)
	}

}

func PrintStructInfo(itemList []FieldInfo){
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	table.SetHeader([]string{"ID", "FieldType", "FieldName", "FieldSize"})

	for _, item := range itemList {
		table.Append([]string{item.StrID,item.FieldType.String(),
			item.FieldName, strconv.Itoa(item.FieldSize)})
	}
	table.Render() // Send output
}

func ColorFormatPrint(buff []byte) {
	for i := 0; i < len(buff); i++ {
		if buff[i] >= 'A' && buff[i] <= 'Z' {
			idx := int(buff[i]-'A') % len(colorList)
			fmt.Printf(colorList[idx](string(buff[i])))
		} else {
			fmt.Printf(string(buff[i]))
		}
		if (i+1)%16 == 0 {
			fmt.Println("|")
		}
	}
}