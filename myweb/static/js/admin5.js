var page = 1
$.ajax({
    type: "POST",
    url: 'http://localhost:8081/api/adminfindclub',
    dataType: "json",
    data: {},
    success: function (res) {
        if (res.code == 1) {
            alert(JSON.stringify(res.message));
            return;
        }
        $("#ul1").append("<li><a name='1' onclick='selectpage(this)'>&laquo;</a></li><li class='active'><a>1</a></li>")
        if (res.totalpages - res.page > 4) {
            for (var i = 2; i <= 5; i++) {
                $("#ul1").append("<li><a name='" + i + "' onclick='selectpage(this)'>" + i + "</a></li>")
            }
        }
        else {
            for (var i = res.page + 1; i <= res.totalpages; i++) {
                $("#ul1").append("<li><a name='" + i + "' onclick='selectpage(this)'>" + i + "</a></li>")
            }
        }
        $("#ul1").append("<li><a name='" + res.totalpages + "' onclick='selectpage(this)'>&raquo;</a></li>")
        if (res.change == 1) {
            $("#div1").append("<label>当前状态：干部换届已关闭</label></br><label id = 'labal1'></label></br><button type='button' class='btn btn-success' onclick='start()'>继续上次报名</button> <button  type='button' class='btn btn-danger'onclick='restart()'>发起新的报名</button>");
            refreshpage();
        }
        if (res.change == 2) {
            $("#div1").append("<label>当前状态：正在进行干部换届报名</label></br><label id = 'labal1'></label></br><button type='button' class='btn btn-danger' onclick='stop()'>关闭报名</button>");
            refreshpage();
        }
        $("#welcome").html("欢迎" + $.cookie('username') + "登录Tkkc社团组织管理系统");
    },
    error: function (e) {
        alert("请求失败");
    }
});
function rTime(date) {
    var json_date = new Date(date).toJSON();
    return new Date(new Date(json_date) + 8 * 3600 * 1000).toISOString().replace(/T/g, ' ').replace(/\.[\d]{3}Z/, '')
}
function resume(val) {
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/findresume',
        dataType: "json",
        data: { username: val.name },
        success: function (res) {
            $("#resume_name").text(res.name)
            $("#resume_sex").text(res.sex)
            $("#resume_studentname").text(res.student_number)
            $("#resume_major").text(res.major)
            $("#resume_outlook").text(res.political_outlook)
            $("#resume_number").text(res.phone_number)
            $("#resume_email").text(res.email)
            $("#resume_introduce").text(res.introduction)
            $("#resume_birthday").text(res.birthday)
            $("#myModal5").modal("show")
        },
        error: function (e) {
            alert("请求失败");
        }
    });
}
function start() {
    if (confirm('确认继续换届报名 ？') == false) {
        return
    }
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/admineditclub',
        dataType: "json",
        data: { recruit: 0 ,change: 2},
        success: function (res) {
            if(res.code==1){
                alert("开始失败")
                return
            }
            alert("成功开始");
            location.reload();
        },
        error: function (e) {
            alert("请求失败");
        }
    });
}
function restart() {
    if (confirm('确认重新开始换届报名？ 这将删除以往的纳新数据！ ') == false) {
        return
    }
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/admindelchangeenroll',
        dataType: "json",
        data: { },
        success: function (res) {
            if(res.code==1){
                alert(res.message)
                return
            }
        },
        error: function (e) {
            alert("请求失败");
        }
    });
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/admineditclub',
        dataType: "json",
        data: {  recruit: 0 ,change: 2},
        success: function (res) {
            if(res.code==1){
                alert("发起失败")
                return
            }
            alert("发起成功")
            location.reload();
        },
        error: function (e) {
            alert("请求失败");
        }
    });
}
function stop() {
    if (confirm('确认停止换届报名 ？') == false) {
        return
    }
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/admineditclub',
        dataType: "json",
        data: {  recruit: 0 ,change: 1},
        success: function (res) {
            if(res.code==1){
                alert("停止失败")
                return
            }
            alert("停止成功");
            location.reload();
           
        },
        error: function (e) {
            alert("请求失败");
        }
    });
}
function selectpage(val) {
    page = val.name
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/adminfindchangeenroll',
        dataType: "json",
        data: { page: val.name, pagesize: 10 },
        success: function (res) {
            if (res.code == 1) {
                alert(JSON.stringify(res.message));
                return;
            }
            $("#labal1").empty();
            $("#labal1").append("已报名人数：",res.total);
            $("#ul1").empty();
            if (res.page == 1) {
                $("#ul1").append("<li><a name='1' onclick='selectpage(this)'>&laquo;</a></li><li class='active'><a>1</a></li>")
                if (res.totalpages - res.page > 4) {
                    for (var i = 2; i <= 5; i++) {
                        $("#ul1").append("<li><a name='" + i + "' onclick='selectpage(this)'>" + i + "</a></li>")
                    }
                }
                else {
                    for (var i = res.page + 1; i <= res.totalpages; i++) {
                        $("#ul1").append("<li><a name='" + i + "' onclick='selectpage(this)'>" + i + "</a></li>")
                    }
                }
                $("#ul1").append("<li><a name='" + res.totalpages + "' onclick='selectpage(this)'>&raquo;</a></li>")
            }
            else if (res.page == 2) {
                $("#ul1").append("<li><a name='1' onclick='selectpage(this)'>&laquo;</a></li><li><a onclick='selectpage(this)' name='1'>1</a></li><li class='active'><a name='2' onclick='selectpage(this)'>2</a></li>")
                if (res.totalpages - res.page > 3) {
                    for (var i = 3; i <= 5; i++) {
                        $("#ul1").append("<li><a name='" + i + "' onclick='selectpage(this)'>" + i + "</a></li>")
                    }
                }
                else {
                    for (var i = res.page + 1; i <= res.totalpages; i++) {
                        $("#ul1").append("<li><a name='" + i + "' onclick='selectpage(this)'>" + i + "</a></li>")
                    }
                }
                $("#ul1").append("<li><a name='" + res.totalpages + "' onclick='selectpage(this)'>&raquo;</a></li>")
            }
            else {
                $("#ul1").append("<li><a name='1' onclick='selectpage(this)'>&laquo;</a></li><li><a onclick='selectpage(this)' name='" + (res.page - 2) + "'>" + (res.page - 2) + "</a></li><li><a name='" + (res.page - 1) + "' onclick='selectpage(this)'>" + (res.page - 1) + "</a></li><li class='active'><a name='" + res.page + "' onclick='selectpage(this)'>" + res.page + "</a></li>")
                if (res.totalpages - res.page > 2) {
                    $("#ul1").append("<li><a name='" + (res.page + 1) + "' onclick='selectpage(this)' >" + (res.page + 1) + "</a></li><li><a name='" + (res.page + 2) + "' onclick='selectpage(this)' >" + (res.page + 2) + "</a></li>")
                }
                else {
                    for (var i = res.page + 1; i <= res.totalpages; i++) {
                        $("#ul1").append("<li><a name='" + i + "' onclick='selectpage(this)'>" + i + "</a></li>")
                    }
                }
                $("#ul1").append("<li><a name='" + res.totalpages + "' onclick='selectpage(this)'>&raquo;</a></li>")
            }
            $("#tbody1").empty();
            for (var i = 0; i < res.data.length; i++) {
                $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].username + "</td><td>" + res.data[i].role + "</td><td>" + res.data[i].name + "</td><td>" + res.data[i].sex + "</td><td>" + res.data[i].student_number + "</td><td>" + res.data[i].major + "</td><td>" + res.data[i].phone_number + "</td><td>" + res.data[i].email + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='resume(this)' name='" + res.data[i].username + "' class='btn btn-default'>简历详情</a></div></td></tr>");
            }
        },
        error: function (e) {
            alert("请求失败");
        }
    });
}
function refreshpage() {
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/adminfindchangeenroll',
        dataType: "json",
        data: { page: page, pagesize: 10},
        success: function (res) {
            if (res.code == 1) {
                alert(JSON.stringify(res.message));
                return;
            }
            $("#labal1").empty();
            $("#labal1").append("已报名人数：",res.total);
            $("#ul1").empty();
            if (res.page == 1) {
                $("#ul1").append("<li><a name='1' onclick='selectpage(this)'>&laquo;</a></li><li class='active'><a>1</a></li>")
                if (res.totalpages - res.page > 4) {
                    for (var i = 2; i <= 5; i++) {
                        $("#ul1").append("<li><a name='" + i + "' onclick='selectpage(this)'>" + i + "</a></li>")
                    }
                }
                else {
                    for (var i = res.page + 1; i <= res.totalpages; i++) {
                        $("#ul1").append("<li><a name='" + i + "' onclick='selectpage(this)'>" + i + "</a></li>")
                    }
                }
                $("#ul1").append("<li><a name='" + res.totalpages + "' onclick='selectpage(this)'>&raquo;</a></li>")
            }
            else if (res.page == 2) {
                $("#ul1").append("<li><a name='1' onclick='selectpage(this)'>&laquo;</a></li><li><a onclick='selectpage(this)' name='1'>1</a></li><li class='active'><a name='2' onclick='selectpage(this)'>2</a></li>")
                if (res.totalpages - res.page > 3) {
                    for (var i = 3; i <= 5; i++) {
                        $("#ul1").append("<li><a name='" + i + "' onclick='selectpage(this)'>" + i + "</a></li>")
                    }
                }
                else {
                    for (var i = res.page + 1; i <= res.totalpages; i++) {
                        $("#ul1").append("<li><a name='" + i + "' onclick='selectpage(this)'>" + i + "</a></li>")
                    }
                }
                $("#ul1").append("<li><a name='" + res.totalpages + "' onclick='selectpage(this)'>&raquo;</a></li>")
            }
            else {
                $("#ul1").append("<li><a name='1' onclick='selectpage(this)'>&laquo;</a></li><li><a onclick='selectpage(this)' name='" + (res.page - 2) + "'>" + (res.page - 2) + "</a></li><li><a name='" + (res.page - 1) + "' onclick='selectpage(this)'>" + (res.page - 1) + "</a></li><li class='active'><a name='" + res.page + "' onclick='selectpage(this)'>" + res.page + "</a></li>")
                if (res.totalpages - res.page > 2) {
                    $("#ul1").append("<li><a name='" + (res.page + 1) + "' onclick='selectpage(this)' >" + (res.page + 1) + "</a></li><li><a name='" + (res.page + 2) + "' onclick='selectpage(this)' >" + (res.page + 2) + "</a></li>")
                }
                else {
                    for (var i = res.page + 1; i <= res.totalpages; i++) {
                        $("#ul1").append("<li><a name='" + i + "' onclick='selectpage(this)'>" + i + "</a></li>")
                    }
                }
                $("#ul1").append("<li><a name='" + res.totalpages + "' onclick='selectpage(this)'>&raquo;</a></li>")
            }
            $("#tbody1").empty();
            for (var i = 0; i < res.data.length; i++) {
                $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].username + "</td><td>" + res.data[i].role + "</td><td>" + res.data[i].name + "</td><td>" + res.data[i].sex + "</td><td>" + res.data[i].student_number + "</td><td>" + res.data[i].major + "</td><td>" + res.data[i].phone_number + "</td><td>" + res.data[i].email + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='resume(this)' name='" + res.data[i].username + "' class='btn btn-default'>简历详情</a></div></td></tr>");
            }
        },
        error: function (e) {
            alert("请求失败");
        }
    });
}
function clearsearch() {
    page = 1;
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/adminfindchangeenroll',
        dataType: "json",
        data: { page: 1, pagesize: 10},
        success: function (res) {
            if (res.code == 1) {
                alert(JSON.stringify(res.message));
                return;
            }
            $("#labal1").empty();
            $("#labal1").append("已报名人数：",res.total);
            $("#ul1").empty();
            $("#ul1").append("<li><a name='1' onclick='selectpage(this)'>&laquo;</a></li><li class='active'><a>1</a></li>")
            if (res.totalpages - res.page > 4) {
                for (var i = 2; i <= 5; i++) {
                    $("#ul1").append("<li><a name='" + i + "' onclick='selectpage(this)'>" + i + "</a></li>")
                }
            }
            else {
                for (var i = res.page + 1; i <= res.totalpages; i++) {
                    $("#ul1").append("<li><a name='" + i + "' onclick='selectpage(this)'>" + i + "</a></li>")
                }
            }
            $("#ul1").append("<li><a name='" + res.totalpages + "' onclick='selectpage(this)'>&raquo;</a></li>")
            $("#tbody1").empty();
            for (var i = 0; i < res.data.length; i++) {
                    $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].username + "</td><td>" + res.data[i].role + "</td><td>" + res.data[i].name + "</td><td>" + res.data[i].sex + "</td><td>" + res.data[i].student_number + "</td><td>" + res.data[i].major + "</td><td>" + res.data[i].phone_number + "</td><td>" + res.data[i].email + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='resume(this)' name='" + res.data[i].username + "' class='btn btn-default'>简历详情</a></div></td></tr>");
            }
        },
        error: function (e) {
            alert("请求失败");
        }
    });
}
function logout() {
    if (confirm("确定退出？") == false) {
        return
    }
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/logout',
        dataType: "json",
        data: { role: "admin" },
        success: function (res) {
            if (res.code == 1) {
                alert(JSON.stringify(res.message));
                return;
            }
            window.location.href = "http://localhost:8081";
        },
        error: function (e) {
            alert("请求失败");
        }
    });
}
$(document).ready(function () {
    $('#ipwd').on('input propertychange', function () {
        var pwd = $.trim($(this).val());
        var rpwd = $.trim($("#i2pwd").val());
        if (rpwd != "") {
            if (pwd == "" || rpwd == "") {
                $("#msg_pwd").html("<font color='red'>密码不能为空</font>");
                $("#btn_changepwd").attr("disabled", true);
            }
            else {
                if (pwd == rpwd) {
                    $("#msg_pwd").html("<font color='green'>两次密码相同</font>");
                    $("btn_changepwd").attr("disabled", false);
                } else {
                    $("#msg_pwd").html("<font color='red'>两次密码不相同</font>");
                    $("#btn_changepwd").attr("disabled", true);
                }
            }
        }
    })
})
$(document).ready(function () {
    $('#i2pwd').on('input propertychange', function () {
        var pwd = $.trim($(this).val());
        var rpwd = $.trim($("#ipwd").val());
        if (pwd == "" || rpwd == "") {
            $("#msg_pwd").html("<font color='red'>密码不能为空</font>");
            $("#btn_changepwd").attr("disabled", true);
        }
        else {
            if (pwd == rpwd) {
                $("#msg_pwd").html("<font color='green'>两次密码相同</font>");
                $("#btn_changepwd").attr("disabled", false);
            } else {
                $("#msg_pwd").html("<font color='red'>两次密码不相同</font>");
                $("#btn_changepwd").attr("disabled", true);
            }
        }
    })

})
function changepassword() {
    var previousPassword = document.getElementById('oldpwd').value;
    var proposedPassword = document.getElementById('i2pwd').value
    document.getElementById('oldpwd').value = ""
    document.getElementById('ipwd').value = ""
    document.getElementById('i2pwd').value = ""
    $("#msg_pwd").html("");
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/adminchangepassword',
        dataType: "json",
        data: { previousPassword: previousPassword, proposedPassword: proposedPassword },
        success: function (res) {
            alert(JSON.stringify(res.message));
            return;
        },
        error: function (e) {
            alert("请求失败");
        }
    });
}
function editintriduce() {
    if ($("#introduce1").val()==""){
        alert("介绍不能为空")
        return;
    }
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/admineditclub',
        dataType: "json",
        data: { introduce: $("#introduce1").val(),recruit: 0, change: 0 },
        success: function (res) {
            alert(res.message);
        },
        error: function (e) {
            alert("请求失败");
        }
    });
    $("#introduce1").val("");
}
function exportexcel() {
    window.open("http://localhost:8081/api/changeenrolldownloadfile", '_blank');

}