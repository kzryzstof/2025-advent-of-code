package abstractions

type Section struct {
	/* Represents a section within a department, consisting of at 2 or 3 rows. */
	Rows []Row
	/* Index of the section within the department to analyze */
	RowIndex int
}
