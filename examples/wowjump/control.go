package main

import "log"

func process() {
	defer func() {
		err := recover()
		log.Println("process end??", err)
		if err != nil {
			go process()
		}
	}()
	var selfLogout = make([]*Logout, 0, 10)
	for {
		mux.Lock()
		if !running || logouts == nil || len(logouts) == 0 {
			mux.Unlock()
			sleep(1000)
			continue
		}
		selfLogout = make([]*Logout, len(logouts))
		copy(selfLogout, logouts)
		mux.Unlock()
		for _, logout := range selfLogout {
			// 处理逻辑
			if !logout.IsValid() {
				continue
			}
			logout.Update()
		}
		randomSleep(5000, 2500)
	}
}
