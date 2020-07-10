package controllers

import (
	"encoding/json"
	"finance/common"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
	"time"
)

type FinanceInvestigateController struct {
	beego.Controller
}

func (c *FinanceInvestigateController) Get() {

}

func (c *FinanceInvestigateController) Post() {
	f := &common.FromData{}
	if err := c.ParseForm(f); err != nil {
		fmt.Println(err)
	} else {
		//fmt.Println(f)

		var purposeStr = c.GetString("purpose")
		var purpose  []*common.Field
		if err := json.Unmarshal([]byte(purposeStr),&purpose);err != nil {
			fmt.Println(err)
		} else {
			f.Purpose = purpose
		}

		var guaranteeStr = c.GetString("guarantee")
		var guarantee  []*common.Field
		if err := json.Unmarshal([]byte(guaranteeStr),&guarantee);err != nil {
			fmt.Println(err)
		} else {
			f.Guarantee = guarantee
		}
		f.Created = time.Now().Format("2006-01-02 15:04:05")
		// 先加入队列
		common.Lock.Lock()
		common.FormDataHistory = append(common.FormDataHistory, f)
		common.Lock.Unlock()

		sql := "insert into f_finance_investigate (company_name,register_time,main_business,registered_capital,paid_capital ,tax_level,tax_amount,assets,liabilities,loan,income,profit,account,purpose,guarantee,contact,phone,created) value(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

		if stmt, err := common.Source.Db.Prepare(sql);err != nil {
			logs.Error("sql err 1 : %v",err)
		} else {
			p := ""
			for _,v := range f.Purpose{
				if v.Checked {
					p += v.Content + ","
				}
			}
			g := ""
			for _,v := range f.Guarantee{
				if v.Checked {
					g += v.Content + ","
				}
			}
			_,err := stmt.Exec(f.CompanyName,
				f.RegisterTime,
				f.MainBusiness,
				f.RegisteredCapital,
				f.PaidCapital,
				f.TaxLevel,
				f.TaxAmount,
				f.Assets,
				f.Liabilities,
				f.Loan,
				f.Income,
				f.Profit,
				f.Account,
				p[:len(p)-1],
				g[:len(g)-1],
				//f.Purpose,
				//f.Guarantee,
				f.Contact,
				f.Phone,
				f.Created)
			if err != nil {
				logs.Error("sql err 2 :%v",err)
			}
		}

		//if len(common.FormDataHistory) >= common.Source.EmailLimit || time.Now().Unix() - common.LastSendTime > int64(common.Source.EmailExpire) {	//大于1小时或者超过30封时，发送邮件
		//	//FormDataHistory = FormDataHistory[len(FormDataHistory)-10:]
		//	common.LastSendTime = time.Now().Unix()
		//	//send2()
		//	common.FormDataHistory = make([]*common.FromData,0)
		//}
	}

	c.Data["json"] = &Data{
		Code:0,
	}
	c.ServeJSON()
}

func send()  {
	//邮件
	html := `<style type="text/css">.table{width:100%;padding:0;margin:0}th{font:bold 12px "Trebuchet MS",Verdana,Arial,Helvetica,sans-serif;color:#4f6b72;border-right:1px solid #c1dad7;border-bottom:1px solid #c1dad7;border-top:1px solid #c1dad7;letter-spacing:2px;text-transform:uppercase;text-align:left;padding:6px 6px 6px 12px;background:#cae8ea no-repeat}td{border-right:1px solid #c1dad7;border-bottom:1px solid #c1dad7;background:#fff;font-size:14px;padding:6px 6px 6px 12px;color:#4f6b72}</style></head><body><table class="table" cellspacing="0" summary="The technical specifications of the Apple PowerMac G5 series"><tr><td>企业名称</td><td>注册时间</td><td>主营业务</td><td>注册资本（单位/万元）</td><td>实收资本（单位/万元）</td><td>纳税等级</td><td>上一年纳税总额（单位/万元）</td><td>上一年度总资产（单位/万元）</td><td>上一年度总负债（单位/万元）</td><td>银行贷款（单位/万元）</td><td>上一年度营业收入（单位/万元）</td><td>上一年度净利润（单位/万元）</td><td>拟融资额度（单位/万元）</td><td>资金用途</td><td>担保方式</td><td>企业联系人</td><td>联系电话</td><td>提交时间</td></tr>`
	for i := len(common.FormDataHistory)-1; i >= 0; i--{
		formData := common.FormDataHistory[i]
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
		html += fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td></tr>",
			formData.CompanyName,
			formData.RegisterTime,
			formData.MainBusiness,
			strconv.FormatFloat(formData.RegisteredCapital, 'g', 1, 64),
			strconv.FormatFloat(formData.PaidCapital, 'g', 1, 64),
			formData.TaxLevel,
			strconv.FormatFloat(formData.TaxAmount, 'g', 1, 64),
			strconv.FormatFloat(formData.Assets, 'g', 1, 64),
			strconv.FormatFloat(formData.Liabilities, 'g', 1, 64),
			strconv.FormatFloat(formData.Loan, 'g', 1, 64),
			strconv.FormatFloat(formData.Income, 'g', 1, 64),
			strconv.FormatFloat(formData.Profit, 'g', 1, 64),
			strconv.FormatFloat(formData.Account, 'g', 1, 64),
			purpose,
			guarantee,
			//formData.Purpose,
			//formData.Guarantee,
			formData.Contact,
			formData.Phone,
			formData.Created,
		)
	}
	html += "</table>"
	//common.EmailCh <- &common.EmailContent{
	//	"山西转型综合改革示范区金融服务平台-企业融资需求表",
	//	html,
	//}
}


