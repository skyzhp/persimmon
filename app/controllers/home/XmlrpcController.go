package home

import (
	"github.com/revel/revel"
	"github.com/cong5/persimmon/app/info"
)

type XmlRpc struct {
	BaseHomeController
}

type Methods struct {
	action    string
	mapAction string
}

func (c XmlRpc) Index() revel.Result {
	/*
	xmlFunc := map[string]string{
		"blogger.getUsersBlogs":     "GetUsersBlogs",
		"blogger.deletePost":        "DeletePost",
		"metaWeblog.newPost":        "NewPost",
		"metaWeblog.editPost":       "EditPost",
		"metaWeblog.getPost":        "GetPost",
		"metaWeblog.getCategories":  "GetCategories",
		"metaWeblog.newMediaObject": "NewMediaObject",
		"metaWeblog.getRecentPosts": "GetRecentPosts",
		"wp.newCategory":            "NewCategory",
	}

	body := c.Request.GetBody()
	methodName, extracts, err, fault := xmlrpc.Unmarshal(body)
	if err != nil {
		revel.INFO.Printf("Decode Error: %s %s", err,fault)
	}*/

	/*
	funcs := xmlrpc.NewFuncs(len(xmlFunc))
	funcs.Bind(methodName, xmlFunc["adduser"])
	retValue, _ := funcs.Call("adduser", extracts)
	revel.INFO.Println(retValue)
	*/
	/*
	revel.INFO.Println(extracts[1])
	revel.INFO.Println(extracts[1])
	revel.INFO.Println(extracts[2])
	revel.INFO.Println(method)
	revel.INFO.Println(fault)
	*/

	return c.RenderJSON("")
}

func (c XmlRpc) ShowMessage() revel.Result {
	return c.RenderJSON(info.Res{Status: 501, Info: "GET request type error. Please use POST request"})
}

func (c XmlRpc) GetUsersBlogs(params []string) {
	revel.INFO.Println("GetUsersBlogs")
	revel.INFO.Println(params)
}

func (c XmlRpc) GetPost(params map[string]string) {

}
func (c XmlRpc) NewPost(params map[string]string) {

}
func (c XmlRpc) EditPost(params map[string]string) {

}
func (c XmlRpc) DeletePost(params map[string]string) {

}
func (c XmlRpc) NewCategory(params map[string]string) {

}
func (c XmlRpc) GetCategories(params map[string]string) {

}
func (c XmlRpc) NewMediaObject(params map[string]string) {

}
func (c XmlRpc) GetRecentPosts(params map[string]string) {

}
