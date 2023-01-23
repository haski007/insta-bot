package transform

import "google.golang.org/api/calendar/v3"

func ToAttendees(emails []string) []*calendar.EventAttendee {
	var attendees []*calendar.EventAttendee

	for _, email := range emails {
		attendees = append(attendees, &calendar.EventAttendee{Email: email})
	}

	return attendees
}
