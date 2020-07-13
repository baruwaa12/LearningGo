package twelve

func twelvedayschristmaslyrics(versenumber int) (verse string) {

	if versenumber >= 1 {
		var days = []string{"", "first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
		var partnumber = []string{""}

		verse = "On the " + days[versenumber] + " day of Christmas my true love gave to me" + partnumber[versenumber]
	}
	return verse
}
