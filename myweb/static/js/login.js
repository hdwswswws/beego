var countdown = 60;
var send = 0;
function settime(val) {
    var username = document.getElementById("r_emial").value;
    var password = document.getElementById("r_password").value;
    if (username == "") {
        alert("邮箱不能为空")
        return;
    }
    if (password == "") {
        alert("密码不能为空")
        return;
    }
    if (send == 0) {
        $.ajax({
            type: "POST",
            url: 'http://localhost:8081/register',
            dataType: "json",
            data: { username: username, password: password },
            success: function (res) {
                if (res.code == 0) { }
                else {
                    alert(res.message)
                    countdown = 0
                }
            },
            error: function (e) {
                alert("请求失败");
            }
        });
        send = 1
    }
    if (countdown == 0) {
        val.removeAttribute("disabled");
        val.value = "获取验证码";
        countdown = 60;
        send = 0;
        return;
    } else {
        val.setAttribute("disabled", true);
        val.value = "重新发送(" + countdown + ")";
        countdown--;
    }
    setTimeout(function () {
        settime(val)
    }, 1000)

}
$(document).ready(function () {
    $('#r_password1').on('input propertychange', function () {
        var pwd = $.trim($(this).val());
        var rpwd = $.trim($("#r_password").val());
        if (rpwd != "") {
            if (pwd == "" || rpwd == "") {
                $("#msg_pwd").html("<font color='red'>密码不能为空</font>");
                $("#create").attr("disabled", true);
                $("#forgot").attr("disabled", true);
                $("#btn").attr("disabled", true);
            }
            else {
                if (pwd == rpwd) {
                    $("#msg_pwd").html("<font color='green'>两次密码相同</font>");
                    $("#create").attr("disabled", false);
                    $("#forgot").attr("disabled", false);
                    $("#btn").attr("disabled", false);
                } else {
                    $("#msg_pwd").html("<font color='red'>两次密码不相同</font>");
                    $("#create").attr("disabled", true);
                    $("#forgot").attr("disabled", true);
                    $("#btn").attr("disabled", true);
                }
            }
        }
    })
})
$(document).ready(function () {
    $('#r_password').on('input propertychange', function () {
        var pwd = $.trim($(this).val());
        var rpwd = $.trim($("#r_password1").val());
        if (pwd == "" || rpwd == "") {
            $("#msg_pwd").html("<font color='red'>密码不能为空</font>");
            $("#create").attr("disabled", true);
            $("#forgot").attr("disabled", true);
            $("#btn").attr("disabled", true);
        }
        else {
            if (pwd == rpwd) {
                $("#msg_pwd").html("<font color='green'>两次密码相同</font>");
                $("#create").attr("disabled", false);
                $("#forgot").attr("disabled", false);
                $("#btn").attr("disabled", false);
            } else {
                $("#msg_pwd").html("<font color='red'>两次密码不相同</font>");
                $("#create").attr("disabled", true);
                $("#forgot").attr("disabled", true);
                $("#btn").attr("disabled", true);
            }
        }
    })

})
$(function () {
    $('.message a').click(function () {
        $('form').animate({
            height: 'toggle',
            opacity: 'toggle'
        }, 'slow');
    });
    $("#create").click(function () {
        var username = document.getElementById('r_emial').value;
        var password = document.getElementById("r_password").value;
        var code = document.getElementById('code').value;
        if (username == "") {
            alert("邮箱不能为空")
            return;
        }
        if (password == "") {
            alert("密码不能为空")
            return;
        }
        if (code == "") {
            alert("验证码不能为空")
            return;
        }
        $.ajax({
            type: "POST",
            url: 'http://localhost:8081/confirmregister',
            dataType: "json",
            data: { username: username, code: code },
            success: function (res) {
                if (res.code == 0) {
                    window.location.href = "http://localhost:8081";
                }
                alert(res.message)

            },
            error: function (e) {
                alert("请求失败");
            }
        });
    });
    $("#login").click(function () {
        var username = document.getElementById('emial').value;
        var password = document.getElementById('password').value;
        var role = document.getElementById('checkedLevel').value
        if (username == "") {
            alert("邮箱不能为空")
            return;
        }
        if (password == "") {
            alert("密码不能为空")
            return;
        }
        $.ajax({
            type: "POST",
            url: 'http://localhost:8081/login',
            dataType: "json",
            data: { username: username, password: password, role: role },
            success: function (res) {
                if (res.code == 0) {
                    if (role == "super_admin") {
                        window.location.href = "http://localhost:8081/superadmin/index";
                    }
                    if (role == "admin") {
                        window.location.href = "http://localhost:8081/admin/index";
                    }
                    if (role == "user") {
                        $.ajax({
                            type: "POST",
                            url: 'http://localhost:8081/api/userfindresume',
                            dataType: "json",
                            data: {},
                            success: function (res) {
                               if(res.code == 1||res.username == ""){
                                alert("请先完善个人简历") 
                                window.location.href = "http://localhost:8081/user/resume";
                               }
                               else{
                                window.location.href = "http://localhost:8081/user/index";
                               }
                            },
                            error: function (e) {
                                alert("请求失败");
                            }
                        });
                    }
                    return
                }
                alert(res.message)

            },
            error: function (e) {
                alert("请求失败");
            }
        });
    });
    $("#forgot").click(function () {
        var username = document.getElementById('f_emial').value;
        var password = document.getElementById("r_password").value;
        var code = document.getElementById('f_code').value;
        if (username == "") {
            alert("邮箱不能为空")
            return;
        }
        if (password == "") {
            alert("密码不能为空")
            return;
        }
        if (code == "") {
            alert("验证码不能为空")
            return;
        }
        $.ajax({
            type: "POST",
            url: 'http://localhost:8081/forgotpassword',
            dataType: "json",
            data: { username: username, password: password, code: code },
            success: function (res) {
                if (res.code == 0) {
                    alert(res.message)
                    window.location.href = "http://localhost:8081";
                }
                else {
                    alert(res.message)
                }

            },
            error: function (e) {
                alert("请求失败");
            }
        });
    });
});
var countdown = 60;
var send = 0;
function forgotsettime(val) {
    var username = document.getElementById("f_emial").value;
    if (username == "") {
        alert("邮箱不能为空")
        return;
    }
    if (send == 0) {
        $.ajax({
            type: "POST",
            url: 'http://localhost:8081/resendconfirmationcode',
            dataType: "json",
            data: { username: username },
            success: function (res) {
                if (res.code == 0) { }
                else {
                    alert(res.message)
                    countdown = 0
                }

            },
            error: function (e) {
                alert("请求失败");
            }
        });
        send = 1
    }
    if (countdown == 0) {
        val.removeAttribute("disabled");
        val.value = "获取验证码";
        countdown = 60;
        send = 0;
        return;
    } else {
        val.setAttribute("disabled", true);
        val.value = "重新发送(" + countdown + ")";
        countdown--;
    }
    setTimeout(function () {
        forgotsettime(val)
    }, 1000)

}     