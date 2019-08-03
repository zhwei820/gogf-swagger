package gogfSwagger

import (
	"github.com/gogf/gf/g/net/ghttp"
	"html/template"
	"os"
	"regexp"
	"strings"

	"github.com/swaggo/swag"
	"golang.org/x/net/webdav"
)

// Config stores ginSwagger configuration variables.
type Config struct {
	//The url pointing to API definition (normally swagger.json or swagger.yaml). Default is `doc.json`.
	URL string
}

// URL presents the url pointing to API definition (normally swagger.json or swagger.yaml).
func URL(url string) func(c *Config) {
	return func(c *Config) {
		c.URL = url
	}
}

// WrapHandler wraps `http.Handler` into `func(r *ghttp.Request)`.
func WrapHandler(h *webdav.Handler, confs ...func(c *Config)) func(r *ghttp.Request) {
	defaultConfig := &Config{
		URL: "doc.json",
	}

	for _, c := range confs {
		c(defaultConfig)
	}

	return CustomWrapHandler(defaultConfig, h)
}

// CustomWrapHandler wraps `http.Handler` into `func(r *ghttp.Request)`
func CustomWrapHandler(config *Config, h *webdav.Handler) func(r *ghttp.Request) {
	//create a template with name
	t := template.New("swagger_index.html")
	index, _ := t.Parse(swagger_index_templ)

	var rexp = regexp.MustCompile(`/([a-z]*)/(.*)[\?|.]*`)

	return func(r *ghttp.Request) {

		type swaggerUIBundle struct {
			URL string
		}

		var matches []string
		if matches = rexp.FindStringSubmatch(r.Request.RequestURI); len(matches) != 3 {
			r.Response.WriteStatus(404, "404 page not found")
			return
		}
		path := matches[2]
		prefix := matches[1]
		h.Prefix = "/" + prefix

		if strings.HasSuffix(path, ".html") {
			r.Response.Header().Set("Content-Type", "text/html; charset=utf-8")
		} else if strings.HasSuffix(path, ".css") {
			r.Response.Header().Set("Content-Type", "text/css; charset=utf-8")
		} else if strings.HasSuffix(path, ".js") {
			r.Response.Header().Set("Content-Type", "application/javascript")
		} else if strings.HasSuffix(path, ".json") {
			r.Response.Header().Set("Content-Type", "application/json")
		}

		switch path {
		case "index.html":
			index.Execute(r.Response.Writer, &swaggerUIBundle{
				URL: config.URL,
			})
		case "doc.json":
			doc, err := swag.ReadDoc()
			if err != nil {
				panic(err)
			}
			r.Response.Writer.Write([]byte(doc))
			return
		default:
			h.ServeHTTP(r.Response.Writer, r.Request)
		}
	}
}

// DisablingWrapHandler turn handler off
// if specified environment variable passed
func DisablingWrapHandler(h *webdav.Handler, envName string) func(r *ghttp.Request) {
	eFlag := os.Getenv(envName)
	if eFlag != "" {
		return func(r *ghttp.Request) {
			// Simulate behavior when route unspecified and
			// return 404 HTTP code
			r.Response.WriteStatus(404, "")
		}
	}

	return WrapHandler(h)
}

// DisablingCustomWrapHandler turn handler off
// if specified environment variable passed
func DisablingCustomWrapHandler(config *Config, h *webdav.Handler, envName string) func(r *ghttp.Request) {
	eFlag := os.Getenv(envName)
	if eFlag != "" {
		return func(r *ghttp.Request) {
			// Simulate behavior when route unspecified and
			// return 404 HTTP code
			r.Response.WriteStatus(404, "")
		}
	}

	return CustomWrapHandler(config, h)
}

