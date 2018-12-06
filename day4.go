package main

import (
	"fmt"
	utils "github.com/Mabrollin/AoC/utils"
	"strconv"
	"strings"
	time "time"
)

const DATE_FORMAT = "2006-01-02 15:04"

type shift struct {
	night        string
	sleepMinutes map[int]bool
}

func (shift shift) getTotalSleepMinutes() int {
	total := 0
	for i := 0; i < 60; i++ {
		if val, ok := shift.sleepMinutes[i]; ok && val {
			total++
		}
	}
	return total
}

func (shift shift) toString() string {
	str := " " + shift.night + " "
	for i := 0; i < 60; i++ {
		if val, ok := shift.sleepMinutes[i]; ok {
			if val {
				str += "#"
			} else {
				str += "-"
			}

		}
	}
	return str
}

func (shift shift) fillMissingMinutes() {
	sleep := false
	for i := 0; i < 60; i++ {
		if val, ok := shift.sleepMinutes[i]; ok {
			sleep = val
		} else {
			shift.sleepMinutes[i] = sleep
		}
	}
}

func main() {
	rawShiftInputs := utils.ReadArgumentFile()
	shiftsByGuards, _ := mapRawShiftsToStructMapByGuardId(rawShiftInputs)
	for i, shiftsByGuard := range shiftsByGuards {
		sleptMinutes := 0
		for _, shift := range shiftsByGuard {
			sleptMinutes += shift.getTotalSleepMinutes()
		}
		a, b := getMostSleepyMinute(shiftsByGuard)
		fmt.Printf("\nGuard: %s sleept %d minutes for %d shifts, mostly under minute %d for %d times", i, sleptMinutes, len(shiftsByGuard), a, b)
	}
}

func mapRawShiftsToStructMapByGuardId(rawShifts []string) (map[string][]shift, error) {
	nightsByGuardId := make(map[string][]string)
	shiftsByNight := make(map[string]shift)
	for _, rawShift := range rawShifts {
		rawDateTime := rawShift[1:17]
		dateTime, err := time.Parse(DATE_FORMAT, rawDateTime)
		if err != nil {
			fmt.Println("can't parse datetime: %s", err)
		}
		night := toYearNightKey(dateTime)
		rawEvent := rawShift[19:]
		if rawEvent == "falls asleep" {
			if _, ok := shiftsByNight[night]; !ok {
				shiftsByNight[night] = shift{
					night:        night,
					sleepMinutes: make(map[int]bool),
				}
			}
			shiftsByNight[night].sleepMinutes[dateTime.Minute()] = true
		}

		if rawEvent == "wakes up" {
			if _, ok := shiftsByNight[night]; !ok {
				shiftsByNight[night] = shift{
					night:        night,
					sleepMinutes: make(map[int]bool),
				}
			}
			shiftsByNight[night].sleepMinutes[dateTime.Minute()] = false
		}

		if strings.HasPrefix(rawEvent, "Guard") && strings.HasSuffix(rawEvent, "begins shift") {
			guardId := strings.TrimPrefix(strings.TrimSuffix(rawEvent, " begins shift"), "Guard ")
			if _, ok := nightsByGuardId[guardId]; !ok {
				nightsByGuardId[guardId] = []string{}
			}
			nightsByGuardId[guardId] = append(nightsByGuardId[guardId], night)

		}
	}
	shiftsByGuardId := make(map[string][]shift)

	for guardId := range nightsByGuardId {
		shiftsByGuardId[guardId] = make([]shift, len(nightsByGuardId[guardId]))
		for i, night := range nightsByGuardId[guardId] {

			if val, ok := shiftsByNight[night]; ok {
				shiftsByGuardId[guardId][i] = val
			} else {
				shiftsByGuardId[guardId][i] = shift{
					night:        night,
					sleepMinutes: make(map[int]bool),
				}
			}
			shiftsByGuardId[guardId][i].fillMissingMinutes()
		}
	}

	return shiftsByGuardId, nil
}

func toYearNightKey(dateTime time.Time) string {
	nightOffset := 0
	if dateTime.Hour() != 0 {
		nightOffset = 1
	}
	return strconv.Itoa(dateTime.Year()) + strconv.Itoa(dateTime.YearDay()+nightOffset)
}

func getMostSleepyMinute(shifts []shift) (int, int) {
	mostSleepyMinute := 0
	mostSleepyMinuteCount := 0
	for i := 0; i < 60; i++ {
		timesSleptThisMinute := 0
		for _, shift := range shifts {
			if val, ok := shift.sleepMinutes[i]; ok && val {
				timesSleptThisMinute++
			}
		}
		if timesSleptThisMinute > mostSleepyMinuteCount {
			mostSleepyMinute = i
			mostSleepyMinuteCount = timesSleptThisMinute
		}
	}
	return mostSleepyMinute, mostSleepyMinuteCount
}
