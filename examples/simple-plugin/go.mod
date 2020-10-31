module github.com/goeven/traepik-simple-plugin-example

go 1.15

require github.com/goeven/traepik v0.0.0-20201031130552-3befc6944c98

// Keep `replace` directives from traefik: https://github.com/traefik/traefik/blob/3e61d1f233ebdf015eaf59e2c253fb8a1f79dbe9/go.mod#L101-L111.

// Docker v19.03.6
replace github.com/docker/docker => github.com/docker/engine v1.4.2-0.20200204220554-5f6d6f3f2203

// Containous forks
replace (
	github.com/abbot/go-http-auth => github.com/containous/go-http-auth v0.4.1-0.20200324110947-a37a7636d23e
	github.com/go-check/check => github.com/containous/check v0.0.0-20170915194414-ca0bf163426a
	github.com/gorilla/mux => github.com/containous/mux v0.0.0-20181024131434-c33f32e26898
	github.com/mailgun/minheap => github.com/containous/minheap v0.0.0-20190809180810-6e71eb837595
	github.com/mailgun/multibuf => github.com/containous/multibuf v0.0.0-20190809014333-8b6c9a7e6bba
)
