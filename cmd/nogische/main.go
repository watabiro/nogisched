package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/watabiro/nogisched"
)

func main() {
	ymd := flag.String("d", time.Now().Format("20060102"), "the date of schedule: yyyyMMdd")
	notify := flag.Bool("notify", false, "notify LINE")
	flag.Parse()
	t, err := time.Parse("20060102", *ymd)
	if err != nil {
		log.Fatal(fmt.Errorf("cannot parse input date '%s' as yyyyMMdd format", *ymd))
	}
	ctx := context.Background()
	html, err := nogisched.Fetch(ctx, (*ymd)[:6])
	if err != nil {
		log.Fatal(err)
	}
	scheds, err := nogisched.Scrape(html)
	if err != nil {
		log.Fatal(err)
	}
	sched := scheds[t.Day()-1].String()
	if *notify {
		_, err = nogisched.Notify(sched)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Fprint(os.Stdout, sched)
}
