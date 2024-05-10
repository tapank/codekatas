package ledger

import (
	"errors"
	"fmt"
	"slices"
	"sort"
	"strings"
	"time"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

var currencies = map[string]string{"EUR": "€", "USD": "$"}
var locales = []string{"nl-NL", "en-US"}
var headerFormat = "%-11s| %-26s| %s\n"
var dateFormats = map[string]string{
	"nl-NL": "02-01-2006",
	"en-US": "01/02/2006",
}

func FormatLedger(currency string, locale string, orignialEntries []Entry) (string, error) {
	// validate currency
	if _, ok := currencies[currency]; !ok {
		return "", errors.New("unknown currency")
	}

	// validate locale
	if !slices.Contains(locales, locale) {
		return "", errors.New("unknown locale")
	}

	// create a copy of entries and sort by date and amount
	entries := make([]Entry, len(orignialEntries))
	copy(entries, orignialEntries)
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Date < entries[j].Date || entries[i].Change < entries[j].Change
	})

	// create header
	var s string
	if locale == "nl-NL" {
		s = fmt.Sprintf(headerFormat, "Datum", "Omschrijving", "Verandering")
	} else if locale == "en-US" {
		s = fmt.Sprintf(headerFormat, "Date", "Description", "Change")
	}

	// process records
	ss := []string{s}
	for _, record := range entries {
		if row, err := formatRow(locale, currency, record); err != nil {
			return "", err
		} else {
			ss = append(ss, row)
		}
	}
	return strings.Join(ss, ""), nil
}

func formatRow(locale, currency string, entry Entry) (string, error) {
	// parse date
	var d string
	if t, err := time.Parse("2006-01-02", entry.Date); err != nil {
		return "", err
	} else {
		d = t.Format(dateFormats[locale])
	}

	// truncate extra long descriptions
	de := entry.Description
	if len(de) > 25 {
		de = de[:22] + "..."
	}

	// format amount
	a := formatAmount(entry.Change, locale, currency)

	return fmt.Sprintf("%10s | %-25s | %13s\n", d, de, a), nil
}

func formatAmount(amount int, locale string, currency string) string {
	sign := 1
	if amount < 0 {
		sign = -1
		amount *= sign
	}
	var s = fmt.Sprintf("%02d", amount%100)
	amount /= 100
	for {
		if amount > 999 {
			s = fmt.Sprintf("%03d,%s", amount%1000, s)
		} else {
			s = fmt.Sprintf("%d,%s", amount%1000, s)
		}
		if amount /= 1000; amount == 0 {
			break
		}
	}
	if locale == "nl-NL" {
		s = strings.ReplaceAll(s, ",", ".")
		s = s[:len(s)-3] + "," + s[len(s)-2:]
		if sign < 0 {
			s += "-"
		}
		s = currencies[currency] + " " + s
	} else {
		s = s[:len(s)-3] + "." + s[len(s)-2:]
		s = currencies[currency] + s
		if sign < 0 {
			s = "(" + s + ")"
		}
	}
	if sign > 0 {
		s += " "
	}
	return s
}
