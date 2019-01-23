package controllers

import (
	"compress/flate"
	"fmt"
	"github.com/kataras/iris"
	"github.com/mholt/archiver"
	"io"
	"io/ioutil"
	"os"
	"qiaoyi_back/model"
	"qiaoyi_back/repositorys"
	"qiaoyi_back/services"
	"qiaoyi_back/utils"
	"strings"
	"time"
)

type FileController struct {
	Service services.FileService
	Ctx iris.Context
}

func NewFileController() *FileController {
	return &FileController{Service: services.NewFileService(repositorys.NewFileRepository())}
}

type FileCondition struct {
	CurrentPage int
	PageSize    int
	Id          uint
}

//图片上传
func (c *FileController) PostImg() (result model.ResultBean) {
	str,_ := utils.UploadFile(c.Ctx)
	if str == "1" || str == "2" || str == "3" {
		result.Status = false
		result.Msg = "系统异常,上传失败"
	} else {
		result.Status = true
		result.Msg = "success"
		result.Data = str
	}
	return
}

//附件/图片上传
func (c *FileController) PostAttachment() (result model.ResultBean) {
	str,fileName:= utils.UploadFile(c.Ctx)
	var multimedia model.Multimedia
	multimedia.Name = fileName
	multimedia.Path = str
	fmt.Println("Upload File:",fileName,"-------Save Path:",str)
	if str == "1" || str == "2" || str == "3" {
		result.Status = false
		result.Msg = "系统异常,上传失败"
	} else {
		result.Status = true
		result.Msg = "success"
		result.Data = multimedia
	}
	return
}

//附件上传
func (c *FileController) PostFiles() (result model.ResultBean) {
	str,_:= utils.UploadFile(c.Ctx)
	if str == "1" || str == "2" || str == "3" {
		result.Status = false
		result.Msg = "系统异常,上传失败"
	} else {
		result.Status = true
		result.Msg = "success"
		result.Data = str
	}
	return
}

//图片显示
func (c *FileController) GetShowImg() {
	fileName := c.Ctx.URLParam("fileName")
	file, _ := ioutil.ReadFile(utils.Conf.FileSavePath + fileName)
	w := c.Ctx.ResponseWriter()
	w.Write(file)
}

//文件下载
func (c *FileController) GetDownfileBy(id uint) {
	multimedia := c.Service.GetMultimediaByID(id)
	fileSavePath := utils.Conf.FileSavePath
	utils.DownLoadFile(c.Ctx, fileSavePath+"/"+multimedia.Path, multimedia.Name)
}

//文件批量下载
func (c *FileController) GetDownfiles() {

	ids := c.Ctx.Request().FormValue("ids")
	idArr := strings.Split(ids, ",")
	fmt.Println("--------------",idArr)
	fileSavePath := utils.Conf.FileSavePath
	// 压缩文件
	z := archiver.Zip{
		CompressionLevel:       flate.DefaultCompression,
		MkdirAll:               true,
		SelectiveCompression:   true,
		ContinueOnError:        false,
		OverwriteExisting:      false,
		ImplicitTopLevelFolder: false,
	}
	multimedias := c.Service.GetMultimediasByIds(idArr)
	//复制并修改文件名
	for i := 0; i < len(multimedias); i++ {
		srcFile, err := os.Open(fileSavePath + "/" + multimedias[i].Path)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer srcFile.Close()
		//创建文件
		fW, err := os.Create(fileSavePath + "/" + multimedias[i].Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer fW.Close()
		_, err = io.Copy(fW, srcFile)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	//压缩文件
	var files []string
	for i := 0; i < len(multimedias); i++ {
		files = append(files, fileSavePath+"/"+multimedias[i].Name)
	}
	fileName := time.Now().Format("20060102150405") + ".zip"
	err := z.Archive(files, fileSavePath+"/"+fileName)
	if err != nil {
		panic(err)
	}
	//删除文件
	//下载压缩文件
	if err := utils.DownLoadFile(c.Ctx, fileSavePath+"/"+fileName, fileName); err != nil{
		fmt.Println(err)
	}

	defer func() {
		//恢复程序的控制权
		err := recover()
		if err==nil{
			err:=os.Remove(fileSavePath+"/"+fileName)
			if err!=nil{
				fmt.Println(err)
			}
			for i := 0; i < len(multimedias); i++ {
				err2:=os.Remove(fileSavePath + "/" + multimedias[i].Name)
				if err2!=nil{
					fmt.Println(err2)
				}
			}
		}
	}()
}

//下载专家文件
func (c *FileController) GetDownMstFile(){
	year := c.Ctx.Request().FormValue("year")
	remark := c.Ctx.Request().FormValue("remark")
	var multimedia model.Multimedia
	multimedia.Year = year
	multimedia.Remark = remark
	//"reviewtoknow"/"operatemanual"
	multimedia = c.Service.GetByQuery(multimedia)
	fileSavePath := utils.Conf.FileSavePath
	if multimedia.Path != "" {
		utils.DownLoadFile(c.Ctx, fileSavePath+"/"+multimedia.Path, multimedia.Name)
	}
}

//下载银行信息
//func (c *FileController) GetDownBankinfo(){
//	ids := c.Ctx.Request().FormValue("ids")
//	idArr := strings.Split(ids, ",")
//	fmt.Println(idArr)
//	mstExperts := serviceMark.GetMarkProInfoByIds(idArr)
//	filePath,fileName := serviceMark.DownBankinfo(mstExperts)
//	if filePath != "" {
//		utils.DownLoadFile(c.Ctx, filePath, fileName)
//	}
//}