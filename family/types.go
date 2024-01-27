package family

type relationship struct {
	Sons      []int `yaml:"sons,omitempty"`
	Daughters []int `yaml:"daughters,omitempty"`
	Wives     []int `yaml:"wives,omitempty"`
	Father    int   `yaml:"father,omitempty"`
}

type person struct {
	Id            int          `yaml:"id"`
	Name          string       `yaml:"name"`
	Relationships relationship `yaml:"relationships,omitempty"`
}

type people struct {
	Person []person `yaml:"person,omitempty"`
}
