package model

type Stream struct {
	*Base
	Name     string `json:"name"`
	IntakeID int64  `json:"intakeID"`
	Intake   Intake
}

type Intake struct {
	*Base
	GraduationYear         int64    `json:"graduationYear"`
	OriginalGraduationYear int64    `json:"OriginalGraduationYear"`
	GraduationMonth        int64    `json:"graduationMonth"`
	Streams                []Stream `gorm:"ForeignKey:IntakeID" json:"streams"`
}

/*func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Intake{}, &Stream{})
	AddForeignKey(db, &Stream{}, "intake_id", "intakes(id)", "RESTRICT", "CASCADE")
}
*/
