package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// run a one on one match
// start with the lowest scored robot and work up
func Match2() {
	// Tournament 2x2

	match := make([]thematch, 0)

	for i := 0; i < GameInfo.NumOfBots-1; i++ {
		for j := i + 1; j < GameInfo.NumOfBots; j++ {

			m := thematch{}
			m.RobotFilename = append(m.RobotFilename, GameInfo.RobotStorage[i].Filename)
			m.RobotFilename = append(m.RobotFilename, GameInfo.RobotStorage[j].Filename)

			buf := new(bytes.Buffer) // buffer to capture the output of the basicbots program

			// print the match
			fmt.Printf("Tournament: %s vs %s\n",
				GameInfo.RobotStorage[i].Filename,
				GameInfo.RobotStorage[j].Filename)

			matches := "17" // number of matches to run
			cmd := exec.Command(GameInfo.BBprogram, "-tt", "-m", matches,
				GameInfo.Robotspath+"/"+GameInfo.RobotStorage[i].Filename,
				GameInfo.Robotspath+"/"+GameInfo.RobotStorage[j].Filename)

			cmd.Stdout = buf

			err := cmd.Run()
			if err != nil {
				fmt.Println("Error on Run")
				log.Fatal(err)
			}
			lines := strings.Split(buf.String(), "\n")

			for _, line := range lines {

				fmt.Println(line)
				if len(line) == 0 {
					continue
				}
				parts1 := strings.Split(line, " ")
				filename := parts1[0]
				if len(filename) == 0 {
					log.Fatalln("Error on filename, during select")
				}
				w := parts1[1]
				t := parts1[2]
				l := parts1[3]
				p := parts1[4]

				// this update just robots table
				sqlstring := "UPDATE robots SET count = count + " + matches + ", win = win + " + w + ", tie = tie + " + t + ", loss = loss + " + l + ", points = points + " + p + " WHERE filename = '" + filename + "';"
				_, err := db.Db.Exec(sqlstring)
				if err != nil {
					fmt.Fprintln(os.Stderr, sqlstring)
					log.Fatal(err)
				}

				//parts2 := strings.Split(parts1[1], " ")
				//parts3 := strings.Split(parts2[0], ":")
				//whoami := hashing.StringHash(hashing.SHA256, parts1[0])

				//tt := db.Db.ReadStr(whoami)
				// ttu := therobots{}
				// err := json.Unmarshal([]byte(tt), &ttu)
				// if err != nil {
				// 	fmt.Println("Error on Unmarshal")
				// 	log.Fatal(err)
				// }
				// ttu.Count++
				// w, _ := strconv.Atoi(parts3[1])
				// t, _ := strconv.Atoi(parts3[3])
				// l, _ := strconv.Atoi(parts3[5])
				// p, _ := strconv.ParseFloat(parts3[7], 64)
				// ttu.Win += w
				// ttu.Tie += t
				// ttu.Loss += l
				// ttu.Points += p

			}

			fmt.Println(buf.String())
		}
	}

}
