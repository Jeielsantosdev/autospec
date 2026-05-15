package autospec

type Option func(*Autospec)

func WithPlugin(p Plugin) Option {
	return func(a *Autospec) { a.plugins = append(a.plugins, p) }
}

func WithAdapterName(name string) Option {
	return func(a *Autospec) { a.preferredAdapter = name }
}
