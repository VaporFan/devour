/*
 * mtStats Devour
 */

package main

import (
	"container/heap"
	"fmt"
	"github.com/mtstats/devour/controllers"
	"github.com/mtstats/devour/utils/config"
	"github.com/mtstats/devour/utils/db"
	"github.com/mtstats/devour/webapi"
	"time"
)

var (
	lastSeqNum = 1520885231
)

func main() {

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	q := make(Queue, 1)
	q[0] = &Job{
		task:     "get_heroes",
		value:    "3423423423",
		priority: 1,
		index:    0,
	}
	heap.Init(&q)

	// Insert a new job and then modify its priority.
	job := &Job{
		task:     "get_items",
		value:    "new_value",
		priority: 2,
	}
	heap.Push(&q, job)

	// Setup Config Settings
	config := config.LoadConfig()

	// Setup Steam API Handler
	api := webapi.LoadApi(config)

	// Setup DB Connection
	db := db.OpenDb(config)

	last, err := controllers.GetLastSeqNum(db)
	if err == nil {
		lastSeqNum = last
	}

	ticker := time.NewTicker(time.Millisecond * 1000) // 1 second
	go func() {
		//func() {
		for range ticker.C {

			//fmt.Println("Tick at ", t)

			// Check if there are any Jobs in the Queue
			if q.Len() > 0 {

				// Got a Job, lets do it.
				job = heap.Pop(&q).(*Job)
				fmt.Printf("%.2d:%s \n", job.priority, job.task)

				if job.task == "get_heroes" {
					fmt.Println("Updating Heroes...")
					err := controllers.GetHeroes(db, api)
					if err != nil {
						continue
					}
				}

				if job.task == "get_items" {
					fmt.Println("Updating Game Items...")
					err := controllers.GetGameItems(db, api)
					if err != nil {
						continue
					}
				}

			} else {

				duration := time.Duration(1) * time.Second
				time.Sleep(duration)

				// No Jobs, lets update matches.
				fmt.Println("Match History Update...")
				err := controllers.ContMatchHistorySeq(db, api, &lastSeqNum)
				if err != nil {
					continue
				}

			}

		}
	}()
	//time.Sleep(time.Second * 20)
	//ticker.Stop()
	//fmt.Println("Ticker stopped")

	ticker2 := time.NewTicker(time.Second * 60) // 60 seconds
	// go func() {
	func() {
		for range ticker2.C {

			/*job := &Job{
				task:     "get_itemszz",
				value:    "12",
				priority: 1,
			}
			heap.Push(&q, job)*/

		}
	}()

}
