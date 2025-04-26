package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gtfo_logger/zone_areas"
)

func main() {
	// NOTE: Search for GTFO log folder
	user_profile := os.Getenv("USERPROFILE")
	if user_profile == "" {
		fmt.Println("USERPROFILE not found")
		return
	}

	folder_path := filepath.Join(user_profile, "AppData", "LocalLow", "10 Chambers Collective/GTFO")

	// NOTE: find most recently modified file
	var newest_file string
	var newest_mod_time time.Time

	err := filepath.WalkDir(folder_path, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return nil
		}

		if info.ModTime().After(newest_mod_time) {
			newest_mod_time = info.ModTime()
			newest_file = path
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error scanning directory:", err)
		return
	}

	if newest_file == "" {
		fmt.Println("No files found")
		return
	}

	fmt.Println("Most recent file:", newest_file)

	// NOTE: scan file for important items
	for {
		// Open the file each time
		file, err := os.Open(newest_file)
		if err != nil {
			panic(err)
		}

		// Read all lines into a fresh slice
		var lines []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		file.Close() // Important: close manually here (can't defer inside loop)

		if err := scanner.Err(); err != nil {
			panic(err)
		}

		// Find matching lines
		var matchIndexes []int
		for i, line := range lines {
			if strings.Contains(line, "Next Batch: Distribution -") {
				matchIndexes = append(matchIndexes, i)
			}
		}

		if len(matchIndexes) == 0 {
			fmt.Println("No matches found")
			time.Sleep(time.Second)
			continue
		}

		// latest run log
		startIndex := matchIndexes[len(matchIndexes)-1]

		// find keys
		var re = regexp.MustCompile(`ZONE(\d+).*ri:\s+(\d+)`)
		// var re = regexp.MustCompile(`(ri):\s*(\d+)`)

		for _, line := range lines[startIndex:] {
			matches := re.FindStringSubmatch(line)

			if len(matches) >= 3 {
				zone, _ := strconv.Atoi(matches[1])
				key, _ := strconv.Atoi(matches[2])

				letter := getLetter(zone, key)

				if zone == 528 || zone == 533 {
					fmt.Printf("E_%d: %s - %d\n", zone, letter, key)
				}
			}
		}

		// Clear output and find new key
		fmt.Print("Press ENTER to refresh...")
		fmt.Scanln()
		// time.Sleep(5 * time.Second)
		fmt.Print("\033[H\033[2J")
	}
}

func getLetter(zone, key int) string {
	var ranges []zone_areas.RangeLetter

	switch zone {
	case 528:
		ranges = zone_areas.E_528
	case 533:
		ranges = zone_areas.E_533
	default:
		ranges = zone_areas.E_DEFAULT
	}

	for _, rl := range ranges {
		if key >= rl.Min && key <= rl.Max {
			return rl.Letter
		}
	}
	return "?"
}