const swagger_index_templ = `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta http-equiv="x-ua-compatible" content="IE=edge">
    <title>Swagger UI</title>
    <link rel="icon" type="image/png" href="images/favicon-32x32.png" sizes="32x32"/>
    <link rel="icon" type="image/png" href="images/favicon-16x16.png" sizes="16x16"/>
    <link href='css/typography.css' media='screen' rel='stylesheet' type='text/css'/>
    <link href='css/reset.css' media='screen' rel='stylesheet' type='text/css'/>
    <link href='css/screen.css' media='screen' rel='stylesheet' type='text/css'/>
    <link href='css/reset.css' media='print' rel='stylesheet' type='text/css'/>
    <link href='css/print.css' media='print' rel='stylesheet' type='text/css'/>

    <script src='lib/object-assign-pollyfill.js' type='text/javascript'></script>
    <script src='lib/jquery-1.8.0.min.js' type='text/javascript'></script>
    <script src='lib/jquery.slideto.min.js' type='text/javascript'></script>
    <script src='lib/jquery.wiggle.min.js' type='text/javascript'></script>
    <script src='lib/jquery.ba-bbq.min.js' type='text/javascript'></script>
    <script src='lib/handlebars-2.0.0.js' type='text/javascript'></script>
    <script src='lib/lodash.min.js' type='text/javascript'></script>
    <script src='lib/backbone-min.js' type='text/javascript'></script>
    <script src='swagger-ui.min.js' type='text/javascript'></script>
    <script src='lib/highlight.9.1.0.pack.js' type='text/javascript'></script>
    <script src='lib/highlight.9.1.0.pack_extended.js' type='text/javascript'></script>
    <script src='lib/jsoneditor.min.js' type='text/javascript'></script>
    <script src='lib/marked.js' type='text/javascript'></script>
    <script src='lib/swagger-oauth.js' type='text/javascript'></script>

    <!-- Some basic translations -->
    <!-- <script src='lang/translator.js' type='text/javascript'></script> -->
    <!-- <script src='lang/ru.js' type='text/javascript'></script> -->
    <!-- <script src='lang/en.js' type='text/javascript'></script> -->

    <script type="text/javascript">
        $(function () {
            var url = window.location.search.match(/url=([^&]+)/);

            url = "{{.URL}}",

            hljs.configure({
                highlightSizeThreshold: 5000
            });

            // Pre load translate...
            if (window.SwaggerTranslator) {
                window.SwaggerTranslator.translate();
            }
            window.swaggerUi = new SwaggerUi({
                url: url,
                dom_id: "swagger-ui-container",
                supportedSubmitMethods: ['get', 'post', 'put', 'delete', 'patch'],
                onComplete: function (swaggerApi, swaggerUi) {
                    if (typeof initOAuth == "function") {

                        initOAuth({
                            clientId: "your-client-id",
                            clientSecret: "your-client-secret-if-required",
                            realm: "your-realms",
                            appName: "your-app-name",
                            scopeSeparator: " ",
                            additionalQueryStringParams: {}
                        });
                    }

                    if (window.SwaggerTranslator) {
                        window.SwaggerTranslator.translate();
                    }
                },
                onFailure: function (data) {
                    log("Unable to Load SwaggerUI");
                },
                docExpansion: "none",
                jsonEditor: false,
                defaultModelRendering: 'schema',
                showRequestHeaders: false,
                showOperationIds: false
            });

            window.swaggerUi.load();

            function log() {
                if ('console' in window) {
                    console.log.apply(console, arguments);
                }
            }
        });
    </script>
</head>

<body class="swagger-section">
<div id='header'>
    <div class="swagger-ui-wrap">
        <a id="logo" href="http://swagger.io"><img class="logo__img" alt="swagger" height="30" width="30"
                                                   src="images/logo_small.png"/><span class="logo__title">swagger</span></a>
        <form id='api_selector'>
            <div class='input'><input placeholder="http://example.com/api" id="input_baseUrl" name="baseUrl"
                                      type="text"/></div>
            <div id='auth_container'></div>
            <div class='input'><a id="explore" class="header__btn" href="#" data-sw-translate>Explore</a></div>

        </form>
    </div>
</div>

<div id="message-bar" class="swagger-ui-wrap" data-sw-translate>&nbsp;</div>
<div id="swagger-ui-container" class="swagger-ui-wrap"></div>
</body>
<script src='jwt.js' type='text/javascript'></script>

</html>

`
