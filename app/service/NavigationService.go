package service

import (
	"github.com/cong5/persimmon/app/info"
	"github.com/cong5/persimmon/app/utils"
	"github.com/revel/revel/cache"
	"github.com/revel/revel"
	"encoding/json"
	"time"
)

type NavigationService struct {
	NavKey string
}

func (this *NavigationService) GetNavigation(real bool) ([]info.Navigation, error) {
	NavigationMenu := make([]info.Navigation, 0)
	cKey := utils.CacheKey("NavigationService", "Navigation")

	if real {
		go cache.Delete(cKey)
	}

	if err := cache.Get(cKey, &NavigationMenu); err != nil {
		navigation, navErr := optionService.GetValueByName(this.GetNavKey(),false)
		if navErr != nil {
			return nil, navErr
		}

		err := json.Unmarshal([]byte(navigation), &NavigationMenu)
		if err != nil {
			revel.INFO.Printf("Json unmarshal failed. data: %s : error: %s", navigation, err)
			return nil, err
		}
		go cache.Set(cKey, NavigationMenu, 30*time.Minute)
	}

	return NavigationMenu, nil
}

func (this *NavigationService) GetNavKey() string {
	return "navigations"
}
