 var page = 1
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/findclubmember',
        dataType: "json",
        data: { page: 1, pagesize: 10, name: "" },
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
                $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].username + "</td><td>" + res.data[i].name + "</td><td>" + res.data[i].role + "</td><td><div class='btn-group'><a onclick='resume(this)' name='" + res.data[i].username + "' class='btn btn-default'>查看简历</a> <a  onclick='change(this)' data-target='#myModal4' data-toggle='modal' name='" + res.data[i].username + "' class='btn btn-default'>更改角色</a><a onclick='del(this)' name='" + res.data[i].username + "' class='btn btn-danger'>移出组织</a></div></td></tr>");
            }
            $("#welcome").html("欢迎" + $.cookie('username') + "登录Tkkc社团组织管理系统");
        },
        error: function (e) {
            alert("请求失败");
        }
    });
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/findclubrole',
        dataType: "json",
        data: {},
        success: function (res) {
            if (res.code == 1) {
                alert(JSON.stringify(res.message));
                return;
            }
            for (var i = 0; i < res.length; i++) {
                $(".role").append("<option value='" + res[i].role + "'>" + res[i].role + "</option>")
            }
        },
        error: function (e) {
            alert("请求失败");
        }
    });
    var changeroleid = ""
    function change(val) {
        changeroleid = val.name
    }
    function changerole() {
        $.ajax({
            type: "POST",
            url: 'http://localhost:8081/api/editclubmember',
            dataType: "json",
            data: { username: changeroleid, role: $("#role3").val() },
            success: function (res) {
                alert(JSON.stringify(res.message));
                refreshrole();
                refreshpage();
            },
            error: function (e) {
                alert("请求失败");
            }
        });

    }
    function resume(val) {

        $.ajax({
            type: "POST",
            url: 'http://localhost:8081/api/findresume',
            dataType: "json",
            data: { username: val.name},
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
    function del(val) {
        if (confirm('确认删除用户: ' + val.name + " ？") == false) {
            return
        }
        $.ajax({
            type: "POST",
            url: 'http://localhost:8081/api/delclubmember',
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
        var name = document.getElementById('name1').value;
        page = val.name
        $.ajax({
            type: "POST",
            url: 'http://localhost:8081/api/findclubmember',
            dataType: "json",
            data: { page: val.name, pagesize: 10, name: name },
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
                    $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].username + "</td><td>" + res.data[i].name + "</td><td>" + res.data[i].role + "</td><td><div class='btn-group'><a onclick='resume(this)' name='" + res.data[i].username + "' class='btn btn-default'>查看简历</a> <a  onclick='change(this)' data-target='#myModal4' data-toggle='modal' name='" + res.data[i].username + "' class='btn btn-default'>更改角色</a><a onclick='del(this)' name='" + res.data[i].username + "' class='btn btn-danger'>移出组织</a></div></td></tr>");
                }
            },
            error: function (e) {
                alert("请求失败");
            }
        });
    }
    function refreshrole() {
        $.ajax({
            type: "POST",
            url: 'http://localhost:8081/api/findclubrole',
            dataType: "json",
            data: {},
            success: function (res) {
                if (res.code == 1) {
                    alert(JSON.stringify(res.message));
                    return;
                }
                $(".role").empty();
                $(".role").append("<option value=''>不选择</option>")
                for (var i = 0; i < res.length; i++) {
                    $(".role").append("<option value='" + res[i].role + "'>" + res[i].role + "</option>")
                }
            },
            error: function (e) {
                alert("请求失败");
            }
        });
    }
    function refreshpage() {
        var name = document.getElementById('name1').value;
        $.ajax({
            type: "POST",
            url: 'http://localhost:8081/api/findclubmember',
            dataType: "json",
            data: { page: page, pagesize: 10, name: name },
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
                    $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].username + "</td><td>" + res.data[i].name + "</td><td>" + res.data[i].role + "</td><td><div class='btn-group'><a onclick='resume(this)' name='" + res.data[i].username + "' class='btn btn-default'>查看简历</a> <a  onclick='change(this)' data-target='#myModal4' data-toggle='modal' name='" + res.data[i].username + "' class='btn btn-default'>更改角色</a><a onclick='del(this)' name='" + res.data[i].username + "' class='btn btn-danger'>移出组织</a></div></td></tr>");
                }
            },
            error: function (e) {
                alert("请求失败");
            }
        });
    }
    function createrole() {
        $.ajax({
            type: "POST",
            url: 'http://localhost:8081/api/createclubrole',
            dataType: "json",
            data: { role: $("#role").val() },
            success: function (res) {
                $("#role").val("") 
                if (res.code == 1) {
                    alert(JSON.stringify(res.message));
                    return;
                }
                alert("创建成功");
                refreshrole();
            },
            error: function (e) {
                $("#role").val("") 
                alert("请求失败");
            }
        });
    }
    function search() {
        var name = document.getElementById('name1').value;
        page = 1
        $.ajax({
            type: "POST",
            url: 'http://localhost:8081/api/findclubmember',
            dataType: "json",
            data: { page: 1, pagesize: 10, name: name },
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
                    $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].username + "</td><td>" + res.data[i].name + "</td><td>" + res.data[i].role + "</td><td><div class='btn-group'><a onclick='resume(this)' name='" + res.data[i].username + "' class='btn btn-default'>查看简历</a> <a  onclick='change(this)' data-target='#myModal4' data-toggle='modal' name='" + res.data[i].username + "' class='btn btn-default'>更改角色</a><a onclick='del(this)' name='" + res.data[i].username + "' class='btn btn-danger'>移出组织</a></div></td></tr>");
                }
            },
            error: function (e) {
                alert("请求失败");
            }
        });
    }
    function clearsearch() {
        document.getElementById('name1').value = "";
        page = 1;
        $.ajax({
            type: "POST",
            url: 'http://localhost:8081/api/findclubmember',
            dataType: "json",
            data: { page: 1, pagesize: 10, name: "" },
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
                    $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].username + "</td><td>" + res.data[i].name + "</td><td>" + res.data[i].role + "</td><td><div class='btn-group'><a onclick='resume(this)' name='" + res.data[i].username + "' class='btn btn-default'>查看简历</a> <a  onclick='change(this)' data-target='#myModal4' data-toggle='modal' name='" + res.data[i].username + "' class='btn btn-default'>更改角色</a><a onclick='del(this)' name='" + res.data[i].username + "' class='btn btn-danger'>移出组织</a></div></td></tr>");
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
    function delrole() {
        var role = document.getElementById('role2').value
        if (confirm("确定删除" + role + "？") == false) {
            return
        }
        $.ajax({
            type: "POST",
            url: 'http://localhost:8081/api/delclubrole',
            dataType: "json",
            data: { role: role },
            success: function (res) {
                refreshrole();
                refreshpage();
                alert(JSON.stringify(res.message));
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