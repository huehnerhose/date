// Copyright 2015 Rick Beton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clock

import "fmt"

func clockHours(cm Clock) Clock {
	return (cm / ClockHour)
}

func clockHours12(cm Clock) (Clock, string) {
	h := clockHours(cm)
	if h < 1 {
		return 12, "am"
	} else if h > 12 {
		return h - 12, "pm"
	} else if h == 12 {
		return 12, "pm"
	}
	return h, "am"
}

func clockMinutes(cm Clock) Clock {
	return (cm % ClockHour) / ClockMinute
}

func clockSeconds(cm Clock) Clock {
	return (cm % ClockMinute) / ClockSecond
}

func clockMillisec(cm Clock) Clock {
	return cm % ClockSecond
}

// Hh gets the clock-face number of hours as a two-digit string.
// It is calculated from the modulo time; see Mod24.
func (c Clock) Hh() string {
	cm := c.Mod24()
	return fmt.Sprintf("%02d", clockHours(cm))
}

// HhMm gets the clock-face number of hours and minutes as a five-character ISO-8601 time string.
// It is calculated from the modulo time; see Mod24.
func (c Clock) HhMm() string {
	cm := c.Mod24()
	return fmt.Sprintf("%02d:%02d", clockHours(cm), clockMinutes(cm))
}

// HhMmSs gets the clock-face number of hours, minutes, seconds as an eight-character ISO-8601 time string.
// It is calculated from the modulo time; see Mod24.
func (c Clock) HhMmSs() string {
	cm := c.Mod24()
	return fmt.Sprintf("%02d:%02d:%02d", clockHours(cm), clockMinutes(cm), clockSeconds(cm))
}

// Hh12 gets the clock-face number of hours as a one- or two-digit string, followed by am or pm.
// Remember that midnight is 12am, noon is 12pm.
// It is calculated from the modulo time; see Mod24.
func (c Clock) Hh12() string {
	cm := c.Mod24()
	h, sfx := clockHours12(cm)
	return fmt.Sprintf("%d%s", h, sfx)
}

// HhMm12 gets the clock-face number of hours and minutes, followed by am or pm.
// Remember that midnight is 12am, noon is 12pm.
// It is calculated from the modulo time; see Mod24.
func (c Clock) HhMm12() string {
	cm := c.Mod24()
	h, sfx := clockHours12(cm)
	return fmt.Sprintf("%d:%02d%s", h, clockMinutes(cm), sfx)
}

// HhMm12 gets the clock-face number of hours, minutes and seconds, followed by am or pm.
// Remember that midnight is 12am, noon is 12pm.
// It is calculated from the modulo time; see Mod24.
func (c Clock) HhMmSs12() string {
	cm := c.Mod24()
	h, sfx := clockHours12(cm)
	return fmt.Sprintf("%d:%02d:%02d%s", h, clockMinutes(cm), clockSeconds(cm), sfx)
}

// String gets the clock-face number of hours, minutes, seconds and milliseconds as a 12-character ISO-8601
// time string (calculated from the modulo time, see Mod24), specified to the nearest millisecond.
func (c Clock) String() string {
	cm := c.Mod24()
	return fmt.Sprintf("%02d:%02d:%02d.%03d", clockHours(cm), clockMinutes(cm), clockSeconds(cm), clockMillisec(cm))
}
