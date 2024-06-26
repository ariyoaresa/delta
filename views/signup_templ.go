// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

// method="post" action="/signup"
func HomePage() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.Raw("<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script src=\"https://unpkg.com/htmx.org@1.9.12\" type=\"text/javascript\"></script><script src=\"https://cdn.tailwindcss.com\"></script><title>Chart Delta Sakinah</title></head><body class=\"bg-gray-100 flex flex-col min-h-screen\"><section class=\"flex-grow\"><div class=\"bg-gray-100 py-12\"><div class=\"container mx-auto px-6 md:px-12 lg:px-24\"><div class=\"flex flex-col md:flex-row items-start md:space-x-12\"><div class=\"md:w-1/2 mb-8 md:mb-0\"><h1 class=\"text-4xl md:text-5xl font-bold mb-6 text-gray-900 text-center md:text-left\">Chart Delta Sakinah</h1><div class=\"text-center md:text-left mb-12\"><h2 class=\"text-4xl md:text-5xl font-extrabold text-gray-900\">Tired of hunting for chart repositories?</h2><p class=\"mt-4 text-lg md:text-xl text-gray-600\">Break free from Helm chart release update chaos - your path to effortless management.</p></div></div><div class=\"md:w-1/2 hidden md:flex items-start py-4\"><img src=\"/assets/image.webp\" alt=\"Illustration\" class=\"rounded-lg shadow-lg max-w-full h-auto\"></div></div><div class=\"flex flex-col md:flex-row md:space-x-8 space-y-8 md:space-y-0\"><div class=\"bg-white rounded-lg shadow-lg p-8 md:flex-1\"><h2 class=\"text-2xl font-bold text-gray-900 mb-4\">One Place for All Your Charts</h2><p class=\"text-gray-700 text-base leading-relaxed\">Centralize all your Helm charts in a single location. Stay updated with notifications on new releases. Easily compare versions and visualize changes.</p></div><div class=\"bg-white rounded-lg shadow-lg p-8 md:flex-1\"><h2 class=\"text-2xl font-bold text-gray-900 mb-4\">Simplify Your Helm Workflow</h2><p class=\"text-gray-700 text-base leading-relaxed\">Regain your sanity. Try our solution today and experience the ease of managing your Helm charts in one place.</p></div></div><div class=\"mt-8 w-full\"><div class=\"text-center md:text-left\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = signupWithGithub().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div></div></section><footer class=\"bg-gray-800 text-white py-4 mt-auto\"><div class=\"container mx-auto px-6 md:px-12 text-center\"><p>Built with ♡ by Saintmalik, Styled by Ariyo Aresa</p></div></footer></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func signupWithGithub() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form class=\"flex flex-col gap-3 max-w-[400px] mx-auto\" method=\"post\" action=\"/signup\"><input class=\"items-center justify-center rounded-md btn p-3 px-6 bg-slate-700 text-white hover:bg-slate-800 transition-opacity\" type=\"submit\" value=\"Login &amp; Sign-up with GitHub\"></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

// templ loginWithGithub() {
//     <form id="todo-form" class="flex flex-col gap-3 max-w-[400px] mx-auto" hx-post="/signup" hx-request='{"noHeaders": true}' hx-swap="outerHTML">
//         <input class="btn p-3 px-6 bg-teal-600 text-white hover:opacity-60" type="submit" value="Sign in With GitHub" form="todo-form"/>
//     </form>
// }

// @templ.Raw(`
//         <form class="flex flex-col gap-3 max-w-\[400px\] mx-auto"
//               hx-post="/signup"
//               hx-trigger="submit"
//               hx-swap="outerHTML"
//               hx-headers='{"X-HTTP-Method-Override": "POST"}'>
//             <input class="btn p-3 px-6 bg-teal-600 text-white hover:opacity-60" type="submit" value="Sign up with GitHub"/>
//         </form>`)

// package views

// templ PackageDash(r releaseDataList) {
// 	<!DOCTYPE html>
// 	<html lang="en">
// 		<head>
// 			@templ.Raw("<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">")
// 			<script src="https://unpkg.com/htmx.org@1.9.12" type="text/javascript"></script>
// 			<script src="https://cdn.tailwindcss.com"></script>
// 			<link rel="stylesheet" href="/assets/styles.css"/>
// 			<title>Page title</title>
// 		</head>
//     <body>
//         {{ range . }}
//         <h2>New release found: {{ r.Owner }}/{{ r.Repo }} {{ r.LatestTag }}</h2>
//         <h3>Release Notes:</h3>
//         {{ if r.ReleaseBody }}
//         <p>{{ r.ReleaseBody }}</p>
//         <hr>
//         <h3>Compared to the previous release: {{ r.CurrentTag }}</h3>
//         <p>{{ r.ReleaseNotesDiff }}</p>
//         {{ else }}
//         <p>No release notes available.</p>
//         {{ end }}
//         <hr>
//         {{ end }}
//     </body>
//     </html>
// }
