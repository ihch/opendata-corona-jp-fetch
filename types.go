package root

/*
 * 型定義
 */
type APICovidItem struct {
  Date string `json:"date"`;
  NameJp string `json:"name_jp"`;
  Npatients string `json:"npatients"`;
}

type APICovidItems = []APICovidItem;

type APICovidItemList struct {
  ItemList APICovidItems `json:"itemList"`;
}

type CovidItem struct {
  Date string `json:"date"`;
  Prefecture string `json:"prefecture"`;
  Patients int `json:"patients"`;
  TotalPatients int `json:"totalPatients"`;
}

type CovidItems = []CovidItem;
