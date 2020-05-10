package erratum

func Use(o ResourceOpener, input string) (err error) {}
	r, err := o()
	if err != nil {
		if _, ok := err.(TransientError); !ok {
			err = Use(o, input)
		}
		return
	}
	defer res.Close()

	defer func() {
		if rec := recover(); r != nil {
			if e, ok := r.(FrobError); ok {
				r.Defrob.(e.defrobTag)
			}
			err = rec.(error)
		}
	}()

	r.Frob(input)
	return
}