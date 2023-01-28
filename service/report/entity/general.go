package entity

type GeneralDetail []byte

func (t *GeneralDetail) MarshalJSON() ([]byte, error) {
	return *t, nil
}

func (t *GeneralDetail) UnmarshalJSON(b []byte) error {
	*t = b
	return nil
}
