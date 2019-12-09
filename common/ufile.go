package common

import (
	"fmt"
	"gin-gorm-demo/conf"
	"net/http"
	"path/filepath"
	"reflect"
	"runtime"
	"time"

	"github.com/Terry-Mao/goconf"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	log "github.com/thinkboy/log4go"
	ufsdk "github.com/ufilesdk-dev/ufile-gosdk"
)

const (
	UFile_Url = "http://test.cn-bj.ufileos.com/"
)

const (
	/*个人测试使用*/
	UCLOUD_PUBLIC_KEY  = "TOKEN_d34523-28d3-4891-81fc-3814fa745"
	UCLOUD_PRIVATE_KEY = "6f23e-2115-4c2f-8c59-336323232e18d"
)

var (
	SuffixAllowed = []string{".jpeg", ".gif", ".jpg", ".png", ".xlsx", ".txt", ".doc", ".docx", ".pdf", ".sql", ".zip", ".log"}
)

func UploadFile(src string, key string) (url string, err error) {

	RemoteFileKey := key
	FilePath := src
	config := &ufsdk.Config{
		PublicKey:  UCLOUD_PUBLIC_KEY,
		PrivateKey: UCLOUD_PRIVATE_KEY,
		BucketName: "test",
		BucketHost: "test.cn-bj.ufileos.com",
		FileHost:   "cn-bj.ufileos.com",
	}
	if err != nil {
		panic(err.Error())
	}

	req, err := ufsdk.NewFileRequest(config, nil)
	if err != nil {
		panic(err.Error())
	}
	log.Info("正在上传文件。。。。")

	err = req.PutFile(FilePath, RemoteFileKey, "")
	if err != nil {
		log.Info("上传文件失败，错误信息为：%s\n", req.DumpResponse(true))
		return
	}

	log.Info("正在下载文件。。。。")
	url = req.GetPublicURL(RemoteFileKey)
	return
}
func checkSuffix(suffix string) bool {
	for _, b := range SuffixAllowed {
		if b == suffix {
			return true
		}
	}
	return true
}

var (
	gconf    *goconf.Config
	Conf     *Config
	confFile string
)

type Config struct {
	// base section
	PidFile string `goconf:"base:pidfile"`
	Dir     string `goconf:"base:dir"`
	Log     string `goconf:"base:log"`
	Upload  string `goconf:"base:upload"`
	// logic push
	HttpPush   string   `goconf:"base:httppush"`
	UBU        string   `goconf:"base:ubu"`
	MaxProc    int      `goconf:"base:maxproc"`
	PprofAddrs []string `goconf:"base:pprof.addrs:,"`
	// rpc
	RPCAddrs         []string      `goconf:"rpc:addrs:,"`
	HTTPAddrs        []string      `goconf:"base:http.addrs:,"`
	HTTPReadTimeout  time.Duration `goconf:"base:http.read.timeout:time"`
	HTTPWriteTimeout time.Duration `goconf:"base:http.write.timeout:time"`
	// bucket
	Bucket            int           `goconf:"bucket:bucket"`
	Server            int           `goconf:"bucket:server"`
	Cleaner           int           `goconf:"bucket:cleaner"`
	BucketCleanPeriod time.Duration `goconf:"bucket:clean.period:time"`
	// session
	Session       int           `goconf:"session:session"`
	SessionExpire time.Duration `goconf:"session:expire:time"`
	// monitor
	MonitorOpen  bool     `goconf:"monitor:open"`
	MonitorAddrs []string `goconf:"monitor:addrs:,"`
	// mysql
	// DBAddrs string `goconf:"mysql:dbaddrs"`
}

func NewConfig() *Config {
	return &Config{
		// base section
		PidFile: "/tmp/uim-router.pid",
		Dir:     "./",
		Log:     "./router-log.xml",
		MaxProc: runtime.NumCPU(),
		//PprofAddrs: []string{"localhost:6971"},
		// rpc
		//RPCAddrs:  []string{"localhost:9090"},
		//HTTPAddrs: []string{"7375"},
		// bucket
		Bucket:            runtime.NumCPU(),
		Server:            5,
		Cleaner:           1000,
		BucketCleanPeriod: time.Hour * 1,
		// session
		Session:       1000,
		SessionExpire: time.Hour * 1,
	}
}

// InitConfig init the global config.
func InitConfig() (err error) {
	Conf = NewConfig()
	gconf = goconf.New()
	if err = gconf.Parse(confFile); err != nil {
		return err
	}
	if err := gconf.Unmarshal(Conf); err != nil {
		return err
	}
	return nil
}

func Upload(c *gin.Context) {

	form, err := c.MultipartForm()

	fmt.Println("FORM格式", reflect.TypeOf(form))

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": conf.UploadFail, "message": fmt.Sprintf("get form err: %s", err.Error())})
		return
	}
	files := form.File["file"]

	var fileName []string
	for _, file := range files {
		//	fmt.Println("filename", file)
		filename := filepath.Base(file.Filename)
		suffix := filepath.Ext(file.Filename)
		fmt.Println("获取文件后缀", suffix)
		// 检查后缀
		if ok := checkSuffix(suffix); !ok {
			c.JSON(http.StatusOK, gin.H{"ret": conf.UploadFail, "Message": fmt.Sprintf("suffix %s not allow err", suffix)})
			return
		}
		//修改文件名称
		u1 := uuid.Must(uuid.NewV4())
		filename = u1.String() + "_" + filename
		dst := Conf.Upload + "/" + filename
		// 下载文件到本地
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusOK, gin.H{"ret": conf.UploadFail, "Message": fmt.Sprintf("upload file err: %s", err.Error())})
			return
		}
		// 上传UFile
		// var url string
		if _, err = UploadFile(dst, filename); err != nil {
			c.JSON(http.StatusOK, gin.H{"ret": conf.UploadFail, "Message": fmt.Sprintf("upload ufile fail: %s ", err.Error())})
			return
		}
		fileName = append(fileName, UFile_Url+filename)
	}
	c.JSON(http.StatusOK, gin.H{"ret": conf.Ret_OK, "urls": fileName})
}
