package controllers
import (
	"finance/common"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"math/rand"
	"time"
)

type QuestionController struct {
	beego.Controller
}
type Data struct {
	Code int
	Msg string
}

func (c *QuestionController) Get() {

}

func (c *QuestionController) Post() {
	sql := "insert into f_financial_demand (uuid,company_name,user_name,user_phone,financing_amount,length_maturity,guaranty_style,create_time) value(?,?,?,?,?,?,?,?)"

	if stmt, err := common.Source.Db.Prepare(sql);err != nil {
		logs.Error("sql err 1 : %v",err)
	} else {
		_,err := stmt.Exec(RandString(36),c.GetString("company_name"),c.GetString("contact"),c.GetString("phone"),c.GetString("amount"),
			c.GetString("expire"),c.GetString("guarantee"),time.Now().Format("2006-01-02 15:04:05"),)
		if err != nil {
			logs.Error("sql err 2 :%v",err)
		}
	}
	c.Data["json"] = &Data{
		Code:0,
	}
	c.ServeJSON()
}

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

// RandString 生成随机字符串
func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
