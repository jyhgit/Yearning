package commom

import (
	"github.com/jinzhu/gorm"
	"reflect"
)

const QueryField = "work_id, username, text, backup, date, real_name, executor, `status`, `type`, `delay`, `source`,`id_c`,`data_base`,`table`,`execute_time`,assigned,current_step,relevant"

func AccordingToWorkId(workId string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if workId == "" {
			 return db
		}
		return db.Where("work_id like ?", "%"+workId+"%")
	}
}

func AccordingToQueryPer() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("`query_per` in (?)", []int{1, 3})
	}
}

func AccordingToOrderState() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("`status` in (?)", []int{1, 4})
	}
}

func AccordingToAssigned(user string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("`assigned` = ?", user)
	}
}

func AccordingToUsername(user string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if user == "" {
			return db
		}
		return db.Where("username like ?", "%"+user+"%")
	}
}

func AccordingToDatetime(time []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if reflect.DeepEqual(time, []string{"", ""}) {
			return db
		}
		return db.Where("time >= ? AND time <= ?", time[0], time[1])
	}
}

func AccordingToDate(time []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if reflect.DeepEqual(time, []string{"", ""}) {
			return db
		}
		return db.Where("date >= ? AND date <= ?", time[0], time[1])
	}
}

func AccordingToRelevant(user string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		return db.Where("JSON_SEARCH(relevant, 'one', ?) IS NOT NULL", user)
	}
}

func AccordingToGuest(user string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("username = ?", user)
	}
}

func AccordingToText(text string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if text == "" {
			return db
		}
		return db.Where("text like ?", "%"+text+"%")
	}
}

func AccordingToOrderName(text string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if text == "" {
			return db
		}
		return db.Where("`name` like ?", "%"+text+"%")
	}
}

func AccordingToOrderIDC(text string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if text == "" {
			return db
		}
		return db.Where("id_c LIKE ? ", "%"+text+"%")
	}
}

func AccordingToOrderSource(text string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if text == "" {
			return db
		}
		return db.Where("`source` LIKE ?", "%"+text+"%")
	}
}

func AccordingToOrderDept(text string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if text == "" {
			return db
		}
		return db.Where("department LIKE ?", "%"+text+"%")
	}
}

func AccordingToRuleSuperOrAdmin() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("rule in (?)", []string{"admin", "super"})
	}
}