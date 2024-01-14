// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.501
package home

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/muhhae/lorem-ipsum/internal/views/util"
import "strconv"

func Post(username string, content string, likeCount int, commentCount int) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-full card rounded-none mb-10 bg-transparent p-4 text-base-content\"><figure class=\"w-full h-full\"><img class=\"w-full h-full object-cover\" src=\"https://daisyui.com/images/stock/photo-1494232410401-ad00d5433cfa.jpg\" alt=\"Album\"></figure><div class=\"card-body \"><h2 class=\"card-title text-base-content font-bold\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(username)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal\views\home\post.templ`, Line: 15, Col: 64}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(content)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal\views\home\post.templ`, Line: 16, Col: 12}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"card-actions flex flex-col sm:flex-row w-full items-center justify-center mb-4\"><button class=\"btn hover:text-primary btn-ghost btn-md text-xl w-40\"><svg class=\"h-3/4\" viewBox=\"0 0 32 32\" version=\"1.1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xmlns:sketch=\"http://www.bohemiancoding.com/sketch/ns\" fill=\"currentColor\"><g id=\"SVGRepo_bgCarrier\" stroke-width=\"0\"></g><g id=\"SVGRepo_tracerCarrier\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></g><g id=\"SVGRepo_iconCarrier\"><title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var4 := `arrow-up-circle`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var4)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><desc>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var5 := `Created with Sketch Beta.`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var5)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</desc> <defs></defs> <g id=\"Page-1\" stroke=\"none\" stroke-width=\"1\" fill=\"currentColor\" fill-rule=\"evenodd\" sketch:type=\"MSPage\"><g id=\"Icon-Set-Filled\" sketch:type=\"MSLayerGroup\" transform=\"translate(-362.000000, -1089.000000)\" fill=\"currentColor\"><path d=\"M384.535,1105.54 C384.145,1105.93 383.512,1105.93 383.121,1105.54 L379,1101.41 L379,1112 C379,1112.55 378.553,1113 378,1113 C377.447,1113 377,1112.55 377,1112 L377,1101.41 L372.879,1105.54 C372.488,1105.93 371.854,1105.93 371.465,1105.54 C371.074,1105.14 371.074,1104.51 371.465,1104.12 L377.121,1098.46 C377.361,1098.22 377.689,1098.15 378,1098.21 C378.311,1098.15 378.639,1098.22 378.879,1098.46 L384.535,1104.12 C384.926,1104.51 384.926,1105.14 384.535,1105.54 L384.535,1105.54 Z M378,1089 C369.163,1089 362,1096.16 362,1105 C362,1113.84 369.163,1121 378,1121 C386.837,1121 394,1113.84 394,1105 C394,1096.16 386.837,1089 378,1089 L378,1089 Z\" id=\"arrow-up-circle\" sketch:type=\"MSShapeGroup\"></path></g></g></g></svg> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var6 := `Yahaha`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var6)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button><div class=\"h-12 w-24 rounded-xl bg-primary font-mono font-black text-2xl text-primary-content text-center flex items-center justify-center\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(util.Format(likeCount))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal\views\home\post.templ`, Line: 24, Col: 28}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><button class=\"btn btn-ghost hover:text-primary btn-md text-xl w-40\"><svg class=\"sm:hidden block h-3/4\" viewBox=\"0 0 32 32\" version=\"1.1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xmlns:sketch=\"http://www.bohemiancoding.com/sketch/ns\" fill=\"currentColor\" transform=\"matrix(1, 0, 0, -1, 0, 0)\"><g id=\"SVGRepo_bgCarrier\" stroke-width=\"0\"></g><g id=\"SVGRepo_tracerCarrier\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></g><g id=\"SVGRepo_iconCarrier\"><title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var8 := `arrow-up-circle`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var8)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><desc>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var9 := `Created with Sketch Beta.`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var9)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</desc> <defs></defs> <g id=\"Page-1\" stroke=\"none\" stroke-width=\"1\" fill=\"currentColor\" fill-rule=\"evenodd\" sketch:type=\"MSPage\"><g id=\"Icon-Set-Filled\" sketch:type=\"MSLayerGroup\" transform=\"translate(-362.000000, -1089.000000)\" fill=\"currentColor\"><path d=\"M384.535,1105.54 C384.145,1105.93 383.512,1105.93 383.121,1105.54 L379,1101.41 L379,1112 C379,1112.55 378.553,1113 378,1113 C377.447,1113 377,1112.55 377,1112 L377,1101.41 L372.879,1105.54 C372.488,1105.93 371.854,1105.93 371.465,1105.54 C371.074,1105.14 371.074,1104.51 371.465,1104.12 L377.121,1098.46 C377.361,1098.22 377.689,1098.15 378,1098.21 C378.311,1098.15 378.639,1098.22 378.879,1098.46 L384.535,1104.12 C384.926,1104.51 384.926,1105.14 384.535,1105.54 L384.535,1105.54 Z M378,1089 C369.163,1089 362,1096.16 362,1105 C362,1113.84 369.163,1121 378,1121 C386.837,1121 394,1113.84 394,1105 C394,1096.16 386.837,1089 378,1089 L378,1089 Z\" id=\"arrow-up-circle\" sketch:type=\"MSShapeGroup\"></path></g></g></g></svg> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var10 := `Ahahay`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var10)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <svg class=\"sm:block hidden h-3/4\" viewBox=\"0 0 32 32\" version=\"1.1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xmlns:sketch=\"http://www.bohemiancoding.com/sketch/ns\" fill=\"currentColor\" transform=\"matrix(1, 0, 0, -1, 0, 0)\"><g id=\"SVGRepo_bgCarrier\" stroke-width=\"0\"></g><g id=\"SVGRepo_tracerCarrier\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></g><g id=\"SVGRepo_iconCarrier\"><title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var11 := `arrow-up-circle`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var11)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><desc>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var12 := `Created with Sketch Beta.`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var12)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</desc> <defs></defs> <g id=\"Page-1\" stroke=\"none\" stroke-width=\"1\" fill=\"currentColor\" fill-rule=\"evenodd\" sketch:type=\"MSPage\"><g id=\"Icon-Set-Filled\" sketch:type=\"MSLayerGroup\" transform=\"translate(-362.000000, -1089.000000)\" fill=\"currentColor\"><path d=\"M384.535,1105.54 C384.145,1105.93 383.512,1105.93 383.121,1105.54 L379,1101.41 L379,1112 C379,1112.55 378.553,1113 378,1113 C377.447,1113 377,1112.55 377,1112 L377,1101.41 L372.879,1105.54 C372.488,1105.93 371.854,1105.93 371.465,1105.54 C371.074,1105.14 371.074,1104.51 371.465,1104.12 L377.121,1098.46 C377.361,1098.22 377.689,1098.15 378,1098.21 C378.311,1098.15 378.639,1098.22 378.879,1098.46 L384.535,1104.12 C384.926,1104.51 384.926,1105.14 384.535,1105.54 L384.535,1105.54 Z M378,1089 C369.163,1089 362,1096.16 362,1105 C362,1113.84 369.163,1121 378,1121 C386.837,1121 394,1113.84 394,1105 C394,1096.16 386.837,1089 378,1089 L378,1089 Z\" id=\"arrow-up-circle\" sketch:type=\"MSShapeGroup\"></path></g></g></g></svg></button></div><div class=\"divider\"><button class=\"btn btn-ghost\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var13 := `Show `
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var13)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var14 string
		templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(util.Format(commentCount))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal\views\home\post.templ`, Line: 33, Col: 66}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var15 := `Comments`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var15)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

type ReactData struct {
	PostID    string
	LikeCount int
	Value     int
}

type PostData struct {
	Username     string
	Content      string
	ImgSrc       []string
	CommentCount int
	ReactStruct  ReactData
}

func postUrl(iteration int) string {
	return "/api/v1/post/Default?iteration=" + strconv.Itoa(iteration)
}

func ManyPost(postDatas []PostData, iteration int) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var16 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var16 == nil {
			templ_7745c5c3_Var16 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		for _, postData := range postDatas {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-full card rounded-none mb-10 bg-transparent p-4 text-base-content\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, imgSrc := range postData.ImgSrc {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<figure class=\"w-full h-full\"><img class=\"w-full h-full object-cover\" src=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(imgSrc))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" alt=\"Meme\"></figure>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"card-body \"><h2 class=\"card-title text-base-content font-bold\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var17 string
			templ_7745c5c3_Var17, templ_7745c5c3_Err = templ.JoinStringErrs(postData.Username)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal\views\home\post.templ`, Line: 69, Col: 74}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var17))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var18 string
			templ_7745c5c3_Var18, templ_7745c5c3_Err = templ.JoinStringErrs(postData.Content)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal\views\home\post.templ`, Line: 70, Col: 22}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var18))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = ReactSection(postData.ReactStruct).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"divider\"><button class=\"btn btn-ghost\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var19 := `Show `
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var19)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var20 string
			templ_7745c5c3_Var20, templ_7745c5c3_Err = templ.JoinStringErrs(util.Format(postData.CommentCount))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal\views\home\post.templ`, Line: 74, Col: 76}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var20))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var21 := `Comments`
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var21)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"loader\" hx-trigger=\"intersect\" hx-swap=\"outerHTML\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(postUrl(iteration + 1)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"loading loading-spinner loading-lg mx-auto\"></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func ReactSection(reactData ReactData) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var22 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var22 == nil {
			templ_7745c5c3_Var22 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"react-section card-actions flex flex-col sm:flex-row w-full items-center justify-center mb-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = LikeButton(reactData.Value == 1, reactData.PostID).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = LikeCount(reactData.LikeCount, reactData.PostID).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = DislikeButton(reactData.Value == -1, reactData.PostID).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func LikeCount(likeCount int, postID string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var23 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var23 == nil {
			templ_7745c5c3_Var23 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString("/api/v1/reaction/count/" + postID))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-trigger=\"intersect, update\" hx-swap=\"innerHTML\" class=\"like-count h-12 w-24 rounded-xl bg-primary font-mono font-black text-2xl text-primary-content text-center flex items-center justify-center\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var24 string
		templ_7745c5c3_Var24, templ_7745c5c3_Err = templ.JoinStringErrs(util.Format(likeCount))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal\views\home\post.templ`, Line: 104, Col: 26}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var24))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func LikeButton(liked bool, postID string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var25 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var25 == nil {
			templ_7745c5c3_Var25 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if liked {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button hx-post=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString("/api/v1/reaction/react/" + postID + "?value=0"))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-swap=\"none\" hx-on::after-on-load=\"reactAndRefresh(event)\" class=\"like btn text-primary btn-ghost btn-md text-xl w-40\"><svg class=\"h-3/4\" viewBox=\"0 0 32 32\" version=\"1.1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xmlns:sketch=\"http://www.bohemiancoding.com/sketch/ns\" fill=\"currentColor\"><g id=\"SVGRepo_bgCarrier\" stroke-width=\"0\"></g><g id=\"SVGRepo_tracerCarrier\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></g><g id=\"SVGRepo_iconCarrier\"><title>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var26 := `arrow-up-circle`
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var26)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><desc>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var27 := `Created with Sketch Beta.`
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var27)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</desc> <defs></defs> <g id=\"Page-1\" stroke=\"none\" stroke-width=\"1\" fill=\"currentColor\" fill-rule=\"evenodd\" sketch:type=\"MSPage\"><g id=\"Icon-Set-Filled\" sketch:type=\"MSLayerGroup\" transform=\"translate(-362.000000, -1089.000000)\" fill=\"currentColor\"><path d=\"M384.535,1105.54 C384.145,1105.93 383.512,1105.93 383.121,1105.54 L379,1101.41 L379,1112 C379,1112.55 378.553,1113 378,1113 C377.447,1113 377,1112.55 377,1112 L377,1101.41 L372.879,1105.54 C372.488,1105.93 371.854,1105.93 371.465,1105.54 C371.074,1105.14 371.074,1104.51 371.465,1104.12 L377.121,1098.46 C377.361,1098.22 377.689,1098.15 378,1098.21 C378.311,1098.15 378.639,1098.22 378.879,1098.46 L384.535,1104.12 C384.926,1104.51 384.926,1105.14 384.535,1105.54 L384.535,1105.54 Z M378,1089 C369.163,1089 362,1096.16 362,1105 C362,1113.84 369.163,1121 378,1121 C386.837,1121 394,1113.84 394,1105 C394,1096.16 386.837,1089 378,1089 L378,1089 Z\" id=\"arrow-up-circle\" sketch:type=\"MSShapeGroup\"></path></g></g></g></svg> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var28 := `Yahaha`
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var28)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button hx-post=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString("/api/v1/reaction/react/" + postID + "?value=1"))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-swap=\"none\" hx-on::after-on-load=\"reactAndRefresh(event)\" class=\"like btn btn-ghost btn-md text-xl w-40\"><svg class=\"h-3/4\" viewBox=\"0 0 32 32\" version=\"1.1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xmlns:sketch=\"http://www.bohemiancoding.com/sketch/ns\" fill=\"currentColor\"><g id=\"SVGRepo_bgCarrier\" stroke-width=\"0\"></g><g id=\"SVGRepo_tracerCarrier\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></g><g id=\"SVGRepo_iconCarrier\"><title>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var29 := `arrow-up-circle`
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var29)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><desc>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var30 := `Created with Sketch Beta.`
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var30)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</desc> <defs></defs> <g id=\"Page-1\" stroke=\"none\" stroke-width=\"1\" fill=\"currentColor\" fill-rule=\"evenodd\" sketch:type=\"MSPage\"><g id=\"Icon-Set-Filled\" sketch:type=\"MSLayerGroup\" transform=\"translate(-362.000000, -1089.000000)\" fill=\"currentColor\"><path d=\"M384.535,1105.54 C384.145,1105.93 383.512,1105.93 383.121,1105.54 L379,1101.41 L379,1112 C379,1112.55 378.553,1113 378,1113 C377.447,1113 377,1112.55 377,1112 L377,1101.41 L372.879,1105.54 C372.488,1105.93 371.854,1105.93 371.465,1105.54 C371.074,1105.14 371.074,1104.51 371.465,1104.12 L377.121,1098.46 C377.361,1098.22 377.689,1098.15 378,1098.21 C378.311,1098.15 378.639,1098.22 378.879,1098.46 L384.535,1104.12 C384.926,1104.51 384.926,1105.14 384.535,1105.54 L384.535,1105.54 Z M378,1089 C369.163,1089 362,1096.16 362,1105 C362,1113.84 369.163,1121 378,1121 C386.837,1121 394,1113.84 394,1105 C394,1096.16 386.837,1089 378,1089 L378,1089 Z\" id=\"arrow-up-circle\" sketch:type=\"MSShapeGroup\"></path></g></g></g></svg> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var31 := `Yahaha`
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var31)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func DislikeButton(disliked bool, postID string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var32 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var32 == nil {
			templ_7745c5c3_Var32 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if disliked {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button hx-post=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString("/api/v1/reaction/react/" + postID + "?value=0"))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-swap=\"none\" hx-on::after-on-load=\"reactAndRefresh(event)\" class=\"dislike btn btn-ghost text-primary btn-md text-xl w-40\"><svg class=\"sm:hidden block h-3/4\" viewBox=\"0 0 32 32\" version=\"1.1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xmlns:sketch=\"http://www.bohemiancoding.com/sketch/ns\" fill=\"currentColor\" transform=\"matrix(1, 0, 0, -1, 0, 0)\"><g id=\"SVGRepo_bgCarrier\" stroke-width=\"0\"></g><g id=\"SVGRepo_tracerCarrier\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></g><g id=\"SVGRepo_iconCarrier\"><title>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var33 := `arrow-up-circle`
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var33)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><desc>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var34 := `Created with Sketch Beta.`
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var34)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</desc> <defs></defs> <g id=\"Page-1\" stroke=\"none\" stroke-width=\"1\" fill=\"currentColor\" fill-rule=\"evenodd\" sketch:type=\"MSPage\"><g id=\"Icon-Set-Filled\" sketch:type=\"MSLayerGroup\" transform=\"translate(-362.000000, -1089.000000)\" fill=\"currentColor\"><path d=\"M384.535,1105.54 C384.145,1105.93 383.512,1105.93 383.121,1105.54 L379,1101.41 L379,1112 C379,1112.55 378.553,1113 378,1113 C377.447,1113 377,1112.55 377,1112 L377,1101.41 L372.879,1105.54 C372.488,1105.93 371.854,1105.93 371.465,1105.54 C371.074,1105.14 371.074,1104.51 371.465,1104.12 L377.121,1098.46 C377.361,1098.22 377.689,1098.15 378,1098.21 C378.311,1098.15 378.639,1098.22 378.879,1098.46 L384.535,1104.12 C384.926,1104.51 384.926,1105.14 384.535,1105.54 L384.535,1105.54 Z M378,1089 C369.163,1089 362,1096.16 362,1105 C362,1113.84 369.163,1121 378,1121 C386.837,1121 394,1113.84 394,1105 C394,1096.16 386.837,1089 378,1089 L378,1089 Z\" id=\"arrow-up-circle\" sketch:type=\"MSShapeGroup\"></path></g></g></g></svg> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var35 := `Ahahay`
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var35)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <svg class=\"sm:block hidden h-3/4\" viewBox=\"0 0 32 32\" version=\"1.1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xmlns:sketch=\"http://www.bohemiancoding.com/sketch/ns\" fill=\"currentColor\" transform=\"matrix(1, 0, 0, -1, 0, 0)\"><g id=\"SVGRepo_bgCarrier\" stroke-width=\"0\"></g><g id=\"SVGRepo_tracerCarrier\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></g><g id=\"SVGRepo_iconCarrier\"><title>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var36 := `arrow-up-circle`
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var36)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><desc>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var37 := `Created with Sketch Beta.`
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var37)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</desc> <defs></defs> <g id=\"Page-1\" stroke=\"none\" stroke-width=\"1\" fill=\"currentColor\" fill-rule=\"evenodd\" sketch:type=\"MSPage\"><g id=\"Icon-Set-Filled\" sketch:type=\"MSLayerGroup\" transform=\"translate(-362.000000, -1089.000000)\" fill=\"currentColor\"><path d=\"M384.535,1105.54 C384.145,1105.93 383.512,1105.93 383.121,1105.54 L379,1101.41 L379,1112 C379,1112.55 378.553,1113 378,1113 C377.447,1113 377,1112.55 377,1112 L377,1101.41 L372.879,1105.54 C372.488,1105.93 371.854,1105.93 371.465,1105.54 C371.074,1105.14 371.074,1104.51 371.465,1104.12 L377.121,1098.46 C377.361,1098.22 377.689,1098.15 378,1098.21 C378.311,1098.15 378.639,1098.22 378.879,1098.46 L384.535,1104.12 C384.926,1104.51 384.926,1105.14 384.535,1105.54 L384.535,1105.54 Z M378,1089 C369.163,1089 362,1096.16 362,1105 C362,1113.84 369.163,1121 378,1121 C386.837,1121 394,1113.84 394,1105 C394,1096.16 386.837,1089 378,1089 L378,1089 Z\" id=\"arrow-up-circle\" sketch:type=\"MSShapeGroup\"></path></g></g></g></svg></button>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button hx-post=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString("/api/v1/reaction/react/" + postID + "?value=-1"))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-swap=\"none\" hx-on::after-on-load=\"reactAndRefresh(event)\" class=\"dislike btn btn-ghost btn-md text-xl w-40\"><svg class=\"sm:hidden block h-3/4\" viewBox=\"0 0 32 32\" version=\"1.1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xmlns:sketch=\"http://www.bohemiancoding.com/sketch/ns\" fill=\"currentColor\" transform=\"matrix(1, 0, 0, -1, 0, 0)\"><g id=\"SVGRepo_bgCarrier\" stroke-width=\"0\"></g><g id=\"SVGRepo_tracerCarrier\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></g><g id=\"SVGRepo_iconCarrier\"><title>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var38 := `arrow-up-circle`
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var38)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><desc>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var39 := `Created with Sketch Beta.`
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var39)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</desc> <defs></defs> <g id=\"Page-1\" stroke=\"none\" stroke-width=\"1\" fill=\"currentColor\" fill-rule=\"evenodd\" sketch:type=\"MSPage\"><g id=\"Icon-Set-Filled\" sketch:type=\"MSLayerGroup\" transform=\"translate(-362.000000, -1089.000000)\" fill=\"currentColor\"><path d=\"M384.535,1105.54 C384.145,1105.93 383.512,1105.93 383.121,1105.54 L379,1101.41 L379,1112 C379,1112.55 378.553,1113 378,1113 C377.447,1113 377,1112.55 377,1112 L377,1101.41 L372.879,1105.54 C372.488,1105.93 371.854,1105.93 371.465,1105.54 C371.074,1105.14 371.074,1104.51 371.465,1104.12 L377.121,1098.46 C377.361,1098.22 377.689,1098.15 378,1098.21 C378.311,1098.15 378.639,1098.22 378.879,1098.46 L384.535,1104.12 C384.926,1104.51 384.926,1105.14 384.535,1105.54 L384.535,1105.54 Z M378,1089 C369.163,1089 362,1096.16 362,1105 C362,1113.84 369.163,1121 378,1121 C386.837,1121 394,1113.84 394,1105 C394,1096.16 386.837,1089 378,1089 L378,1089 Z\" id=\"arrow-up-circle\" sketch:type=\"MSShapeGroup\"></path></g></g></g></svg> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var40 := `Ahahay`
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var40)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <svg class=\"sm:block hidden h-3/4\" viewBox=\"0 0 32 32\" version=\"1.1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xmlns:sketch=\"http://www.bohemiancoding.com/sketch/ns\" fill=\"currentColor\" transform=\"matrix(1, 0, 0, -1, 0, 0)\"><g id=\"SVGRepo_bgCarrier\" stroke-width=\"0\"></g><g id=\"SVGRepo_tracerCarrier\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></g><g id=\"SVGRepo_iconCarrier\"><title>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var41 := `arrow-up-circle`
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var41)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><desc>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var42 := `Created with Sketch Beta.`
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var42)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</desc> <defs></defs> <g id=\"Page-1\" stroke=\"none\" stroke-width=\"1\" fill=\"currentColor\" fill-rule=\"evenodd\" sketch:type=\"MSPage\"><g id=\"Icon-Set-Filled\" sketch:type=\"MSLayerGroup\" transform=\"translate(-362.000000, -1089.000000)\" fill=\"currentColor\"><path d=\"M384.535,1105.54 C384.145,1105.93 383.512,1105.93 383.121,1105.54 L379,1101.41 L379,1112 C379,1112.55 378.553,1113 378,1113 C377.447,1113 377,1112.55 377,1112 L377,1101.41 L372.879,1105.54 C372.488,1105.93 371.854,1105.93 371.465,1105.54 C371.074,1105.14 371.074,1104.51 371.465,1104.12 L377.121,1098.46 C377.361,1098.22 377.689,1098.15 378,1098.21 C378.311,1098.15 378.639,1098.22 378.879,1098.46 L384.535,1104.12 C384.926,1104.51 384.926,1105.14 384.535,1105.54 L384.535,1105.54 Z M378,1089 C369.163,1089 362,1096.16 362,1105 C362,1113.84 369.163,1121 378,1121 C386.837,1121 394,1113.84 394,1105 C394,1096.16 386.837,1089 378,1089 L378,1089 Z\" id=\"arrow-up-circle\" sketch:type=\"MSShapeGroup\"></path></g></g></g></svg></button>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func EndOfFeed() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var43 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var43 == nil {
			templ_7745c5c3_Var43 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"divider\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var44 := `You reached the bottom, get a Job!`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var44)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
