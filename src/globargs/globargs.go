package globargs

type Args []string

func (i *Args) String() string {
    return "test"
}

func (i *Args) Set(value string) error {
    *i = append(*i, value)
    return nil
}