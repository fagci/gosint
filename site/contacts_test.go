package site

import (
	"fmt"
	"testing"
)

func TestContactsByType(t *testing.T) {
	text := `+7 (999) 999-99-99
    Email: test@site.org, Telegram: t.me/fagci
    <a href="tel:+12215555555">+1(221) 555-55-55</a>
    <a href="mailto:mail@example.com">Email link example</a>
    <a href="https://telegram.me/tarmanau">Telegram the author</a>
    <a href="https://wa.me/2217777777">Chat with a WhatsApp user</a>
    <a href="https://wa.me/2217777777?text=Hello!">Chat+message on WhatsApp</a>
    <a href="skype:LOGIN?userinfo">Skype</a>
    <a href="skype:LOGIN?call">Call on the skype user</a>
    <a href="https://www.messenger.com/t/jack.malbon.3">Facebook Messenger</a>`
	for t, c := range ContactsByType(text) {
		fmt.Println(t+":", c)
	}
}
