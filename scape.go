package nogisched

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

// chromedpOptはchromedpの設定をする
func chromedpOpt() []func(*chromedp.ExecAllocator) {
	// chromedp設定
	opts := chromedp.DefaultExecAllocatorOptions[:]
	opts = append(opts,
		chromedp.DisableGPU,
		chromedp.NoSandbox,
		chromedp.Headless,
		chromedp.Flag("lang", "ja-JP"),
		chromedp.Flag("no-zygote", true),
		chromedp.Flag("single-process", true),
		chromedp.Flag("disable-setuid-sandbox", true),
		chromedp.Flag("enable-webgl", true),
		chromedp.Flag("use-gl", "osmesa"),
		chromedp.Flag("homedir", "/tmp"),
		chromedp.Flag("data-path", "/tmp/data-path"),
		chromedp.Flag("disk-cache-dir", "/tmp/cache-dir"),
		chromedp.Flag("remote-debugging-port", "9222"),
		chromedp.Flag("remote-debugging-address", "0.0.0.0"),
		chromedp.Flag("disable-dev-shm-usage", true),
	)
	return opts
}

// Fetchはブラウザを立ち上げ乃木坂のサイトから指定の年月のスケジュールのHTMLを取得する
func Fetch(ctx context.Context, ym string) (string, error) {
	os.Setenv("LANG", "ja_JP.UTF-8")
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	ctx, cancel = chromedp.NewExecAllocator(ctx, chromedpOpt()...)
	defer cancel()
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	q := url.Values{}
	q.Set("ima", "1")
	q.Set("dy", ym)
	url := url.URL{
		Scheme:   "https",
		Host:     "www.nogizaka46.com",
		Path:     "s/n46/media/list",
		RawQuery: q.Encode(),
	}

	var html string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url.String()),
		chromedp.OuterHTML(".sc--lists", &html, chromedp.ByQuery),
	); err != nil {
		return "", err
	}
	return html, nil
}

// Scrapeは与えられた乃木坂の予定のHTMLを使ってスクレイピングを実行
func Scrape(html string) ([]Schedule, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}
	scheds := make([]Schedule, 0, 32)
	doc.Find(".sc--day").
		Each(func(i int, s *goquery.Selection) {
			sched := Schedule{}
			sched.Date = fmt.Sprintf("%s %s",
				// 	// 日付(01/02/.../31)
				s.Find(".sc--day__d").First().Text(),
				// 	// 曜日(Mon/Tue/.../Sat)
				s.Find(".sc--day__w").First().Text())

			sched.Appearances = make([]Appearance, 0)
			s.Find("a").Each(func(i int, s *goquery.Selection) {
				apear := Appearance{
					// カテゴリ(WEB/TV/ラジオ/雑誌....)
					Category: s.Find(".m--scone__cat__name").First().Text(),
					// 放送時間(24:00~24:30)
					Time: s.Find(".m--scone__start").First().Text(),
					// タイトル(テレビ東京系「乃木坂工事中」出演者)
					Title: s.Find(".m--scone__ttl").First().Text(),
				}
				sched.Appearances = append(sched.Appearances, apear)
			})
			scheds = append(scheds, sched)
		})
	return scheds, nil
}
