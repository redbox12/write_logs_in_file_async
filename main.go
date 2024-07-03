package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

var actions = []string{
	"logged in",
	"logged out",
	"create record",
	"delete record",
	"update record",
}

type logItem struct {
	action    string
	timestamp time.Time
}

type User struct {
	id    int
	email string
	logs  []logItem
}

func (u User) getInfoAction() string {
	out := fmt.Sprintf("ID:%d | Email: %s\nActivity Log:\n", u.id, u.email)
	for i, item := range u.logs {
		out += fmt.Sprintf("%d. [%s] at %s\n", i+1, item.action, item.timestamp)
	}
	return out
}

/*
before using the goroutine time: 14.778768872s
after using the goroutine, time: 2.482559738s.
*/

func main() {
	rand.Seed(time.Now().Unix())
	t := time.Now()

	wg := &sync.WaitGroup{}

	users := generateUsers(1000)
	for _, user := range users{
		wg.Add(1)
		go saveUserInfo(user, wg)	
		
	}

	wg.Wait()

	fmt.Println("TIME ELAPSED:", time.Since(t).String())
}

func saveUserInfo(user User, wg *sync.WaitGroup) error{
	time.Sleep(time.Millisecond *10) //simulating a network request
	fmt.Printf("WRITING FILE FOR USER ID: %d\n", user.id)
	filename := fmt.Sprintf("logs/uid_%d", user.id)

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	_, err = file.WriteString(user.getInfoAction())
	if err != nil {
		return err
	}

	wg.Done()
	return nil

}

func generateUsers(count int) []User{
	users := make([]User, count)
	for i:=0; i<count; i++{
		users[i] = User{
			id: i+1,
			email: fmt.Sprintf("user_email%d@gmail.com", i+1),
			logs: generateLogs(rand.Intn(1000)),
		}
	}
	return users
}

func generateLogs(count int) []logItem {
	logs := make([]logItem, count)

	for i := 0; i < count; i++{
		logs[i] = logItem{
			timestamp: time.Now(),
			action: actions[rand.Intn(len(actions)-1)],
		}
	}
	return logs
}

