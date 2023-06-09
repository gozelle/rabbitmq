package message

func NewMessage(value any) *Message {
	return &Message{value: value}
}

type Message struct {
	value       any
	exchange    *string
	routeKey    string
	mandatory   bool
	immediate   bool
	contentType string
	persistent  bool
}

func (m *Message) Exchange() *string {
	return m.exchange
}

func (m *Message) RouteKey() string {
	return m.routeKey
}

func (m *Message) Mandatory() bool {
	return m.mandatory
}

func (m *Message) Immediate() bool {
	return m.immediate
}

func (m *Message) ContentType() string {
	return m.contentType
}

func (m *Message) Persistent() bool {
	return m.persistent
}

func (m *Message) Value() any {
	return m.value
}

func (m *Message) WitRouteKey(routeKey string) *Message {
	m.routeKey = routeKey
	return m
}

func (m *Message) WitMandatory(mandatory bool) *Message {
	m.mandatory = mandatory
	return m
}

func (m *Message) WitImmediate(immediate bool) *Message {
	m.immediate = immediate
	return m
}

func (m *Message) WitContentType(contentType string) *Message {
	m.contentType = contentType
	return m
}

// WitPersistent 开启消息持久化
func (m *Message) WitPersistent(persistent bool) *Message {
	m.persistent = persistent
	return m
}

func (m *Message) WithExchange(exchange string) *Message {
	m.exchange = &exchange
	return m
}
