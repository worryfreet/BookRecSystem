package upload

import (
	"BookRecSystem/global"
	"BookRecSystem/utils"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"

	"go.uber.org/zap"
)

type Local struct{}

// UploadFile
//@object: *Local
//@description: 上传文件
//@param: file *multipart.FileHeader
//@return: string, string, error
func (*Local) UploadFile(file *multipart.FileHeader) (string, string, error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V([]byte(name))
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	// 尝试创建此路径
	mkdirErr := os.MkdirAll(global.GSD_CONFIG.Local.Path, os.ModePerm)
	if mkdirErr != nil {
		global.GSD_LOG.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	p := global.GSD_CONFIG.Local.Path + "/" + filename

	f, openError := file.Open() // 读取文件
	defer f.Close()             // 创建文件 defer 关闭
	if openError != nil {
		global.GSD_LOG.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}

	out, createErr := os.Create(p)
	defer out.Close() // 创建文件 defer 关闭
	if createErr != nil {
		global.GSD_LOG.Error("function os.Create() Filed", zap.Any("err", createErr.Error()))
		return "", "", errors.New("function os.Create() Filed, err:" + createErr.Error())
	}

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		global.GSD_LOG.Error("function io.Copy() Filed", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return p, filename, nil
}

// DeleteFile
// @object: *Local
// @description: 删除文件
// @param: key string
// @return: error
func (*Local) DeleteFile(key string) error {
	p := global.GSD_CONFIG.Local.Path + "/" + key
	if strings.Contains(p, global.GSD_CONFIG.Local.Path) {
		if err := os.Remove(p); err != nil {
			return errors.New("本地文件删除失败, err:" + err.Error())
		}
	}
	return nil
}
