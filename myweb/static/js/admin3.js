var page = 1
$.ajax({
    type: "POST",
    url: 'http://localhost:8081/api/findclubactivity',
    dataType: "json",
    data: { page: 1, pagesize: 10 },
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
            if (res.data[i].state == 1) {
                $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].activity_name + "</td><td>待审核</td><td>" + res.data[i].club_name + "</td><td>" + res.data[i].remarks + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='details(this)' name='" + res.data[i].activity_name + "' class='btn btn-default'>活动详情</a><a onclick='editactivity1(this)' name='" + res.data[i].activity_name + "' class='btn btn-default'>修改活动</a><a onclick='delactivity(this)' name='" + res.data[i].activity_name + "' class='btn btn-danger'>删除</a></div></td></tr>");
            }
            if (res.data[i].state == 2) {
                $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].activity_name + "</td><td>审核失败</td><td>" + res.data[i].club_name + "</td><td>" + res.data[i].remarks + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='details(this)' name='" + res.data[i].activity_name + "' class='btn btn-default'>活动详情</a><a onclick='editactivity1(this)' name='" + res.data[i].activity_name + "' class='btn btn-default'>修改活动</a><a onclick='editactivity3(this)' name='" + res.data[i].activity_name + "' class='btn btn-success'>重新发起</a><a onclick='delactivity(this)' name='" + res.data[i].activity_name + "' class='btn btn-danger'>删除</a></div></td></tr>");
            }
            if (res.data[i].state == 3) {
                $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].activity_name + "</td><td>审核通过</td><td>" + res.data[i].club_name + "</td><td>" + res.data[i].remarks + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='details(this)' name='" + res.data[i].activity_name + "' class='btn btn-default'>活动详情</a><a onclick='enrolldetails(this)' name='" + res.data[i].activity_name + "' class='btn btn-default'>报名详情</a><a onclick='editactivity2(this)' name='" + res.data[i].activity_name + "' class='btn btn-danger'>关闭</a></div></td></tr>");
            }
            if (res.data[i].state == 4) {
                $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].activity_name + "</td><td>活动已关闭</td><td>" + res.data[i].club_name + "</td><td>" + res.data[i].remarks + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='details(this)' name='" + res.data[i].activity_name + "' class='btn btn-default'>活动详情</a><a onclick='enrolldetails(this)' name='" + res.data[i].activity_name + "' class='btn btn-default'>报名详情</a><a onclick='editactivity1(this)' name='" + res.data[i].activity_name + "' class='btn btn-default'>修改活动</a><a onclick='editactivity3(this)' name='" + res.data[i].activity_name + "' class='btn btn-success'>重新发起</a><a onclick='delactivity(this)' name='" + res.data[i].activity_name + "' class='btn btn-danger'>删除</a></div></td></tr>");
            }
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
var changeactivityid = ""
function editactivity1(val) {
    changeactivityid = val.name
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/findclubactivityintroduce',
        dataType: "json",
        data: { activityname: val.name },
        success: function (res) {
            $('#introduce2').val(res.introduce);
        },
        error: function (e) {
            alert("请求失败");
        }
    });
    $("#myModal4").modal("show")
}
function editactivity2(val) {
    if (confirm('确认关闭活动：' + val.name + ' ？') == false) {
        return
    }
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/editclubactivity',
        dataType: "json",
        data: { activityname: val.name, state: 4 },
        success: function (res) {
            alert(JSON.stringify(res.message));
            refreshpage();
        },
        error: function (e) {
            alert("请求失败");
        }
    });
}
function editactivity3(val) {
    if (confirm('确认发起活动：' + val.name + ' ？') == false) {
        return
    }
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/editclubactivity',
        dataType: "json",
        data: { activityname: val.name, state: 1 },
        success: function (res) {
            alert(JSON.stringify(res.message));
            refreshpage();
        },
        error: function (e) {
            alert("请求失败");
        }
    });
}
function changeactivity() {
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/editclubactivity',
        dataType: "json",
        data: { activityname: changeactivityid, introduce: $("#introduce2").val(), state: 0 },
        success: function (res) {
            alert(JSON.stringify(res.message));
            refreshpage();
        },
        error: function (e) {
            alert("请求失败");
        }
    });

}
function createactivity() {
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
        url: 'http://localhost:8081/api/createclubactivity?activityname='+ $("#activityname1").val()+"&introduce="+document.getElementById("introduce1").value.replace(/\n/g,"<br/>"),
        dataType: "json",
        data: formData,
        processData: false,
        contentType: false,
        success: function (res) {
            if (res.code == 1){
            alert(JSON.stringify(res.message));
            return;
            }
            $("#activityname1").val("");
            $("#introduce1").val("");
            $("#fileName1").val("");
            $('#myModal').modal('hide');
            alert("创建成功");
            refreshpage();
        },
        error: function (e) {
            alert("请求失败");
        }
    });
}
function delactivity(val) {
    if (confirm("确认删除活动" + val.name + " ？") == false) {
        return
    }
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/delclubactivity',
        dataType: "json",
        data: { activityname: val.name },
        success: function (res) {
            alert(JSON.stringify(res.message));
            refreshpage();
        },
        error: function (e) {
            alert("请求失败");
        }
    });
}
function details(val) {
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/findclubactivityintroduce',
        dataType: "json",
        data: { activityname: val.name },
        success: function (res) {
            $('#activitytitle').empty();
            $('#activityintroduce').empty();
            $('#activitytitle').append("活动名：" + val.name);
            $('#activityintroduce').append(res.introduce);
            $('#myModal2').modal('show');
        },
        error: function (e) {
            alert("请求失败");
        }
    });

}
function exportexcel(val) {
    window.open("http://localhost:8081/api/activityenrolldownloadfile?activityname="+val.name, '_blank');

}
var enrollactivityname = ""
function enrolldetails(val) {
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/adminfindclubenroll',
        dataType: "json",
        data: { page: 1, pagesize: 10, activityname: val.name },
        success: function (res) {
            enrollactivityname = val.name
            $('#li1').empty();
            $('#div1').empty();
            $('#tr1').empty();
            $('#li1').append("<a href='http://localhost:8081/admin/activity'>活动管理</a>");
            $('#div1').append("<div class='panel-heading'>报名人数：" + res.total + " </div> <div class='panel-body'><button class='btn btn-danger' type='button' onclick='black()'>返回</button><br/><br/><button class='btn btn-default' name='"+val.name+"' type='button' onclick='exportexcel(this)'>导出报名表</button></div>")
            $('#ol1').append("<li class='active'>报名详情</li>");
            $('#tr1').append("<th>序号</th><th>用户名</th><th>姓名</th><th>性别</th><th>学号</th><th>专业</th><th>电话</th><th>电子邮件</th><th>报名时间</th>")
            if (res.code == 1) {
                alert(JSON.stringify(res.message));
                return;
            }
            $("#ul1").empty();
            $("#ul1").append("<li><a name='1' onclick='enrollselectpage(this)'>&laquo;</a></li><li class='active'><a>1</a></li>")
            if (res.totalpages - res.page > 4) {
                for (var i = 2; i <= 5; i++) {
                    $("#ul1").append("<li><a name='" + i + "' onclick='enrollselectpage(this)'>" + i + "</a></li>")
                }
            }
            else {
                for (var i = res.page + 1; i <= res.totalpages; i++) {
                    $("#ul1").append("<li><a name='" + i + "' onclick='enrollselectpage(this)'>" + i + "</a></li>")
                }
            }
            $("#ul1").append("<li><a name='" + res.totalpages + "' onclick='enrollselectpage(this)'>&raquo;</a></li>")
            $("#tbody1").empty();
            for (var i = 0; i < res.data.length; i++) {
                $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].username + "</td><td>" + res.data[i].name + "</td><td>" + res.data[i].sex + "</td><td>" + res.data[i].student_number + "</td><td>" + res.data[i].major + "</td><td>" + res.data[i].phone_number + "</td><td>" + res.data[i].email + "</td><td>" + rTime(res.data[i].created_at) + "</td></tr>");
            }
        },
        error: function (e) {
            alert("请求失败");
        }
    });

}
function black() {
    $('#ol1').empty();
    $('#div1').empty();
    $('#tr1').empty();
    $('#ol1').append("<li class='active'>活动详情</li>");
    $('#div1').append("<div class='panel-heading'>添加活动</div><div class='panel-body'><label for='name'>添加活动</label><button class='btn btn-success' data-target='#myModal' data-toggle='modal'> 点击添加</button></div>")
    $('#tr1').append("<th>编号</th><th>活动名</th> <th>状态</th><th>发布组织</th><th>备注</th><th>创建时间</th><th>操作</th>");
    refreshpage();
}
function enrollselectpage(val) {
    $.ajax({
        type: "POST",
        url: 'http://localhost:8081/api/adminfindclubenroll',
        dataType: "json",
        data: { page: val.name, pagesize: 10, activityname: enrollactivityname },
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
                $("#ul1").append("<li><a name='1' onclick=''enrollselectpage(this)'>&laquo;</a></li><li class='active'><a>1</a></li>")
                if (res.totalpages - res.page > 4) {
                    for (var i = 2; i <= 5; i++) {
                        $("#ul1").append("<li><a name='" + i + "' onclick=''enrollselectpage(this)'>" + i + "</a></li>")
                    }
                }
                else {
                    for (var i = res.page + 1; i <= res.totalpages; i++) {
                        $("#ul1").append("<li><a name='" + i + "' onclick=''enrollselectpage(this)'>" + i + "</a></li>")
                    }
                }
                $("#ul1").append("<li><a name='" + res.totalpages + "' onclick=''enrollselectpage(this)'>&raquo;</a></li>")
            }
            else if (res.page == 2) {
                $("#ul1").append("<li><a name='1' onclick=''enrollselectpage(this)'>&laquo;</a></li><li><a onclick='enrollselectpage(this)' name='1'>1</a></li><li class='active'><a name='2' onclick='enrollselectpage(this)'>2</a></li>")
                if (res.totalpages - res.page > 3) {
                    for (var i = 3; i <= 5; i++) {
                        $("#ul1").append("<li><a name='" + i + "' onclick='enrollselectpage(this)'>" + i + "</a></li>")
                    }
                }
                else {
                    for (var i = res.page + 1; i <= res.totalpages; i++) {
                        $("#ul1").append("<li><a name='" + i + "' onclick='enrollselectpage(this)'>" + i + "</a></li>")
                    }
                }
                $("#ul1").append("<li><a name='" + res.totalpages + "' onclick='enrollselectpage(this)'>&raquo;</a></li>")
            }
            else {
                $("#ul1").append("<li><a name='1' onclick='enrollselectpage(this)'>&laquo;</a></li><li><a onclick='enrollselectpage(this)' name='" + (res.page - 2) + "'>" + (res.page - 2) + "</a></li><li><a name='" + (res.page - 1) + "' onclick='enrollselectpage(this)'>" + (res.page - 1) + "</a></li><li class='active'><a name='" + res.page + "' onclick='enrollselectpage(this)'>" + res.page + "</a></li>")
                if (res.totalpages - res.page > 2) {
                    $("#ul1").append("<li><a name='" + (res.page + 1) + "' onclick='enrollselectpage(this)' >" + (res.page + 1) + "</a></li><li><a name='" + (res.page + 2) + "' onclick='enrollselectpage(this)' >" + (res.page + 2) + "</a></li>")
                }
                else {
                    for (var i = res.page + 1; i <= res.totalpages; i++) {
                        $("#ul1").append("<li><a name='" + i + "' onclick='enrollselectpage(this)'>" + i + "</a></li>")
                    }
                }
                $("#ul1").append("<li><a name='" + res.totalpages + "' onclick='enrollselectpage(this)'>&raquo;</a></li>")
            }
            $("#tbody1").empty();
            for (var i = 0; i < res.data.length; i++) {
                $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].username + "</td><td>" + res.data[i].name + "</td><td>" + res.data[i].sex + "</td><td>" + res.data[i].student_number + "</td><td>" + res.data[i].major + "</td><td>" + res.data[i].phone_number + "</td><td>" + res.data[i].email + "</td><td>" + rTime(res.data[i].created_at) + "</td></tr>");
            }
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
        url: 'http://localhost:8081/api/findclubactivity',
        dataType: "json",
        data: { page: val.name, pagesize: 10 },
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
                if (res.data[i].state == 1) {
                    $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].activity_name + "</td><td>待审核</td><td>" + res.data[i].club_name + "</td><td>" + res.data[i].remarks + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='details(this)' name='" + res.data[i].activity_name + "' class='btn btn-default'>活动详情</a><a onclick='editactivity1(this)' name='" + res.data[i].activity_name + "' class='btn btn-default'>修改活动</a><a onclick='delactivity(this)' name='" + res.data[i].activity_name + "' class='btn btn-danger'>删除</a></div></td></tr>");
                }
                if (res.data[i].state == 2) {
                    $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].activity_name + "</td><td>审核失败</td><td>" + res.data[i].club_name + "</td><td>" + res.data[i].remarks + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='details(this)' name='" + res.data[i].activity_name + "' class='btn btn-default'>活动详情</a><a onclick='editactivity1(this)' name='" + res.data[i].activity_name + "' class='btn btn-default'>修改活动</a><a onclick='editactivity3(this)' name='" + res.data[i].activity_name + "' class='btn btn-success'>重新发起</a><a onclick='delactivity(this)' name='" + res.data[i].activity_name + "' class='btn btn-danger'>删除</a></div></td></tr>");
                }
                if (res.data[i].state == 3) {
                    $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].activity_name + "</td><td>审核通过</td><td>" + res.data[i].club_name + "</td><td>" + res.data[i].remarks + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='details(this)' name='" + res.data[i].activity_name + "' class='btn btn-default'>活动详情</a><a onclick='enrolldetails(this)' name='" + res.data[i].activity_name + "' class='btn btn-default'>报名详情</a><a onclick='editactivity2(this)' name='" + res.data[i].activity_name + "' class='btn btn-danger'>关闭</a></div></td></tr>");
                }
                if (res.data[i].state == 4) {
                    $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].activity_name + "</td><td>活动已关闭</td><td>" + res.data[i].club_name + "</td><td>" + res.data[i].remarks + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='details(this)' name='" + res.data[i].activity_name + "' class='btn btn-default'>活动详情</a><a onclick='enrolldetails(this)' name='" + res.data[i].activity_name + "' class='btn btn-default'>报名详情</a><a onclick='editactivity1(this)' name='" + res.data[i].activity_name + "' class='btn btn-default'>修改活动</a><a onclick='editactivity3(this)' name='" + res.data[i].activity_name + "' class='btn btn-success'>重新发起</a><a onclick='delactivity(this)' name='" + res.data[i].activity_name + "' class='btn btn-danger'>删除</a></div></td></tr>");
                }
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
        url: 'http://localhost:8081/api/findclubactivity',
        dataType: "json",
        data: { page: page, pagesize: 10 },
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
                if (res.data[i].state == 1) {
                $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].activity_name + "</td><td>待审核</td><td>" + res.data[i].club_name + "</td><td>" + res.data[i].remarks + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='details(this)' name='" + res.data[i].activity_name + "' class='btn btn-default'>活动详情</a><a onclick='editactivity1(this)' name='" + res.data[i].activity_name + "' class='btn btn-default'>修改活动</a><a onclick='delactivity(this)' name='" + res.data[i].activity_name + "' class='btn btn-danger'>删除</a></div></td></tr>");
            }
            if (res.data[i].state == 2) {
                $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].activity_name + "</td><td>审核失败</td><td>" + res.data[i].club_name + "</td><td>" + res.data[i].remarks + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='details(this)' name='" + res.data[i].activity_name + "' class='btn btn-default'>活动详情</a><a onclick='editactivity1(this)' name='" + res.data[i].activity_name + "' class='btn btn-default'>修改活动</a><a onclick='editactivity3(this)' name='" + res.data[i].activity_name + "' class='btn btn-success'>重新发起</a><a onclick='delactivity(this)' name='" + res.data[i].activity_name + "' class='btn btn-danger'>删除</a></div></td></tr>");
            }
            if (res.data[i].state == 3) {
                $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].activity_name + "</td><td>审核通过</td><td>" + res.data[i].club_name + "</td><td>" + res.data[i].remarks + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='details(this)' name='" + res.data[i].activity_name + "' class='btn btn-default'>活动详情</a><a onclick='enrolldetails(this)' name='" + res.data[i].activity_name + "' class='btn btn-default'>报名详情</a><a onclick='editactivity2(this)' name='" + res.data[i].activity_name + "' class='btn btn-danger'>关闭</a></div></td></tr>");
            }
            if (res.data[i].state == 4) {
                $("#tbody1").append("<tr><td>" + (i + 1) + "</td><td>" + res.data[i].activity_name + "</td><td>活动已关闭</td><td>" + res.data[i].club_name + "</td><td>" + res.data[i].remarks + "</td><td>" + rTime(res.data[i].created_at) + "</td><td><div class='btn-group'> <a onclick='details(this)' name='" + res.data[i].activity_name + "' class='btn btn-default'>活动详情</a><a onclick='enrolldetails(this)' name='" + res.data[i].activity_name + "' class='btn btn-default'>报名详情</a><a onclick='editactivity1(this)' name='" + res.data[i].activity_name + "' class='btn btn-default'>修改活动</a><a onclick='editactivity3(this)' name='" + res.data[i].activity_name + "' class='btn btn-success'>重新发起</a><a onclick='delactivity(this)' name='" + res.data[i].activity_name + "' class='btn btn-danger'>删除</a></div></td></tr>");
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