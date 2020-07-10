(function() {
    //贷款额度
    document.getElementById('tax-level').addEventListener("click",
        function() {
            weui.picker([{
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
                onConfirm: function(result) {
                    console.log(result);
                    document.getElementById('tax-level').innerHTML = result[0].label;
                },
                title: '纳税登记'
            });
        })

    //提交
    document.getElementById('submit').addEventListener("click", function() {
        var company_name = document.getElementById('company_name').value;
        var contact = document.getElementById('contact').value;
        var phone = document.getElementById('phone').value;

        if (company_name == "" || contact == "" || phone == "") {
            weui.dialog({
                title: '请完善调查表',
                // content: 'dialog内容',
                className: 'custom-classname',
                buttons: [{
                    label: '确定',
                    type: 'primary',
                    // onClick: function() {
                    //     alert('确定')
                    // }
                }]
            });
            return false;
        } else {
            var amount = document.getElementById("amount").innerHTML;
            var expire = document.getElementById("expire").innerHTML;
            var guarantee = document.getElementById("guarantee").innerHTML;
            // var url = "http://www.sfqjr.com/financial/frontAPI/front/news/listNewsType";
            var url = "http://localhost:8080/question";
            var data = {
                company_name,
                contact,
                phone,
                amount,
                expire,
                guarantee,
            };
            console.log(data);
            // return false;
            var xhr = new XMLHttpRequest();
            xhr.open("POST", url, true);
            xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
            xhr.onreadystatechange = function() {
                if (xhr.readyState == 4 && (xhr.status == 200 || xhr.status == 304)) {
                    // console.log(123)
                    window.location.href = "success.html";
                }
            };

            var str = [];
            for (var p in data)
                if (data.hasOwnProperty(p)) {
                    str.push(encodeURIComponent(p) + "=" + encodeURIComponent(data[p]));
                }
            data = str.join("&");

            xhr.send(data);

            // document.getElementById("js_toast").style.display = "block";
            // setTimeout(function() {
            //     document.getElementById('js_toast').style.opacity = '1';
            // }, 0)
            // setTimeout(() => {
            //     document.getElementById("js_toast").style.display = "none";
            //     setTimeout(function() {
            //         document.getElementById('js_toast').style.opacity = '0';
            //     }, 0)
            // }, 2000);
        }
    })
}())