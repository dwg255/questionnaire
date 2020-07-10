<template>
  <div id="container" class="page form_page js_show">
    <div class="weui-form">
        <div class="weui-form__text-area" id="header">
            <h2 class="weui-form__title">企业融资需求表</h2>
            <img src="../assets/banner.png" alt="">
        </div>
        <div id="description">
            <span>&nbsp;&nbsp;为了更好的为您提供金融服务，请认真填写《企业融资需求调查表》，保证数据的真实及完整性！</span>
        </div>
        <div id="question">
            <div class="field" :class="err.companyName != '' ? 'has-error' : ''">
                <div class="field-label">
                    1. 企业名称
                    <span class="req">*</span>
                </div>
                <div class="field-input">
                    <div class="weui-cell__bd">
                        <!-- <input name="company-name" v-model="formData.companyName" class="weui-input" placeholder=""> -->
                        <input name="company-name" v-validate="'required|min:3|max:10'" v-model="formData.companyName" class="weui-input" placeholder="">
                    </div>
                </div>
                <transition name="fade">
                    <div class="err-msg" v-if="err.companyName">
                        {{err.companyName}}
                    </div>
                </transition>
            </div>
            <div class="field"  :class="err.registerTime != '' ? 'has-error' : ''">
                <div class="field-label">
                    2. 注册时间
                    <span class="req">*</span>
                </div>
                <div class="field-input">
                    <div class="weui-cell__bd" id="showDatePicker" v-on:click="showDatePicker">
                        <input readonly name="register-time" v-model="formData.registerTime" class="weui-input" placeholder="">
                        <img src="../assets/date@2x.png" alt="">
                    </div>
                </div>
                <transition name="fade">
                    <div class="err-msg" v-if="err.registerTime">
                        {{err.registerTime}}
                    </div>
                </transition>
            </div>
            <div class="field"  :class="err.mainBusiness != '' ? 'has-error' : ''">
                <div class="field-label">
                    3. 主营业务
                    <span class="req">*</span>
                </div>
                <div class="field-input">
                    <div class="weui-cell__bd">
                        <input name="main-bussiness" v-model="formData.mainBusiness" class="weui-input" placeholder="">
                    </div>
                </div>
                <transition name="fade">
                    <div class="err-msg" v-if="err.mainBussiness">
                        {{err.mainBussiness}}
                    </div>
                </transition>
                
            </div>
            <div class="field"  :class="err.registeredCapital != '' ? 'has-error' : ''">
                <div class="field-label">
                    4. 注册资本（单位/万元）
                    <span class="req">*</span>
                </div>
                <div class="field-input">
                    <div class="weui-cell__bd">
                        <input type="number" v-model="formData.registeredCapital" name="regist-capital" class="weui-input" placeholder="">
                    </div>
                </div>
                <transition name="fade">
                    <div class="err-msg" v-if="err.registCapital">
                        {{err.registCapital}}
                    </div>
                </transition>
                
            </div>
            <div class="field"  :class="err.paidCapital != '' ? 'has-error' : ''">
                <div class="field-label">
                    5. 实收资本（单位/万元）
                    <span class="req">*</span>
                </div>
                <div class="field-input">
                    <div class="weui-cell__bd">
                        <input type="number" v-model="formData.paidCapital" name="paid-capital" class="weui-input" placeholder="">
                    </div>
                </div>
                <transition name="fade">
                    <div class="err-msg" v-if="err.paidCapital">
                        {{err.paidCapital}}
                    </div>
                </transition>
                
            </div>
            <div class="field"  :class="err.taxLevel != '' ? 'has-error' : ''">
                <div class="field-label">
                    6. 纳税等级
                    <span class="req">*</span>
                </div>
                <div class="field-input">
                    <div class="weui-cell__bd" @click="taxLevel">
                        <input readonly name="tax-level" v-model="formData.taxLevel" class="weui-input" placeholder="A" value="A">
                    </div>
                </div>
                <transition name="fade">
                    <div class="err-msg" v-if="err.taxLevel">
                        {{err.taxLevel}}
                    </div>
                </transition>
                
            </div>
            <div class="field"  :class="err.taxAmount != '' ? 'has-error' : ''">
                <div class="field-label">
                    7. 上一年纳税总额（单位/万元）
                    <span class="req">*</span>
                </div>
                <div class="field-input">
                    <div class="weui-cell__bd">
                        <input type="number" v-model="formData.taxAmount" name="tax-amount" class="weui-input" placeholder="">
                    </div>
                </div>
                <transition name="fade">
                    <div class="err-msg" v-if="err.taxAmount">
                        {{err.taxAmount}}
                    </div>
                </transition>
                
            </div>
            <div class="field"  :class="err.assets != '' ? 'has-error' : ''">
                <div class="field-label">
                    8. 上一年度总资产（单位/万元）
                    <span class="req">*</span>
                </div>
                <div class="field-input">
                    <div class="weui-cell__bd">
                        <input type="number" v-model="formData.assets" name="assets" class="weui-input" placeholder="">
                    </div>
                </div>
                <transition name="fade">
                    <div class="err-msg" v-if="err.assets">
                        {{err.assets}}
                    </div>
                </transition>
                
            </div>
            <div class="field"  :class="err.liabilities != '' ? 'has-error' : ''">
                <div class="field-label">
                    9. 上一年度总负债（单位/万元）
                    <span class="req">*</span>
                </div>
                <div class="field-input">
                    <div class="weui-cell__bd">
                        <input type="number" v-model="formData.liabilities" name="liabilities" class="weui-input" placeholder="">
                    </div>
                </div>
                <transition name="fade">
                    <div class="err-msg" v-if="err.liabilities">
                        {{err.liabilities}}
                    </div>
                </transition>
                
            </div>
            <div class="field"  :class="err.loan != '' ? 'has-error' : ''">
                <div class="field-label">
                    10. 银行贷款（单位/万元）
                    <span class="req">*</span>
                </div>
                <div class="field-input">
                    <div class="weui-cell__bd">
                        <input type="number" v-model="formData.loan" name="loan" class="weui-input" placeholder="">
                    </div>
                </div>
                <transition name="fade">
                    <div class="err-msg" v-if="err.loan">
                        {{err.loan}}
                    </div>
                </transition>
                
            </div>
            <div class="field"  :class="err.income != '' ? 'has-error' : ''">
                <div class="field-label">
                    11. 上一年度营业收入（单位/万元）
                    <span class="req">*</span>
                </div>
                <div class="field-input">
                    <div class="weui-cell__bd">
                        <input type="number" v-model="formData.income" name="income" class="weui-input" placeholder="">
                    </div>
                </div>
                <transition name="fade">
                    <div class="err-msg" v-if="err.income">
                        {{err.income}}
                    </div>
                </transition>
                
            </div>
            <div class="field"  :class="err.profit != '' ? 'has-error' : ''">
                <div class="field-label">
                    12. 上一年度净利润（单位/万元）
                    <span class="req">*</span>
                </div>
                <div class="field-input">
                    <div class="weui-cell__bd">
                        <input type="number" v-model="formData.profit" name="profit" class="weui-input" placeholder="">
                    </div>
                </div>
                <transition name="fade">
                     <div class="err-msg" v-if="err.profit">
                        {{err.profit}}
                    </div>
                </transition>
               
            </div>
            <div class="field"  :class="err.account != '' ? 'has-error' : ''">
                <div class="field-label">
                    13. 拟融资额度（单位/万元）
                    <span class="req">*</span>
                </div>
                <div class="field-input">
                    <div class="weui-cell__bd">
                        <input type="number" v-model="formData.account" name="account" class="weui-input" placeholder="">
                    </div>
                </div>
                <transition name="fade">
                    <div class="err-msg" v-if="err.account">
                        {{err.account}}
                    </div>
                </transition>
                
            </div>

            <div class="field"  :class="err.purpose != '' ? 'has-error' : ''">
                <div class="field-label">
                    14. 资金用途 <span class="qtypetip">【多选】</span>
                    <span class="req">*</span>
                </div>
                <div class="field-input">
                    <div class="weui-cells__group weui-cells__group_form">
                        <div class="weui-cells weui-cells_checkbox">

                          <label v-for='(item,index) in formData.purpose' :key="index" class="weui-cell weui-cell_active weui-check__label" :for="item.label">
                            <div class="weui-cell__hd">
                                <input type="checkbox" class="weui-check" name="purpose" :id="item.label" v-model="item.checked">
                                <i class="weui-icon-checked"></i>
                            </div>
                            <div class="weui-cell__bd">
                                <p>{{item.content}}</p>
                            </div>
                          </label>
                        </div>
                    </div>
                </div>
                <transition name="fade">
                    <div class="err-msg" v-if="err.purpose">
                        {{err.purpose}}
                    </div>
                </transition>
                
            </div>
            <div class="field"  :class="err.guarantee != '' ? 'has-error' : ''">
                <div class="field-label">
                    15. 担保方式 <span class="qtypetip">【多选】</span>
                    <span class="req">*</span>
                </div>
                <div class="field-input">
                    <div class="weui-cells__group weui-cells__group_form">
                      <div class="weui-cells weui-cells_checkbox">
                        <label v-for='(item,index) in formData.guarantee' :key="index" class="weui-cell weui-cell_active weui-check__label" :for="item.label">
                            <div class="weui-cell__hd">
                                <input type="checkbox" class="weui-check" name="guarantee" :id="item.label" v-model="item.checked">
                                <i class="weui-icon-checked"></i>
                            </div>
                            <div class="weui-cell__bd">
                                <p>{{item.content}}</p>
                            </div>
                        </label>
                          <!-- <div class="field-input">
                              <div class="weui-cell__bd">
                                  <input name="contact" v-model="formData.contact" class="weui-input" placeholder="">
                              </div>
                          </div> -->
                        </div>
                    </div>
                </div>
                <transition name="fade">
                    <div class="err-msg" v-if="err.guarantee">
                        {{err.guarantee}}
                    </div>
                </transition>
                
            </div>
            <div class="field"  :class="err.contact != '' ? 'has-error' : ''">
                <div class="field-label">
                    16. 企业联系人
                    <span class="req">*</span>
                </div>
                <div class="field-input">
                    <div class="weui-cell__bd">
                        <input name="contact" v-model="formData.contact" class="weui-input" placeholder="">
                    </div>
                </div>
                <transition name="fade">
                    <div class="err-msg" v-if="err.contact">
                        {{err.contact}}
                    </div>
                </transition>
                
            </div>
            <div class="field"  :class="err.phone != '' ? 'has-error' : ''">
                <div class="field-label">
                    17. 联系电话
                    <span class="req">*</span>
                </div>
                <div class="field-input">
                    <div class="weui-cell__bd">
                        <input type="number" v-model="formData.phone" name="phone" class="weui-input" placeholder="">
                    </div>
                </div>
                <transition name="fade">
                    <div class="err-msg" v-if="err.phone">
                        {{err.phone}}
                    </div>
                </transition>
                
            </div>

        </div>
        <div class="weui-form__opr-area">
            <a class="weui-btn weui-btn_primary" href="javascript:" id="submit" @click="submit">确定</a>
        </div>
    </div>
    <div class="footer">
        <span>
            山西转型综合改革示范区金融服务平台
        </span>
    </div>
  </div>
