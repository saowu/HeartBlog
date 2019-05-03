layui.use(['laypage', 'layer'], function () {
    var laypage = layui.laypage;
    var $ = layui.jquery, layer = layui.layer;

    //自定义样式
    laypage.render({
        elem: 'pageable'
        , count: 70
        , theme: '#009688'
        , limit: 6
    });
    //自定义样式
    laypage.render({
        elem: 'pageable1'
        , count: 100
        , theme: '#009688'
    });

//触发事件
    //触发事件
    var active = {
        setTop: function () {
            var that = this;
            //多窗口模式，层叠置顶
            index = layer.open({
                type: 1
                , title: '<i class="layui-icon layui-icon-edit"></i>添加类型'
                , area: ['400px', '290px']
                , shade: 0
                , maxmin: true
                , offset: 'auto'
                , shade: 0.3
                , content: '<div class="site-text site-block" style="margin: 13%">\n' +
                '<form class="layui-form" action="" lay-filter="example">\n' +
                '  <div class="layui-form-item">\n' +
                '    <label class="layui-form-label">类型</label>\n' +
                '    <div class="layui-input-inline">\n' +
                '      <input type="text" name="atype" id="a_type" lay-verify="title" autocomplete="off" placeholder="请输入添加类型" class="layui-input">\n' +
                '    </div>\n' +
                '  </div>\n' +
                '</form>\n' +
                '</div>'
                , btn: ['保存']
                , yes: function () {
                    var a_type = $("#a_type").val()
                    var data = {"name": a_type}
                    if (a_type != "") {
                        $.ajax({
                            type: "Post",
                            url: "/categorys",
                            contentType: 'application/json;charset=utf-8',
                            dataType: 'json',
                            data: JSON.stringify(data),
                            success: function (data) {
                                layer.msg("保存成功！")
                                close_this()
                                location.replace("/category")
                            },
                            err: function (err) {
                                layer.msg("保存失败！")
                                close_this()
                            }
                        });
                    } else {
                        layer.msg("数据不能为空！")
                        close_this()

                    }
                }

                , zIndex: layer.zIndex
                , success: function (layero) {
                    layer.setTop(layero);
                }
            });

            function close_this() {
                layer.close(index)
            }
        },
        addArtcle: function () {
            var that = this;
            //多窗口模式，层叠置顶
            layer.open({
                type: 2
                , title: '<i class="layui-icon layui-icon-edit"></i>添加文章'
                , area: ['900px', '600px']
                , shade: 0
                , maxmin: true
                , offset: 'auto'
                , shade: 0.3
                , content: '/static/popup/addarticle.html'
                , zIndex: layer.zIndex
                , success: function (layero) {
                    layer.setTop(layero);
                }
            });
        },
        addLink: function () {
            var that = this;
            //多窗口模式，层叠置顶
            layer.open({
                type: 1
                , title: '<i class="layui-icon layui-icon-edit"></i>添加文章'
                , area: ['400px', '290px']
                , shade: 0
                , maxmin: true
                , offset: 'auto'
                , shade: 0.3
                , btn: ['保存']
                , content: '<div style="margin: 13%">\n' +
                '<div class="layui-form-item">\n' +
                '        <label class="layui-form-label">链接名</label>\n' +
                '        <div class="layui-input-inline">\n' +
                '            <input type="text" name="title" required lay-verify="required" placeholder="请输入名称" autocomplete="off"\n' +
                '                   class="layui-input">\n' +
                '        </div>\n' +
                '    </div>\n' +
                '    <div class="layui-form-item">\n' +
                '        <label class="layui-form-label">链接</label>\n' +
                '        <div class="layui-input-inline">\n' +
                '            <input type="password" name="password" required lay-verify="required" placeholder="请输入链接" autocomplete="off"\n' +
                '                   class="layui-input">\n' +
                '        </div>\n' +
                '    </div>\n' +
                '    </div>'
                , zIndex: layer.zIndex
                , success: function (layero) {
                    layer.setTop(layero);
                }
            });
        },
        addFLink: function () {
            var that = this;
            //多窗口模式，层叠置顶
            layer.open({
                type: 1
                , title: '<i class="layui-icon layui-icon-edit"></i>添加文章'
                , area: ['400px', '290px']
                , shade: 0
                , maxmin: true
                , offset: 'auto'
                , shade: 0.3
                , btn: ['保存']
                , content: '<div style="margin: 13%">\n' +
                '<div class="layui-form-item">\n' +
                '        <label class="layui-form-label">链接名</label>\n' +
                '        <div class="layui-input-inline">\n' +
                '            <input type="text" name="title" required lay-verify="required" placeholder="请输入名称" autocomplete="off"\n' +
                '                   class="layui-input">\n' +
                '        </div>\n' +
                '    </div>\n' +
                '    <div class="layui-form-item">\n' +
                '        <label class="layui-form-label">链接</label>\n' +
                '        <div class="layui-input-inline">\n' +
                '            <input type="password" name="password" required lay-verify="required" placeholder="请输入链接" autocomplete="off"\n' +
                '                   class="layui-input">\n' +
                '        </div>\n' +
                '    </div>\n' +
                '    </div>'
                , zIndex: layer.zIndex
                , success: function (layero) {
                    layer.setTop(layero);
                }
            });
        },
        addAlbum: function () {
            var that = this;
            //多窗口模式，层叠置顶
            layer.open({
                type: 2
                , title: '<i class="layui-icon layui-icon-edit"></i>添加相册'
                , area: ['900px', '600px']
                , shade: 0
                , maxmin: true
                , offset: 'auto'
                , shade: 0.3
                , content: '/static/popup/addalbum.html'
                , zIndex: layer.zIndex
                , success: function (layero) {
                    layer.setTop(layero);
                }
            });
        },
        addTab: function () {
            var that = this;
            //多窗口模式，层叠置顶
            layer.open({
                type: 1
                , title: '<i class="layui-icon layui-icon-edit"></i>添加标签'
                , area: ['400px', '290px']
                , shade: 0
                , maxmin: true
                , offset: 'auto'
                , shade: 0.3
                , content: '<div class="site-text site-block" style="margin: 13%">\n' +
                '<form class="layui-form" action="" lay-filter="example">\n' +
                '  <div class="layui-form-item">\n' +
                '    <label class="layui-form-label">标签</label>\n' +
                '    <div class="layui-input-inline">\n' +
                '      <input type="text" name="username" lay-verify="title" autocomplete="off" placeholder="请输入添加标签" class="layui-input">\n' +
                '    </div>\n' +
                '  </div>\n' +
                '</form>\n' +
                '</div>'
                , btn: ['保存'] //只是为了演示
                , yes: function () {
                    layer.closeAll();
                }

                , zIndex: layer.zIndex
                , success: function (layero) {
                    layer.setTop(layero);
                }
            });
        }
    };
    $('.dropdown-toggle').on('click', function () {
        var othis = $(this), method = othis.data('method');
        active[method] ? active[method].call(this, othis) : '';
    });
});