package common

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"strconv"
	"sync"
)
type  Field struct {
	Label string `form:"label"`
	Content string `form:"content"`
	Checked bool `form:"checked"`
}

type FromData struct {
	CompanyName string `form:"companyName"`
	RegisterTime string `form:"registerTime"`
	MainBusiness string `form:"mainBusiness"`
	RegisteredCapital  float64 `form:"registeredCapital"`
	PaidCapital float64 `form:"paidCapital"`
	TaxLevel string `form:"taxLevel"`
	TaxAmount float64 `form:"taxAmount"`
	Assets float64 `form:"assets"`
	Liabilities float64 `form:"liabilities"`
	Loan float64 `form:"loan"`
	Income float64 `form:"income"`
	Profit float64 `form:"profit"`
	Account float64 `form:"account"`
	//Purpose string `form:"purpose"`
	//Guarantee string `form:"guarantee"`
	Purpose []*Field
	Guarantee []*Field
	Contact string `form:"contact"`
	Phone string `form:"phone"`
	Created string
}
var Lock = sync.Mutex{}
//var LastSendTime = time.Now().Unix() - 86400
var FormDataHistory = make([]*FromData,0)

func send()  {
	//邮件
	html := `<style type="text/css">.table{width:100%;padding:0;margin:0}th{font:bold 12px "Trebuchet MS",Verdana,Arial,Helvetica,sans-serif;color:#4f6b72;border-right:1px solid #c1dad7;border-bottom:1px solid #c1dad7;border-top:1px solid #c1dad7;letter-spacing:2px;text-transform:uppercase;text-align:left;padding:6px 6px 6px 12px;background:#cae8ea no-repeat}td{border-right:1px solid #c1dad7;border-bottom:1px solid #c1dad7;background:#fff;font-size:14px;padding:6px 6px 6px 12px;color:#4f6b72}</style></head><body><table class="table" cellspacing="0" summary="The technical specifications of the Apple PowerMac G5 series">`
	var tdsCompanyName,tdsRegisterTime,tdsMainBusiness,tdsRegistCapital,tdsPaidCapital,tdsTaxLevel,tdsTaxAmount,tdsAssets,tdsLiabilities,tdsLoan,tdsIncome,tdsProfit,tdsAccount,tdspurpose,tdsguarantee,tdsContact,tdsPhone,tdsCreated string
	for i := len(FormDataHistory)-1; i >= 0; i--{
		formData := FormDataHistory[i]
		purpose := ""
		for _,v := range formData.Purpose{
			if v.Checked {
				purpose += fmt.Sprintf("<p>%s</p>",v.Content)
			}
		}

		guarantee := ""
		for _,v := range formData.Guarantee{
			if v.Checked {
				guarantee += fmt.Sprintf("<p>%s</p>",v.Content)
			}
		}
		tdsCompanyName += fmt.Sprintf("<td>%s</td>",formData.CompanyName)
		tdsRegisterTime += fmt.Sprintf("<td>%s</td>",formData.RegisterTime)
		tdsMainBusiness += fmt.Sprintf("<td>%s</td>",formData.MainBusiness)
		tdsRegistCapital += fmt.Sprintf("<td>%s</td>",strconv.FormatFloat(formData.RegisteredCapital, 'g', 6, 64))
		tdsPaidCapital += fmt.Sprintf("<td>%s</td>",strconv.FormatFloat(formData.PaidCapital, 'g', 6, 64))
		tdsTaxLevel += fmt.Sprintf("<td>%s</td>",formData.TaxLevel)
		tdsTaxAmount += fmt.Sprintf("<td>%s</td>",strconv.FormatFloat(formData.TaxAmount, 'g', 6, 64))
		tdsAssets += fmt.Sprintf("<td>%s</td>",strconv.FormatFloat(formData.Assets, 'g', 6, 64))
		tdsLiabilities += fmt.Sprintf("<td>%s</td>",strconv.FormatFloat(formData.Liabilities, 'g', 6, 64))
		tdsLoan += fmt.Sprintf("<td>%s</td>",strconv.FormatFloat(formData.Loan, 'g', 6, 64))
		tdsIncome += fmt.Sprintf("<td>%s</td>",strconv.FormatFloat(formData.Income, 'g', 6, 64))
		tdsProfit += fmt.Sprintf("<td>%s</td>",strconv.FormatFloat(formData.Profit, 'g', 6, 64))
		tdsAccount += fmt.Sprintf("<td>%s</td>",strconv.FormatFloat(formData.Account, 'g', 6, 64))
		tdspurpose += fmt.Sprintf("<td>%s</td>",purpose)
		tdsguarantee += fmt.Sprintf("<td>%s</td>",guarantee)
		tdsContact += fmt.Sprintf("<td>%s</td>",formData.Contact)
		tdsPhone += fmt.Sprintf("<td>%s</td>",formData.Phone)
		tdsCreated += fmt.Sprintf("<td>%s</td>",formData.Created)
	}
	html += fmt.Sprintf(`<tr><td>企业名称</td>%s</tr><tr><td>注册时间</td>%s</tr><tr><td>主营业务</td>%s</tr><tr><td>注册资本（单位/万元）</td>%s</tr><tr><td>实收资本（单位/万元）</td>%s</tr><tr><td>纳税等级</td>%s</tr><tr><td>上一年纳税总额（单位/万元）</td>%s</tr><tr><td>上一年度总资产（单位/万元）</td>%s</tr><tr><td>上一年度总负债（单位/万元）</td>%s</tr><tr><td>银行贷款（单位/万元）</td>%s</tr><tr><td>上一年度营业收入（单位/万元）</td>%s</tr><tr><td>上一年度净利润（单位/万元）</td>%s</tr><tr><td>拟融资额度（单位/万元）</td>%s</tr><tr><td>资金用途</td>%s</tr><tr><td>担保方式</td>%s</tr><tr><td>企业联系人</td>%s</tr><tr><td>联系电话</td>%s</tr><tr><td>提交时间</td>%s</tr></html>`,
		tdsCompanyName,
		tdsRegisterTime,
		tdsMainBusiness,
		tdsRegistCapital,
		tdsPaidCapital,
		tdsTaxLevel,
		tdsTaxAmount,
		tdsAssets,
		tdsLiabilities,
		tdsLoan,
		tdsIncome,
		tdsProfit,
		tdsAccount,
		tdspurpose,
		tdsguarantee,
		tdsContact,
		tdsPhone,
		tdsCreated)
	var ep = &EmailParam{
		Source.EmailStmpHost,
		Source.EmailStmpPort,
		Source.EmailFrom,
		Source.EmailPassword,
		Source.EmailToers,
		Source.EmailCcers,
	}
	InitEmail(ep)
	SendEmail("山西转型综合改革示范区金融服务平台-企业融资需求表",html)

	logs.Debug("send email success...")
}
