var page = 1
$.ajax({
    type: "POST",
    url: 'http://localhost:8081/api/findadmin',
    dataType: "json",
    data: { page: 1, pagesize: 10, username: "", state: 0, clubname: "" },
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
            if (res.data[i].active_state == 1) {
                $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].username + "</td><td>正常</td><td>" + res.data[i].club_name + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='editpwd(this)' name='" + res.data[i].username + "' class='btn btn-default'>修改密码</a><a onclick='editstate1(this)' name='" + res.data[i].username + "' class='btn btn-default'>禁用</a><a onclick='del(this)' name='" + res.data[i].username + "' class='btn btn-danger'>删除</a></div></td></tr>");
            }
            else {
                $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].username + "</td><td>冻结</td><td>" + res.data[i].club_name + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='editpwd(this)' name='" + res.data[i].username + "' class='btn btn-default'>修改密码</a><a onclick='editstate2(this)' name='" + res.data[i].username + "' class='btn btn-default'>激活</a><a onclick='del(this)' name='" + res.data[i].username + "' class='btn btn-danger'>删除</a></div></td></tr>");
            }
        }
        $("#welcome").html("欢迎"+$.cookie('username')+"登录Tkkc社团组织管理系统");
    },
    error: function (e) {
        alert("请求失败");
    }
});
function rTime(date) {
    var json_date = new Date(date).toJSON();
    return new Date(new Date(json_date) + 8 * 3600 * 1000).toISOString().replace(/T/g, ' ').replace(/\.[\d]{3}Z/, '') 
}
$.ajax({
    type: "POST",
    url: 'http://localhost:8081/api/findallclub',
    dataType: "json",
    data: {},
    success: function (res) {
        if (res.code == 1) {
            alert(JSON.stringify(res.message));
            return;
        }
        for (var i = 0; i < res.length; i++) {
            $(".clubname").append("<option value='" + res[i].club_name + "'>" + res[i].club_name + "</option>")
        }
    },
    error: function (e) {
        alert("请求失败");
    }
});
function editpwd(val) {
    var pwd = prompt("新的密码：");
    if (pwd == null) {
        return
    }
    if (pwd == "") {
        alert("密码不能为空")
        return
    }
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/editadmin',
        dataType: "json",
        data: { username: val.name, password: pwd, state: 0 },
        success: function (res) {
            alert(JSON.stringify(res.message));
            refreshpage();
        },
        error: function (e) {
            alert("请求失败");
        }
    });

}
function editstate1(val) {
    if (confirm('确认禁用: ' + val.name + " ？") == false) {
        return
    }
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/editadmin',
        dataType: "json",
        data: { username: val.name, state: 2 },
        success: function (res) {
            alert(JSON.stringify(res.message));
            refreshpage();
        },
        error: function (e) {
            alert("请求失败");
        }
    });

}
function editstate2(val) {
    if (confirm('确认激活用户: ' + val.name + " ？") == false) {
        return
    }
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/editadmin',
        dataType: "json",
        data: { username: val.name, state: 1 },
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
    if (confirm('确认删除用户: ' + val.name + " ？") == false) {
        return
    }
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/deladmin',
        dataType: "json",
        data: { username: val.name },
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
    var username = document.getElementById('username1').value;
    var state = document.getElementById('state1').value
    var clubname = document.getElementById('clubname1').value
    page = val.name
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/findadmin',
        dataType: "json",
        data: { page: val.name, pagesize: 10, username: username, state: state, clubname: clubname },
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
                if (res.data[i].active_state == 1) {
                    $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].username + "</td><td>正常</td><td>" + res.data[i].club_name + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='editpwd(this)' name='" + res.data[i].username + "' class='btn btn-default'>修改密码</a><a onclick='editstate1(this)' name='" + res.data[i].username + "' class='btn btn-default'>禁用</a><a onclick='del(this)' name='" + res.data[i].username + "' class='btn btn-danger'>删除</a></div></td></tr>");
                }
                else {
                    $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].username + "</td><td>冻结</td><td>" + res.data[i].club_name + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='editpwd(this)' name='" + res.data[i].username + "' class='btn btn-default'>修改密码</a><a onclick='editstate2(this)' name='" + res.data[i].username + "' class='btn btn-default'>激活</a><a onclick='del(this)' name='" + res.data[i].username + "' class='btn btn-danger'>删除</a></div></td></tr>");
                }
            }
        },
        error: function (e) {
            alert("请求失败");
        }
    });
}
function refreshpage() {
    var username = document.getElementById('username1').value;
    var state = document.getElementById('state1').value
    var clubname = document.getElementById('clubname1').value
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/findadmin',
        dataType: "json",
        data: { page: page, pagesize: 10, username: username, state: state, clubname: clubname },
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
                if (res.data[i].active_state == 1) {
                    $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].username + "</td><td>正常</td><td>" + res.data[i].club_name + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='editpwd(this)' name='" + res.data[i].username + "' class='btn btn-default'>修改密码</a><a onclick='editstate1(this)' name='" + res.data[i].username + "' class='btn btn-default'>禁用</a><a onclick='del(this)' name='" + res.data[i].username + "' class='btn btn-danger'>删除</a></div></td></tr>");
                }
                else {
                    $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].username + "</td><td>冻结</td><td>" + res.data[i].club_name + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='editpwd(this)' name='" + res.data[i].username + "' class='btn btn-default'>修改密码</a><a onclick='editstate2(this)' name='" + res.data[i].username + "' class='btn btn-default'>激活</a><a onclick='del(this)' name='" + res.data[i].username + "' class='btn btn-danger'>删除</a></div></td></tr>");
                }
            }
        },
        error: function (e) {
            alert("请求失败");
        }
    });
}
function createadmin() {
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/createadmin',
        dataType: "json",
        data: { username: $("#username").val(), password: $("#password").val(), club: $("#clubname").val() },
        success: function (res) {
            if (res.code == 1) {
                alert(JSON.stringify(res.message));
                return;
            }
            alert("创建成功");
            refreshpage();
        },
        error: function (e) {
            alert("请求失败");
        }
    });
}
function search() {
    var username = document.getElementById('username1').value;
    var state = document.getElementById('state1').value;
    var clubname = document.getElementById('clubname1').value
    page = 1
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/findadmin',
        dataType: "json",
        data: { page: 1, pagesize: 10, username: username, state: state, clubname: clubname },
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
                if (res.data[i].active_state == 1) {
                    $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].username + "</td><td>正常</td><td>" + res.data[i].club_name + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='editpwd(this)' name='" + res.data[i].username + "' class='btn btn-default'>修改密码</a><a onclick='editstate1(this)' name='" + res.data[i].username + "' class='btn btn-default'>禁用</a><a onclick='del(this)' name='" + res.data[i].username + "' class='btn btn-danger'>删除</a></div></td></tr>");
                }
                else {
                    $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].username + "</td><td>冻结</td><td>" + res.data[i].club_name + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='editpwd(this)' name='" + res.data[i].username + "' class='btn btn-default'>修改密码</a><a onclick='editstate2(this)' name='" + res.data[i].username + "' class='btn btn-default'>激活</a><a onclick='del(this)' name='" + res.data[i].username + "' class='btn btn-danger'>删除</a></div></td></tr>");
                }
            }
        },
        error: function (e) {
            alert("请求失败");
        }
    });
}
function clearsearch() {
    document.getElementById('username1').value = "";
    document.getElementById('state1').value = 0;
    document.getElementById('clubname1').value = ""
    page = 1;
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/findadmin',
        dataType: "json",
        data: { page: 1, pagesize: 10, username: "", state: 0, clubname: clubname },
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
                if (res.data[i].active_state == 1) {
                    $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].username + "</td><td>正常</td><td>" + res.data[i].club_name + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='editpwd(this)' name='" + res.data[i].username + "' class='btn btn-default'>修改密码</a><a onclick='editstate1(this)' name='" + res.data[i].username + "' class='btn btn-default'>禁用</a><a onclick='del(this)' name='" + res.data[i].username + "' class='btn btn-danger'>删除</a></div></td></tr>");
                }
                else {
                    $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].username + "</td><td>冻结</td><td>" + res.data[i].club_name + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='editpwd(this)' name='" + res.data[i].username + "' class='btn btn-default'>修改密码</a><a onclick='editstate2(this)' name='" + res.data[i].username + "' class='btn btn-default'>激活</a><a onclick='del(this)' name='" + res.data[i].username + "' class='btn btn-danger'>删除</a></div></td></tr>");
                }
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
            if (pwd == ""  || rpwd == "") {
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