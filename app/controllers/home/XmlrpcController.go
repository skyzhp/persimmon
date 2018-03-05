package home

import (
	"github.com/revel/revel"
	"github.com/divan/gorilla-xmlrpc/xml"
	"github.com/gorilla/rpc"
	"github.com/cong5/persimmon/app/service"
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

	methods := map[string]string{
		"blogger.getUsersBlogs":"getUsersBlogs",
		"blogger.deletePost":"deletePost",
		"metaWeblog.newPost":"newPost",
		"metaWeblog.editPost":"editPost",
		"metaWeblog.getPost":"getPost",
		"metaWeblog.getCategories":"getCategories",
		"metaWeblog.newMediaObject":"newMediaObject",
		"metaWeblog.getRecentPosts":"getRecentPosts",
		"wp.newCategory":"newCategory",
	}

	RPC := rpc.NewServer()
	xmlRpcCodec := xml.NewCodec()
	RPC.RegisterCodec(xmlRpcCodec, "text/xml")
	RPC.RegisterService(new(service.XmlRpcService), "shizidXmlRpc")
	return c.RenderXML(methods)
}

func (c XmlRpc) ShowMessage() revel.Result {
	return c.RenderJSON(info.Res{Status: 501, Info: "GET request type error. Please use POST request"})
}
