package gui

import (
	"fmt"
	"strings"

	"github.com/yetone/magpie/internal/gateway"
)

// trayZh is the tray's text in Chinese; the page has its own words
// (i18n.js), and these are all the menu, the usage cards' tooltip and the
// idle tip show (#301).
var trayZh = map[string]string{
	"Open magpie":                          "打开 magpie",
	"Version %s":                           "版本 %s",
	"Restart to Update":                    "重启以更新",
	"Restart to Update to %s":              "重启以更新到 %s",
	"Quit magpie":                          "退出 magpie",
	"Restart Now to Update":                "立即重启以更新",
	"Restart Now to Update (%d in flight)": "立即重启以更新（%d 个进行中）",
	"%s used":                              "已用 %s",
	"%s left":                              "剩余 %s",
	"pace %s":                              "配速 %s",
	"resets in %s":                         "%s后重置",
	"as of %s, couldn't be read just now":  "%s 的数据，暂时无法读取",
	"Today in magpie":                      "magpie · 今天",
	"Last 7 days in magpie":                "magpie · 最近 7 天",
	"%s tokens (%s in · %s out":            "%s Token（%s 输入 · %s 输出",
	" · %s cached":                         " · %s 缓存",
	"%d requests":                          "%d 个请求",
	"%d request":                           "%d 个请求",
	", %d with no known price":             "，其中 %d 个没有已知价格",
	"Today: no requests yet":               "今天还没有请求",
	"Today: %s":                            "今天：%s",
	"5 hours":                              "5 小时",
	"7 days":                               "7 天",
	"1 day":                                "1 天",
	"Weekly":                               "每周",
	"Month":                                "本月",
	"Session":                              "整个会话",
	"Credits":                              "积分",
	"Allowance":                            "额度",
	"Limit":                                "限额",
	"Extra usage":                          "额外用量",
}

// trayZhTW is the tray's text in Traditional Chinese, with Taiwan's words.
var trayZhTW = map[string]string{
	"Open magpie":                          "開啟 magpie",
	"Version %s":                           "版本 %s",
	"Restart to Update":                    "重新啟動以更新",
	"Restart to Update to %s":              "重新啟動以更新到 %s",
	"Quit magpie":                          "結束 magpie",
	"Restart Now to Update":                "立即重新啟動以更新",
	"Restart Now to Update (%d in flight)": "立即重新啟動以更新（%d 個進行中）",
	"%s used":                              "已用 %s",
	"%s left":                              "剩餘 %s",
	"pace %s":                              "配速 %s",
	"resets in %s":                         "%s後重置",
	"as of %s, couldn't be read just now":  "%s 的資料，暫時無法讀取",
	"Today in magpie":                      "magpie · 今天",
	"Last 7 days in magpie":                "magpie · 最近 7 天",
	"%s tokens (%s in · %s out":            "%s Token（%s 輸入 · %s 輸出",
	" · %s cached":                         " · %s 快取",
	"%d requests":                          "%d 個請求",
	"%d request":                           "%d 個請求",
	", %d with no known price":             "，其中 %d 個沒有已知價格",
	"Today: no requests yet":               "今天尚無請求",
	"Today: %s":                            "今天：%s",
	"5 hours":                              "5 小時",
	"7 days":                               "7 天",
	"1 day":                                "1 天",
	"Weekly":                               "每週",
	"Month":                                "本月",
	"Session":                              "整個會話",
	"Credits":                              "積分",
	"Allowance":                            "額度",
	"Limit":                                "限額",
	"Extra usage":                          "額外用量",
}

// trayJa is the tray's text in Japanese.
var trayJa = map[string]string{
	"Open magpie":                          "magpie を開く",
	"Version %s":                           "バージョン %s",
	"Restart to Update":                    "再起動してアップデート",
	"Restart to Update to %s":              "再起動して %s にアップデート",
	"Quit magpie":                          "magpie を終了",
	"Restart Now to Update":                "今すぐ再起動してアップデート",
	"Restart Now to Update (%d in flight)": "今すぐ再起動してアップデート（%d 件処理中）",
	"%s used":                              "使用済み %s",
	"%s left":                              "残り %s",
	"pace %s":                              "ペース %s",
	"resets in %s":                         "%s後にリセット",
	"as of %s, couldn't be read just now":  "%s 時点の値（現在は読み取れません）",
	"Today in magpie":                      "magpie · 今日",
	"Last 7 days in magpie":                "magpie · 過去 7 日間",
	"%s tokens (%s in · %s out":            "%s トークン（%s 入力 · %s 出力",
	" · %s cached":                         " · %s キャッシュ",
	"%d requests":                          "%d 件のリクエスト",
	"%d request":                           "%d 件のリクエスト",
	", %d with no known price":             "、%d 件は価格不明",
	"Today: no requests yet":               "今日はまだリクエストがありません",
	"Today: %s":                            "今日：%s",
	"5 hours":                              "5 時間",
	"7 days":                               "7 日",
	"1 day":                                "1 日",
	"Weekly":                               "週",
	"Month":                                "今月",
	"Session":                              "セッション",
	"Credits":                              "クレジット",
	"Allowance":                            "利用枠",
	"Limit":                                "上限",
	"Extra usage":                          "追加使用量",
}

