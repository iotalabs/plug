package pioneer

type Plugger interface {
	Plug(Handler) Handler
}
