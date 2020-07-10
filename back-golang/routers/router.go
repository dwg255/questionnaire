package routers

import (
	"finance/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/question",&controllers.QuestionController{})
    beego.Router("/finance",&controllers.FinanceInvestigateController{})

}
