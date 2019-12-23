//获取id，加载数据
layui.use(['form', 'table', 'element', 'tableSelect', 'apiTemplate', 'jquery', 'layer'], function () {
    var $ = layui.jquery,
        layer = layui.layer,
        apiTemplate = layui.apiTemplate;
    ;

    var id = $('input[name="id"]').val();
    $.get('/v1/scenes/getScenesById?id=' + id, function (data) {
        if (data.code == 1) {
            layer.msg_error(data.errMsg);
            var loading = layer.load(0, {shade: false, time: 2 * 1000});
            layuimini.initPageTitle("/v1/home/scenesList", "场景列表");
            layuimini.initConten("/v1/home/scenesList");
            layuimini.initDevice();
            layer.close(loading);
        } else {
            var scenes = data.data;
            var fileDataList = scenes.fileDataList;
            if (typeof fileDataList !='undefined' && fileDataList.length > 0) {
                $('#scenes_datafile_json').val(JSON.stringify(fileDataList));
            }
            $('input[name="name"]').val(scenes.name);
            $('input[name="agentCount"]').val(scenes.agentCount);
            $('input[name="vuserPerAgent"]').val(scenes.vuserPerAgent);
            $('input[name="processes"]').val(scenes.processes);
            $('input[name="threads"]').val(scenes.threads);

            $("#total_vuser").text($("#agent_count").val() * $("#vuser_per_agent").val());

            //设置threshold
            $('input[type="radio"][name="threshold"]').each(function () {
                if (this.value == scenes.threshold) {
                    $(this).attr('checked', true);
                }
                if (this.value == 'D') {
                    $("#duration").parent().parent().show();
                    $("#run_count").parent().parent().hide();
                } else {
                    $("#duration").parent().parent().hide();
                    $("#run_count").parent().parent().show();
                }
            });

            $('input[name="duration"]').val(typeof scenes.duration == 'undefined' ? '0' : scenes.duration);
            $('input[name="runCount"]').val(typeof scenes.runCount == 'undefined' ? '0' : scenes.runCount);

            //高级设置
            $('input[type="checkbox"][name="advancedSwitch"]').attr("checked", true);
            $("#advanced_config").show();
            if (typeof scenes.samplingInterval != 'undefined') {
                $('select[name="samplingInterval"]').each(function () {
                    if (this.value == scenes.samplingInterval) {
                        $(this).attr("selected", true);
                    }
                });
            }
            $('input[name="ignoreSampleCount"]').val(typeof scenes.ignoreSampleCount == 'undefined' ? '0' : scenes.ignoreSampleCount);
            $('input[name="param"]').val(scenes.param);

            //设置ramp-up
            if (scenes.useRampUp == 'T') {
                $('input[type="checkbox"][name="useRampUp"]').attr("checked", true);
            }
            if (typeof scenes.rampUpType != 'undefined') {
                $('select[name="rampUpType"]').each(function () {
                    if (this.value == scenes.rampUpType) {
                        $(this).attr("selected", true);
                    }
                });
            }
            $('input[name="rampUpInitCount"]').val(typeof scenes.rampUpInitCount == 'undefined' ? '0' : scenes.rampUpInitCount);
            $('input[name="rampUpStep"]').val(typeof scenes.rampUpStep == 'undefined' ? '0' : scenes.rampUpStep);
            $('input[name="rampUpInitSleepTime"]').val(typeof scenes.rampUpInitSleepTime == 'undefined' ? '0' : scenes.rampUpInitSleepTime);
            $('input[name="rampUpIncrementInterval"]').val(typeof scenes.rampUpIncrementInterval == 'undefined' ? '0' : scenes.rampUpIncrementInterval);

            $('input[name="statusCode"]').val(typeof scenes.statusCode == 'undefined' ? '' : scenes.statusCode);
            $('input[name="targetHosts"]').val(scenes.targetHosts);

            //动态设置api部分
            for (var i = 0; i < scenes.requestPmsList.length; i++) {
                var api = scenes.requestPmsList[i];
                if (api.type == 0 || api.type == 1) {
                    apiTemplate.APITemplate($('#api_template'));

                    var $li = $('#api_template li[name="scenes_li_element"]:eq(' + i + ")");
                    //基本请求信息
                    $li.find('input[name="apiName"]').val(api.apiName);
                    $li.find('textarea[name="desc"]').val(api.url);

                    $li.find('select[name="requestMethod"]').find('option[value="' + api.method + '"]').attr("selected", true);
                    $li.find('input[name="requestTimeout"]').val(typeof api.timeout == 'undefined' ? 5000 : api.timeout);

                    if (api.type == 0) {
                        $('input[type="checkbox"][name="type"]').attr("checked", true);
                    }
                    //设置body
                    if (api.method == 'POST' || api.method == 'PUT') {
                        $li.find('li[name="params_tab_body"]').show();
                        $li.find('input[name="params_method_content_type"]').val(api.contentType);
                        if (api.contentType == 'application/json') {
                            $li.find('div[name="params_form_table"]').hide();
                            $li.find('textarea[name="paramsBody"]').parent().show();

                            var $this = $li.find('button[name="contentType"][sampcontenttype="1"]');
                            var $other = $li.find('button[name="contentType"][sampcontenttype="0"]');
                            $this.removeClass('layui-btn-primary');
                            $this.removeClass('layui-btn-normal');
                            $this.addClass('layui-btn-normal');
                            $other.removeClass('layui-btn-primary');
                            $other.removeClass('layui-btn-normal');
                            $other.addClass('layui-btn-primary');
                        }
                    }
                    if (typeof api.paramList != 'undefined' && api.paramList.length > 0) {
                        for (var k = 0; k < api.paramList.length; k++) {
                            var $tr = $li.find('table[name="params_table"] tr:eq(' + (k + 1) + ')');
                            var param = api.paramList[k];
                            if ($tr.length > 0) {
                                $tr.find('input[name="paramsName"]').val(param.name);
                                $tr.find('input[name="paramsValue"]').val(param.value);
                            } else {
                                $li.find('button[name="params_add_btn"]').trigger("click");
                                var $tr = $li.find('table[name="params_table"] tr:eq(' + (k + 1) + ')');
                                $tr.find('input[name="paramsName"]').val(param.name);
                                $tr.find('input[name="paramsValue"]').val(param.value);
                            }
                        }
                    }
                    if (typeof api.body != 'undefined' && api.body.length > 0) {
                        $li.find('textarea[name="paramsBody"]').val(api.body);
                    }
                    //设置header
                    if (typeof api.headerList != 'undefined' && api.headerList.length > 0) {
                        for (var k = 0; k < api.headerList.length; k++) {
                            var $tr = $li.find('table[name="header_table"] tr:eq(' + (k + 1) + ')');
                            var header = api.headerList[k];
                            if ($tr.length > 0) {
                                $tr.find('input[name="headerName"]').val(header.name);
                                $tr.find('input[name="headerValue"]').val(header.value);
                            } else {
                                $li.find('button[name="header_add_btn"]').trigger("click");
                                var $tr = $li.find('table[name="header_table"] tr:eq(' + (k + 1) + ')');
                                $tr.find('input[name="headerName"]').val(header.name);
                                $tr.find('input[name="headerValue"]').val(header.value);
                            }
                        }
                    }
                    //设置出参
                    if (typeof api.outParamsList != 'undefined' && api.outParamsList.length > 0) {
                        for (var k = 0; k < api.outParamsList.length; k++) {
                            var $tr = $li.find('table[name="outpms_table"] tr:eq(' + (k + 1) + ')');
                            var outPms = api.outParamsList[k];
                            if ($tr.length > 0) {
                                $tr.find('input[name="outName"]').val(outPms.name);
                                $tr.find('select[name="outSource"]').find('option[value="' + outPms.source + '"]').attr("selected", true);
                                $tr.find('input[name="outResolveExpress"]').val(outPms.resolveExpress);
                                $tr.find('input[name="outIndex"]').val(typeof outPms.index == 'undefined' ? '0' : outPms.index);
                            } else {
                                $li.find('button[name="outpms_add_btn"]').trigger("click");
                                var $tr = $li.find('table[name="outpms_table"] tr:eq(' + (k + 1) + ')');
                                $tr.find('input[name="outName"]').val(outPms.name);
                                $tr.find('select[name="outSource"]').find('option[value="' + outPms.source + '"]').attr("selected", true);
                                $tr.find('input[name="outResolveExpress"]').val(outPms.resolveExpress);
                                $tr.find('input[name="outIndex"]').val(typeof outPms.index == 'undefined' ? '0' : outPms.index);
                            }
                        }
                    }
                    //设置检查点
                    if (typeof api.assertionList != 'undefined' && api.assertionList.length > 0) {
                        for (var k = 0; k < api.assertionList.length; k++) {
                            var $tr = $li.find('table[name="assertion_table"] tr:eq(' + (k + 1) + ')');
                            var assertion = api.assertionList[k];
                            if ($tr.length > 0) {
                                $tr.find('select[name="assertionType"]').find('option[value="' + assertion.type + '"]').attr("selected", true);
                                $tr.find('input[name="assertionName"]').val(assertion.name);
                                $tr.find('select[name="assertionFactor"]').find('option[value="' + assertion.factor + '"]').attr("selected", true);
                                $tr.find('input[name="assertionContent"]').val(assertion.content);
                            } else {
                                $li.find('button[name="assertion_add_btn"]').trigger("click");
                                var $tr = $li.find('table[name="assertion_table"] tr:eq(' + (k + 1) + ')');
                                $tr.find('select[name="assertionType"]').find('option[value="' + assertion.type + '"]').attr("selected", true);
                                $tr.find('input[name="assertionName"]').val(assertion.name);
                                $tr.find('select[name="assertionFactor"]').find('option[value="' + assertion.factor + '"]').attr("selected", true);
                                $tr.find('input[name="assertionContent"]').val(assertion.content);
                            }
                        }
                    }
                } else if (api.type == 2) {
                    apiTemplate.APIConsolidationPointTemplate($('#api_template'));
                    var $li = $('#api_template li[name="scenes_li_element"]:eq(' + i + ")");
                    var $this, $other;
                    if (typeof api.waitTime != 'undefined' && api.waitTime > 0) {
                        $li.find('input[type="text"][name="waitTime"]').parent().parent().show();
                        $li.find('input[type="text"][name="waitVuserNum"]').parent().parent().hide();
                        $li.find('input[type="text"][name="waitTime"]').val(api.waitTime);
                        $this = $li.find('button[name="pointType"][samppoint="0"]');
                        $other = $li.find('button[name="pointType"][samppoint="1"]');
                    } else {
                        $li.find('input[type="text"][name="waitTime"]').parent().parent().hide();
                        $li.find('input[type="text"][name="waitVuserNum"]').parent().parent().show();
                        $li.find('input[type="text"][name="waitTime"]').val(api.waitVuserNum);
                        $this = $li.find('button[name="pointType"][samppoint="1"]');
                        $other = $li.find('button[name="pointType"][samppoint="0"]');
                    }
                    $this.removeClass('layui-btn-primary');
                    $this.removeClass('layui-btn-normal');
                    $this.addClass('layui-btn-normal');
                    $other.removeClass('layui-btn-primary');
                    $other.removeClass('layui-btn-normal');
                    $other.addClass('layui-btn-primary');
                }
            }

            layui.form.render();
        }
    });
});