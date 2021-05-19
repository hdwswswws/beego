var page = 1
$.ajax({
    type: "POST",
    url: 'http://localhost:8081/api/findclub',
    dataType: "json",
    data: { page: 1, pagesize: 10, clubname: "" },
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
        for (var i = 0; i < res.data.length; i++) {
            $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].club_name + "</td><td>" + res.data[i].introduce + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='editclubname(this)' name='" + res.data[i].club_name + "' class='btn btn-default'>修改名字</a><a onclick='del(this)' name='" + res.data[i].club_name + "' class='btn btn-danger'>删除</a></div></td></tr>");
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
function editclubname(val) {
    var clubname = prompt("新的组织名称：");
    if (clubname == null) {
        return
    }
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/editclubname',
        dataType: "json",
        data: { clubname: val.name, newclubname: clubname },
        success: function (res) {
            alert(JSON.stringify(res.message));
            refreshpage();
        },
        error: function (e) {
            alert("请求失败");
        }
    });

}
function del(val) {
    if (confirm('确认删除组织: ' + val.name + " ？") == false) {
        return
    }
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/delclub',
        dataType: "json",
        data: { clubname: val.name },
        success: function (res) {
            alert(JSON.stringify(res.message));
            refreshpage();
        },
        error: function (e) {
            alert("请求失败");
        }
    });
}
function selectpage(val) {
    var clubname = document.getElementById('clubname1').value;
    page = val.name
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/findclub',
        dataType: "json",
        data: { page: val.name, pagesize: 10, clubname: clubname },
        success: function (res) {
            if (res.code == 1) {
                alert(JSON.stringify(res.message));
                return;
            }
            $("#ul1").empty();
            if (res.code == 1) {
                alert(JSON.stringify(res.message));
                return;
            }
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
                $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].club_name + "</td><td>" + res.data[i].introduce + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='editclubname(this)' name='" + res.data[i].club_name + "' class='btn btn-default'>修改名字</a><a onclick='del(this)' name='" + res.data[i].club_name + "' class='btn btn-danger'>删除</a></div></td></tr>");
            }
        },
        error: function (e) {
            alert("请求失败");
        }
    });
}
function refreshpage() {
    var clubname = document.getElementById('clubname1').value;
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/findclub',
        dataType: "json",
        data: { page: page, pagesize: 10, clubname: clubname },
        success: function (res) {
            if (res.code == 1) {
                alert(JSON.stringify(res.message));
                return;
            }
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
                $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].club_name + "</td><td>" + res.data[i].introduce + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='editclubname(this)' name='" + res.data[i].club_name + "' class='btn btn-default'>修改名字</a><a onclick='del(this)' name='" + res.data[i].club_name + "' class='btn btn-danger'>删除</a></div></td></tr>");
            }
        },
        error: function (e) {
            alert("请求失败");
        }
    });
}
function createclub() {
    var file1 = $("#fileName1").val();//用户文件内容(文件)
    // 判断文件是否为空 
    if (file1 == "") {
        alert("请选择上传的目标文件! ")
        $("#fileName1").val("");
        return;
    }
    //判断文件类型
    var fileName1 = file1.substring(file1.lastIndexOf(".") + 1).toLowerCase();
    if (fileName1 != "jpg" && fileName1 != "png") {
        alert("请选择jpg或png格式!");
        $("#fileName1").val("");
        return;
    }
    //判断文件大小
    var size1 = $("#fileName1")[0].files[0].size;
    if (size1 > 10485760) {
        alert("上传文件不能大于10M!");
        $("#fileName1").val("");
        return;
    }
    var formData = new FormData();//这里需要实例化一个FormData来进行文件上传
    formData.append("fileName1", $("#fileName1")[0].files[0]);
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/createclub',
        dataType: "json",
        data: { clubname: $("#clubname").val(), introduce: $("#introduce").val() },
        success: function (res) {
            if (res.code == 1) {
                alert(JSON.stringify(res.message));
                return;
            }
            $.ajax({
                type: "post",
                url: "http://localhost:8081/api/upclublogo?clubname="+$("#clubname").val(),
                data: formData,
                processData: false,
                contentType: false,
                success: function (res) {
                    $("#fileName1").val("");
                    $("#clubname").val("");
                    $("#introduce").val("");
                }
            });
            $('#myModal').modal('hide');
            alert("创建成功");
            refreshpage();
        },
        error: function (e) {
            alert("请求失败");
        }
    });
}
function search() {
    var clubname = document.getElementById('clubname1').value;
    page = 1
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/findclub',
        dataType: "json",
        data: { page: 1, pagesize: 10, clubname: clubname },
        success: function (res) {
            if (res.code == 1) {
                alert(JSON.stringify(res.message));
                return;
            }
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
                $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].club_name + "</td><td>" + res.data[i].introduce + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='editclubname(this)' name='" + res.data[i].club_name + "' class='btn btn-default'>修改名字</a><a onclick='del(this)' name='" + res.data[i].club_name + "' class='btn btn-danger'>删除</a></div></td></tr>");
            }
        },
        error: function (e) {
            alert("请求失败");
        }
    });
}
function clearsearch() {
    document.getElementById('clubname1').value = "";
    page = 1;
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/findclub',
        dataType: "json",
        data: { page: 1, pagesize: 10, clubname: "" },
        success: function (res) {
            if (res.code == 1) {
                alert(JSON.stringify(res.message));
                return;
            }
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
                $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].club_name + "</td><td>" + res.data[i].introduce + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='editclubname(this)' name='" + res.data[i].club_name + "' class='btn btn-default'>修改名字</a><a onclick='del(this)' name='" + res.data[i].club_name + "' class='btn btn-danger'>删除</a></div></td></tr>");
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
        data: { role: "super_admin" },
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
    page = 1
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/sadminchangepassword',
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