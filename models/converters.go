package models

// LessonPlayerToPlayer converts a LessonPlayer to a Player
func LessonPlayerToPlayer(lp *LessonPlayer) Player {
	return Player{
		BasePlayer: BasePlayer{
			UserID:     lp.UserID,
			LevelValue: lp.LevelValue,
			Picture:    lp.Picture,
		},
		Name: lp.FullName,
		// Other fields left with zero values
	}
}

// LessonTenantToTenant converts a LessonTenant to a Tenant
func LessonTenantToTenant(lt *LessonTenant) Tenant {
	return Tenant{
		TenantID:   lt.TenantID,
		TenantName: lt.TenantName,
		Address:    lt.TenantAddress,
		Images:     lt.TenantImages,
		Properties: lt.Properties,
		// PlaytomicStatus is omitted as it's not available in LessonTenant
	}
}
