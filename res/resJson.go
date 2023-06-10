package res

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rxlwin/g-pkg/id"
	"github.com/rxlwin/g-pkg/ip"
	"github.com/rxlwin/g-pkg/util"
	"log"
	"net/http"
	"os"
	"strings"
)

var MachineFlag string

func init() {
	ip0, err0 := ip.GetLocalIP()
	if err0 != nil {
		log.Fatalln(err0.Error())
	}
	ipStr := ip0.String()
	ipList := strings.Split(ipStr, ".")
	MachineFlag = ipList[len(ipList)-1]
	fmt.Println("machineFlag: ", MachineFlag)
}

type result struct {
	ReqId string `json:"req_id"`
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Data  Data   `json:"data"`
}

type Data map[string]interface{}

func Success(c *gin.Context, data Data) {
	export(c, SUCCESS, data, "ok")
}

func Error(c *gin.Context, err Errs) {
	data := Data{}
	if os.Getenv("ENV") != "prod" {
		data = Data{
			"err_msg": err.Error(),
			"notice":  "err_msg只在测试环境显示,方便前端调试",
		}
	}
	fmt.Println("err_msg：", err.Error())
	export(c, err.GetCode(), data, err.GetShowMsg())
}

func export(c *gin.Context, code int, data Data, msg string) {
	if c.Value("request_end_flag") == 1 {
		return
	}
	result := result{
		ReqId: getImpressionId(c),
		Code:  code,
		Msg:   msg,
		Data:  data,
	}
	c.Set("res", result)
	c.Set("request_end_flag", 1)

	fmt.Println("返回 to 用户：", result)
	c.JSON(http.StatusOK, result)
}

func getImpressionId(c *gin.Context) string {
	value := c.Value("impression_id")
	imId, _ := value.(string)
	if imId != "" {
		return imId
	}
	idInt := id.GetId()
	idStr := util.Val2string(idInt)
	h := md5.New()
	h.Write([]byte(idStr))
	imId = hex.EncodeToString(h.Sum(nil)) + "-" + MachineFlag
	c.Set("impression_id", imId)
	return imId
}
