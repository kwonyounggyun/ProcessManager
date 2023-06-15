package db

import (
	"ProcessManager/db/dbtask"
	"container/list"
	"database/sql"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type DBManager struct {
	tasks *list.List
	db    *sql.DB
	lock  *sync.Mutex
}

func CreateManager(driver_name, connect_str string) (*DBManager, error) {
	manager := &DBManager{}
	db, err := sql.Open(driver_name, connect_str)
	manager.db = db
	manager.tasks = list.New()
	manager.lock = &sync.Mutex{}
	return manager, err
}

func (m *DBManager) Run(count int) {
	go func() {
		wg := &sync.WaitGroup{}
		wg.Add(count)
		for i := 0; i < count; i++ {
			go func() {
				defer wg.Done()

				for {
					task := m.popTask()
					if task != nil {
						//connection close
						if task.Execute(m.db) == sql.ErrConnDone {
							return
						}
						task.Complete()
					}
				}
			}()
		}
		wg.Wait()
	}()
}

func (m *DBManager) popTask() dbtask.IDBTask {
	m.lock.Lock()
	defer m.lock.Unlock()
	va := m.tasks.Front()
	if va != nil {
		return m.tasks.Remove(va).(dbtask.IDBTask)
	}
	return nil
}

func (m *DBManager) PushTask(task dbtask.IDBTask) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.tasks.PushBack(task)
}
