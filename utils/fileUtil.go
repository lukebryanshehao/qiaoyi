package utils

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/json-iterator/go"
	"github.com/kataras/iris"
	"github.com/tealeg/xlsx"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type Confs struct {
	Port          string `json:"Port"`
	StaticPath    string `json:"StaticPath"`
	TokenKey      string `json:"TokenKey"`
	RedisHost     string `json:"RedisHost"`
	RedisPort     string `json:"RedisPort"`
	RedisPassword string `json:"RedisPassword"`
	SmsAccount    string `json:"SmsAccount"`
	SmsPassword   string `json:"SmsPassword"`
	SmsUrl        string `json:"SmsUrl"`
	SysUrl        string `json:"SysUrl"`
	FileSavePath  string `json:"FileSavePath"`
	Ip            string `json:"Ip"`
}

var Conf = &Confs{}

func init() {
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic("DB config read err")
	}
	//can only unmarshal into pointer
	err = jsoniter.Unmarshal(b, Conf)
	if err != nil {
		panic(err)
	}
}

//复制文件	目标文件，源文件
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

//文件上传
func UploadFile(u iris.Context) (string,string) {
	//获取文件内容 要这样获取
	file, head, err := u.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return "1","1"
	}
	defer file.Close()
	fmt.Println(head.Filename)
	//获取当前时间戳作为文件名
	fileName := GetTimeUnix()
	//截取后缀
	suffix := head.Filename[strings.LastIndex(head.Filename, "."):]

	//创建文件夹:当前时间年月日
	nowTime := time.Now()
	timeStr := nowTime.Format("2006") + "/" + nowTime.Format("200601")
	path := Conf.FileSavePath + "/" + timeStr
	exist, _ := PathExists(path)
	if !exist {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		}
	}

	//创建文件
	fW, err := os.Create(path + "/" + fileName + suffix)
	if err != nil {
		fmt.Println("文件创建失败")
		return "2","2"
	}
	defer fW.Close()
	_, err = io.Copy(fW, file)
	if err != nil {
		fmt.Println("文件保存失败")
		return "3","2"
	}
	return timeStr + "/" + fileName + suffix,head.Filename
}

//文件下载
func DownLoadFile(u iris.Context, allFileName string, newFileName string)error {
	a:=u.ResponseWriter()
	a.Header().Set("Content-Disposition", "attachment; filename="+newFileName)
	a.Header().Set("Content-Type", "application/octet-stream")
	return u.SendFile(allFileName, newFileName)
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//读取Excel
func ReadExcel(path string, fileName string) ([][]string) {
	var data [][]string
	xlsx, err := excelize.OpenFile(path + fileName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// Get value from cell by given worksheet name and axis.
	xlsx.GetCellValue("Sheet1", "B2")
	// Get all the rows in the Sheet1.
	rows := xlsx.GetRows("Sheet1")
	for rindex, row := range rows {
		if rindex < 5 {
			continue
		}
		var rdata []string
		for cindex, colCell := range row {
			if cindex < 1 {
				continue
			}
			rdata = append(rdata, colCell)
		}
		data = append(data, rdata)
	}
	return data
}
func ExportExcel1(header []string, data interface{}) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row, row1 *xlsx.Row
	var cell *xlsx.Cell
	var err error
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	row.SetHeightCM(1)
	for _, v := range header {
		cell = row.AddCell()
		cell.Value = v
	}
	y, ok := data.([][]string)
	for _, v := range y {
		row1 = sheet.AddRow()
		if ok {
			//二维string切片
			var cv []string = v
			for _, cellValue := range cv {
				cell = row1.AddCell()
				cell.Value = cellValue
			}
		} else {
			//模型切片
			/*r := reflect.ValueOf(v)
			fmt.Println(r)
			count := r.NumField()
			for i := 0; i < count; i++ {
				cell = row1.AddCell()
				f := r.Field(i)
				switch f.Kind() {
				case reflect.String:
					if f.String() != "" {
						cell.Value = f.String()
					}
				case reflect.Int:
					if f.Int() != 0 {
						cell.Value = strconv.FormatInt(f.Int(), 10)
					}
				}
			}*/
		}
	}
}
