package cases

import (
	"bytes"
	"dm"
	"github.com/sipin/gorazor/gorazor"
	"zfw/models"
	. "zfw/tplhelper"
)

func Scope(obj *models.Widget) string {
	var _buffer bytes.Buffer

	data, dmType := dm.GetData(obj.PlaceHolder)

	if dmType == "simple" {
		obj.StringList = data.([]string)

		_buffer.WriteString("<div>")
		_buffer.WriteString((SelectPk(obj)))
		_buffer.WriteString("</div>")

	} else {
		node := data.(*dm.DMTree)

		_buffer.WriteString("<div class=\"form-group ")
		_buffer.WriteString(gorazor.HTMLEscape(GetErrorClass(obj)))
		_buffer.WriteString("\">\n  <label for=\"")
		_buffer.WriteString(gorazor.HTMLEscape(obj.Name))
		_buffer.WriteString("\" class=\"col-sm-2 control-label\">")
		_buffer.WriteString(gorazor.HTMLEscape(obj.Label))
		_buffer.WriteString("</label>\n  <div class=\"col-sm-10\">\n    <select class=\"form-control\" name=\"")
		_buffer.WriteString(gorazor.HTMLEscape(obj.Name))
		_buffer.WriteString("\" ")
		_buffer.WriteString(gorazor.HTMLEscape(BoolStr(obj.Disabled, "disabled")))
		_buffer.WriteString(">\n      ")
		for _, option := range node.Keys {
			if values, ok := node.Values[option]; ok {

				_buffer.WriteString("<optgroup label=\"")
				_buffer.WriteString(gorazor.HTMLEscape(option))
				_buffer.WriteString("\">\n        ")
				for _, value := range values {
					if value == obj.Value {

						_buffer.WriteString("<option selected>")
						_buffer.WriteString(gorazor.HTMLEscape(value))
						_buffer.WriteString("</option>")

					} else {

						_buffer.WriteString("<option>")
						_buffer.WriteString(gorazor.HTMLEscape(value))
						_buffer.WriteString("</option>")

					}
				}
				_buffer.WriteString("\n      </optgroup>")

			} else {
				if option == obj.Value {

					_buffer.WriteString("<option selected>")
					_buffer.WriteString(gorazor.HTMLEscape(option))
					_buffer.WriteString("</option>")

				} else {

					_buffer.WriteString("<option>")
					_buffer.WriteString(gorazor.HTMLEscape(option))
					_buffer.WriteString("</option>")

				}
			}
		}
		_buffer.WriteString("\n    </select>\n    ")
		if obj.ErrorMsg != "" {

			_buffer.WriteString("<span class=\"label label-danger\">")
			_buffer.WriteString(gorazor.HTMLEscape(obj.ErrorMsg))
			_buffer.WriteString("</span>")

		}
		_buffer.WriteString("\n  </div>\n</div>")
	}

	return _buffer.String()
}
