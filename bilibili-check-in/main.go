package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
)

type Response struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	TTL     int64  `json:"ttl"`
	Data    RData   `json:"data"`
}

type RData struct {
	Following    int64 `json:"following"`
	Follower     int64 `json:"follower"`
	DynamicCount int64 `json:"dynamic_count"`
}

// Generated by https://quicktype.io

type NavResponse struct {
	Code    int64  `json:"code"`   
	Message string `json:"message"`
	TTL     int64  `json:"ttl"`    
	Data    Data   `json:"data"`   
}

type Data struct {
	IsLogin            bool           `json:"isLogin"`             
	EmailVerified      int64          `json:"email_verified"`      
	Face               string         `json:"face"`                
	FaceNft            int64          `json:"face_nft"`            
	LevelInfo          LevelInfo      `json:"level_info"`          
	Mid                int64          `json:"mid"`                 
	MobileVerified     int64          `json:"mobile_verified"`     
	Money              float64        `json:"money"`               
	Moral              int64          `json:"moral"`               
	Official           Official       `json:"official"`            
	OfficialVerify     OfficialVerify `json:"officialVerify"`      
	Pendant            Pendant        `json:"pendant"`             
	Scores             int64          `json:"scores"`              
	Uname              string         `json:"uname"`               
	VipDueDate         int64          `json:"vipDueDate"`          
	VipStatus          int64          `json:"vipStatus"`           
	VipType            int64          `json:"vipType"`             
	VipPayType         int64          `json:"vip_pay_type"`        
	VipThemeType       int64          `json:"vip_theme_type"`      
	VipLabel           Label          `json:"vip_label"`           
	VipAvatarSubscript int64          `json:"vip_avatar_subscript"`
	VipNicknameColor   string         `json:"vip_nickname_color"`  
	Vip                Vip            `json:"vip"`                 
	Wallet             Wallet         `json:"wallet"`              
	HasShop            bool           `json:"has_shop"`            
	ShopURL            string         `json:"shop_url"`            
	AllowanceCount     int64          `json:"allowance_count"`     
	AnswerStatus       int64          `json:"answer_status"`       
}

type LevelInfo struct {
	CurrentLevel int64 `json:"current_level"`
	CurrentMin   int64 `json:"current_min"`  
	CurrentExp   int64 `json:"current_exp"`  
	NextExp      int64 `json:"next_exp"`     
}

type Official struct {
	Role  int64  `json:"role"` 
	Title string `json:"title"`
	Desc  string `json:"desc"` 
	Type  int64  `json:"type"` 
}

type OfficialVerify struct {
	Type int64  `json:"type"`
	Desc string `json:"desc"`
}

type Pendant struct {
	PID               int64  `json:"pid"`                
	Name              string `json:"name"`               
	Image             string `json:"image"`              
	Expire            int64  `json:"expire"`             
	ImageEnhance      string `json:"image_enhance"`      
	ImageEnhanceFrame string `json:"image_enhance_frame"`
}

type Vip struct {
	Type               int64  `json:"type"`                
	Status             int64  `json:"status"`              
	DueDate            int64  `json:"due_date"`            
	VipPayType         int64  `json:"vip_pay_type"`        
	ThemeType          int64  `json:"theme_type"`          
	Label              Label  `json:"label"`               
	AvatarSubscript    int64  `json:"avatar_subscript"`    
	NicknameColor      string `json:"nickname_color"`      
	Role               int64  `json:"role"`                
	AvatarSubscriptURL string `json:"avatar_subscript_url"`
}

type Label struct {
	Path        string `json:"path"`        
	Text        string `json:"text"`        
	LabelTheme  string `json:"label_theme"` 
	TextColor   string `json:"text_color"`  
	BgStyle     int64  `json:"bg_style"`    
	BgColor     string `json:"bg_color"`    
	BorderColor string `json:"border_color"`
}

type Wallet struct {
	Mid           int64 `json:"mid"`            
	BcoinBalance  int64 `json:"bcoin_balance"`  
	CouponBalance int64 `json:"coupon_balance"` 
	CouponDueTime int64 `json:"coupon_due_time"`
}


var cookie string

func init() {
	flag.StringVar(&cookie, "c", "", "cookie")
}

func data() error {
	req, err := http.NewRequest(http.MethodGet, "https://api.bilibili.com/x/web-interface/nav", nil)
	if err != nil {
		return err
	}
	req.Header.Add("Cookie", cookie)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var res NavResponse

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return err
	}

	if res.Code != 0 {
		return errors.New("Not login, please provide lastest cookie again")
	}
	
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent(" ", "  ")
	return encoder.Encode(res)
}

func CheckIn() error {
	req, err := http.NewRequest(http.MethodGet, "https://api.bilibili.com/x/web-interface/nav/stat", nil)
	if err != nil {
		return err
	}
	req.Header.Add("Cookie", cookie)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var res Response

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return err
	}

	if res.Code != 0 {
		return errors.New("Not login, please provide lastest cookie again")
	}
	
	fmt.Printf("Res: %#v\n", res)
	return nil
}

func main() {
	flag.Parse()
	err := CheckIn()
	if err != nil {
		fmt.Println("Check in failed: ", err)
		return
	}
	fmt.Println("Check In Success")
	data()
}
