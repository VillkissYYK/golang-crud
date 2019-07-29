package conf
// import (

// 	"go-crud/model"
// 	"os"

// 	"github.com/joho/godotenv"
// )
//原版用了Redis，我目前没有装，所以先不用
import (
	"golang-crud/cache"
	"golang-crud/model"
	"os"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		panic(err)
	}

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	//建立Redis缓存
	cache.Redis()
}

