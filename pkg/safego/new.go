package safego

func New(run func(), callback func(err any)) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				callback(err)
				return
			}
		}()
		run()
	}()
}
