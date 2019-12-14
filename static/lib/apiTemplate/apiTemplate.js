layui.define(['jquery', 'form', 'layer', 'upload'], function (exports) {
    "use strict";

    var MOD_NAME = 'apiTemplate',
        layer = layui.layer,
        form = layui.form,
        upload = layui.upload,
        $ = layui.jquery;

    var apiTemplate = new function () {

        //配置信息
        this.config = function () {
            var config = {};
            return config;
        };

        //增加api模板
        this.APITemplate = function (elem) {
            var api = '<li class="layui-timeline-item" name="scenes_li_element">'
                + '	<i class="layui-icon layui-timeline-axis">&#xe63f;</i>'
                + '	<div class="layui-timeline-content layui-text" style="padding-top: 10px;">'
                + '		<hr style="margin-bottom: 25px;"/>'
                + '		<div class="layui-form-item">'
                + '			<div class="layui-input-inline">'
                + '				<input type="text" class="layui-input" required lay-verify="required" placeholder="请输入压测api名称" name="apiName"/>'
                + '				<input type="hidden" name="id" value=""/>'
                + '			</div>'
                + '			<div class="layui-form-mid layui-word-aux">请配置压测api</div>'
                + '			<div class="layui-input-inline" style="float: right;width:285px;">'
                + '				<button type="button" class="layui-btn" name="api_add_btn">新增API</button>'
                + '				<button type="button" class="layui-btn" name="api_ponit_add_btn">新增集合点</button>'
                + '				<button type="button" class="layui-btn layui-btn-danger" name="api_delete_btn">删除</button>'
                + '			</div>'
                + '		</div>'
                + '		<div class="layui-tab layui-tab-brief">'
                + '			<ul class="layui-tab-title">'
                + '				<li class="layui-this">基本请求信息</li>'
                + '				<li name="params_tab_body" style="display:none;">Body定义</li>'
                + '				<li>Header定义</li>'
                + '				<li>出参定义</li>'
                + '				<li>检查点（断言）</li>'
                + '			</ul>'
                + '			<div class="layui-tab-content">'
                + '				<div class="layui-tab-item layui-show">'
                + '					<div class="layui-form-item">'
                + '						<label class="layui-form-label">压测url</label>'
                + '						<div class="layui-input-block">'
                + '							<textarea name="desc" required lay-verify="required|url" placeholder="请输入一个完整的URL，不能包含空格、换行等非法字符"'
                + '									  class="layui-textarea"></textarea>'
                + '						</div>'
                + '					</div>'
                + '					<div class="layui-form-item">'
                + '						<label class="layui-form-label">请求方式</label>'
                + '						<div class="layui-input-inline">'
                + '							<select name="requestMethod" lay-filter="apiMethodFilter">'
                + '								<option value="GET">GET</option>'
                + '								<option value="POST">POST</option>'
                + '								<option value="PUT">PUT</option>'
                + '								<option value="DELETE">DELETE</option>'
                + '							</select>'
                + '						</div>'
                + '						<label class="layui-form-label">超时时间</label>'
                + '						<div class="layui-input-inline">'
                + '							<input type="text" class="layui-input" lay-verify="number" name="requestTimeout"'
                + '								   value="5000"/>'
                + '						</div>'
                + '						<div class="layui-form-mid layui-word-aux">ms</div>'
                + '						<label class="layui-form-label">登陆请求</label>'
                + '						<div class="layui-input-inline">'
                + '							<input type="checkbox" lay-skin="switch" lay-text="ON|OFF"'
                + '								    name="type" value="1" lay-filter="apiTypeSwitchFilter"/>'
                + '						</div>'
                + '						<div class="layui-form-mid layui-word-aux">链路中只有一个登陆请求</div>'
                + '					</div>'
                + '				</div>'
                + '				<div class="layui-tab-item">'
                + '                  <input type="hidden" name="params_method_content_type" value="application/x-www-form-urlencoded"/>'
                + '					<label class="layui-form-label" style="width: 100px;">Content-Type：</label>'
                + '					<div class="layui-input-inline" style="width: 190px;">'
                + '						<button type="button" class="layui-btn layui-btn-normal" name="contentType" sampContentType="0">x-www-form-urlencode</button>'
                + '					</div>'
                + '					<div class="layui-input-inline" style="width: 120px;">'
                + '						<button type="button" class="layui-btn layui-btn-primary" name="contentType" sampContentType="1">raw</button>'
                + '					</div>'
                + '					<div name="params_form_table">'
                + '						<table class="layui-table" name="params_table">'
                + '						  <colgroup>'
                + '							<col>'
                + '							<col>'
                + '							<col>'
                + '						  </colgroup>'
                + '						  <thead>'
                + '							<tr>'
                + '							  <th>key</th>'
                + '							  <th>value</th>'
                + '							  <th>操作</th>'
                + '							</tr> '
                + '						  </thead>'
                + '						  <tbody>'
                + '							<tr>'
                + '							  <td><input type="text" class="layui-input" name="paramsName" placeholder="输入Key"/></td>'
                + '							  <td><input type="text" class="layui-input" name="paramsValue" placeholder="输入Value"/></td>'
                + '							  <td><button type="button" class="layui-btn" name="params_delete_btn">删除</button></td>'
                + '							</tr>'
                + '						  </tbody>'
                + '						</table>'
                + '						<button type="button" class="layui-btn" name="params_add_btn">增加</button>'
                + '					</div>'
                + '					<div class="layui-input-block" style="display:none;margin: 30px 0 0 10px;">'
                + '						<textarea name="paramsBody" placeholder="如果服务端（被压测端）需要强校验换行符（\n）或者待加密的部分需要有换行符，请使用unescape解码函数对包含换行符的字符串进行反转义：${sys.unescapeJava(text)}"'
                + '								  class="layui-textarea"></textarea>'
                + '					</div>'
                + '				</div>'
                + '				<div class="layui-tab-item">'
                + '					<table class="layui-table" name="header_table">'
                + '					  <colgroup>'
                + '						<col>'
                + '						<col>'
                + '						<col>'
                + '					  </colgroup>'
                + '					  <thead>'
                + '						<tr>'
                + '						  <th>key</th>'
                + '						  <th>value</th>'
                + '						  <th>操作</th>'
                + '						</tr> '
                + '					  </thead>'
                + '					  <tbody>'
                + '						<tr>'
                + '						  <td><input type="text" class="layui-input" name="headerName" placeholder="输入Header Key"/></td>'
                + '						  <td><input type="text" class="layui-input" name="headerValue" placeholder="输入Header Value"/></td>'
                + '						  <td><button type="button" class="layui-btn" name="header_delete_btn">删除</button></td>'
                + '						</tr>'
                + '					  </tbody>'
                + '					</table>'
                + '					<button type="button" class="layui-btn" name="header_add_btn">增加</button>'
                + '				</div>'
                + '				<div class="layui-tab-item">'
                + '					<table class="layui-table" name="outpms_table">'
                + '					  <colgroup>'
                + '						<col>'
                + '						<col>'
                + '						<col>'
                + '						<col>'
                + '						<col>'
                + '					  </colgroup>'
                + '					  <thead>'
                + '						<tr>'
                + '						  <th>出参名</th>'
                + '						  <th>来源</th>'
                + '						  <th>解析表达式</th>'
                + '						  <th>第几个匹配项</th>'
                + '						  <th>操作</th>'
                + '						</tr> '
                + '					  </thead>'
                + '					  <tbody>'
                + '						<tr>'
                + '						  <td><input type="text" class="layui-input" placeholder="输入出参名" name="outName"/></td>'
                + '						  <td>'
                + '							 <select name="outSource">'
                + '								<option value="0">Body:TEXT</option>'
                + '								<option value="1">Body:JSON</option>'
                + '								<option value="2">Header:K/V</option>'
                + '								<option value="3">Cookie:K/V</option>'
                + '								<option value="4">响应状态码</option>'
                + '							</select>'
                + '						  </td>'
                + '						  <td><input type="text" class="layui-input" placeholder="输入出参提取表达式" name="outResolveExpress"/></td>'
                + '						  <td><input type="text" class="layui-input" lay-verify="number" value="0" name="outIndex"/></td>'
                + '						  <td><button type="button" class="layui-btn" name="outpms_delete_btn">删除</button></td>'
                + '						</tr>'
                + '					  </tbody>'
                + '					</table>'
                + '					<button type="button" class="layui-btn" name="outpms_add_btn">增加</button>'
                + '				</div>'
                + '				<div class="layui-tab-item">'
                + '					<table class="layui-table" name="assertion_table">'
                + '					  <colgroup>'
                + '						<col>'
                + '						<col>'
                + '						<col>'
                + '						<col>'
                + '						<col>'
                + '					  </colgroup>'
                + '					  <thead>'
                + '						<tr>'
                + '						  <th>检查点类型</th>'
                + '						  <th>检查对象</th>'
                + '						  <th>检查条件</th>'
                + '						  <th>检查内容</th>'
                + '						  <th>操作</th>'
                + '						</tr> '
                + '					  </thead>'
                + '					  <tbody>'
                + '						<tr>'
                + '						  <td>'
                + '							<select name="assertionType">'
                + '								<option value="0">响应header</option>'
                + '								<option value="1">响应状态码</option>'
                + '								<option value="2">响应body</option>'
                + '								<option value="3">出参</option>'
                + '							</select>'
                + '						  </td>'
                + '						  <td>'
                + '							 <input type="text" class="layui-input" placeholder="输入检查对象" name="assertionName"/>'
                + '						  </td>'
                + '						  <td>'
                + '							<select name="assertionFactor">'
                + '								<option value=">">大于</option>'
                + '								<option value=">=">大于等于</option>'
                + '								<option value="<">小于</option>'
                + '								<option value="<=">小于等于</option>'
                + '								<option value="=">等于</option>'
                + '							</select>'
                + '						  </td>'
                + '						  <td><input type="text" class="layui-input" placeholder="请输入期望值，可与本API实用到的参数变量对比" name="assertionContent"/></td>'
                + '						  <td><button type="button" class="layui-btn" name="assertion_delete_btn">删除</button></td>'
                + '						</tr>'
                + '					  </tbody>'
                + '					</table>'
                + '					<button type="button" class="layui-btn" name="assertion_add_btn">增加</button>'
                + '				</div>'
                + '			</div>'
                + '		</div>'
                + '	</div>'
                + '</li>';

            var index = 0;
            if (elem.is('ul')) {
                index = elem.find("ul > li").length;
                elem.append(api);
            } else if (elem.is('button')) {
                var $li = elem.closest('li');
                $li.after(api);
            }
            layui.form.render();
            apiTemplate.APITemplateBind();
            layer.msg('添加成功', {time: 700});
        };

        //增加集合点模板
        this.APIConsolidationPointTemplate = function (elem) {
            var api = '<li class="layui-timeline-item" name="scenes_li_element">'
                + '	<i class="layui-icon layui-timeline-axis">&#xe63f;</i>'
                + '	<div class="layui-timeline-content layui-text" style="padding-top: 10px;">'
                + '		<hr style="margin-bottom: 25px;"/>'
                + '		<div class="layui-form-item">'
                + '			<label class="layui-form-label">集合点</label>'
                + '			<div class="layui-form-mid layui-word-aux">请配置有效集合点</div>'
                + '			<div class="layui-input-inline" style="float: right;width:285px;">'
                + '				<button type="button" class="layui-btn" name="api_add_btn">新增API</button>'
                + '				<button type="button" class="layui-btn" name="api_ponit_add_btn">新增集合点</button>'
                + '				<button type="button" class="layui-btn layui-btn-danger" name="api_delete_btn">删除</button>'
                + '			</div>'
                + '		</div>'
                + '		<div class="layui-form-item">'
                + '			<label class="layui-form-label">集合点类型</label>'
                + '			<div class="layui-input-inline" style="width: 120px;">'
                + '				<button type="button" class="layui-btn layui-btn-normal" name="pointType" sampPoint="0">根据时间集合</button>'
                + '			</div>'
                + '			<div class="layui-input-inline" style="width: 120px;">'
                + '				<button type="button" class="layui-btn layui-btn-primary" name="pointType" sampPoint="1">根据用户量集合</button>'
                + '			</div>'
                + '		</div>'
                + '		<div class="layui-form-item">'
                + '			<label class="layui-form-label">等待时间</label>'
                + '			<div class="layui-input-inline">'
                + '				<input name="waitTime" type="text" class="layui-input" lay-verify="number" placeholder="取值范围:(0,3600]" value="0"/>'
                + '			</div>'
                + '			<div class="layui-form-mid layui-word-aux">s</div>'
                + '		</div>'
                + '		<div class="layui-form-item" style="display: none;">'
                + '			<label class="layui-form-label">用户量</label>'
                + '			<div class="layui-input-inline">'
                + '				<input name="waitVuserNum" type="text" class="layui-input" lay-verify="number" placeholder="取值范围(0,1000000]" value="0"/>'
                + '			</div>'
                + '			<div class="layui-form-mid layui-word-aux">分布式施压系统会出现低比例的施压机异常失联的情况，剩余可用机器量级准备完成之后进行放量。</div>'
                + '		</div>'
                + '	</div>'
                + '</li>';

            var $li = elem.closest('li');
            $li.after(api);
            layui.form.render();
            apiTemplate.APITemplateBind();
            layer.msg('添加成功', {time: 700});
        };

        //绑定模板的事件
        this.APITemplateBind = function () {
            //api btn 新增
            $('button[name="api_add_btn"]').off("click").on('click', function () {
                apiTemplate.APITemplate($(this));
            });

            //api btn 删除
            $('button[name="api_delete_btn"]').off("click").on('click', function () {
                var idx = $(this).closest('li').siblings().length;
                var $ul = $(this).closest('ul');
                var $li = $(this).closest('li');
                layer.confirm('确认要删除吗?', {
                    btn: ['删除', '取消'] //按钮
                }, function () {
                    $li.remove();
                    if (idx == 0) {
                        apiTemplate.APITemplate($ul);
                    }
                    layer.msg('删除成功');
                }, function () {
                });
            });

            //api btn 集合点
            $('button[name="api_ponit_add_btn"]').off("click").on('click', function () {
                apiTemplate.APIConsolidationPointTemplate($(this));
            });

            //是否为登陆请求
            form.on('switch(apiTypeSwitchFilter)', function () {
                $(this).parent().find("input[name='type']").val(this.checked ? 0 : 1);
            })
            //header事件
            $('button[name="header_add_btn"]').off("click").on('click', function () {
                var tr = $(this).prev('table[name="header_table"]').find('tbody tr:eq(0)').prop("outerHTML");
                $(this).prev('table[name="header_table"]').append(tr);
                clearData($(this).prev('table[name="header_table"] tr:last'));

                deleteBtnPms($('button[name="header_delete_btn"]'));
            });

            deleteBtnPms($('button[name="header_delete_btn"]'));

            //出参事件
            $('button[name="outpms_add_btn"]').off("click").on('click', function () {
                var tr = $(this).prev('table[name="outpms_table"]').find('tbody tr:eq(0)').prop("outerHTML");
                $(this).prev('table[name="outpms_table"]').append(tr);
                clearData($(this).prev('table[name="outpms_table"] tr:last'));

                deleteBtnPms($('button[name="outpms_delete_btn"]'));
                layui.form.render();
            });

            deleteBtnPms($('button[name="outpms_delete_btn"]'));

            //检查点事件
            $('button[name="assertion_add_btn"]').off("click").on('click', function () {
                var tr = $(this).prev('table[name="assertion_table"]').find('tbody tr:eq(0)').prop("outerHTML");
                $(this).prev('table[name="assertion_table"]').append(tr);
                clearData($(this).prev('table[name="assertion_table"] tr:last'));

                deleteBtnPms($('button[name="assertion_delete_btn"]'));

                layui.form.render();
            });

            deleteBtnPms($('button[name="assertion_delete_btn"]'));

            //删除事件绑定
            function deleteBtnPms(event) {
                event.off("click").on('click', function () {
                    var idx = $(this).parent().parent().siblings().length;
                    if (idx == 0) {
                        clearData($(this).parent().parent());
                    } else {
                        $(this).parent().parent().remove();
                    }
                });
            }

            //清除单元内数据
            function clearData(event) {
                event.find('input').val('');
                event.find('select').val('');
                event.find('checkbox').attr('checked', '');
                event.find('radio').attr('checked', '');
            }

            //radio绑定事件
            $('button[name="pointType"]').off('click').on('click', function () {
                var val = $(this).attr('sampPoint');
                var $div = $(this).closest('li');

                var $other;
                if (val == 0) {
                    $other = $(this).parent().next().find('button');
                    $div.find('input[type="text"][name="waitTime"]').parent().parent().show();
                    $div.find('input[type="text"][name="waitVuserNum"]').parent().parent().hide();
                    $div.find('input[type="text"][name="waitVuserNum"]').val("0");
                } else if (val == '1') {
                    $other = $(this).parent().prev().find('button');
                    $div.find('input[type="text"][name="waitTime"]').parent().parent().hide();
                    $div.find('input[type="text"][name="waitVuserNum"]').parent().parent().show();
                    $div.find('input[type="text"][name="waitTime"]').val("0");
                }
                $(this).removeClass('layui-btn-primary');
                $(this).removeClass('layui-btn-normal');
                $(this).addClass('layui-btn-normal');
                $other.removeClass('layui-btn-primary');
                $other.removeClass('layui-btn-normal');
                $other.addClass('layui-btn-primary');
            });

            //body绑定事件
            $('button[name="contentType"]').off('click').on('click', function () {
                var val = $(this).attr('sampcontenttype');
                var $div = $(this).closest('li');
                var $other;
                if (val == 0) {
                    $other = $(this).parent().next().find('button');
                    $div.find('div[name="params_form_table"]').show();
                    $div.find('textarea[name="paramsBody"]').parent().hide();
                    $div.find('input[type="hidden"][name="params_method_content_type"]').val('application/x-www-form-urlencoded');
                } else if (val == '1') {
                    $other = $(this).parent().prev().find('button');
                    $div.find('div[name="params_form_table"]').hide();
                    $div.find('textarea[name="paramsBody"]').parent().show();
                    $div.find('input[type="hidden"][name="params_method_content_type"]').val('application/json');
                }
                $(this).removeClass('layui-btn-primary');
                $(this).removeClass('layui-btn-normal');
                $(this).addClass('layui-btn-normal');
                $other.removeClass('layui-btn-primary');
                $other.removeClass('layui-btn-normal');
                $other.addClass('layui-btn-primary');
            });
            //body事件
            $('button[name="params_add_btn"]').off("click").on('click', function () {
                var tr = $(this).prev('table[name="params_table"]').find('tbody tr:eq(0)').prop("outerHTML");
                $(this).prev('table[name="params_table"]').append(tr);
                clearData($(this).prev('table[name="params_table"] tr:last'));

                deleteBtnPms($('button[name="params_delete_btn"]'));
                layui.form.render();
            });
            deleteBtnPms($('button[name="params_delete_btn"]'));

            //api method select事件
            form.on('select(apiMethodFilter)', function (data) {
                var $div = $(this).closest('li');
                if (data.value == 'POST' || data.value == 'PUT') {
                    $div.find('li[name="params_tab_body"]').show();
                } else {
                    $div.find('li[name="params_tab_body"]').hide();
                }
            });
        };
    };

    //数据源文件
    $('body').on('click', '[data-file-upload]', function () {
        var loading = layer.load(0, {shade: false, time: 2 * 1000});
        var clientHeight = (document.documentElement.clientHeight) - 95;
        var clientWidth = (document.documentElement.clientWidth) / 2;
        var bgColorHtml = '<form class="layui-form">'
            + '	<div class="layui-form-item">'
            + '		<label class="layui-form-label">文件参数</label>'
            + '		<div class="layui-input-inline" style="width: 120px;float: right;">'
            + '			<button type="button" class="layui-btn layui-btn-primary" name="upload_save_btn">保存</button>'
            + '		</div>'
            + '	</div>'
            + '	<div class="layui-form-item">'
            + '		<div class="layui-input-inline" style="width: 120px;">'
            + '			<button type="button" class="layui-btn" id="dataUpload">'
            + '			  <i class="layui-icon">&#xe67c;</i>上传图片'
            + '			</button>'
            + '		</div>'
            + '		<div class="layui-form-mid layui-word-aux" style="font-size: 12px;">'
            + '            文件单行长度不能超过2w个字符。新增的文件参数，需在要串联链路中“数据配置节点”后才可使用。'
            + '		</div>'
            + '	</div>'
            + '	<div class="layui-form-item">'
            + '		<table class="layui-table" name="upload_table">'
            + '		  <colgroup>'
            + '			<col>'
            + '			<col>'
            + '			<col>'
            + '		  </colgroup>'
            + '		  <thead>'
            + '			<tr>'
            + '			  <th>文件名</th>'
            + '			  <th>上传进度</th>'
            + '			  <th>操作</th>'
            + '			</tr> '
            + '		  </thead>'
            + '		  <tbody>'
            + '		  </tbody>'
            + '		</table>'
            + '	</div>'
            + '	<div class="layui-form-item">'
            + '		<table class="layui-table" name="upload_params_table" style="display:none;">'
            + '		  <colgroup>'
            + '			<col>'
            + '			<col>'
            + '			<col>'
            + '		  </colgroup>'
            + '		  <thead>'
            + '			<tr>'
            + '			  <th>参数名</th>'
            + '			  <th>来源文件</th>'
            + '			  <th>列索引</th>'
            + '			</tr> '
            + '		  </thead>'
            + '		  <tbody>'
            + '		  </tbody>'
            + '		</table>'
            + '	</div>'
            + '</form>';
        var html = '<div class="layuimini-color">\n' +
            '<div class="color-title">\n' +
            '<span>数据源管理</span>\n' +
            '</div>\n' +
            '<div class="color-content">\n' +
            '<ul>\n' + bgColorHtml + '</ul>\n' +
            '</div>\n' +
            '</div>';
        layer.open({
            type: 1,
            title: false,
            closeBtn: 0,
            shade: 0.2,
            anim: 2,
            shadeClose: true,
            id: 'apiTemplateUpload',
            area: [clientWidth + 'px', clientHeight + 'px'],
            offset: 'rb',
            content: html,
            success: function (layero, index) {
                //执行实例
                var uploadInst = upload.render({
                    elem: '#dataUpload' //绑定元素
                    , accept: 'file'
                    , exts: 'csv'
                    , field: 'uploadFile'
                    , url: '/v1/home/uploadDataFile' //上传接口
                    , done: function (res) {
                        //上传完毕回调
                        if (res.code == 1) {
                            layer.msg(res.errMsg);
                        } else {
                            handleUploadSuccess(res.fileName, res.data, res.count, '');
                        }
                    }
                    , error: function () {
                        //请求异常回调
                        layer.msg('上传失败，请刷新后尝试');
                    }
                });
                uploadSaveBtn();
                dataUploadRevert();
            },
            end: function () {
                $('.layuimini-select-bgcolor').removeClass('layui-this');
            }
        });
        layer.close(loading);
    });

    //文件上传成功后更新表格
    function handleUploadSuccess(fileName, path, count, paramsList) {
        var tr = '<tr>'
            + '  <td><input type="hidden" name="filePath" file-name="' + fileName + '" value="' + path + '"/>' + fileName + '</td>'
            + '  <td>完成</td>'
            + '  <td><button type="button" class="layui-btn" name="file_delete_btn" data-file="' + path + '">删除</button></td>'
            + '</tr>';
        $('table[name="upload_table"]').append(tr);
        handleParamTable(fileName, path, count, paramsList);
    }

    function handleParamTable(fileName, path, count, paramsList) {
        var tablestr = '';
        for (var i = 0; i < count; i++) {
            var vul = paramsList != '' ? paramsList[i].name : '';
            tablestr += '<tr name="' + path + '">'
                + '  <td><input type="text" class="layui-input" name="dataKey" placeholder="请输入参数名" value="' + vul + '"/></td>'
                + '  <td>' + fileName + '</td>'
                + '  <td>第' + (i + 1) + '列</td>'
                + '</tr>';
        }
        $('table[name="upload_params_table"]').append(tablestr);
        $('table[name="upload_params_table"]').show();
        $('button[name="file_delete_btn"]').off('click').on('click', function () {
            var file = $(this).attr('data-file');
            var $this = $(this);
            layer.confirm('确认要删除吗?', {
                btn: ['删除', '取消'] //按钮
            }, function () {
                $this.parent().parent().remove();
                $('table[name="upload_params_table"]').find('tr[name="' + file + '"]').remove();
                if ($('table[name="upload_params_table"]').find('tr').length == 0) {
                    $('table[name="upload_params_table"]').hide();
                }
                layer.msg('删除成功');
            }, function () {
            });
        })
    }

    //文件上传以后的保存按钮事件
    function uploadSaveBtn() {
        $('button[name="upload_save_btn"]').off('click').on('click', function () {
            var number = $('input[type="hidden"][name="filePath"]').length;
            if (number == 0) {
                layer.msg('保存成功，请点击<保存配置>来保存数据');
                return;
            }
            var fileDataList = [];
            var filePathEle = $('input[type="hidden"][name="filePath"]');
            for (var i = 0; i < number; i++) {
                var fileData = {};
                fileData.path = $(filePathEle[i]).val();
                fileData.name = $(filePathEle[i]).attr('file-name');
                //收集参数，并校验
                var paramsList = [];
                var paramEle = $('tr[name="' + fileData.path + '"] input[name="dataKey"]');
                for (var k = 0; k < paramEle.length; k++) {
                    var param = {};
                    param.name = $(paramEle[k]).val();
                    param.value = k;
                    paramsList[k] = param;
                }
                fileData.paramsList = paramsList;
                fileDataList[i] = fileData;
            }
            $('#scenes_datafile_json').val(JSON.stringify(fileDataList));
            layer.msg('保存成功，请点击<保存配置>来保存数据');
        });
    }

    //重新打开文件上传时，将保存的数据还原到页面
    function dataUploadRevert() {
        var str = $('#scenes_datafile_json').val();
        if (typeof str == 'undefined' || str == '') {
            return;
        }
        var fileDataList = JSON.parse(str);
        for (var i = 0; i < fileDataList.length; i++) {
            handleUploadSuccess(fileDataList[i].name, fileDataList[i].path, fileDataList[i].paramsList.length, fileDataList[i].paramsList);
        }
    }

    exports(MOD_NAME, apiTemplate);
});