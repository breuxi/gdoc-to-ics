package converter

import (
	"encoding/csv"
	"errors"
	"fmt"
	"strings"
	"time"

	ics "github.com/arran4/golang-ical"

	"github.com/goodsign/monday"
	"github.com/google/uuid"
)

const (
	dateFormatUtc = "20060102"

	propertyDtStart ics.Property = "DTSTART;VALUE=DATE"
	propertyDtEnd   ics.Property = "DTEND;VALUE=DATE"

	componentPropertyDtStart = ics.ComponentProperty(propertyDtStart)
	componentPropertyDtEnd   = ics.ComponentProperty(propertyDtEnd)
)

func ConvertGDocCSVToIcs(csv_content string, filename string) (string, error) {
	csv_reader := csv.NewReader(strings.NewReader(csv_content))

	records, err := csv_reader.ReadAll()

	if err != nil {
		return "", err
	}

	ics := ics.NewCalendarFor("gdoc-to-ics")
	ics.SetName(filename)
	ics.SetXWRCalName(filename)

	if len(records) >= 2 {
		for rec_index := range records {
			record := records[rec_index]
			if !isEmpty(record) {
				if len(record) > 0 && strings.HasPrefix(strings.ToLower(record[0]), "januar") {
					// Januar  Februar  März  April  Mai  Juni  Juli  August  September  Oktober  November  Dezember
					for month_index := range record {
						month_name := record[month_index]

						if len(month_name) > 0 && month_index+1 < len(record) {
							for day_rec_index := range records {
								day_rec := records[day_rec_index][month_index]
								names_rec := records[day_rec_index][month_index+1]

								names := strings.Split(names_rec, ",")

								for _, name := range names {
									if len(name) > 0 {
										event := ics.AddEvent(filename + "_gdoc-to-ics@" + uuid.NewString())
										event.SetSummary(name)
										event.SetCreatedTime(time.Now())
										event.SetDtStampTime(time.Now())
										event.SetModifiedAt(time.Now())

										loc, _ := time.LoadLocation("Europe/Berlin")

										parsed_date, err := monday.ParseInLocation("2. January", fmt.Sprintf("%s. %s", day_rec, month_name), loc, monday.LocaleDeDE)
										if err == nil {
											date := time.Date(time.Now().Year(), parsed_date.Month(), parsed_date.Day(), 0, 0, 0, 0, parsed_date.Location())

											event.SetProperty(componentPropertyDtStart, date.Format(dateFormatUtc))
										}

										event.AddRrule("FREQ=YEARLY")

										event.SetURL("https://twitter.com/" + name)
									}
								}
							}
						}
					}
					break
				}
			}
		}

		return ics.Serialize(), nil
	} else {
		return "", errors.New("Error! Not enough records!")
	}
}

func isEmpty(record []string) bool {
	for i := range record {
		if record[i] != "" {
			return false
		}
	}
	return true
}
