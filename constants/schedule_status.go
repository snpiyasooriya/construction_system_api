package constants

type ScheduleStatus string

const (
	// ScheduleStatusDraft indicates the schedule is in draft state
	ScheduleStatusDraft ScheduleStatus = "DRAFT"
	// ScheduleStatusPending indicates the schedule is pending review
	ScheduleStatusPending ScheduleStatus = "PENDING"
	// ScheduleStatusReviewed indicates the schedule has been reviewed
	ScheduleStatusReviewed ScheduleStatus = "REVIEWED"
	// ScheduleStatusApproved indicates the schedule has been approved
	ScheduleStatusApproved ScheduleStatus = "APPROVED"
	// ScheduleStatusRejected indicates the schedule has been rejected
	ScheduleStatusRejected ScheduleStatus = "REJECTED"
	// ScheduleStatusCancelled indicates the schedule has been cancelled
	ScheduleStatusCancelled ScheduleStatus = "CANCELLED"
)
