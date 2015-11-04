package gcm

// The field meaning explained at [GCM Architectural Overview](http://developer.android.com/guide/google/gcm/gcm.html#send-msg)
type Message struct {
	To             string                 `json:"to"`
	CollapseKey    string                 `json:"collapse_key,omitempty"`
	DelayWhileIdle bool                   `json:"delay_while_idle,omitempty"`
	Notification   map[string]interface{} `json:"notification,omitempty"`
	Data           map[string]interface{} `json:"data,omitempty"`
	TimeToLive     int                    `json:"time_to_live,omitempty"`
}

func NewMessage(ids string) *Message {
	return &Message{
		To:   ids,
		Data: make(map[string]interface{}),
	}
}

func (m *Message) SetNotification(key string, value interface{}) {
	if m.Notification == nil {
		m.Notification = make(map[string]interface{})
	}
	m.Notification[key] = value
}

func (m *Message) SetData(key string, value interface{}) {
	if m.Data == nil {
		m.Data = make(map[string]interface{})
	}
	m.Data[key] = value
}

// The field meaning explained at [GCM Architectural Overview](http://developer.android.com/guide/google/gcm/gcm.html#send-msg)
type Response struct {
	MulticastID  int64 `json:"multicast_id"`
	Success      int   `json:"success"`
	Failure      int   `json:"failure"`
	CanonicalIDs int   `json:"canonical_ids"`
	Results      []struct {
		MessageID      string `json:"message_id"`
		RegistrationID string `json:"registration_id"`
		Error          string `json:"error"`
	} `json:"results"`
}

// Return the indexes of succeed sent registration ids
func (r *Response) SuccessIndexes() []int {
	ret := make([]int, 0, r.Success)
	for i, result := range r.Results {
		if result.Error == "" {
			ret = append(ret, i)
		}
	}
	return ret
}

// Return the indexes of failed sent registration ids
func (r *Response) ErrorIndexes() []int {
	ret := make([]int, 0, r.Failure)
	for i, result := range r.Results {
		if result.Error != "" {
			ret = append(ret, i)
		}
	}
	return ret
}

// Return the indexes of registration ids which need update
func (r *Response) RefreshIndexes() []int {
	ret := make([]int, 0, r.CanonicalIDs)
	for i, result := range r.Results {
		if result.RegistrationID != "" {
			ret = append(ret, i)
		}
	}
	return ret
}