// trayDe is the tray's text in German.
var trayDe = map[string]string{
	"Open magpie":                          "magpie öffnen",
	"Version %s":                           "Version %s",
	"Restart to Update":                    "Zum Aktualisieren neu starten",
	"Restart to Update to %s":              "Zum Aktualisieren auf %s neu starten",
	"Quit magpie":                          "magpie beenden",
	"Restart Now to Update":                "Jetzt neu starten und aktualisieren",
	"Restart Now to Update (%d in flight)": "Jetzt neu starten und aktualisieren (%d laufend)",
	"%s used":                              "%s verbraucht",
	"%s left":                              "%s übrig",
	"pace %s":                              "Tempo %s",
	"resets in %s":                         "Reset in %s",
	"as of %s, couldn't be read just now":  "Stand %s, gerade nicht lesbar",
	"Today in magpie":                      "Heute in magpie",
	"Last 7 days in magpie":                "Letzte 7 Tage in magpie",
	"%s tokens (%s in · %s out":            "%s Token (%s ein · %s aus",
	" · %s cached":                         " · %s im Cache",
	"%d requests":                          "%d Anfragen",
	"%d request":                           "%d Anfrage",
	", %d with no known price":             ", %d ohne bekannten Preis",
	"Today: no requests yet":               "Heute noch keine Anfragen",
	"Today: %s":                            "Heute: %s",
	"5 hours":                              "5 Stunden",
	"7 days":                               "7 Tage",
	"1 day":                                "1 Tag",
	"Weekly":                               "Wöchentlich",
	"Month":                                "Monat",
	"Session":                              "Sitzung",
	"Credits":                              "Credits",
	"Allowance":                            "Kontingent",
	"Limit":                                "Limit",
	"Extra usage":                          "Zusätzliche Nutzung",
}

// trayWords are the tray's translations by language.
var trayWords = map[string]map[string]string{"zh": trayZh, "zh-TW": trayZhTW, "ja": trayJa, "de": trayDe}

// onLang relabels the tray menu when the Settings page changes the
// language; set by the process that has the tray. The usage cards are
// redrawn through onTrayUsage, whose words have their own refresh loop.
var onLang func()

// trayLang is the language the tray's text is in, as the page picks its own:
// the setting, or with "system" (or none) the system's: Traditional Chinese
// for Taiwan, Hong Kong, Macau and any zh-Hant (zh_TW.UTF-8, zh-Hant-TW),
// Simplified for any other zh, Japanese for any ja, German for any de.
func trayLang(pref string, system func() string) string {
	switch pref {
	case "en", "zh", "zh-TW", "ja", "de":
		return pref
	}
	switch sys := strings.ReplaceAll(strings.ToLower(system()), "_", "-"); {
	case zhTraditional(sys):
		return "zh-TW"
	case strings.HasPrefix(sys, "zh"):
		return "zh"
	case strings.HasPrefix(sys, "ja"):
		return "ja"
	case strings.HasPrefix(sys, "de"):
		return "de"
	}
	return "en"
}

// zhTraditional says a lowercased, dashed system tag is Traditional
// Chinese, as the page's sysLocale does.
func zhTraditional(sys string) bool {
	for _, p := range []string{"zh-tw", "zh-hk", "zh-mo", "zh-hant"} {
		if sys == p || strings.HasPrefix(sys, p+"-") || strings.HasPrefix(sys, p+".") || strings.HasPrefix(sys, p+"@") {
			return true
		}
	}
	return false
}

// notesLang is the language release notes are asked in for lang: there are
// no German ones, so German asks for the English, and Traditional Chinese
// asks for the Chinese ones.
func notesLang(lang string) string {
	switch lang {
	case "de":
		return "en"
	case "zh-TW":
		return "zh"
	}
	return lang
}

// trayText is a tray line in the language, English where it has none.
func trayText(lang, key string, args ...any) string {
	if s, ok := trayWords[lang][key]; ok {
		key = s
	}
	if len(args) == 0 {
		return key
	}
	return fmt.Sprintf(key, args...)
}

// trayLabels are the tray menu's lines: open, version, restart and quit.
type trayLabels struct{ open, version, restart, quit string }

// trayMenuLabels are the lines in the language; update is the version
// waiting for a restart, if there is one.
func trayMenuLabels(lang, version, update string) trayLabels {
	l := trayLabels{
		open:    trayText(lang, "Open magpie"),
		version: trayText(lang, "Version %s", version),
		restart: trayText(lang, "Restart to Update"),
		quit:    trayText(lang, "Quit magpie"),
	}
	if update != "" {
		l.restart = trayText(lang, "Restart to Update to %s", update)
	}
	return l
}

// trayRestartNow is the restart item while a restart waits for the
// gateway (#577): a click restarts at once, cutting short what b has in
// flight.
func trayRestartNow(lang string, b gateway.Busy) string {
	if n := b.Requests + b.Tools; n > 0 {
		return trayText(lang, "Restart Now to Update (%d in flight)", n)
	}
	return trayText(lang, "Restart Now to Update")
}

// envLang is the language the environment names: the first of LC_ALL,
// LC_MESSAGES, LANG and LANGUAGE set to one.
func envLang(getenv func(string) string) string {
	for _, k := range []string{"LC_ALL", "LC_MESSAGES", "LANG", "LANGUAGE"} {
		if v := getenv(k); v != "" && v != "C" && v != "POSIX" && !strings.HasPrefix(v, "C.") {
			return v
		}
	}
	return ""
}

// firstAppleLang is the first entry `defaults read -g AppleLanguages`
// prints: ( "zh-Hans-CN", "en-US" ), one to a line, quoted or not.
func firstAppleLang(out string) string {
	for _, line := range strings.Split(out, "\n") {
		s := strings.Trim(strings.TrimSpace(line), `",`)
		if s != "" && s != "(" && s != ")" {
			return s
		}
	}
	return ""
}
