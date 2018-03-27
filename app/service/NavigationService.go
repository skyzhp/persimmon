package service

import (
	"encoding/json"
	"github.com/cong5/persimmon/app/info"
)

type NavigationService struct{}

func (this *NavigationService) GetNavigation() ([]info.Navigation, error) {
	navigation, navErr := optionService.GetValueByName(this.GetNavKey())
	if navErr != nil {
		return nil, navErr
	}

	NavigationMenu := make([]info.Navigation, 0)
	err := json.Unmarshal([]byte(navigation), &NavigationMenu)
	if err != nil {
		//revel.INFO.Printf("Json unmarshal failed. data: %s : error: %s", navigation, err)
		return nil, err
	}

	return NavigationMenu, nil
}

func (this *NavigationService) GetNavKey() string {
	return "navigations"
}
