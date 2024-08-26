package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// run a one on one match
// start with the lowest scored robot and work up
func Match2() {
	// Tournament 2x2

	match := make([]match2, 0)

	for i := 0; i < GameInfo.NumOfBots-1; i++ {
		for j := i + 1; j < GameInfo.NumOfBots; j++ {

			m := match2{}
			m.Robot1 = GameInfo.RobotStorage[i].ID
			m.Robot2 = GameInfo.RobotStorage[j].ID

			buf := new(bytes.Buffer) // buffer to capture the output of the basicbots program

			// print the match
			fmt.Printf("Tournament: %s vs %s\n",
				GameInfo.RobotStorage[i].Filename,
				GameInfo.RobotStorage[j].Filename)

			m.Robot1 = GameInfo.RobotStorage[i].Filename
			m.Robot2 = GameInfo.RobotStorage[j].Filename

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

			for pid, line := range lines {

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

				switch pid {
				case 0:
					m.Win1, _ = strconv.Atoi(w)
					m.Tie1, _ = strconv.Atoi(t)
					m.Lose1, _ = strconv.Atoi(l)
					m.Points1, _ = strconv.ParseFloat(p, 64)
				case 1:
					m.Win2, _ = strconv.Atoi(w)
					m.Tie2, _ = strconv.Atoi(t)
					m.Lose2, _ = strconv.Atoi(l)
					m.Points2, _ = strconv.ParseFloat(p, 64)
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

			sqlstring := "UPDATE match2 SET Robot1='" + m.Robot1 + "', Robot2='" + m.Robot2 + "', Win1=" + strconv.Itoa(m.Win1) + ", Win2=" + strconv.Itoa(m.Win2) + ", Tie1=" + strconv.Itoa(m.Tie1) + ", Tie2=" + strconv.Itoa(m.Tie2) + ", Lose1=" + strconv.Itoa(m.Lose1) + ", Lose2=" + strconv.Itoa(m.Lose2) + ", Points1=" + strconv.FormatFloat(m.Points1, 'f', -1, 64) + ", Points2=" + strconv.FormatFloat(m.Points2, 'f', -1, 64) + " WHERE Robot1='" + m.Robot1 + "' AND Robot2='" + m.Robot2 + "';"
			_, err := db.Db.Exec(sqlstring)
			if err != nil {
				fmt.Fprintln(os.Stderr, sqlstring)
				log.Fatal(err)
			}

		}
	}

}
