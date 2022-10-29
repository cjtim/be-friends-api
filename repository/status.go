package repository

type Status string

var (
	NEW       Status = "NEW"
	REVIEWING Status = "กำลังตรวจสอบข้อมูล"
	ADOPED    Status = "เสร็จสิ้น"
)
