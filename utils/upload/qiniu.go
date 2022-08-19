package upload

import (
	"BookRecSystem/global"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"go.uber.org/zap"
)

type Qiniu struct{}

// UploadFile
// @object: *Qiniu
// @description: 上传文件
// @param: file *multipart.FileHeader
// @return: string, string, error
func (*Qiniu) UploadFile(file *multipart.FileHeader) (string, string, error) {
	putPolicy := storage.PutPolicy{Scope: global.GSD_CONFIG.Qiniu.Bucket}
	mac := qbox.NewMac(global.GSD_CONFIG.Qiniu.AccessKey, global.GSD_CONFIG.Qiniu.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := qiniuConfig()
	formUploader := storage.NewFormUploader(cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{Params: map[string]string{"x:name": "github logo"}}

	f, openError := file.Open()
	if openError != nil {
		global.GSD_LOG.Error("function file.Open() Filed", zap.Any("err", openError.Error()))

		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close()                                                  // 创建文件 defer 关闭
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename) // 文件名格式 自己可以改 建议保证唯一性
	putErr := formUploader.Put(context.Background(), &ret, upToken, fileKey, f, file.Size, &putExtra)
	if putErr != nil {
		global.GSD_LOG.Error("function formUploader.Put() Filed", zap.Any("err", putErr.Error()))
		return "", "", errors.New("function formUploader.Put() Filed, err:" + putErr.Error())
	}
	return global.GSD_CONFIG.Qiniu.ImgPath + "/" + ret.Key, ret.Key, nil
}

// DeleteFile
// @object: *Qiniu
// @description: 删除文件
// @param: key string
// @return: error
func (*Qiniu) DeleteFile(key string) error {
	mac := qbox.NewMac(global.GSD_CONFIG.Qiniu.AccessKey, global.GSD_CONFIG.Qiniu.SecretKey)
	cfg := qiniuConfig()
	bucketManager := storage.NewBucketManager(mac, cfg)
	if err := bucketManager.Delete(global.GSD_CONFIG.Qiniu.Bucket, key); err != nil {
		global.GSD_LOG.Error("function bucketManager.Delete() Filed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}
	return nil
}

// qiniuConfig
// @object: *Qiniu
// @description: 根据配置文件进行返回七牛云的配置
// @return: *storage.Config
func qiniuConfig() *storage.Config {
	cfg := storage.Config{
		UseHTTPS:      global.GSD_CONFIG.Qiniu.UseHTTPS,
		UseCdnDomains: global.GSD_CONFIG.Qiniu.UseCdnDomains,
	}
	switch global.GSD_CONFIG.Qiniu.Zone { // 根据配置文件进行初始化空间对应的机房
	case "ZoneHuadong":
		cfg.Zone = &storage.ZoneHuadong
	case "ZoneHuabei":
		cfg.Zone = &storage.ZoneHuabei
	case "ZoneHuanan":
		cfg.Zone = &storage.ZoneHuanan
	case "ZoneBeimei":
		cfg.Zone = &storage.ZoneBeimei
	case "ZoneXinjiapo":
		cfg.Zone = &storage.ZoneXinjiapo
	}
	return &cfg
}
