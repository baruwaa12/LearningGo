package erratum

func Use(o ResourceOpener, input string) (err error) {
	var res Resource
	for {
		res, err = o()
		if err == nil {
			break
		}
		if _, ok := err.(TransientError); !ok {
			return
		}
	}
	defer res.Close()

	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(FrobError); ok {
				res.Defrob.(e.defrobTag)
			}
			err = r.(error)
		}
	}()
	res.Frob(input)

	return

}