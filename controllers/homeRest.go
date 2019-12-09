package controllers

//http://layuimini-onepage.99php.cn/#/page/icon.html

type HomeController struct {
	BaseController
}

// @router	/index	[get]
func (h *HomeController) Index() {
	h.Data["userId"] = h.GetSession(SESSION_NAME)
	h.TplName = "index.html"
}

// @router	/scenesList	[get]
func (h *HomeController) ScenesList() {
	h.TplName = "scenes_list.html"
}
