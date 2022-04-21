package site

import "regexp"

var contactPatterns = map[string]string{
	"Facebook":  `facebook\.com[-.A-Za-z0-9/]+`,
	"Facebook2": `fb\.me[-.A-Za-z0-9/]+`,
	"Github":    `github\.com/[^"'/]+`,
	"Instagram": `instagram\.com/[^"'/]+`,
	"Linkedin":  `linkedin.com[-._A-Za-z0-9/]+`,
	"Mails":     `[\w.-]+@[\w.-]+\.\w{2,5}`,
	"OK":        `ok\.ru/[^"'/]+`,
	"Phones":    `\+\d{0,3}\s?0?\d{7,10}`,
	"Phones2":   `\+?\d{0,3}?\s?0?\d{3}\s\d{3}\s\d{3}`,
	"Phones3":   `\+?\(?\d{0,3}\)?\s?0?\d{3}\s\d{4}`,
	"Phones4":   `\+?\d{1,2}\s?\(?\d{0,3}\)?\s?\d{3}[\s-]\d{2}[\s-]\d{2}`,
	"Telegram":  `t\.me/[-._A-Za-z0-9/]+`,
	"Twitter":   `twitter\.com[-._A-Za-z0-9/]+`,
	"VK":        `vk\.com/[^"'/]+`,
	"Whatsapp":  `api\.whatsapp\.com/send\?phone=([\d]+)`,
	"Whatsapp2": `web\.whatsapp\.com/send\?phone=([\d]+)`,
	"Whatsapp3": `wa\.me/([\d]+)`,
	"YouTube":   `youtube\.\w+?/channel/[^"']+`,
	"tel":       `tel:(\+?[^'"<>]+)`,
	"mailto":    `mailto:(\+?[^'"<>]+)`,
}

var contactRegexps = make(map[string]*regexp.Regexp)

func init() {
	for k, v := range contactPatterns {
		contactRegexps[k] = regexp.MustCompile(v)
	}
}

func ContactsByType(text string) (contacts map[string][]string) {
	contacts = make(map[string][]string)
	for k, r := range contactRegexps {
		matches := r.FindAllString(text, -1)
		if len(matches) != 0 {
			contacts[k] = append(contacts[k], matches...)
		}
	}
	return
}
