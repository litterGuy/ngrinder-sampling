/**
 * date:2019/08/16
 * author:Mr.Chung
 * description:此处放layui自定义扩展
 */

window.rootPath = (function (src) {
    src = document.scripts[document.scripts.length - 1].src;
    return src.substring(0, src.lastIndexOf("/") + 1);
})();

layui.config({
    base: rootPath + "/",
    version: true
}).extend({
    layuimini: "layuimini/layuimini", // layuimini扩展
    tableSelect: "tableSelect/tableSelect", // tableSelect
    apiTemplate: "apiTemplate/apiTemplate",
});