// Package kickscopes provides helpers and types for managing OAuth scopes
// in the Kick API. It supports marshaling/unmarshaling to JSON and
// joining scopes into strings for query parameters.
package kickscopes

import (
	"encoding/json"
	"strings"
)

type Scope string

const (
	UserRead                    Scope = "user:read"
	ChannelRead                 Scope = "channel:read"
	ChannelWrite                Scope = "channel:write"
	ChatWrite                   Scope = "chat:write"
	StreamKeyRead               Scope = "streamkey:read"
	EventsSubscribe             Scope = "events:subscribe"
	ModerationBan               Scope = "moderation:ban"
	ModerationChatMessageManage Scope = "moderation:chat_message:manage"
	KicksRead                   Scope = "kicks:read"
	ChannelRewardsRead          Scope = "channel:rewards:read"
	ChannelRewardsWrite         Scope = "channel:rewards:write"
)

type Scopes []Scope

func (s Scopes) Join(separator string) string {
	if len(s) == 0 {
		return ""
	}

	sepLen := len(separator)
	size := 0
	for _, scope := range s {
		size += len(scope)
	}
	size += sepLen * (len(s) - 1)

	var b strings.Builder
	b.Grow(size)

	first := true
	for _, scope := range s {
		if !first {
			b.WriteString(separator)
		} else {
			first = false
		}
		b.Write([]byte(scope))
	}

	return b.String()
}

func (s Scopes) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Join(" "))
}

func (s *Scopes) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	if str == "" {
		*s = nil
		return nil
	}

	parts := strings.Fields(str)
	scopes := make(Scopes, len(parts))
	for i, p := range parts {
		scopes[i] = Scope(p)
	}
	*s = scopes
	return nil
}
