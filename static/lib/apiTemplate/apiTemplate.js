layui.define(['jquery','form','layer'], function (exports) {
    "use strict";

    var MOD_NAME = 'apiTemplate',
        layer = layui.layer,
        form = layui.form,
        $ = layui.jquery;

    var apiTemplate = new function () {

        //配置信息
        this.config = function () {
            var config = {};
            return config;
        };

        //增加api模板
        this.APITemplate = function (elem) {
            var api = '<li class="layui-timeline-item">'
                +'	<i class="layui-icon layui-timeline-axis">&#xe63f;</i>'
                +'	<div class="layui-timeline-content layui-text" style="padding-top: 10px;">'
                +'		<hr style="margin-bottom: 25px;"/>'
                +'		<div class="layui-form-item">'
                +'			<div class="layui-input-inline">'
                +'				<input type="text" class="layui-input" placeholder="请输入压测api名称">'
                +'			</div>'
                +'			<div class="layui-form-mid layui-word-aux">请配置压测api</div>'
                +'			<div class="layui-input-inline" style="float: right;width:285px;">'
                +'				<button type="button" class="layui-btn" name="api_add_btn">新增API</button>'
                +'              <button type="button" class="layui-btn" name="api_ponit_add_btn">新增集合点</button>'
                +'				<button type="button" class="layui-btn layui-btn-danger" name="api_delete_btn">删除</button>'
                +'			</div>'
                +'		</div>'
                +'		<div class="layui-tab layui-tab-brief">'
                +'			<ul class="layui-tab-title">'
                +'				<li class="layui-this">基本请求信息</li>'
                +'				<li>Header定义</li>'
                +'				<li>出参定义</li>'
                +'				<li>检查点（断言）</li>'
                +'			</ul>'
                +'			<div class="layui-tab-content">'
                +'				<div class="layui-tab-item layui-show">'
                +'					<div class="layui-form-item">'
                +'						<label class="layui-form-label">压测url</label>'
                +'						<div class="layui-input-block">'
                +'							<textarea name="desc" placeholder="请输入一个完整的URL，不能包含空格、换行等非法字符"'
                +'									  class="layui-textarea"></textarea>'
                +'						</div>'
                +'					</div>'
                +'					<div class="layui-form-item">'
                +'						<label class="layui-form-label">请求方式</label>'
                +'						<div class="layui-input-inline">'
                +'							<select name="requestMethod">'
                +'								<option value="GET">GET</option>'
                +'								<option value="POST">POST</option>'
                +'								<option value="PUT">PUT</option>'
                +'								<option value="DELETE">DELETE</option>'
                +'							</select>'
                +'						</div>'
                +'						<label class="layui-form-label">超时时间</label>'
                +'						<div class="layui-input-inline">'
                +'							<input type="text" class="layui-input" name="requestTimeout"'
                +'								   value="5000"/>'
                +'						</div>'
                +'						<div class="layui-form-mid layui-word-aux">ms</div>'
                +'						<label class="layui-form-label">登陆请求</label>'
                +'						<div class="layui-input-inline">'
                +'							<input type="checkbox" lay-skin="switch" lay-text="ON|OFF"'
                +'								   lay-filter=""/>'
                +'						</div>'
                +'						<div class="layui-form-mid layui-word-aux">链路中只有一个登陆请求</div>'
                +'					</div>'
                +'				</div>'
                +'				<div class="layui-tab-item">'
                +'					<table class="layui-table" name="header_table">'
                +'					  <colgroup>'
                +'						<col>'
                +'						<col>'
                +'						<col>'
                +'					  </colgroup>'
                +'					  <thead>'
                +'						<tr>'
                +'						  <th>key</th>'
                +'						  <th>value</th>'
                +'						  <th>操作</th>'
                +'						</tr> '
                +'					  </thead>'
                +'					  <tbody>'
                +'						<tr>'
                +'						  <td><input type="text" class="layui-input" placeholder="输入Header Key"/></td>'
                +'						  <td><input type="text" class="layui-input" placeholder="输入Header Value"/></td>'
                +'						  <td><button type="button" class="layui-btn" name="header_delete_btn">删除</button></td>'
                +'						</tr>'
                +'					  </tbody>'
                +'					</table>'
                +'					<button type="button" class="layui-btn" name="header_add_btn">增加</button>'
                +'				</div>'
                +'				<div class="layui-tab-item">'
                +'					<table class="layui-table" name="outpms_table">'
                +'					  <colgroup>'
                +'						<col>'
                +'						<col>'
                +'						<col>'
                +'						<col>'
                +'						<col>'
                +'					  </colgroup>'
                +'					  <thead>'
                +'						<tr>'
                +'						  <th>出参名</th>'
                +'						  <th>来源</th>'
                +'						  <th>解析表达式</th>'
                +'						  <th>第几个匹配项</th>'
                +'						  <th>操作</th>'
                +'						</tr> '
                +'					  </thead>'
                +'					  <tbody>'
                +'						<tr>'
                +'						  <td><input type="text" class="layui-input" placeholder="输入出参名"/></td>'
                +'						  <td>'
                +'							 <select name="requestSource">'
                +'								<option value="0">Body:TEXT</option>'
                +'								<option value="1">Body:JSON</option>'
                +'								<option value="2">Header:K/V</option>'
                +'								<option value="3">Cookie:K/V</option>'
                +'								<option value="4">响应状态码</option>'
                +'							</select>'
                +'						  </td>'
                +'						  <td><input type="text" class="layui-input" placeholder="输入出参提取表达式"/></td>'
                +'						  <td><input type="text" class="layui-input"/></td>'
                +'						  <td><button type="button" class="layui-btn" name="outpms_delete_btn">删除</button></td>'
                +'						</tr>'
                +'					  </tbody>'
                +'					</table>'
                +'					<button type="button" class="layui-btn" name="outpms_add_btn">增加</button>'
                +'				</div>'
                +'				<div class="layui-tab-item">'
                +'					<table class="layui-table" name="assertion_table">'
                +'					  <colgroup>'
                +'						<col>'
                +'						<col>'
                +'						<col>'
                +'						<col>'
                +'						<col>'
                +'					  </colgroup>'
                +'					  <thead>'
                +'						<tr>'
                +'						  <th>检查点类型</th>'
                +'						  <th>检查对象</th>'
                +'						  <th>检查条件</th>'
                +'						  <th>检查内容</th>'
                +'						  <th>操作</th>'
                +'						</tr> '
                +'					  </thead>'
                +'					  <tbody>'
                +'						<tr>'
                +'						  <td>'
                +'							<select name="requestAssertType">'
                +'								<option value="0">响应header</option>'
                +'								<option value="1">响应状态码</option>'
                +'								<option value="2">响应body</option>'
                +'								<option value="3">出参</option>'
                +'							</select>'
                +'						  </td>'
                +'						  <td>'
                +'							 <input type="text" class="layui-input" placeholder="输入检查对象"/>'
                +'						  </td>'
                +'						  <td>'
                +'							<select name="requestAssertOpt">'
                +'								<option value=">">大于</option>'
                +'								<option value=">=">大于等于</option>'
                +'								<option value="<">小于</option>'
                +'								<option value="<=">小于等于</option>'
                +'								<option value="=">等于</option>'
                +'							</select>'
                +'						  </td>'
                +'						  <td><input type="text" class="layui-input" placeholder="请输入期望值，可与本API实用到的参数变量对比"/></td>'
                +'						  <td><button type="button" class="layui-btn" name="assertion_delete_btn">删除</button></td>'
                +'						</tr>'
                +'					  </tbody>'
                +'					</table>'
                +'					<button type="button" class="layui-btn" name="assertion_add_btn">增加</button>'
                +'				</div>'
                +'			</div>'
                +'		</div>'
                +'	</div>'
                +'</li>';

            var index = 0;
            if(elem.is('ul')){
                index = elem.find("ul > li").length;
                elem.append(api);
            }else if(elem.is('button')){
                var $li = elem.closest('li');
                $li.after(api);
            }
            layui.form.render();
            apiTemplate.APITemplateBind();
            layer.msg('添加成功',{time:700});
        };

        //增加集合点模板
        this.APIConsolidationPointTemplate = function (elem) {
            var api ='<li class="layui-timeline-item">'
                +'	<i class="layui-icon layui-timeline-axis">&#xe63f;</i>'
                +'	<div class="layui-timeline-content layui-text" style="padding-top: 10px;">'
                +'		<hr style="margin-bottom: 25px;"/>'
                +'		<div class="layui-form-item">'
                +'			<label class="layui-form-label">集合点</label>'
                +'			<div class="layui-form-mid layui-word-aux">请配置有效集合点</div>'
                +'			<div class="layui-input-inline" style="float: right;width:285px;">'
                +'				<button type="button" class="layui-btn" name="api_add_btn">新增API</button>'
                +'				<button type="button" class="layui-btn" name="api_ponit_add_btn">新增集合点</button>'
                +'				<button type="button" class="layui-btn layui-btn-danger" name="api_delete_btn">删除</button>'
                +'			</div>'
                +'		</div>'
                +'		<div class="layui-form-item">'
                +'			<label class="layui-form-label">集合点类型</label>'
                +'			<div class="layui-input-inline" style="width: 120px;">'
                +'				<button type="button" class="layui-btn layui-btn-normal" name="pointType" sampPoint="0">根据时间集合</button>'
                +'			</div>'
                +'			<div class="layui-input-inline" style="width: 120px;">'
                +'				<button type="button" class="layui-btn layui-btn-primary" name="pointType" sampPoint="1">根据用户量集合</button>'
                +'			</div>'
                +'		</div>'
                +'		<div class="layui-form-item">'
                +'			<label class="layui-form-label">等待时间</label>'
                +'			<div class="layui-input-inline">'
                +'				<input name="waitTime" type="text" class="layui-input" placeholder="取值范围:(0,3600]" />'
                +'			</div>'
                +'			<div class="layui-form-mid layui-word-aux">s</div>'
                +'		</div>'
                +'		<div class="layui-form-item" style="display: none;">'
                +'			<label class="layui-form-label">用户量</label>'
                +'			<div class="layui-input-inline">'
                +'				<input name="waitVuserNum" type="text" class="layui-input" placeholder="取值范围(0,1000000]"/>'
                +'			</div>'
                +'			<div class="layui-form-mid layui-word-aux">分布式施压系统会出现低比例的施压机异常失联的情况，剩余可用机器量级准备完成之后进行放量。</div>'
                +'		</div>'
                +'	</div>'
                +'</li>';

            var $li = elem.closest('li');
            $li.after(api);
            layui.form.render();
            apiTemplate.APITemplateBind();
            layer.msg('添加成功',{time:700});
        };

        //绑定模板的事件
        this.APITemplateBind = function () {
            //api btn 新增
            $('button[name="api_add_btn"]').off("click").on('click',function () {
                apiTemplate.APITemplate($(this));
            });

            //api btn 删除
            $('button[name="api_delete_btn"]').off("click").on('click',function () {
                var idx = $(this).closest('li').siblings().length;
                var $ul = $(this).closest('ul');
                var $li = $(this).closest('li');
                layer.confirm('确认要删除吗?', {
                    btn: ['删除','取消'] //按钮
                }, function(){
                    $li.remove();
                    if(idx == 0){
                        apiTemplate.APITemplate($ul);
                    }
                    layer.msg('删除成功');
                }, function(){
                });
            });

            //api btn 集合点
            $('button[name="api_ponit_add_btn"]').off("click").on('click',function () {
                apiTemplate.APIConsolidationPointTemplate($(this));
            });

            //header事件
            $('button[name="header_add_btn"]').off("click").on('click',function () {
                var tr = $(this).prev('table[name="header_table"]').find('tbody tr:eq(0)').prop("outerHTML");
                $(this).prev('table[name="header_table"]').append(tr);
                clearData($(this).prev('table[name="header_table"] tr:last'));

                deleteBtnPms($('button[name="header_delete_btn"]'));
            });

            deleteBtnPms($('button[name="header_delete_btn"]'));

            //出参事件
            $('button[name="outpms_add_btn"]').off("click").on('click',function () {
                var tr = $(this).prev('table[name="outpms_table"]').find('tbody tr:eq(0)').prop("outerHTML");
                $(this).prev('table[name="outpms_table"]').append(tr);
                clearData($(this).prev('table[name="outpms_table"] tr:last'));

                deleteBtnPms($('button[name="outpms_delete_btn"]'));
                layui.form.render();
            });

            deleteBtnPms($('button[name="outpms_delete_btn"]'));

            //检查点事件
            $('button[name="assertion_add_btn"]').off("click").on('click',function () {
                var tr = $(this).prev('table[name="assertion_table"]').find('tbody tr:eq(0)').prop("outerHTML");
                $(this).prev('table[name="assertion_table"]').append(tr);
                clearData($(this).prev('table[name="assertion_table"] tr:last'));

                deleteBtnPms($('button[name="assertion_delete_btn"]'));

                layui.form.render();
            });

            deleteBtnPms($('button[name="assertion_delete_btn"]'));

            //删除事件绑定
            function deleteBtnPms(event) {
                event.off("click").on('click',function () {
                    var idx = $(this).parent().parent().siblings().length;
                    if(idx == 0){
                        clearData($(this).parent().parent());
                    }else{
                        $(this).parent().parent().remove();
                    }
                });
            }
            //清除单元内数据
            function clearData(event) {
                event.find('input').val('');
                event.find('select').val('');
                event.find('checkbox').attr('checked','');
                event.find('radio').attr('checked','');
            }

            //radio绑定事件
            $('button[name="pointType"]').off('click').on('click',function () {
                var val = $(this).attr('sampPoint');
                var $div = $(this).closest('li');

                var $other;
                if (val == 0) {
                    $other = $(this).parent().next().find('button');
                    $div.find('input[type="text"][name="waitTime"]').parent().parent().show();
                    $div.find('input[type="text"][name="waitVuserNum"]').parent().parent().hide();
                    $div.find('input[type="text"][name="waitVuserNum"]').val("");
                } else if (val == '1') {
                    $other = $(this).parent().prev().find('button');
                    $div.find('input[type="text"][name="waitTime"]').parent().parent().hide();
                    $div.find('input[type="text"][name="waitVuserNum"]').parent().parent().show();
                    $div.find('input[type="text"][name="waitTime"]').val("");
                }
                $(this).removeClass('layui-btn-primary');
                $(this).removeClass('layui-btn-normal');
                $(this).addClass('layui-btn-normal');
                $other.removeClass('layui-btn-primary');
                $other.removeClass('layui-btn-normal');
                $other.addClass('layui-btn-primary');
            });
        };
    }

    exports(MOD_NAME, apiTemplate);
});