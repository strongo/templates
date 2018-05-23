package layout

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
	"tpl/admin/helper"
)

func Base(body string, title string, js string) string {
	var _buffer bytes.Buffer

	companyName := "深圳思品科技有限公司"

	_buffer.WriteString("\n<!DOCTYPE html>\n<html>\n<head>\n\t<meta charset=\"utf-8\" />\n    <meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n\t<link rel=\"stylesheet\" href=\"/css/bootstrap.min.css\">\n\t<link rel=\"stylesheet\" href=\"/css/dashboard.css\">\n    <!-- HTML5 shim and Respond.js IE8 support of HTML5 elements and media queries -->\n    <!--[if lt IE 9]>\n      <script src=\"https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js\"></script>\n      <script src=\"https://oss.maxcdn.com/libs/respond.js/1.4.2/respond.min.js\"></script>\n    <![endif]-->\n\t<title>")
	_buffer.WriteString((title))
	_buffer.WriteString("</title>\n</head>\n<body>\n    <div class=\"navbar navbar-inverse navbar-fixed-top\" role=\"navigation\">\n      <div class=\"container-fluid\">\n        <div class=\"navbar-header\">\n          <button type=\"button\" class=\"navbar-toggle\" data-toggle=\"collapse\" data-target=\".navbar-collapse\">\n            <span class=\"sr-only\">Toggle navigation</span>\n            <span class=\"icon-bar\"></span>\n            <span class=\"icon-bar\"></span>\n            <span class=\"icon-bar\"></span>\n          </button>\n          <a class=\"navbar-brand\" href=\"http://wethinkwith.com\">")
	_buffer.WriteString(gorazor.HTMLEscape(companyName))
	_buffer.WriteString("</a>我们在<a href=\"http://www.v2ex.com/t/109162\">招聘</a>\n        </div>\n        <div class=\"navbar-collapse collapse\">\n          <ul class=\"nav navbar-nav navbar-right\">\n            <li><a href=\"/admin/setting\">设置</a></li>\n            <li><a href=\"/admin/help\">帮助</a></li>\n            <li><a href=\"/admin/logout\">退出</a></li>\n          </ul>\n          <form class=\"navbar-form navbar-right\">\n            <input type=\"text\" class=\"form-control\" placeholder=\"搜索...\">\n          </form>\n        </div>\n      </div>\n    </div>\n\n    <div class=\"container-fluid\">\n      <div class=\"row\">\n        <div class=\"col-sm-3 col-md-2 sidebar\">\n\t\t\t")
	_buffer.WriteString((helper.Menu()))
	_buffer.WriteString("\n        </div>\n        <div class=\"col-sm-9 col-sm-offset-3 col-md-10 col-md-offset-2 main\">\n          ")
	_buffer.WriteString((body))
	_buffer.WriteString("\n        </div>\n      </div>\n    </div>\n\t<script src=\"/js/jquery.min.js\"></script>\n\t<script src=\"/js/bootstrap.min.js\"></script>\n\t")
	_buffer.WriteString((js))
	_buffer.WriteString("\n  </body>\n</html>")

	return _buffer.String()
}
