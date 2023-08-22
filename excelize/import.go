package main

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CheckMachineAccount struct {
	No                   uint   `gorm:"column:no;type:int(10) unsigned;comment:序号;NOT NULL" json:"no"`
	DeviceNo             string `gorm:"column:device_no;type:varchar(128);comment:设备编号" json:"device_no"`
	DeviceName           string `gorm:"column:device_name;type:varchar(128);comment:设备名称" json:"device_name"`
	CheckStandardNo      string `gorm:"column:check_standard_no;type:varchar(128);comment:点检标准号" json:"check_standard_no"`
	CheckProjectName     string `gorm:"column:check_project_name;type:varchar(128);comment:点检项目名" json:"check_project_name"`
	CheckContent         string `gorm:"column:check_content;type:varchar(128);comment:点检内容" json:"check_content"`
	StandardNature       string `gorm:"column:standard_nature;type:varchar(128);comment:点检性质" json:"standard_nature"`
	PlanType             string `gorm:"column:plan_type;type:varchar(128);comment:计划类型" json:"plan_type"`
	PlanDt               string `gorm:"column:plan_dt;type:varchar(128);comment:计划日期" json:"plan_dt"`
	TodoJobNo            string `gorm:"column:todo_job_no;type:varchar(128);comment:待办岗号" json:"todo_job_no"`
	TodoJobName          string `gorm:"column:todo_job_name;type:varchar(128);comment:待办岗位名称" json:"todo_job_name"`
	ResponsibleAreaCode  string `gorm:"column:responsible_area_code;type:varchar(128);comment:责任区域编码" json:"responsible_area_code"`
	ResponsibleAreaName  string `gorm:"column:responsible_area_name;type:varchar(128);comment:责任区域名称" json:"responsible_area_name"`
	InvalidDt            string `gorm:"column:invalid_dt;type:varchar(128);comment:失效日期" json:"invalid_dt"`
	CompleteDt           string `gorm:"column:complete_dt;type:varchar(128);comment:交工日期" json:"complete_dt"`
	StandardType         string `gorm:"column:standard_type;type:varchar(128);comment:标准类型" json:"standard_type"`
	Implementor          string `gorm:"column:implementor;type:varchar(128);comment:实施方" json:"implementor"`
	CheckPeriod          string `gorm:"column:check_period;type:varchar(128);comment:点检周期" json:"check_period"`
	PeriodUnit           string `gorm:"column:period_unit;type:varchar(128);comment:周期单位" json:"period_unit"`
	DeviceStatus         string `gorm:"column:device_status;type:varchar(128);comment:设备状态" json:"device_status"`
	CheckMethod          string `gorm:"column:check_method;type:varchar(128);comment:点检方法" json:"check_method"`
	Status               string `gorm:"column:status;type:varchar(128);comment:状态" json:"status"`
	CheckPerson          string `gorm:"column:check_person;type:varchar(128);comment:点检人" json:"check_person"`
	CheckPersonSignature string `gorm:"column:check_person_signature;type:varchar(128);comment:点检人签名" json:"check_person_signature"`
}

func (m *CheckMachineAccount) TableName() string {
	return "check_machine_account"
}

func main() {

	dsn := "root:163766@tcp(yuhan.bestzyx.com:3306)/excel_to_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	excel, err := excelize.OpenFile("/Users/zhangyongxiang/Downloads/点检台账.xlsx")
	if err != nil {
		panic("fail to open excel file: " + "/Users/zhangyongxiang/Downloads/点检台账.xlsx, reason: " + err.Error())
	}
	sheets := excel.GetSheetList()
	for _, sheet := range sheets {
		rows, err := excel.GetRows(sheet)
		if err != nil {
			panic(err.Error())
		}
		for rowNum, row := range rows {
			if rowNum < 2 {
				continue
			}
			no, _ := strconv.Atoi(row[0])
			e := CheckMachineAccount{
				No:                   uint(no),
				DeviceNo:             row[1],
				DeviceName:           row[2],
				CheckStandardNo:      row[3],
				CheckProjectName:     row[4],
				CheckContent:         row[5],
				StandardNature:       row[6],
				PlanType:             row[7],
				PlanDt:               row[8],
				TodoJobNo:            row[9],
				TodoJobName:          row[10],
				ResponsibleAreaCode:  row[11],
				ResponsibleAreaName:  row[12],
				InvalidDt:            row[13],
				CompleteDt:           row[14],
				StandardType:         row[15],
				Implementor:          row[16],
				CheckPeriod:          row[17],
				PeriodUnit:           row[18],
				DeviceStatus:         row[19],
				CheckMethod:          row[20],
				Status:               row[21],
				CheckPerson:          row[22],
				CheckPersonSignature: "",
			}
			err := db.Create(&e).Error
			if err != nil {
				fmt.Printf("current sheet %s row %d processd failed, %s", sheet, rowNum, err)
			}
		}
	}

}