</template>

<script>
import axios from  "axios"
import qs from "qs"
export default {
  name: 'HelloWorld',
  props: {
    
  },
  data(){
    return {
      // bannerImg:require('../assets/banner.png'),
      // dateImg:require('../assets/date@2x.png'),
      formData:{
        companyName:"",//1. 企业名称
        registerTime:"",//2. 注册时间
        mainBusiness:"",//3. 主营业务
        registeredCapital:"",// 4. 注册资本（单位/万元）
        paidCapital:"",// 5. 实收资本（单位/万元）
        taxLevel:"A",//6. 纳税等级
        taxAmount:"",//7. 上一年纳税总额（单位/万元）
        assets:"",//8. 上一年度总资产（单位/万元）
        liabilities:"",//9. 上一年度总负债（单位/万元）
        loan:"",//10. 银行贷款（单位/万元）
        income:"",//11. 上一年度营业收入（单位/万元）
        profit:"",//12. 上一年度净利润（单位/万元）
        account:"",//13. 拟融资额度（单位/万元）
        purpose:[
            {label:"p1",content:"流动资金贷款",checked:true,},
            {label:"p2",content:"固定资产贷款",checked:false,},
            {label:"p3",content:"项目贷款",checked:false,},
            {label:"p4",content:"其他",checked:false,},
        ],// 14. 资金用途
        guarantee:[
          {label:"s1",content:"房产抵押",checked:true,},
          {label:"s2",content:"土地抵押",checked:false},
          {label:"s3",content:"股权质押",checked:false},
          {label:"s4",content:"存货质押",checked:false},
          {label:"s5",content:"存单质押",checked:false},
          {label:"s6",content:"股权质押",checked:false},
          {label:"s7",content:"订单融资",checked:false},
          {label:"s8",content:"知识产权质押",checked:false},
          {label:"s9",content:"应收账款质押",checked:false},
          {label:"s10",content:"保证",checked:false},
          {label:"s11",content:"信用",checked:false},
          {label:"s12",content:"其他",checked:false},
        ],//15. 担保方式
        contact:"",//16. 企业联系人
        phone:"",// 17 联系电话
      },
      err:{
        companyName: "",//1. 企业名称
        registerTime:"",//2. 注册时间
        mainBusiness:"",//3. 主营业务
        registeredCapital:"",// 4. 注册资本（单位/万元）
        paidCapital:"",// 5. 实收资本（单位/万元）
        taxLevel:"",//6. 纳税等级
        taxAmount:"",//7. 上一年纳税总额（单位/万元）
        assets:"",//8. 上一年度总资产（单位/万元）
        liabilities:"",//9. 上一年度总负债（单位/万元）
        loan:"",//10. 银行贷款（单位/万元）
        income:"",//11. 上一年度营业收入（单位/万元）
        profit:"",//12. 上一年度净利润（单位/万元）
        account:"",//13. 拟融资额度（单位/万元）
        purpose:"",// 14. 资金用途
        guarantee:"",//15. 担保方式
        contact:"",//16. 企业联系人
        phone:"",// 17 联系电话
      },
    }
  },
  methods:{
      isPhone(phoneNumber){
        let checkTel = (s)=>{
            if(!/^((0\d{2,3})-)?(\d{7,8})(-(\d{3,}))?$/.test(s))return   false  ;
            return   true  ;
        }
        let checkPhone = (s)=>{ 
            if(!(/^1[3456789]\d{9}$/.test(s)))return false; 
            return true;
        }
        if (checkTel(phoneNumber) || checkPhone(phoneNumber)) return true;
        return false
      },
        isNumber(val) {
            var regPos = /^\d+(\.\d+)?$/;//非负浮点数
            var regNeg = /^(-(([0-9]+\.[0-9]*[1-9][0-9]*)|([0-9]*[1-9][0-9]*\.[0-9]+)|([0-9]*[1-9][0-9]*)))$/; //负浮点数
            if (regPos.test(val) || regNeg.test(val)) {
                return true;
            } else {
                return false;
            }
        },
      //验证表单值
      validate(){
          return {
            companyName:()=>{this.err.companyName = this.formData.companyName == ""  ?  "请输入企业名称" : '';},
            registerTime:()=>{this.err.registerTime = this.formData.registerTime == ""  ?  "请选择注册时间": '';},
            mainBusiness:()=>{this.err.mainBusiness = this.formData.mainBusiness == ""  ?  "请输入主营业务": '';},
            registeredCapital:()=>{
                this.err.registeredCapital = this.formData.registeredCapital == ""  ?  "请输入注册资本": '';
                // this.err.registCapital = this.isNumber(this.formData.registCapital) ? "" : "请输入数值";
            },
            paidCapital:()=>{this.err.paidCapital = this.formData.paidCapital == ""  ?  "请输入实收资本": '';},
            taxAmount:()=>{this.err.taxAmount = this.formData.taxAmount == ""  ?  "请输入上一年纳税总额": '';},
            assets:()=>{this.err.assets = this.formData.assets == ""  ?  "请输入上一年度总资产": '';},
            liabilities:()=>{this.err.liabilities = this.formData.liabilities == ""  ?  "请输入上一年度总负债": '';},
            loan:()=>{this.err.loan = this.formData.loan == ""  ?  "请输入银行贷款": '';},
            income:()=>{this.err.income = this.formData.income == ""  ?  "请输入上一年度营业收入": '';},
            profit:()=>{this.err.profit = this.formData.profit == ""  ?  "请输入上一年度净利润": '';},
            account:()=>{this.err.account = this.formData.account == ""  ?  "请输入拟融资额度": '';},
            purpose:()=>{
                console.log("purpose validate...")
                let flag = false;
                this.formData.purpose.forEach(element => {
                    if (element.checked) flag = true;
                });
                this.err.purpose = flag ? '' : "请选择资金用途";
            },
             guarantee:()=>{
                let flag = false;
                this.formData.guarantee.forEach(element => {
                    if (element.checked) flag = true;
                });
                this.err.guarantee = flag ? '' : "请选择担保方式" ;},
            contact:()=>{this.err.contact = this.formData.contact == ""  ?  "请输入企业联系人": '';},
            phone:()=>{
                this.err.phone = this.formData.phone == ""  ?  "请输入联系电话": '';
                // this.err.phone = this.isPhone(this.formData.phone) ?  "": '请输入有效联系电话';
                }
          }
      },
      //显示日期控件
    showDatePicker(){
      this.$weui.datePicker({
            start: 1990,
            end: new Date().getFullYear(),
            // onChange: function(result) {
            //     console.log(result);
            // },
            onConfirm: (result)=> {
                let r = result[0].value + "-" + result[1].value + "-" + result[2].value;
                // console.log(r)
                // console.log(el)
                this.formData.registerTime = r;
            },
            title: '注册时间'
        });
    },
    //纳税等级
    taxLevel(){
      // console.log(this)
       this.$weui.picker([{
                label: 'A',
                value: 0
            }, {
                label: 'B',
                value: 1
            }, {
                label: 'C',
                value: 2
            }, {
                label: 'D',
                value: 3
            }, {
                label: 'M',
                value: 4
            }], {
                onConfirm: (result) =>{
                    console.log(this);
                    this.formData.taxLevel = result[0].label;
                },
                title: '纳税登记'
            });
    },
    //提交表单
    submit(){
        for (let key in this.validate()) {
            this.validate()[key]()
        }
        this.err.phone = this.isPhone(this.formData.phone) ?  "": '请输入有效联系电话';
        let hasError = false
        for(let key  in this.err){
            if (this.err[key] != "" ) {
                hasError =  true;
            } 
        }
        if (hasError) {
           setTimeout(() => {
                window.scrollTo({ 
                    top: document.getElementsByClassName("has-error")[0].offsetTop, 
                    behavior: "smooth" 
                }); 
           }, 200);
        } else {
            //提交
            // console.log(JSON.stringify(this.formData));
            let data = JSON.parse(JSON.stringify(this.formData))
            data.purpose = JSON.stringify(data.purpose)
            data.guarantee = JSON.stringify(data.guarantee)
            let vue = this
            axios.post('http://localhost:8081/finance', qs.stringify(data)).then(function (response) {
                // console.log(response)
                if(response.status == 200 ){
                    vue.$router.push({path:'/success'})
                }
            }).catch(function (error) {
                console.log(error);
            });
        }
    },
 
  },
  watch:{
    "formData.companyName": {
        handler() {
            this.validate().companyName()
        },
        deep: true
    },
    "formData.registerTime": {
        handler() {
            this.validate().registerTime()
        },
        deep: true
    },
    "formData.mainBusiness": {
        handler() {
            this.validate().mainBusiness()
        },
        deep: true
    },
    "formData.registeredCapital": {
        handler() {
            this.validate().registeredCapital()
        },
        deep: true
    },
    "formData.paidCapital": {
        handler() {
            this.validate().paidCapital()
        },
        deep: true
    },
    "formData.taxAmount": {
        handler() {
            this.validate().taxAmount()
        },
        deep: true
    },
    "formData.assets": {
        handler() {
            this.validate().assets()
        },
        deep: true
    },
    "formData.liabilities": {
        handler() {
            this.validate().liabilities()
        },
        deep: true
    },
    "formData.loan": {
        handler() {
            this.validate().loan()
        },
        deep: true
    },
    "formData.income": {
        handler() {
            this.validate().income()
        },
        deep: true
    },
    "formData.profit": {
        handler() {
            this.validate().profit()
        },
        deep: true
    },
    "formData.account": {
        handler() {
            this.validate().account()
        },
        deep: true
    },
    "formData.purpose": {
         handler() {
             this.validate().purpose()
        },
        deep: true
    },
    "formData.guarantee": {
        handler() {
            this.validate().guarantee()
        },
        deep: true
    },
    "formData.contact": {
        handler() {
            this.validate().contact()
        },
        deep: true
    },
    "formData.phone": {
        handler() {
            this.validate().phone()
        },
        deep: true
    },
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="less">
.fade-enter-active, .fade-leave-active {
      transition: opacity .5s
}
.fade-enter, .fade-leave-active {
      opacity: 0
}
#container{
    overflow: hidden;
    #description{
        margin: 0 10px 16px;
        padding: 10px 0 26px;
        border-bottom: 1px dashed #ccc;
        font-size: 14px;
        line-height: 24px;
        color: #333;
    }
    #question{
        display: flex; 
        flex-direction: column;
        
        .field{
            display: flex;
            flex-direction: column;
            margin: 10px 6px;
            padding: 10px 14px;
            border: 1px solid #fff;
            &.has-error{   //表单错误
                border: 1px dashed rgb(222, 103, 82);;
            }
            .field-label{
                position: relative;
                font-size: 16px;
                font-weight: 700;
                padding: 0 0 8px;
                .req{
                    position: absolute;
                    left: -12px;
                    color: red;
                }
            }   
            .field-input{
                display: flex;
                margin: 5px 0;
                border: 1px solid #e0e0e0;
                margin: 5px 0;
                background-color: #fff;
                padding: 0;
                border-radius: 4px;
                input{
                    flex: 1;
                    padding: 10px;
                    width: 100%;
                    border-radius: 4px;
                    border: none!important;
                }
                &>div{
                    width: 100%;
                }
                
            }
            .err-msg{   //错误信息块    
                padding-left: 12px;
                color: Red;
                line-height: 38px;
                background-color: #ffe5e0;
                border-radius: 4px;
                font-size: 12px;
                margin-top: 9px;
                transition: all 1s ease 0s;
            }
            #showDatePicker{
                width: 100%;
                position: relative;
                input{
                    padding-left: 40px;
                }
                img{
                    position: absolute;
                    top: 12px;
                    left: 12px;
                    width: 18px;
                    height: 18px;
                }
            }
        }
        .qtypetip{
            color: #999;
            font-weight: normal;
        }
    }
    .weui-form {
        padding:  0;
    }
    
    #header {
        padding: 10px 0;
    }
    
    #header {
        padding-top: 0;
        .weui-form__title {
            background-color: skyblue;
            padding: 5px 0;
            color: #fff;
        }
    }
    #header {
        img{
            width: 100%;
        }
    }
    .weui-form__opr-area {
        margin-bottom: 50px!important;
    }
    .footer{
        margin: 0 auto;
        background-color: #f0f0f0;
        height: 50px;
        display: flex;
        align-items: center;
        justify-content: center;
        flex-direction: column;
        text-align: center;
        color: #999;
        font-size: 14px;
        span{
            font-size: 14px;
            color: #999;
        }
    }
}

</style>
